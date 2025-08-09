package config

type Config struct {
	ComponentName string
	Team          TeamConfig
	Template      TemplateConfig
}

type GlobalConfig struct {
	Team         string            `yaml:"team"`
	Organization string            `yaml:"organization"`
	Template     TemplateInfo      `yaml:"template"`
	Variables    ResolvedVariables `yaml:"variables"`
	Files        Files             `yaml:"files"`
}

// ResolvedVariables contains the final variable values after applying team config
type ResolvedVariables struct {
	ComponentName ResolvedVariable[string] `yaml:"componentName"`
	ComponentPath ResolvedVariable[string] `yaml:"componentPath"`
	IncludeTests  ResolvedVariable[bool]   `yaml:"includeTests"`
	IncludeProps  ResolvedVariable[bool]   `yaml:"includeProps"`
	Styling       ResolvedVariable[string] `yaml:"styling"`
	ExportType    ResolvedVariable[string] `yaml:"exportType"`
}

// ResolvedVariable contains the final variable configuration
type ResolvedVariable[T any] struct {
	Value       T      `yaml:"value"`
	Required    bool   `yaml:"required"`
	Description string `yaml:"description"`
	Enforced    bool   `yaml:"enforced"` // true if team config overrides this
}
