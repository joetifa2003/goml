package main

import (
	"os"
	"text/template"
)

type Element struct {
	Name   string
	GoName string
	Void   bool
}

const ElementsTemplate = `package goml;

import (
	"io"
)

{{- range .}}

{{- if .Void}}

type {{.GoName}}Element []GomlAttr
{{- else}}

type {{.GoName}}Element []GomlElement
{{- end}}

func (e {{.GoName}}Element) RenderElement(w io.Writer) error {
	_, err := w.Write([]byte("<{{.Name}}"))
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
	{{- if .Void}}
	_, err = w.Write([]byte("/>"))
	if err != nil {
		return err	
	}
	{{- else}}
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
	_, err = w.Write([]byte("</{{.Name}}>"))
	if err != nil {
		return err
	}
	{{- end}}
	return nil
}

{{- if .Void}}

func {{.GoName}}(attrs ...GomlAttr) {{.GoName}}Element {
	return {{.GoName}}Element(attrs)
}

{{- else}}

func {{.GoName}}(children ...GomlElement) {{.GoName}}Element {
	return {{.GoName}}Element(children)
}

{{- end}}

{{- end}}
`

const AttrsTemplate = `package goml

import (
	"io"
)


{{- range . -}}
{{- if .Void}}

type {{.GoName}}Attribute struct{}

func (a {{.GoName}}Attribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" {{.Name}}")) 
	if err != nil {
		return err
	}

	return nil
}

func {{.GoName}}Attr() {{.GoName}}Attribute {
	return {{.GoName}}Attribute{}
} 
{{- else}}

type {{.GoName}}Attribute struct{
	Value string
}

func (a {{.GoName}}Attribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" {{.Name}}='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func {{.GoName}}Attr(value string) {{.GoName}}Attribute {
	return {{.GoName}}Attribute{Value: value}
} 
{{- end}}

func (a {{.GoName}}Attribute) RenderElement(w io.Writer) error { return nil }
{{- end}}
`

func main() {
	elements := []Element{
		{Name: "html", GoName: "Html"},
		{Name: "head", GoName: "Head"},
		{Name: "body", GoName: "Body"},
		{Name: "div", GoName: "Div"},
		{Name: "span", GoName: "Span"},
		{Name: "template", GoName: "Template"},
		{Name: "a", GoName: "A"},
		{Name: "p", GoName: "P"},
		{Name: "strong", GoName: "Strong"},
		{Name: "em", GoName: "Em"},
		{Name: "h1", GoName: "H1"},
		{Name: "h2", GoName: "H2"},
		{Name: "h3", GoName: "H3"},
		{Name: "h4", GoName: "H4"},
		{Name: "h5", GoName: "H5"},
		{Name: "h6", GoName: "H6"},
		{Name: "ul", GoName: "Ul"},
		{Name: "ol", GoName: "Ol"},
		{Name: "li", GoName: "Li"},
		{Name: "button", GoName: "Button"},
		{Name: "label", GoName: "Label"},
		{Name: "form", GoName: "Form"},
		{Name: "select", GoName: "Select"},
		{Name: "option", GoName: "Option"},
		{Name: "header", GoName: "Header"},
		{Name: "footer", GoName: "Footer"},
		{Name: "nav", GoName: "Nav"},
		{Name: "table", GoName: "Table"},
		{Name: "tr", GoName: "Tr"},
		{Name: "td", GoName: "Td"},
		{Name: "th", GoName: "Th"},
		{Name: "video", GoName: "Video"},
		{Name: "audio", GoName: "Audio"},
		{Name: "iframe", GoName: "Iframe"},
		{Name: "style", GoName: "Style"},
		{Name: "script", GoName: "Script"},
		{Name: "section", GoName: "Section"},
		{Name: "article", GoName: "Article"},
		{Name: "aside", GoName: "Aside"},
		{Name: "figure", GoName: "Figure"},
		{Name: "figcaption", GoName: "Figcaption"},
		{Name: "cite", GoName: "Cite"},
		{Name: "time", GoName: "Time"},
		{Name: "abbr", GoName: "Abbr"},
		{Name: "acronym", GoName: "Acronym"},
		{Name: "blockquote", GoName: "Blockquote"},

		{Name: "br", GoName: "Br", Void: true},
		{Name: "hr", GoName: "Hr", Void: true},
		{Name: "img", GoName: "Img", Void: true},
		{Name: "input", GoName: "Input", Void: true},
		{Name: "meta", GoName: "Meta", Void: true},
		{Name: "link", GoName: "Link", Void: true},
	}

	attrs := []Element{
		{Name: "id", GoName: "ID"},
		{Name: "class", GoName: "Class"},
		{Name: "href", GoName: "Href"},
		{Name: "src", GoName: "Src"},
		{Name: "alt", GoName: "Alt"},
		{Name: "title", GoName: "Title"},
		{Name: "width", GoName: "Width"},
		{Name: "height", GoName: "Height"},
		{Name: "placeholder", GoName: "Placeholder"},
		{Name: "value", GoName: "Value"},
		{Name: "name", GoName: "Name"},
		{Name: "type", GoName: "Type"},
		{Name: "style", GoName: "Style"},
		{Name: "target", GoName: "Target"},
		{Name: "rel", GoName: "Rel"},
		{Name: "role", GoName: "Role"},
		{Name: "form", GoName: "Form"},
		{Name: "max", GoName: "Max"},
		{Name: "min", GoName: "Min"},
		{Name: "step", GoName: "Step"},
		{Name: "autocomplete", GoName: "AutoComplete"},
		{Name: "pattern", GoName: "Pattern"},
		{Name: "for", GoName: "For"},
		{Name: "label", GoName: "Label"},
		{Name: "disabled", GoName: "Disabled", Void: true},
		{Name: "readonly", GoName: "ReadOnly", Void: true},
		{Name: "checked", GoName: "Checked", Void: true},
		{Name: "required", GoName: "Required", Void: true},
		{Name: "multiple", GoName: "Multiple", Void: true},
	}

	generate(elements, ElementsTemplate, "elements.generated.go")
	generate(attrs, AttrsTemplate, "attrs.generated.go")
}

func generate(arr []Element, templateSrc string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmpl, err := template.New("").Parse(templateSrc)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, arr)
	if err != nil {
		panic(err)
	}
}
