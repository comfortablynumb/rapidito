package generator

import "text/template"

// Structs

type File struct {
	RelativePath string
	SkipIfExists bool
	Template     *template.Template
	TemplateData interface{}
}
