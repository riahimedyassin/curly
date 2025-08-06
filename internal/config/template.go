package config

type TemplateInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}
type Variable struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
	Description string `yaml:"description"`
	Default     any    `yaml:"default"`
}

type Variables = []Variable

type File struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
	Condition   string `yaml:"condition"`
}

type Files = []File

type TemplateConfig struct {
	Variables Variables `yaml:"variables"`
	Files     Files     `yaml:"files"`
}
