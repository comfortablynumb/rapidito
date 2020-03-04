package generator

import (
	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/mitchellh/mapstructure"
)

// Structs

type GeneratorContext struct {
	ErrorHandler *errorhandler.ErrorHandler
	Options      map[string]interface{}
}

func (g *GeneratorContext) PopulateOptions(options interface{}) {
	err := mapstructure.Decode(g.Options, &options)

	if err != nil {
		g.ErrorHandler.Handle(err, "Could not decode generator options.")
	}
}

// Static functions

func NewGeneratorContext(errorHandler *errorhandler.ErrorHandler, options map[string]interface{}) *GeneratorContext {
	return &GeneratorContext{
		ErrorHandler: errorHandler,
		Options:      options,
	}
}
