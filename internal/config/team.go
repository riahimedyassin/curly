package config

// Can be overriden by users (flags or tui options)
type Defaults struct {
	ComponentPath string `yaml:"componentPath"`
	IncludeTests  bool   `yaml:"includeTests"`
	IncludeProps  bool   `yaml:"includeProps"`
	Styling       string `yaml:"styling"`
	ExportType    string `yaml:"exportType"`
}

// Mandatory for the project acceptance criatiria. (no user override will work here.)
type Enforcement struct {
	IncludeTests  bool   `yaml:"includeTests"`
	Typescript    bool   `yaml:"typescript"`
	ComponentPath string `yaml:"componentPath"`
	IncludeStyles bool   `yaml:"includeStyles"`
	IncludeProps  bool   `yaml:"includeProps"`
}

type StylingRestriction struct {
	Allowed   []string `yaml:"allowed"`
	Forbidden []string `yaml:"forbidden"`
}

type ExportTypeRestriction struct {
	Allowed   []string `yaml:"allowed"`
	Forbidden []string `yaml:"forbidden"`
}

type Restrictions struct {
	Styling    StylingRestriction    `yaml:"styling"`
	ExportType ExportTypeRestriction `yaml:"exportType"`
}

type ComponentNameValidation struct {
	Pattern   string   `yaml:"pattern"` // Regex
	MaxLength int      `yaml:"maxLength"`
	Forbidden []string `yaml:"forbidden"`
}

type ComponentPathValidation struct {
	MustFolder     string   `yaml:"mustFolder"` // must live within this folder.
	ForbiddenPaths []string `yaml:"forbiddenPaths"`
}

type Validations struct {
	ComponentName ComponentNameValidation `yaml:"componentName"`
	ComponentPath ComponentPathValidation `yaml:"componentPath"`
}

// ! Main Struct
type TeamConfig struct {
	Team         string       `yaml:"team"`
	Organization string       `yaml:"organization"`
	Defaults     Defaults     `yaml:"defaults"`
	Enforcement  Enforcement  `yaml:"enforcement"`
	Restrictions Restrictions `yaml:"restrictions"`
	Validations  Validations  `yaml:"validations"`
}
