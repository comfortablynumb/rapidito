package generator

import template2 "text/template"

// Structs

type FileCollection struct {
	files map[string]File
}

func (f *FileCollection) AddFile(relativePath string, template *template2.Template, templateData map[string]interface{}) {
	file := File{
		RelativePath: relativePath,
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
