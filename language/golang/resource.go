package golang

import (
	"github.com/comfortablynumb/rapidito/utils"
)

// Types

type GolangValidationType string

// Constants

const (
	ValidationTypeRequired GolangValidationType = "required"
	ValidationTypeMin      GolangValidationType = "min"
	ValidationTypeMax      GolangValidationType = "max"
	ValidationTypeCustom   GolangValidationType = "custom"
)

// Structs

type GolangResourceFile struct {
	ResourceCollection GolangResourceCollection
	Imports            []string
}

type GolangResourceCollection struct {
	Filename  string
	Model     GolangModel
	Resources []GolangResource
}

func (r *GolangResourceCollection) AddResource(resource GolangResource) *GolangResourceCollection {
	r.Resources = append(r.Resources, resource)

	return r
}

func (r *GolangResourceCollection) GetImports() []string {
	importPackages := make([][]string, 0)

	for _, r := range r.Resources {
		for _, f := range r.Fields {
			importPackages = append(importPackages, f.Type.GetImports())
		}
	}

	return utils.SliceUnionStrings(importPackages...)
}

type GolangResource struct {
	Name                     string
	StructName               string
	BuilderName              string
	Fields                   []GolangResourceField
	EmbeddedStructs          []string
	IncludeFactoryFunction   bool
	IncludeFromModelFunction bool
}

func (r *GolangResource) AddField(field GolangResourceField) *GolangResource {
	r.Fields = append(r.Fields, field)

	return r
}

func (r *GolangResource) AddEmbeddedStruct(structName string) *GolangResource {
	r.EmbeddedStructs = append(r.EmbeddedStructs, structName)

	return r
}

type GolangResourceField struct {
	Name             string
	StructFieldName  string
	BuilderFieldName string
	ExportedName     string
	Type             GolangType
	IsPointer        bool
	Validations      []GolangResourceFieldValidation
}

func (r *GolangResourceField) AddValidation(validation GolangResourceFieldValidation) *GolangResourceField {
	r.Validations = append(r.Validations, validation)

	return r
}

type GolangResourceFieldValidation struct {
	Type GolangValidationType
}
