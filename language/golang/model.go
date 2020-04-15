package golang

import (
	"fmt"

	"github.com/comfortablynumb/rapidito/model"
	"github.com/stoewer/go-strcase"
)

// Structs

type GolangModel struct {
	Name        string
	Filename    string
	StructName  string
	BuilderName string
	Fields      map[string]GolangModelField
}

type GolangModelField struct {
	Name             string
	StructFieldName  string
	BuilderFieldName string
	Type             GolangType
	CustomType       string
}

// Static functions

func NewGolangModelFromModel(m model.Model) GolangModel {
	golangModel := GolangModel{
		Name:        m.Name,
		StructName:  strcase.UpperCamelCase(m.Name),
		BuilderName: fmt.Sprintf("%sBuilder", strcase.UpperCamelCase(m.Name)),
		Filename:    strcase.SnakeCase(m.Name),
		Fields:      make(map[string]GolangModelField),
	}

	for _, field := range m.GetFields() {
		golangModel.Fields[field.Name] = GolangModelField{
			Name:             field.Name,
			StructFieldName:  strcase.UpperCamelCase(field.Name),
			BuilderFieldName: strcase.LowerCamelCase(field.Name),
			Type:             GetTypeFromModelType(field.Type),
		}
	}

	return golangModel
}
