package goml

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type SuspenseContext struct {
	out chan GomlElement
	wg  *sync.WaitGroup
}

func (s *SuspenseContext) Stream(w http.ResponseWriter) {
	w.(http.Flusher).Flush()

	go func() {
		s.wg.Wait()
		close(s.out)
	}()

	for v := range s.out {
		v.RenderElement(w)
		w.(http.Flusher).Flush()
	}
}

func NewSuspenseContext() SuspenseContext {
	return SuspenseContext{
		out: make(chan GomlElement),
		wg:  &sync.WaitGroup{},
	}
}

func suspensePlaceholderWrapper(placeholder GomlElement, id string) GomlElement {
	return Div(
		Attr("id", id),
		placeholder,
	)
}

func suspenseResult(element GomlElement, id string) GomlElement {
	return Fragment(
		Template(
			IDAttr(fmt.Sprintf("%s-templ", id)),
			StyleAttr("display: none;"),
			element,
		),
		Script(
			IDAttr(fmt.Sprintf("%s-script", id)),
			Raw(fmt.Sprintf(`
(() => {
const tmpl = document.getElementById("%s-templ");
document.getElementById("%s").outerHTML = tmpl.innerHTML;
tmpl.remove();
document.getElementById("%s-script").remove();
})()
	    `, id, id, id)),
		),
	)
}

type SuspenseElement struct {
	sc          SuspenseContext
	placeholder GomlElement
	blocking    func() GomlElement
}

func (e SuspenseElement) RenderElement(w io.Writer) error {
	e.sc.wg.Add(1)
	id := uuid.New().String()

	go func() {
		defer e.sc.wg.Done()
		e.sc.out <- suspenseResult(e.blocking(), id)
	}()

	return suspensePlaceholderWrapper(e.placeholder, id).RenderElement(w)
}

func Suspense(sc SuspenseContext, placeholder GomlElement, blocking func() GomlElement) GomlElement {
	return SuspenseElement{sc: sc, placeholder: placeholder, blocking: blocking}
}
