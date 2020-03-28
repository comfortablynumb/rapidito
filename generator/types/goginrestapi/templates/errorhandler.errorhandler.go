package templates

// Constants

const (
	ErrorHandlerErrorHandler = `package errorhandler

import (
	"{{ .Package.Name }}/internal/apperror"
	"{{ .Package.Name }}/internal/context"
	hooks2 "{{ .Package.Name }}/internal/hooks"
	"github.com/rs/zerolog"
)

// Struct

type ErrorHandler struct {
	hooks  *hooks2.Hooks
	logger *zerolog.Logger
}

func (e *ErrorHandler) HandleFatal(err error, message string) {
	e.logger.Error().Msgf("[ERROR] %s - Error: %s", message, err)

	panic(err)
}

func (e *ErrorHandler) HandleFatalIfError(err error, message string) {
	if err == nil {
		return
	}

	e.HandleFatal(err, message)
}

func (e *ErrorHandler) MapAppErrorToHttpError(ctx *context.RequestContext, err *apperror.AppError) *apperror.HttpError {
	switch err.Code {
	case apperror.BindingErrorCode:
		return apperror.NewBindingHttpError(ctx, err.Err, err.Source, err.Data)
	case apperror.ValidationErrorCode:
		return apperror.NewValidationHttpError(ctx, err.Err, err.Source, err.Data)
	case apperror.DbErrorCode:
		return apperror.NewDbHttpError(ctx, err.Err, err.Source, err.Data)
	case apperror.ModelNotFoundErrorCode:
		return apperror.NewNotFoundHttpError(ctx, err.Err, err.Source, err.Data)
	default:
		return apperror.NewInternalServerHttpError(ctx, err.Err, err.Source, err.Data)
	}
}

func (e *ErrorHandler) CreateHttpErrorFromErr(ctx *context.RequestContext, err error, MapAppErrorToHttpError string) *apperror.HttpError {
	var controllerError *apperror.HttpError

	switch err.(type) {
	case *apperror.AppError:
		appError := err.(*apperror.AppError)

		e.logger.Error().Msgf("[Application Error] %s", appError.String())

		controllerError = e.MapAppErrorToHttpError(ctx, err.(*apperror.AppError))
	case *apperror.HttpError:
		controllerError = err.(*apperror.HttpError)
	default:
		controllerError = apperror.NewInternalServerHttpError(ctx, err, "ErrorHandler", nil)
	}

	return controllerError
}

// Static functions

func NewErrorHandler(logger *zerolog.Logger, hooks *hooks2.Hooks) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
		hooks:  hooks,
	}
}
`
)
