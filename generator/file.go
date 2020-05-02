package generator

import "text/template"

// Types

type PreFileWriteFunc func(path string, contents string) (string, error)

// Structs

type File struct {
	RelativePath     string
	SkipIfExists     bool
	Template         *template.Template
	TemplateData     interface{}
	PreFileWriteFunc PreFileWriteFunc
}
