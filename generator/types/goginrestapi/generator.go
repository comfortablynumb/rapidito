package goginrestapi

import (
	"github.com/comfortablynumb/rapidito/generator"
	"github.com/comfortablynumb/rapidito/generator/types/goginrestapi/templates"
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
	options := NewGoGinRestApiOptions()

	context.PopulateOptions(options)

	r.generateCommonFiles(fileCollection, options, helper)

	return nil
}

func (r *goGinRestApiGenerator) PostGeneration(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	helper *generator.GeneratorHelper,
) error {

	return nil
}

func (r *goGinRestApiGenerator) generateCommonFiles(
	fileCollection *generator.FileCollection,
	options *GoGinRestApiOptions,
	helper *generator.GeneratorHelper,
) {
	// :: Root files

	fileCollection.AddFile("main.go", false, helper.ParseTemplate(templates.MainGo), options)
	fileCollection.AddFile("go.mod", false, helper.ParseTemplate(templates.GoMod), options)
	fileCollection.AddFile("go.sum", false, helper.ParseTemplate(templates.GoSum), options)
	fileCollection.AddFile(".gitignore", false, helper.ParseTemplate(templates.GitIgnore), options)

	// :: Generic services / utilities

	// Package app

	fileCollection.AddFile("internal/app/app.go", false, helper.ParseTemplate(templates.AppApp), options)

	// Package apperror

	fileCollection.AddFile("internal/apperror/apperror.go", false, helper.ParseTemplate(templates.AppErrorAppError), options)
	fileCollection.AddFile("internal/apperror/common.go", false, helper.ParseTemplate(templates.AppErrorCommon), options)
	fileCollection.AddFile("internal/apperror/constants.go", false, helper.ParseTemplate(templates.AppErrorConstants), options)

	// Package componentregistry

	fileCollection.AddFile("internal/componentregistry/componentregistry.go", false, helper.ParseTemplate(templates.ComponentRegistryComponentRegistry), options)

	// Package config

	fileCollection.AddFile("internal/config/config.go", false, helper.ParseTemplate(templates.ConfigConfig), options)

	// Package context

	fileCollection.AddFile("internal/context/requestcontext.go", false, helper.ParseTemplate(templates.ContextRequestContext), options)
	fileCollection.AddFile("internal/context/requestcontextfactory.go", false, helper.ParseTemplate(templates.ContextRequestContextFactory), options)

	// Package errorhandler

	fileCollection.AddFile("internal/errorhandler/errorhandler.go", false, helper.ParseTemplate(templates.ErrorHandlerErrorHandler), options)

	// Package hooks

	fileCollection.AddFile("internal/hooks/hooks.go", false, helper.ParseTemplate(templates.HooksHooks), options)
	fileCollection.AddFile("internal/hooks/hooks_custom.go", false, helper.ParseTemplate(templates.HooksHooksCustom), options)

	// Package middleware

	fileCollection.AddFile("internal/middleware/errorhandler.go", false, helper.ParseTemplate(templates.MiddlewareErrorHandler), options)

	// Package mock

	fileCollection.AddFile("internal/mock/app.go", false, helper.ParseTemplate(templates.MockApp), options)

	// Package module

	fileCollection.AddFile("internal/module/common.go", false, helper.ParseTemplate(templates.ModuleCommon), options)

	// Package resource

	fileCollection.AddFile("internal/resource/common.go", false, helper.ParseTemplate(templates.ResourceCommon), options)

	// Package service

	fileCollection.AddFile("internal/service/time.go", false, helper.ParseTemplate(templates.ServiceTime), options)

	// Package validation

	fileCollection.AddFile("internal/validation/validationerror.go", false, helper.ParseTemplate(templates.ValidationValidationError), options)
}

func (r *goGinRestApiGenerator) GetName() string {
	return Type
}

func NewGoGinRestApiGenerator() generator.Generator {
	return &goGinRestApiGenerator{}
}
