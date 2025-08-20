package templates

import (
	"text/template"

	"github.com/riahimedyassin/curly/internal/interfaces"
)

type ComponentTemplateParser struct {
	configLoader *interfaces.ConfigResolver
}

func NewComponentTemplateParser(configLoader *interfaces.ConfigResolver) *ComponentTemplateParser {
	return &ComponentTemplateParser{
		configLoader: configLoader,
	}
}

func (c *ComponentTemplateParser) ParseTemplate() (string, error) {
	return "", nil
}

// todo : Hardocded value should be parsed to dynamic.
// todo : Setup a clear folder structutre for templates in the future. (files folder)
func (c *ComponentTemplateParser) load() (*template.Template, error) {
	template := template.Must(template.ParseFiles("templates/react/files/component.tmpl"))
	return template, nil
}
