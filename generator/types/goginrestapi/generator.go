package goginrestapi

import (
	"github.com/comfortablynumb/rapidito/generator"
)

const (
	Type = "go_gin_rest_api"
)

type goGinRestApiGenerator struct {
}

func (r *goGinRestApiGenerator) Generate(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	helper *generator.GeneratorHelper,
) error {
	options := NewRestOptions()

	context.PopulateOptions(options)

	return nil
}

func (r *goGinRestApiGenerator) GetName() string {
	return Type
}

func NewGoGinRestApiGenerator() generator.Generator {
	return &goGinRestApiGenerator{}
}
