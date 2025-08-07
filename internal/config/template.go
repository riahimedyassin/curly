package config

type TemplateInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

type File struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
	Condition   string `yaml:"condition"`
	Name        string `yaml:"name"`
}

type Variable[T any] struct {
	Required    bool   `yaml:"required,omitempty"`
	Description string `yaml:"description,omitempty"`
	Default     T      `yaml:"default,omitempty"`
	YamlType    string `yaml:"type"`
	GoType      T
}

type Variables struct {
	ComponentName Variable[string] `yaml:"componentName"`
	ComponentPath Variable[string] `yaml:"componentPath"`
	IncludeTests  Variable[bool]   `yaml:"includeTests"`
	IncludeProps  Variable[bool]   `yaml:"includeProps"`
	Styling       Variable[string] `yaml:"styling"`
	ExportType    Variable[string] `yaml:"exportType"`
}

type Files = []File

type TemplateConfig struct {
	Variables Variables `yaml:"variables"`
	Files     Files     `yaml:"files"`
}
