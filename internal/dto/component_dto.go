package dto

type CreateComponentTest struct {
	Path    string
	Suffix  string
	Content string
}

type CreateComponentStyle struct {
	Path      string
	Extension string
	Content   string
}

type CreateComponent struct {
	Path    string
	Content string
	Test    *CreateComponentTest
	Style   *CreateComponentStyle
}

type ComponentArgs struct {
	Name string `mapstructure:"name"`
	// in case the team config is not enforcing the rule, this will be significant.
	IncludeTests bool `mapstructure:"includeTests"`
	// in case the team config is not enforcing the rule, this will be significant.
	IncludeStyles bool `mapstructure:"includeStyles"`
}
