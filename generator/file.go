package generator

import "text/template"

// Structs

type File struct {
	RelativePath string
	Template     *template.Template
	TemplateData map[string]interface{}
}
