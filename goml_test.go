package goml_test

import (
	"bytes"
	"testing"

	g "github.com/joetifa2003/goml"

	gm "github.com/maragudk/gomponents"
	gmh "github.com/maragudk/gomponents/html"
)

func generateSlice() []int {
	s := make([]int, 100)

	for i := 0; i < len(s); i++ {
		s[i] = i
	}

	return s
}

func BenchmarkVsGomponents(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		b.Run("goml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bytes.Buffer{}

				el := g.Div(
					g.Div(g.Img(g.Attr("src", "foobar"))),
					g.Div(g.Div(g.Span(g.Text("barbaz")))),
				)

				err := el.RenderElement(&b)
				if err != nil {
					panic(err)
				}
			}
		})

		b.Run("gomponents", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bytes.Buffer{}

				el := gmh.Div(
					gmh.Div(gmh.Img(gmh.Src("goobar"))),
					gmh.Div(gmh.Div(gmh.Span(gm.Text("barbaz")))),
				)

				err := el.Render(&b)
				if err != nil {
					panic(err)
				}
			}
		})
	})

	b.Run("maps", func(b *testing.B) {
		s := generateSlice()

		b.Run("goml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bytes.Buffer{}

				el := g.Div(
					g.Img(g.Attr("src", "hi")),
					g.Text("Hi"),
					g.Div(
						g.Map(s, func(i, v int) g.GomlElement {
							return g.Div(
								g.Text("child"),
								g.If(i%2 == 0, g.Div(g.Text("only even"))),
							)
						}),
					),
				)
				err := el.RenderElement(&b)
				if err != nil {
					panic(err)
				}
			}
		})

		b.Run("gomponents", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bytes.Buffer{}

				el := gmh.Div(
					gmh.Img(gmh.Src("hi")),
					gm.Text("Hi"),
					gmh.Div(
						gm.Group(gm.Map(s, func(i int) gm.Node {
							return gmh.Div(
								gm.Text("child"),
								gm.If(i%2 == 0, gmh.Div(gm.Text("only even"))),
							)
						})),
					),
				)

				err := el.Render(&b)
				if err != nil {
					panic(err)
				}
			}
		})
	})
}
