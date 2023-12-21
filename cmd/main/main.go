package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	g "github.com/joetifa2003/goml"
)

type Todo struct {
	Title string
	Tags  []string
}

var todos = []Todo{
	{
		Title: "Todo 1",
		Tags:  []string{"Tag 1", "Tag 2", "Tag 3"},
	},
	{
		Title: "Todo 2",
		Tags:  []string{"Tag 1", "Tag 2", "Tag 3", "Tag 4"},
	},
}

func mainLayout(body g.GomlElement) g.GomlElement {
	return g.Page(g.Fragment(
		g.Raw(`<script src="https://unpkg.com/htmx.org@1.9.6" integrity="sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni" crossorigin="anonymous"></script>`),
	), body)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		sc := g.NewSuspenseContext()

		mainLayout(g.Div(
			g.A(g.HrefAttr("/another"), g.Text("another page")),
			g.For(1, 50, func(i int) g.GomlElement {
				return g.Suspense(sc,
					g.Div(
						g.Text(fmt.Sprintf("Loading item %d", i)),
					),
					func() g.GomlElement {
						time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
						return g.Div(g.Text(fmt.Sprintf("Item %d loaded!", i)))
					},
				)
			}),
		)).RenderElement(w)

		sc.Stream(w)
	}))

	mux.Handle("/another", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		sc := g.NewSuspenseContext()

		mainLayout(
			g.Div(
				g.A(g.HrefAttr("/"), g.Text("Go back")),
				g.Suspense(sc, g.Text("wait for it"), func() g.GomlElement {
					time.Sleep(2 * time.Second)

					return g.Div(g.Text("there you go!"))
				}),
			),
		).RenderElement(w)

		sc.Stream(w)
	}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
