package commonfiles

import (
	"github.com/comfortablynumb/rapidito/generator"
	"github.com/comfortablynumb/rapidito/generator/types/commonfiles/templates"
)

const (
	Type = "common_files"
)

type commonFilesGenerator struct {
}

func (r *commonFilesGenerator) Generate(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	helper *generator.GeneratorHelper,
) error {
	options := NewCommonFilesOptions()

	context.PopulateOptions(options)

	fileCollection.AddFile("README.md", true, helper.ParseTemplate(templates.Readme), options)

	return nil
}

func (r *commonFilesGenerator) GetName() string {
	return Type
}

func NewCommonFilesGenerator() generator.Generator {
	return &commonFilesGenerator{}
}
