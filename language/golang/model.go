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
	PrimaryKey  []string
}

type GolangModelField struct {
	Name             string
	StructFieldName  string
	BuilderFieldName string
	Type             GolangType
	CustomType       string
	HideOnSearch     bool
	HideOnCreate     bool
	HideOnUpdate     bool
	HideOnResponse   bool
}

func (m *GolangModel) GetField(name string) *GolangModelField {
	field, found := m.Fields[name]

	if !found {
		return nil
	}

	return &field
}

// Static functions

func NewGolangModelFromModel(m model.Model) GolangModel {
	golangModel := GolangModel{
		Name:        m.Name,
		StructName:  strcase.UpperCamelCase(m.Name),
		BuilderName: fmt.Sprintf("%sBuilder", strcase.UpperCamelCase(m.Name)),
		Filename:    strcase.SnakeCase(m.Name),
		Fields:      make(map[string]GolangModelField),
		PrimaryKey:  m.PrimaryKey,
	}

	for _, field := range m.GetFields() {
		golangModel.Fields[field.Name] = GolangModelField{
			Name:             field.Name,
			StructFieldName:  strcase.UpperCamelCase(field.Name),
			BuilderFieldName: strcase.LowerCamelCase(field.Name),
			Type:             GetTypeFromModelType(field.Type),
			HideOnSearch:     field.HideOnSearch,
			HideOnCreate:     field.HideOnCreate,
			HideOnUpdate:     field.HideOnUpdate,
			HideOnResponse:   field.HideOnResponse,
		}
	}

	return golangModel
}
