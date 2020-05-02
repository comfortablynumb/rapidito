package helper

import (
	"github.com/comfortablynumb/rapidito/configuration"
	"text/template"

	"github.com/comfortablynumb/rapidito/errorhandler"
)

// Structs

type TemplateHelper struct {
	errorHandler *errorhandler.ErrorHandler
}

func (t *TemplateHelper) ParseTemplate(tpl string) *template.Template {
	parsedTemplate, err := template.New("tpl").Parse(tpl)

	t.errorHandler.HandleIfError(err, "An error occurred while parsing a template.")

	return parsedTemplate
}

type TemplateData struct {
	GlobalConfig    configuration.Config
	GeneratorConfig interface{}
	ExtraData       interface{}
}

// Static functions

func NewTemplateHelper(errorHandler *errorhandler.ErrorHandler) *TemplateHelper {
	return &TemplateHelper{
		errorHandler: errorHandler,
	}
}
