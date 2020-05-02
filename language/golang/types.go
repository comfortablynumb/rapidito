package golang

import (
	"github.com/comfortablynumb/rapidito/model"
	"log"
)

func init() {
	// Initialize Core Golang Types

	golangTypes.AddType(TypeString, TypeString, []model.ModelType{model.TypeString})
	golangTypes.AddType(TypeBool, TypeBool, []model.ModelType{model.TypeBool})
	golangTypes.AddType(TypeInt64, TypeInt64, []model.ModelType{model.TypeInt, model.TypeLong})
	golangTypes.AddType(TypeFloat64, TypeFloat64, []model.ModelType{model.TypeFloat, model.TypeDouble})
	golangTypes.AddType(TypeTime, TypeTime, []model.ModelType{model.TypeDateTime}, "time")
}

// Constants

const (
	TypeString  = "string"
	TypeFloat64 = "float64"
	TypeBool    = "bool"
	TypeInt64   = "int64"
	TypeTime    = "time.Time"
	TypeCustom  = "custom"
)

// Vars

var golangTypes = &GolangTypes{
	types: make(map[string]GolangType),
}

// Structs

type GolangTypes struct {
	types map[string]GolangType
}

func (g *GolangTypes) GetByModelType(modelType model.ModelType) *GolangType {
	for _, t := range g.types {
		for _, mt := range t.GetModelTypes() {
			if mt == modelType {
				tCopy := t

				return &tCopy
			}
		}
	}

	return nil
}

func (g *GolangTypes) AddType(id string, typeName string, modelTypes []model.ModelType, importPackage ...string) {
	g.types[id] = GolangType{
		id:         id,
		typeName:   typeName,
		importList: importPackage,
		modelTypes: modelTypes,
	}
}

type GolangType struct {
	id         string
	typeName   string
	importList []string
	modelTypes []model.ModelType
}

func (g GolangType) GetTypeName() string {
	return g.typeName
}

func (g GolangType) GetImports() []string {
	return g.importList
}

func (g GolangType) GetModelTypes() []model.ModelType {
	return g.modelTypes
}

func (g GolangType) String() string {
	return g.GetTypeName()
}

// Static functions

func GetTypeFromModelType(modelType model.ModelType) GolangType {
	golangType := golangTypes.GetByModelType(modelType)

	if golangType == nil {
		log.Panicf("[Golang] Model type '%s' does not have a golang type configured.", modelType)
	}

	return *golangType
}
