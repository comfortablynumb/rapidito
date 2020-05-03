package golang

import (
	"github.com/comfortablynumb/rapidito/model"
)

// Constants

const (
	TypeString  = "string"
	TypeFloat64 = "float64"
	TypeBool    = "bool"
	TypeInt64   = "int64"
	TypeTime    = "time.Time"
	TypeCustom  = "custom"
)

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
