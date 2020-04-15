package golang

import "github.com/comfortablynumb/rapidito/model"

// Types

type GolangType string

// Constants

const (
	TypeString  GolangType = "string"
	TypeFloat64 GolangType = "float64"
	TypeBool    GolangType = "bool"
	TypeInt64   GolangType = "int64"
	TypeTime    GolangType = "time.Time"
	TypeCustom  GolangType = "custom"
)

// Static functions

func GetTypeFromModelType(modelType model.ModelType) GolangType {
	switch modelType {
	case model.TypeBool:
		return TypeBool
	case model.TypeString:
		return TypeString
	case model.TypeInt:
		return TypeInt64
	case model.TypeLong:
		return TypeInt64
	case model.TypeFloat:
		return TypeFloat64
	case model.TypeDouble:
		return TypeFloat64
	case model.TypeDateTime:
		return TypeTime
	case model.TypeCustom:
		return TypeCustom
	default:
		return ""
	}
}
