//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/iancoleman/strcase"
)

var (
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	blue   = color.New(color.FgBlue)
	yellow = color.New(color.FgYellow)
	grey   = color.New(color.FgHiBlack)
)

type templateInfos struct {
	TypeName string
}

//go:embed templates/simple_type.go.tmpl
var templateSimpleType string

//go:embed templates/simple_value.go.tmpl
var templateSimpleValue string

//go:embed templates/array_type.go.tmpl
var templateArrayType string

//go:embed templates/array_value.go.tmpl
var templateArrayValue string

//go:embed templates/nested_type.go.tmpl
var templateNestedType string

//go:embed templates/nested_value.go.tmpl
var templateNestedValue string

//go:embed templates/object_type.go.tmpl
var templateObjectType string

//go:embed templates/object_value.go.tmpl
var templateObjectValue string

//go:embed templates/single_nested_type.go.tmpl
var templateSingleNestedType string

//go:embed templates/single_nested_value.go.tmpl
var templateSingleNestedValue string

var (
	typesSimple      = []string{"string", "bool", "float64", "int64", "number", "object"}
	typesArray       = []string{"list", "set", "map"}
	typesNested      = []string{"list_nested", "set_nested", "map_nested"}
	typeObject       = []string{"object"}
	typeSingleNested = []string{"single_nested"}
)

func main() {
	yellow.Println("Starting generation of supertypes")
	yellow.Println("=================================")

	blue.Println("-> Generating simple types")
	// * Simple types
	for _, t := range typesSimple {
		generateTemplate(t, templateSimpleType, templateSimpleValue)
	}

	blue.Println("-> Generating array types")
	// * Array types
	for _, t := range typesArray {
		generateTemplate(t, templateArrayType, templateArrayValue)
	}

	blue.Println("-> Generating nested types")
	// * Nested types
	for _, t := range typesNested {
		generateTemplate(t, templateNestedType, templateNestedValue)
	}

	blue.Println("-> Generating object type")
	// * Object types
	for _, t := range typeObject {
		generateTemplate(t, templateObjectType, templateObjectValue)
	}

	blue.Println("-> Generating single nested type")
	// * Single nested types
	for _, t := range typeSingleNested {
		generateTemplate(t, templateSingleNestedType, templateSingleNestedValue)
	}
}

func generateTemplate(typeName, templateType, templateValue string) {
	grey.Printf("	* Generating %s ...", typeName)
	infos := templateInfos{
		TypeName: strcase.ToCamel(typeName),
	}

	if strings.HasSuffix(typeName, "_nested") {
		infos.TypeName = strcase.ToCamel(strings.TrimSuffix(typeName, "_nested"))
	}

	tmplType, err := template.New("template").Parse(templateType)
	if err != nil {
		fmt.Printf("error from template parse : %v\n", err)
		os.Exit(1)
	}

	tmplValue, err := template.New("template").Parse(templateValue)
	if err != nil {
		fmt.Printf("error from template parse : %v\n", err)
		os.Exit(1)
	}

	var (
		tplType  bytes.Buffer
		tplValue bytes.Buffer
	)

	if err := tmplType.Execute(&tplType, infos); err != nil {
		fmt.Printf("error from template execute : %v\n", err)
		os.Exit(1)
	}

	if err := tmplValue.Execute(&tplValue, infos); err != nil {
		fmt.Printf("error from template execute : %v\n", err)
		os.Exit(1)
	}

	// format the code
	formattedTmplType, errFormat := format.Source(tplType.Bytes())
	if errFormat != nil {
		fmt.Printf("error from format : %v\n", errFormat)
		os.Exit(1)
	}

	formattedTmplValue, errFormat := format.Source(tplValue.Bytes())
	if errFormat != nil {
		fmt.Printf("error from format : %v\n", errFormat)
		os.Exit(1)
	}

	if err := os.WriteFile(typeName+"_type.go", formattedTmplType, 0644); err != nil {
		fmt.Printf("write to file error : %v\n", err)
		red.Println("✘")
	}

	if err := os.WriteFile(typeName+"_value.go", formattedTmplValue, 0644); err != nil {
		fmt.Printf("write to file error : %v\n", err)
		red.Println("✘")
	}

	green.Println("✔")
}
