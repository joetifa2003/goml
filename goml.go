package goml

import (
	"html"
	"io"

	"golang.org/x/exp/constraints"
)

type GomlElement interface {
	RenderElement(w io.Writer) error
}

type GomlAttr interface {
	RenderAttr(w io.Writer) error
}

type AttrElement struct {
	Name, Value string
}

func (e AttrElement) RenderElement(w io.Writer) error {
	return nil
}

func (e AttrElement) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" " + e.Name + "='" + e.Value + "'"))
	if err != nil {
		return err
	}

	return nil
}

func Attr(name, value string) AttrElement {
	return AttrElement{Name: name, Value: value}
}

type Text string

func (e Text) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte(html.EscapeString(string(e))))
	if err != nil {
		return err
	}

	return nil
}

type Raw string

func (e Raw) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte(e))
	if err != nil {
		return err
	}

	return nil
}

type FragmentElement []GomlElement

func (e FragmentElement) RenderElement(w io.Writer) error {
	for _, ch := range e {
		if ch == nil {
			continue
		}

		err := ch.RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func Fragment(children ...GomlElement) FragmentElement {
	return FragmentElement(children)
}

type MapElement[T any] struct {
	Arr    []T
	Mapper func(i int, v T) GomlElement
}

func (e MapElement[T]) RenderElement(w io.Writer) error {
	for i, c := range e.Arr {
		el := e.Mapper(i, c)
		if el == nil {
			continue
		}

		err := el.RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func Map[T any](arr []T, mapper func(i int, v T) GomlElement) MapElement[T] {
	return MapElement[T]{Arr: arr, Mapper: mapper}
}

type ForElement[T constraints.Integer] struct {
	Start, End T
	F          func(i T) GomlElement
}

func (e ForElement[T]) RenderElement(w io.Writer) error {
	for i := e.Start; i < e.End; i++ {
		el := e.F(i)
		if el == nil {
			continue
		}

		err := el.RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func For[T constraints.Integer](start T, end T, f func(i T) GomlElement) GomlElement {
	return ForElement[T]{Start: start, End: end, F: f}
}

func If(condition bool, children ...GomlElement) GomlElement {
	if condition {
		return Fragment(children...)
	}

	return nil
}

func Page(head, body GomlElement) GomlElement {
	return Fragment(
		Raw("<!DOCTYPE HTML>"),
		Html(
			Head(
				head,
			),
			Body(
				Attr("hx-boost", "true"),
				body,
			),
		),
	)
}
