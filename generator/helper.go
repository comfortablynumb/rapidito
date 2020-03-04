package generator

import (
	"text/template"

	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/comfortablynumb/rapidito/helper"
	"github.com/comfortablynumb/rapidito/logger"
)

type GeneratorHelper struct {
	ErrorHandler   *errorhandler.ErrorHandler
	FileHelper     *helper.FileHelper
	TemplateHelper *helper.TemplateHelper
	Logger         *logger.Logger
}

func (r *GeneratorHelper) ParseTemplate(tpl string) *template.Template {
	return r.TemplateHelper.ParseTemplate(tpl)
}

func (r *GeneratorHelper) HandleError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.Handle(err, msg, args...)
}

func (r *GeneratorHelper) HandleIfError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.HandleIfError(err, msg, args...)
}

func NewGeneratorHelper(
	errorHandler *errorhandler.ErrorHandler,
	fileHelper *helper.FileHelper,
	templateHelper *helper.TemplateHelper,
	logger *logger.Logger,
) *GeneratorHelper {
	return &GeneratorHelper{
		ErrorHandler:   errorHandler,
		FileHelper:     fileHelper,
		TemplateHelper: templateHelper,
		Logger:         logger,
	}
}
