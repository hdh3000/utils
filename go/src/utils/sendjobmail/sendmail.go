package sendjobmail

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
)

// TmplArgs are the set of arguments available to all templates.
type TmplArgs struct {
	To   *BasicInfo
	From *BasicInfo
}

type BasicInfo struct {
	First       string
	Last        string
	CompanyName string
	Title       string
}

func ParseTmpl(name string, tmplDir string, args *TmplArgs) (string, error) {
	path := path.Join(tmplDir, fmt.Sprintf("%s.tmpl", name))

	buf := &bytes.Buffer{}
	if err := template.Must(template.ParseFiles(path)).Execute(buf, args); err != nil {
		return "", err
	}

	return buf.String(), nil
}
