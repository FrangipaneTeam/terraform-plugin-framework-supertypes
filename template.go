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

	"github.com/iancoleman/strcase"
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

var (
	// TODO "object", "single_nested"
	typesSimple = []string{"string", "bool", "float64", "int64", "number"}
	typesArray  = []string{"list", "set", "map"}
	typesNested = []string{"list_nested", "set_nested", "map_nested"}
)

func main() {
	fmt.Println("generating types files...")

	// * Simple types
	for _, t := range typesSimple {
		generateTemplate(t, templateSimpleType, templateSimpleValue)
	}

	// // * Array types
	// for _, t := range typesArray {
	// 	generateTemplate(t, templateArrayType, templateArrayValue)
	// }

	// // * Nested types
	// for _, t := range typesNested {
	// 	generateTemplate(t, templateNestedType, templateNestedValue)
	// }
}

func generateTemplate(typeName, templateType, templateValue string) {
	infos := templateInfos{
		TypeName: strcase.ToCamel(typeName),
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
	}

	if err := os.WriteFile(typeName+"_value.go", formattedTmplValue, 0644); err != nil {
		fmt.Printf("write to file error : %v\n", err)
	}
}
