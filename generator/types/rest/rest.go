package rest

import (
	"text/template"

	"github.com/comfortablynumb/rapidito/generator"
)

const (
	RestApiType = "rest_api"
)

type restApiGenerator struct {
}

func (r *restApiGenerator) Generate(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	helper *generator.GeneratorHelper,
) error {
	options := NewRestOptions()

	context.PopulateOptions(options)

	tpl, err := template.New("rest").Parse("{{ .message }}")

	helper.HandleIfError(err, "An error occurred while parsing a template.")

	tplData := make(map[string]interface{})

	tplData["message"] = "Hello World!"

	fileCollection.AddFile("main.go", tpl, tplData)

	return nil
}

func (r *restApiGenerator) GetName() string {
	return RestApiType
}

func NewRestApiGenerator() generator.Generator {
	return &restApiGenerator{}
}
