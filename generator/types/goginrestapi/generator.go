package goginrestapi

import (
	"fmt"
	helper2 "github.com/comfortablynumb/rapidito/helper"
	"github.com/comfortablynumb/rapidito/language/sql"

	"github.com/comfortablynumb/rapidito/generator"
	"github.com/comfortablynumb/rapidito/generator/types/goginrestapi/templates"
	"github.com/comfortablynumb/rapidito/language/golang"
	"go/format"
)

const (
	Type = "go_gin_rest_api"
)

type goGinRestApiGenerator struct {
	golangManager *golang.GolangManager
	sqlManager    *sql.SqlManager
}

func (r *goGinRestApiGenerator) Generate(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	helper *generator.GeneratorHelper,
) error {
	options := NewGoGinRestApiOptions()

	context.PopulateOptions(options)

	r.generateCommonFiles(fileCollection, options, helper)
	r.generateApiModules(fileCollection, context, options, helper)

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

	fileCollection.AddFile(
		"main.go",
		false,
		helper.ParseTemplate(templates.MainGo),
		options,
		formatGoCode,
	)
	fileCollection.AddFile(
		"go.mod",
		false,
		helper.ParseTemplate(templates.GoMod),
		options,
		nil,
	)
	fileCollection.AddFile(
		"go.sum",
		false,
		helper.ParseTemplate(templates.GoSum),
		options,
		nil,
	)
	fileCollection.AddFile(
		".gitignore",
		false,
		helper.ParseTemplate(templates.GitIgnore),
		options,
		nil,
	)

	// :: Generic services / utilities

	// Package app

	fileCollection.AddFile(
		"internal/app/app.go",
		false,
		helper.ParseTemplate(templates.AppApp),
		options,
		formatGoCode,
	)

	// Package apperror

	fileCollection.AddFile(
		"internal/apperror/apperror.go",
		false,
		helper.ParseTemplate(templates.AppErrorAppError),
		options,
		formatGoCode,
	)
	fileCollection.AddFile(
		"internal/apperror/common.go",
		false,
		helper.ParseTemplate(templates.AppErrorCommon),
		options,
		formatGoCode,
	)
	fileCollection.AddFile(
		"internal/apperror/constants.go",
		false,
		helper.ParseTemplate(templates.AppErrorConstants),
		options,
		formatGoCode,
	)

	// Package componentregistry

	fileCollection.AddFile(
		"internal/componentregistry/componentregistry.go",
		false,
		helper.ParseTemplate(templates.ComponentRegistryComponentRegistry),
		options,
		formatGoCode,
	)

	// Package config

	fileCollection.AddFile(
		"internal/config/config.go",
		false,
		helper.ParseTemplate(templates.ConfigConfig),
		options,
		formatGoCode,
	)

	// Package context

	fileCollection.AddFile(
		"internal/context/requestcontext.go",
		false,
		helper.ParseTemplate(templates.ContextRequestContext),
		options,
		formatGoCode,
	)
	fileCollection.AddFile(
		"internal/context/requestcontextfactory.go",
		false,
		helper.ParseTemplate(templates.ContextRequestContextFactory),
		options,
		formatGoCode,
	)

	// Package errorhandler

	fileCollection.AddFile(
		"internal/errorhandler/errorhandler.go",
		false,
		helper.ParseTemplate(templates.ErrorHandlerErrorHandler),
		options,
		formatGoCode,
	)

	// Package hooks

	fileCollection.AddFile(
		"internal/hooks/hooks.go",
		false,
		helper.ParseTemplate(templates.HooksHooks),
		options,
		formatGoCode,
	)
	fileCollection.AddFile(
		"internal/hooks/hooks_custom.go",
		false,
		helper.ParseTemplate(templates.HooksHooksCustom),
		options,
		formatGoCode,
	)

	// Package middleware

	fileCollection.AddFile(
		"internal/middleware/errorhandler.go",
		false,
		helper.ParseTemplate(templates.MiddlewareErrorHandler),
		options,
		formatGoCode,
	)

	// Package mock

	fileCollection.AddFile(
		"internal/mock/app.go",
		false,
		helper.ParseTemplate(templates.MockApp),
		options,
		formatGoCode,
	)

	// Package module

	fileCollection.AddFile(
		"internal/module/common.go",
		false,
		helper.ParseTemplate(templates.ModuleCommon),
		options,
		formatGoCode,
	)

	// Package resource

	fileCollection.AddFile(
		"internal/resource/common.go",
		false,
		helper.ParseTemplate(templates.ResourceCommon),
		options,
		formatGoCode,
	)

	// Package service

	fileCollection.AddFile(
		"internal/service/time.go",
		false,
		helper.ParseTemplate(templates.ServiceTime),
		options,
		formatGoCode,
	)

	// Package validation

	fileCollection.AddFile(
		"internal/validation/validationerror.go",
		false,
		helper.ParseTemplate(templates.ValidationValidationError),
		options,
		formatGoCode,
	)
}

func (r *goGinRestApiGenerator) generateApiModules(
	fileCollection *generator.FileCollection,
	context *generator.GeneratorContext,
	options *GoGinRestApiOptions,
	helper *generator.GeneratorHelper,
) {
	helper.LogInfo("Generating REST API modules...")

	models := context.GlobalConfig.GetModels()

	if len(models) < 1 {
		helper.LogInfo("No models found on the configuration! No REST API modules to generate.")

		return
	}

	helper.LogInfo("Found %d models. Generating REST API modules.", len(models))

	for _, model := range models {
		// Model

		golangModel := r.golangManager.NewGolangModelFromModel(model)

		fileCollection.AddFile(
			fmt.Sprintf("internal/model/%s.go", golangModel.Filename),
			false,
			helper.ParseTemplate(templates.Model),
			golangModel,
			formatGoCode,
		)

		// Resources

		golangResourceCollection := r.golangManager.NewGolangResourceCollectionFromGolangModel(golangModel)
		golangResourceFile := golang.NewGolangResourceFile(golangResourceCollection)

		templateData := helper2.TemplateData{
			GlobalConfig:    context.GlobalConfig,
			GeneratorConfig: options,
			ExtraData:       golangResourceFile,
		}

		fileCollection.AddFile(
			fmt.Sprintf("internal/resource/%s.go", golangResourceCollection.Filename),
			false,
			helper.ParseTemplate(templates.ResourceSpecific),
			templateData,
			formatGoCode,
		)
	}
}

func (r *goGinRestApiGenerator) GetName() string {
	return Type
}

// Static functions

func NewGoGinRestApiGenerator(golangManager *golang.GolangManager, sqlManager *sql.SqlManager) generator.Generator {
	return &goGinRestApiGenerator{
		golangManager: golangManager,
		sqlManager:    sqlManager,
	}
}

func formatGoCode(filename string, contents string) (string, error) {
	contentBytes, err := format.Source([]byte(contents))

	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}
