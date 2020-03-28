package templates

// Constants

const (
	AppErrorCommon = `package apperror

import (
	"{{ .Package.Name }}/internal/context"
	"{{ .Package.Name }}/internal/validation"
	"gopkg.in/go-playground/validator.v9"
)

// Static functions

func AddValidationErrorsToMap(ctx *context.RequestContext, err error, data map[string]interface{}) {
	fieldErrors, ok := err.(validator.ValidationErrors)

	if ok {
		errors := make([]*validation.ValidationError, 0)
		trans := ctx.GetTranslator()

		for _, fieldError := range fieldErrors {
			errors = append(errors, validation.NewValidationError(fieldError.Namespace(), fieldError.Tag(), fieldError.Translate(*trans)))
		}

		data["errors"] = errors
	}
}
`
)
