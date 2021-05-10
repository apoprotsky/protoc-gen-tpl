package template

import (
	"text/template"

	gs "github.com/apoprotsky/go-services"
)

// Service struct
type Service struct {
	templates map[string]*template.Template
}

// GoService initializes service
func (svc *Service) GoService() {
	svc.templates = map[string]*template.Template{
		DefaultGoTemplate: template.Must(
			template.New(DefaultGoTemplate).Funcs(funcs).Parse(defaultGoTemplate),
		),
	}
}

// GetService returns instance of service
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
