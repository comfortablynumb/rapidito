package golang

import (
	"fmt"
	"github.com/comfortablynumb/rapidito/utils"
	"github.com/stoewer/go-strcase"
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

// Static Functions

func NewGolangResourceFile(resourceCollection GolangResourceCollection) GolangResourceFile {
	return GolangResourceFile{
		ResourceCollection: resourceCollection,
		Imports:            resourceCollection.GetImports(),
	}
}

func NewGolangResourceCollectionFromGolangModel(model GolangModel) GolangResourceCollection {
	resourceCollection := GolangResourceCollection{
		Filename:  model.Filename,
		Model:     model,
		Resources: make([]GolangResource, 0),
	}

	// Find resource

	resourceCollection.AddResource(NewGolangFindResourceFromGolangModel(model))

	// Create resource

	resourceCollection.AddResource(NewGolangCreateResourceFromGolangModel(model))

	// Update resource

	resourceCollection.AddResource(NewGolangUpdateResourceFromGolangModel(model))

	// Delete resource

	resourceCollection.AddResource(NewGolangDeleteResourceFromGolangModel(model))

	// Resource

	resourceCollection.AddResource(NewGolangResourceFromGolangModel(model))

	return resourceCollection
}

func NewGolangFindResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sFindResource", model.StructName)
	resource := NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnSearch {
			continue
		}

		resource.AddField(NewResourceFieldFromModelField(field, true))
	}

	resource.AddEmbeddedStruct("CommonFindResource")

	return resource
}

func NewGolangCreateResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sCreateResource", model.StructName)
	resource := NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnCreate {
			continue
		}

		resource.AddField(NewResourceFieldFromModelField(field, true))
	}

	return resource
}

func NewGolangUpdateResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sUpdateResource", model.StructName)
	resource := NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnUpdate {
			continue
		}

		resource.AddField(NewResourceFieldFromModelField(field, true))
	}

	return resource
}

func NewGolangDeleteResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sDeleteResource", model.StructName)
	resource := NewResource(resourceName, "", false, false)

	for _, fieldName := range model.PrimaryKey {
		field := model.GetField(fieldName)

		resource.Fields = append(resource.Fields, NewResourceFieldFromModelField(*field, true))
	}

	return resource
}

func NewGolangResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sResource", model.StructName)
	resource := NewResource(resourceName, fmt.Sprintf("%sBuilder", resourceName), true, true)

	for _, field := range model.Fields {
		if field.HideOnResponse {
			continue
		}

		resource.AddField(NewResourceFieldFromModelField(field, false))
	}

	return resource
}

func NewResource(
	resourceName string,
	builderName string,
	includeFactoryFunction bool,
	includeFromModelFunction bool,
) GolangResource {
	return GolangResource{
		Name:                     strcase.SnakeCase(resourceName),
		StructName:               resourceName,
		BuilderName:              builderName,
		Fields:                   make([]GolangResourceField, 0),
		EmbeddedStructs:          make([]string, 0),
		IncludeFactoryFunction:   includeFactoryFunction,
		IncludeFromModelFunction: includeFromModelFunction,
	}
}

func NewResourceFieldFromModelField(field GolangModelField, isPointer bool) GolangResourceField {
	return GolangResourceField{
		Name:             strcase.SnakeCase(field.StructFieldName),
		StructFieldName:  field.StructFieldName,
		BuilderFieldName: field.BuilderFieldName,
		ExportedName:     strcase.SnakeCase(field.StructFieldName),
		Type:             field.Type,
		IsPointer:        isPointer,
		Validations:      make([]GolangResourceFieldValidation, 0),
	}
}
