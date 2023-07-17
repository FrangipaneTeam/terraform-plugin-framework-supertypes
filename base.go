package supertypes

import (
	// These are all the packages that are used in the templates.
	_ "github.com/fatih/color"
	_ "github.com/iancoleman/strcase"
)

//go:generate go run template.go
