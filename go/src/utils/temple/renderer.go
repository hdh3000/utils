// package temple provides a means of stringing together predefined Components + data into
// web pages. It provides an api to produce a top level Page, and the Components that live on it.
// It's capable of doing anything you could do with go-templates. Just a bit easier to use as a view
// when apps become more complex.
//
package temple

import (
	"bytes"
	"html/template"
)

// Page represents a single page found at some url + its resources
// Generally the default page should be used for all pages, unless you have
// some good reason not to.
type Page struct {
	tmpl     *template.Template
	Funcs    template.FuncMap
	Title    string
	Children []*Component
	JS       []string
	CSS      []string
	MetaTags []struct {
		Name    string
		Content string
	}
}

func NewPage(title string) (*Page, error) {
	tmpl, err := template.New("page").Funcs(DefaultFuncs).Parse(baseTmpl)
	if err != nil {
		return nil, err
	}

	return &Page{
		tmpl:  tmpl,
		Title: title,
		Funcs: DefaultFuncs,
	}, nil
}

func (p *Page) AddJS(imports ...string) {
	p.JS = append(p.JS, imports...)
}

func (p *Page) AddCSS(imports ...string) {
	p.JS = append(p.CSS, imports...)
}

func (p *Page) AddHTMLMetaTag(name, content string) {
	p.MetaTags = append(p.MetaTags, struct {
		Name    string
		Content string
	}{Name: name, Content: content})
}

func (p *Page) Render() (string, error) {
	// We want to be careful not to actually change the Page itself,
	// in the act of rendering it.
	css := append([]string{}, p.CSS...)
	js := append([]string{}, p.JS...)

	for _, child := range p.Children {
		css = append(css, child.collectCSS()...)
		js = append(js, child.collectJS()...)
	}

	return "", nil
}

// Component encapsulates a distinct component and its child components (e.g. a widget on a Page).
// the only requirement of a component is that it produce a valid html fragment (e.g. '<div>...</div>')
// js, and style sheet paths are passed up and imported by the page.
type Component struct {
	Template *template.Template
	Funcs    template.FuncMap
	JS       []string
	CSS      []string
	Name     string
	Args     interface{}
	Children []*Component
}

func (p *Page) AddChildComponent(component *Component) {
	p.Children = append(p.Children, component)
}

func (c *Component) AddChildComponent(component *Component) {
	c.Children = append(c.Children, component)
}

func NewComponent(name, tmpl string, css, js []string, funcs template.FuncMap, children ...*Component) (*Component, error) {
	t, err := template.New(name).Funcs(funcs).Parse(tmpl)
	if err != nil {
		return nil, err
	}

	return &Component{
		Template: t,
		Funcs:    funcs,
		JS:       js,
		CSS:      css,
		Children: children,
	}, nil
}

func (c *Component) render(args interface{}) (string, error) {
	out := &bytes.Buffer{}
	if err := c.Template.Execute(out, args); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (c *Component) collectCSS() []string {
	out := append([]string{}, c.CSS...)
	for _, child := range c.Children {
		out = append(out, child.collectCSS()...)
	}
	return out
}

func (c *Component) collectJS() []string {
	out := append([]string{}, c.JS...)
	for _, child := range c.Children {
		out = append(out, child.collectJS()...)
	}
	return out
}
