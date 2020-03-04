package generator

import (
	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/comfortablynumb/rapidito/helper"
	"github.com/comfortablynumb/rapidito/logger"
)

type GeneratorHelper struct {
	ErrorHandler *errorhandler.ErrorHandler
	FileHelper   *helper.FileHelper
	Logger       *logger.Logger
}

func (r *GeneratorHelper) HandleError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.Handle(err, msg, args...)
}

func (r *GeneratorHelper) HandleIfError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.HandleIfError(err, msg, args...)
}

func NewGeneratorHelper(errorHandler *errorhandler.ErrorHandler, fileHelper *helper.FileHelper, logger *logger.Logger) *GeneratorHelper {
	return &GeneratorHelper{
		ErrorHandler: errorHandler,
		FileHelper:   fileHelper,
		Logger:       logger,
	}
}
