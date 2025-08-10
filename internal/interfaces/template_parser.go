package interfaces

type TemplateParser interface {
	ParseTemplate() (string, error)
}
