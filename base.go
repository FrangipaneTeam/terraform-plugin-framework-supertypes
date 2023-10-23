package supertypes

import (
	// These are all the packages that are used in the templates.
	_ "github.com/fatih/color"
	_ "github.com/iancoleman/strcase"
)

const (
	errorTest = "An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n"
)

//go:generate go run template.go
