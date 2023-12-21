package goml;

import (
	"io"
)

type HtmlElement []GomlElement

func (e HtmlElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<html"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</html>"))
	if err != nil {
		return err
	}
	return nil
}

func Html(children ...GomlElement) HtmlElement {
	return HtmlElement(children)
}

type HeadElement []GomlElement

func (e HeadElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<head"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</head>"))
	if err != nil {
		return err
	}
	return nil
}

func Head(children ...GomlElement) HeadElement {
	return HeadElement(children)
}

type BodyElement []GomlElement

func (e BodyElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<body"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</body>"))
	if err != nil {
		return err
	}
	return nil
}

func Body(children ...GomlElement) BodyElement {
	return BodyElement(children)
}

type DivElement []GomlElement

func (e DivElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<div"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</div>"))
	if err != nil {
		return err
	}
	return nil
}

func Div(children ...GomlElement) DivElement {
	return DivElement(children)
}

type SpanElement []GomlElement

func (e SpanElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<span"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</span>"))
	if err != nil {
		return err
	}
	return nil
}

func Span(children ...GomlElement) SpanElement {
	return SpanElement(children)
}

type TemplateElement []GomlElement

func (e TemplateElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<template"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</template>"))
	if err != nil {
		return err
	}
	return nil
}

func Template(children ...GomlElement) TemplateElement {
	return TemplateElement(children)
}

type AElement []GomlElement

func (e AElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<a"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</a>"))
	if err != nil {
		return err
	}
	return nil
}

func A(children ...GomlElement) AElement {
	return AElement(children)
}

type PElement []GomlElement

func (e PElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<p"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</p>"))
	if err != nil {
		return err
	}
	return nil
}

func P(children ...GomlElement) PElement {
	return PElement(children)
}

type StrongElement []GomlElement

func (e StrongElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<strong"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</strong>"))
	if err != nil {
		return err
	}
	return nil
}

func Strong(children ...GomlElement) StrongElement {
	return StrongElement(children)
}

type EmElement []GomlElement

func (e EmElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<em"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</em>"))
	if err != nil {
		return err
	}
	return nil
}

func Em(children ...GomlElement) EmElement {
	return EmElement(children)
}

type H1Element []GomlElement

func (e H1Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h1"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h1>"))
	if err != nil {
		return err
	}
	return nil
}

func H1(children ...GomlElement) H1Element {
	return H1Element(children)
}

type H2Element []GomlElement

func (e H2Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h2"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h2>"))
	if err != nil {
		return err
	}
	return nil
}

func H2(children ...GomlElement) H2Element {
	return H2Element(children)
}

type H3Element []GomlElement

func (e H3Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h3"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h3>"))
	if err != nil {
		return err
	}
	return nil
}

func H3(children ...GomlElement) H3Element {
	return H3Element(children)
}

type H4Element []GomlElement

func (e H4Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h4"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h4>"))
	if err != nil {
		return err
	}
	return nil
}

func H4(children ...GomlElement) H4Element {
	return H4Element(children)
}

type H5Element []GomlElement

func (e H5Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h5"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h5>"))
	if err != nil {
		return err
	}
	return nil
}

func H5(children ...GomlElement) H5Element {
	return H5Element(children)
}

type H6Element []GomlElement

func (e H6Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<h6"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</h6>"))
	if err != nil {
		return err
	}
	return nil
}

func H6(children ...GomlElement) H6Element {
	return H6Element(children)
}

type UlElement []GomlElement

func (e UlElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<ul"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</ul>"))
	if err != nil {
		return err
	}
	return nil
}

func Ul(children ...GomlElement) UlElement {
	return UlElement(children)
}

type OlElement []GomlElement

func (e OlElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<ol"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</ol>"))
	if err != nil {
		return err
	}
	return nil
}

func Ol(children ...GomlElement) OlElement {
	return OlElement(children)
}

type LiElement []GomlElement

func (e LiElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<li"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</li>"))
	if err != nil {
		return err
	}
	return nil
}

func Li(children ...GomlElement) LiElement {
	return LiElement(children)
}

type ButtonElement []GomlElement

func (e ButtonElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<button"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</button>"))
	if err != nil {
		return err
	}
	return nil
}

func Button(children ...GomlElement) ButtonElement {
	return ButtonElement(children)
}

type LabelElement []GomlElement

func (e LabelElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<label"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</label>"))
	if err != nil {
		return err
	}
	return nil
}

func Label(children ...GomlElement) LabelElement {
	return LabelElement(children)
}

type FormElement []GomlElement

func (e FormElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<form"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</form>"))
	if err != nil {
		return err
	}
	return nil
}

func Form(children ...GomlElement) FormElement {
	return FormElement(children)
}

type SelectElement []GomlElement

func (e SelectElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<select"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</select>"))
	if err != nil {
		return err
	}
	return nil
}

func Select(children ...GomlElement) SelectElement {
	return SelectElement(children)
}

type OptionElement []GomlElement

func (e OptionElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<option"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</option>"))
	if err != nil {
		return err
	}
	return nil
}

func Option(children ...GomlElement) OptionElement {
	return OptionElement(children)
}

type HeaderElement []GomlElement

func (e HeaderElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<header"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</header>"))
	if err != nil {
		return err
	}
	return nil
}

func Header(children ...GomlElement) HeaderElement {
	return HeaderElement(children)
}

type FooterElement []GomlElement

func (e FooterElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<footer"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</footer>"))
	if err != nil {
		return err
	}
	return nil
}

func Footer(children ...GomlElement) FooterElement {
	return FooterElement(children)
}

type NavElement []GomlElement

func (e NavElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<nav"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</nav>"))
	if err != nil {
		return err
	}
	return nil
}

func Nav(children ...GomlElement) NavElement {
	return NavElement(children)
}

type TableElement []GomlElement

func (e TableElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<table"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</table>"))
	if err != nil {
		return err
	}
	return nil
}

func Table(children ...GomlElement) TableElement {
	return TableElement(children)
}

type TrElement []GomlElement

func (e TrElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<tr"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</tr>"))
	if err != nil {
		return err
	}
	return nil
}

func Tr(children ...GomlElement) TrElement {
	return TrElement(children)
}

type TdElement []GomlElement

func (e TdElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<td"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</td>"))
	if err != nil {
		return err
	}
	return nil
}

func Td(children ...GomlElement) TdElement {
	return TdElement(children)
}

type ThElement []GomlElement

func (e ThElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<th"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</th>"))
	if err != nil {
		return err
	}
	return nil
}

func Th(children ...GomlElement) ThElement {
	return ThElement(children)
}

type VideoElement []GomlElement

func (e VideoElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<video"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</video>"))
	if err != nil {
		return err
	}
	return nil
}

func Video(children ...GomlElement) VideoElement {
	return VideoElement(children)
}

type AudioElement []GomlElement

func (e AudioElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<audio"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</audio>"))
	if err != nil {
		return err
	}
	return nil
}

func Audio(children ...GomlElement) AudioElement {
	return AudioElement(children)
}

type IframeElement []GomlElement

func (e IframeElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<iframe"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</iframe>"))
	if err != nil {
		return err
	}
	return nil
}

func Iframe(children ...GomlElement) IframeElement {
	return IframeElement(children)
}

type StyleElement []GomlElement

func (e StyleElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<style"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</style>"))
	if err != nil {
		return err
	}
	return nil
}

func Style(children ...GomlElement) StyleElement {
	return StyleElement(children)
}

type ScriptElement []GomlElement

func (e ScriptElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<script"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</script>"))
	if err != nil {
		return err
	}
	return nil
}

func Script(children ...GomlElement) ScriptElement {
	return ScriptElement(children)
}

type SectionElement []GomlElement

func (e SectionElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<section"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</section>"))
	if err != nil {
		return err
	}
	return nil
}

func Section(children ...GomlElement) SectionElement {
	return SectionElement(children)
}

type ArticleElement []GomlElement

func (e ArticleElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<article"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</article>"))
	if err != nil {
		return err
	}
	return nil
}

func Article(children ...GomlElement) ArticleElement {
	return ArticleElement(children)
}

type AsideElement []GomlElement

func (e AsideElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<aside"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</aside>"))
	if err != nil {
		return err
	}
	return nil
}

func Aside(children ...GomlElement) AsideElement {
	return AsideElement(children)
}

type FigureElement []GomlElement

func (e FigureElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<figure"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</figure>"))
	if err != nil {
		return err
	}
	return nil
}

func Figure(children ...GomlElement) FigureElement {
	return FigureElement(children)
}

type FigcaptionElement []GomlElement

func (e FigcaptionElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<figcaption"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</figcaption>"))
	if err != nil {
		return err
	}
	return nil
}

func Figcaption(children ...GomlElement) FigcaptionElement {
	return FigcaptionElement(children)
}

type CiteElement []GomlElement

func (e CiteElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<cite"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</cite>"))
	if err != nil {
		return err
	}
	return nil
}

func Cite(children ...GomlElement) CiteElement {
	return CiteElement(children)
}

type TimeElement []GomlElement

func (e TimeElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<time"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</time>"))
	if err != nil {
		return err
	}
	return nil
}

func Time(children ...GomlElement) TimeElement {
	return TimeElement(children)
}

type AbbrElement []GomlElement

func (e AbbrElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<abbr"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</abbr>"))
	if err != nil {
		return err
	}
	return nil
}

func Abbr(children ...GomlElement) AbbrElement {
	return AbbrElement(children)
}

type AcronymElement []GomlElement

func (e AcronymElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<acronym"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</acronym>"))
	if err != nil {
		return err
	}
	return nil
}

func Acronym(children ...GomlElement) AcronymElement {
	return AcronymElement(children)
}

type BlockquoteElement []GomlElement

func (e BlockquoteElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<blockquote"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte(">"))
	if err != nil {
		return err
	}
	for _, el := range e {
		if el == nil {
			continue
		}
		err = el.RenderElement(w)
		if err != nil {
			return err
		}	
	}
	_, err = w.Write([]byte("</blockquote>"))
	if err != nil {
		return err
	}
	return nil
}

func Blockquote(children ...GomlElement) BlockquoteElement {
	return BlockquoteElement(children)
}

type BrElement []GomlAttr

func (e BrElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<br"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Br(attrs ...GomlAttr) BrElement {
	return BrElement(attrs)
}

type HrElement []GomlAttr

func (e HrElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<hr"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Hr(attrs ...GomlAttr) HrElement {
	return HrElement(attrs)
}

type ImgElement []GomlAttr

func (e ImgElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<img"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Img(attrs ...GomlAttr) ImgElement {
	return ImgElement(attrs)
}

type InputElement []GomlAttr

func (e InputElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<input"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Input(attrs ...GomlAttr) InputElement {
	return InputElement(attrs)
}

type MetaElement []GomlAttr

func (e MetaElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<meta"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Meta(attrs ...GomlAttr) MetaElement {
	return MetaElement(attrs)
}

type LinkElement []GomlAttr

func (e LinkElement) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<link"))
	if err != nil {
		return err
	}
	for _, ch := range e {
		if ch == nil {
			continue
		}
		if ch, ok := ch.(GomlAttr); ok {
			ch.RenderAttr(w)
		}
	}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	return nil
}

func Link(attrs ...GomlAttr) LinkElement {
	return LinkElement(attrs)
}
