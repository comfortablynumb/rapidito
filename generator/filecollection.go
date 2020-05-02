package generator

import template2 "text/template"

// Structs

type FileCollection struct {
	files map[string]File
}

func (f *FileCollection) AddFromCollection(fileCollection *FileCollection) {
	for _, file := range fileCollection.GetFiles() {
		f.AddFile(file.RelativePath, file.SkipIfExists, file.Template, file.TemplateData, file.PreFileWriteFunc)
	}
}

func (f *FileCollection) AddFile(
	relativePath string,
	skipIfExists bool,
	template *template2.Template,
	templateData interface{},
	preFileWriteFunc PreFileWriteFunc,
) {
	file := File{
		RelativePath:     relativePath,
		SkipIfExists:     skipIfExists,
		Template:         template,
		TemplateData:     templateData,
		PreFileWriteFunc: preFileWriteFunc,
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
