package generator

import template2 "text/template"

// Structs

type FileCollection struct {
	files map[string]File
}

func (f *FileCollection) AddFile(relativePath string, skipIfExists bool, template *template2.Template, templateData interface{}) {
	file := File{
		RelativePath: relativePath,
		SkipIfExists: skipIfExists,
		Template:     template,
		TemplateData: templateData,
	}

	f.files[file.RelativePath] = file
}

func (f *FileCollection) GetFiles() map[string]File {
	return f.files
}

// Static functions

func NewFileCollection() *FileCollection {
	return &FileCollection{
		files: make(map[string]File),
	}
}
