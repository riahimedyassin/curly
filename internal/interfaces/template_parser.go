package interfaces

// TemplateParser parse load and parse templates from the config.
type TemplateParser interface {
	ParseTemplate() (string, error)
}
