package errorhandler

import (
	"github.com/comfortablynumb/rapidito/logger"
)

type ErrorHandler struct {
	Logger *logger.Logger
}

func (r *ErrorHandler) HandleIfError(err error, msg string, args ...interface{}) {
	if err != nil {
		r.Handle(err, msg, args...)
	}
}

func (r *ErrorHandler) Handle(err error, msg string, args ...interface{}) {
	r.Logger.Panic(err, msg, args...)
}

func NewErrorHandler(logger *logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		Logger: logger,
	}
}
