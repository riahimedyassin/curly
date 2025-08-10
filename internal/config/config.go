package config

type primaryConfig struct {
	ComponentName string
	Team          TeamConfig
	Template      TemplateConfig
}

// Config is the final resolved project config
type Config struct {
	Team         string            `yaml:"team"`
	Organization string            `yaml:"organization"`
	Template     TemplateInfo      `yaml:"template"`
	Variables    resolvedVariables `yaml:"variables"`
	Files        Files             `yaml:"files"`
}

// ResolvedVariables contains the final variable values after applying team config
type resolvedVariables struct {
	ComponentName resolvedVariable[string] `yaml:"componentName"`
	ComponentPath resolvedVariable[string] `yaml:"componentPath"`
	IncludeTests  resolvedVariable[bool]   `yaml:"includeTests"`
	IncludeProps  resolvedVariable[bool]   `yaml:"includeProps"`
	Styling       resolvedVariable[string] `yaml:"styling"`
	ExportType    resolvedVariable[string] `yaml:"exportType"`
}

// ResolvedVariable contains the final variable configuration
type resolvedVariable[T any] struct {
	Value       T      `yaml:"value"`
	Required    bool   `yaml:"required"`
	Description string `yaml:"description"`
	Enforced    bool   `yaml:"enforced"` // true if team config overrides this
}
