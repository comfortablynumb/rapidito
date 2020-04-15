package templates

// Constants

const (
	Model = `package model
{{ $model := . }}

// Structs

type {{ $model.StructName }} struct {
{{ range $model.Fields }}	{{ .StructFieldName }} {{ .Type }}
{{ end }}
}

type {{ $model.BuilderName }} struct {
{{ range $model.Fields }}	{{ .BuilderFieldName }} {{ .Type }}
{{ end }}
}

{{ range $model.Fields }}
func (b *{{ $model.BuilderName }}) With{{ .StructFieldName }}({{ .StructFieldName }} {{ .Type }}) *{{ $model.BuilderName }} {
	b.{{ .BuilderFieldName }} = {{ .StructFieldName }}

	return b
}
{{ end }}

func (b *{{ $model.BuilderName }}) Build() *{{ $model.StructName }} {
	return &{{ $model.StructName }}{
{{ range $model.Fields }}		{{ .StructFieldName }}:        b.{{ .BuilderFieldName }},
{{ end }}
	}
}

// Static functions

func New{{ $model.BuilderName }}() *{{ $model.BuilderName }} {
	return &{{ $model.BuilderName }}{}
}

`
)
