package templates

// Constants

const (
	ModuleCommon = `package module

import (
	"{{ .Package.Name }}/internal/componentregistry"
	"{{ .Package.Name }}/internal/config"
	"{{ .Package.Name }}/internal/errorhandler"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Interfaces

type Module interface {
	GetName() string
	SetUpComponents(appConfig config.AppConfig, errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry)
	SetUpRouter(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, router *gin.Engine)
	SetUpValidator(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, validator *validator.Validate)
}
`
)
