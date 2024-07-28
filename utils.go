package utils

import (
	// https://pkg.go.dev/golang.org/x/text/cases#Title
	// https://pkg.go.dev/golang.org/x/text/cases#example-package
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToTitle(s string) string {
	toTitle := cases.Title(language.English)
	titleCase := toTitle.String(s)

	return titleCase
}
