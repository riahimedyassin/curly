package utils

import (
	"fmt"
	"strings"
	"unicode"
)

// Component -> Component
//
// COMPONENT -> Component
//
// component -> Component.
func ToTitleCase(value string) string {
	first := unicode.ToUpper(rune(value[0]))
	res := fmt.Sprintf("%s %s", string(first), strings.ToLower(value[1:]))
	return res
}
