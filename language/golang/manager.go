package golang

import (
	"fmt"

	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/comfortablynumb/rapidito/logger"
	"github.com/comfortablynumb/rapidito/model"
	"github.com/stoewer/go-strcase"
)

// Structs

type GolangManager struct {
	logger       *logger.Logger
	errorHandler *errorhandler.ErrorHandler
	golangTypes  *GolangTypes
}

func (g *GolangManager) NewGolangModelFromModel(m model.Model) GolangModel {
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
			Type:             g.GetTypeFromModelType(field.Type),
			HideOnSearch:     field.HideOnSearch,
			HideOnCreate:     field.HideOnCreate,
			HideOnUpdate:     field.HideOnUpdate,
			HideOnResponse:   field.HideOnResponse,
		}
	}

	return golangModel
}

func (g *GolangManager) GetTypeFromModelType(modelType model.ModelType) GolangType {
	golangType := g.golangTypes.GetByModelType(modelType)

	if golangType == nil {
		g.errorHandler.Handle(fmt.Errorf("[Golang] Model type '%s' does not have a golang type configured.", modelType), "Missing Golang type for the given model type.")
	}

	return *golangType
}

func NewGolangResourceFile(resourceCollection GolangResourceCollection) GolangResourceFile {
	return GolangResourceFile{
		ResourceCollection: resourceCollection,
		Imports:            resourceCollection.GetImports(),
	}
}

func (g *GolangManager) NewGolangResourceCollectionFromGolangModel(model GolangModel) GolangResourceCollection {
	resourceCollection := GolangResourceCollection{
		Filename:  model.Filename,
		Model:     model,
		Resources: make([]GolangResource, 0),
	}

	// Find resource

	resourceCollection.AddResource(g.NewGolangFindResourceFromGolangModel(model))

	// Create resource

	resourceCollection.AddResource(g.NewGolangCreateResourceFromGolangModel(model))

	// Update resource

	resourceCollection.AddResource(g.NewGolangUpdateResourceFromGolangModel(model))

	// Delete resource

	resourceCollection.AddResource(g.NewGolangDeleteResourceFromGolangModel(model))

	// Resource

	resourceCollection.AddResource(g.NewGolangResourceFromGolangModel(model))

	return resourceCollection
}

func (g *GolangManager) NewGolangFindResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sFindResource", model.StructName)
	resource := g.NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnSearch {
			continue
		}

		resource.AddField(g.NewResourceFieldFromModelField(field, true))
	}

	resource.AddEmbeddedStruct("CommonFindResource")

	return resource
}

func (g *GolangManager) NewGolangCreateResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sCreateResource", model.StructName)
	resource := g.NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnCreate {
			continue
		}

		resource.AddField(g.NewResourceFieldFromModelField(field, true))
	}

	return resource
}

func (g *GolangManager) NewGolangUpdateResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sUpdateResource", model.StructName)
	resource := g.NewResource(resourceName, "", false, false)

	for _, field := range model.Fields {
		if field.HideOnUpdate {
			continue
		}

		resource.AddField(g.NewResourceFieldFromModelField(field, true))
	}

	return resource
}

func (g *GolangManager) NewGolangDeleteResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sDeleteResource", model.StructName)
	resource := g.NewResource(resourceName, "", false, false)

	for _, fieldName := range model.PrimaryKey {
		field := model.GetField(fieldName)

		resource.Fields = append(resource.Fields, g.NewResourceFieldFromModelField(*field, true))
	}

	return resource
}

func (g *GolangManager) NewGolangResourceFromGolangModel(model GolangModel) GolangResource {
	resourceName := fmt.Sprintf("%sResource", model.StructName)
	resource := g.NewResource(resourceName, fmt.Sprintf("%sBuilder", resourceName), true, true)

	for _, field := range model.Fields {
		if field.HideOnResponse {
			continue
		}

		resource.AddField(g.NewResourceFieldFromModelField(field, false))
	}

	return resource
}

func (g *GolangManager) NewResource(
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

func (g *GolangManager) NewResourceFieldFromModelField(field GolangModelField, isPointer bool) GolangResourceField {
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

func (g *GolangManager) initialize() {
	// Initialize Core Golang Types

	g.golangTypes.AddType(TypeString, TypeString, []model.ModelType{model.TypeString})
	g.golangTypes.AddType(TypeBool, TypeBool, []model.ModelType{model.TypeBool})
	g.golangTypes.AddType(TypeInt64, TypeInt64, []model.ModelType{model.TypeInt, model.TypeLong})
	g.golangTypes.AddType(TypeFloat64, TypeFloat64, []model.ModelType{model.TypeFloat, model.TypeDouble})
	g.golangTypes.AddType(TypeTime, TypeTime, []model.ModelType{model.TypeDateTime}, "time")
}

// Static functions

func NewGolangManager(logger *logger.Logger, errorHandler *errorhandler.ErrorHandler) *GolangManager {
	golangManager := &GolangManager{
		logger:       logger,
		errorHandler: errorHandler,
		golangTypes: &GolangTypes{
			types: make(map[string]GolangType),
		},
	}

	golangManager.initialize()

	return golangManager
}
