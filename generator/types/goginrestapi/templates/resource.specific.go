package templates

// Constants

const (
	ResourceSpecific = `package resource
{{ $model := .ExtraData.ResourceCollection.Model }}
import (
	{{ range .ExtraData.Imports }}"{{ . }}"
	{{ end }}

	"{{ .GeneratorConfig.Package.Name }}/internal/model"
)

// Structs
{{ range .ExtraData.ResourceCollection.Resources }}
{{ $resource := . }}
// {{ .StructName }}

type {{ .StructName }} struct {
{{ range .EmbeddedStructs }}    {{ . }}	
{{ end }}
{{ range .Fields }}    {{ .StructFieldName }} {{ if .IsPointer }}*{{ end }}{{ .Type }} ` + "`json:\"{{ .ExportedName }}\"`" + `
{{ end }}
}
{{ if (ne .BuilderName "" ) }}
// {{ .BuilderName }}

type {{ .BuilderName }} struct {
{{ range .Fields }}    {{ .BuilderFieldName }} {{ if .IsPointer }}*{{ end }}{{ .Type }}
{{ end }}
}

{{ range .Fields }}func (b *{{ $resource.BuilderName }}) With{{ .StructFieldName }}({{ .BuilderFieldName }} {{ .Type }}) *{{ $resource.BuilderName }} {
	b.{{ .BuilderFieldName }} = {{ .BuilderFieldName }}

	return b
}
{{ end }}
func (b *{{ $resource.BuilderName }}) Build() *{{ $resource.StructName }} {
	return New{{ $resource.StructName }}(
		{{ range .Fields }}b.{{ .BuilderFieldName }},
		{{ end }}
	)
}

// Static functions

func New{{ $resource.BuilderName }}() *{{ $resource.BuilderName }} {
	return &{{ $resource.BuilderName }}{}
}

{{ end }}
{{ if $resource.IncludeFactoryFunction }}
func New{{ $resource.StructName }}(
	{{ range .Fields }}{{ .BuilderFieldName }} {{ .Type }},
	{{ end }}
) *{{ $resource.StructName }} {
	return &{{ $resource.StructName }}{
		{{ range .Fields }}{{ .StructFieldName }}: {{ .BuilderFieldName }},
		{{ end }}
	}
}
{{ end }}
{{ if $resource.IncludeFromModelFunction }}
func New{{ $resource.StructName }}From{{ $model.StructName }}(modelInstance model.{{ $model.StructName }}) *{{ $resource.StructName }} {
	return New{{ $resource.BuilderName }}().
		{{ range $resource.Fields }}With{{ .StructFieldName }}(modelInstance.{{ .StructFieldName }}).
		{{ end }}Build()
}
{{ end }}
{{ end }}
`
)
