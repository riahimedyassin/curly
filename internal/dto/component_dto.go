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
