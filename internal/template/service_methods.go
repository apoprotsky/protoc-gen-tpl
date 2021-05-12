package template

import (
	"bytes"
	"errors"
	"text/template"
)

// ExecuteTemplate by name
func (svc *Service) ExecuteTemplate(name string, data interface{}) string {
	tmpl := svc.getTemplate(name)
	var buffer bytes.Buffer
	tmpl.Execute(&buffer, data)
	return buffer.String()
}

func (svc *Service) getTemplate(name string) *template.Template {
	tmpl, ok := svc.templates[name]
	if !ok {
		panic(errors.New("template " + name + " not found"))
	}
	return tmpl
}
