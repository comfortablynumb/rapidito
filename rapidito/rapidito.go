package rapidito

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/comfortablynumb/rapidito/configuration"
	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/comfortablynumb/rapidito/generator"
	"github.com/comfortablynumb/rapidito/generator/types/goginrestapi"
	"github.com/comfortablynumb/rapidito/helper"
	"github.com/comfortablynumb/rapidito/language/golang"
	"github.com/comfortablynumb/rapidito/logger"
)

// Rapidito This is the rapidito instance.
type Rapidito struct {
	ErrorHandler   *errorhandler.ErrorHandler
	FileHelper     *helper.FileHelper
	TemplateHelper *helper.TemplateHelper
	Logger         *logger.Logger
	Generators     map[string]generator.Generator
	GolangManager  *golang.GolangManager
}

func (r *Rapidito) Generate(configFile string) error {
	config := r.parseConfig(configFile)
	generatorHelper := generator.NewGeneratorHelper(r.ErrorHandler, r.FileHelper, r.TemplateHelper, r.Logger)
	executedGenerators := make([]*generator.ExecutedGenerator, 0)
	completeFileCollection := generator.NewFileCollection()

	// Run each generator

	for _, generatorConfig := range config.Generators {
		executedGenerator := r.runGenerator(*config, generatorConfig, generatorHelper)

		completeFileCollection.AddFromCollection(executedGenerator.FileCollection)

		executedGenerators = append(executedGenerators, executedGenerator)
	}

	// Now, execute the PostGeneration method on each generator

	for _, executedGenerator := range executedGenerators {
		err := executedGenerator.Generator.PostGeneration(completeFileCollection, executedGenerator.GeneratorContext, generatorHelper)

		r.ErrorHandler.HandleIfError(err, "An error occurred while execution the 'PostGeneration' method on generator: %s", executedGenerator.Generator.GetName())
	}

	return nil
}

func (r *Rapidito) RegisterGenerator(generator generator.Generator) {
	r.Generators[generator.GetName()] = generator
}

func (r *Rapidito) parseConfig(configFile string) *configuration.Config {
	contents := r.FileHelper.GetFileContents(configFile)

	config := &configuration.Config{}

	r.FileHelper.ParseYaml(contents, config)

	if config.TargetPath == "" {
		// Use current directory by default as the target path

		path, err := os.Executable()

		if err != nil {
			r.Logger.Panic(err, "Could not obtain the executable path.")
		}

		config.TargetPath = filepath.Dir(path)

		r.Logger.Info("No target path defined. Defaults to: %s", config.TargetPath)
	}

	return config
}

func (r *Rapidito) runGenerator(
	globalConfig configuration.Config,
	generatorConfig configuration.GeneratorConfig,
	generatorHelper *generator.GeneratorHelper,
) *generator.ExecutedGenerator {
	gen, found := r.Generators[generatorConfig.Type]

	if !found {
		r.HandleError(errors.New("Unknown generator type"), "Generator type is invalid. Please check that the value '%s' is correct", generatorConfig.Type)
	}

	generatorContext := generator.NewGeneratorContext(r.ErrorHandler, globalConfig, generatorConfig.Options)
	fileCollection := generator.NewFileCollection()

	err := gen.Generate(fileCollection, generatorContext, generatorHelper)

	r.HandleIfError(err, "Generator of type '%s' returned an error.", generatorConfig.Type)

	finalPath := filepath.Clean(fmt.Sprintf("%s/%s", globalConfig.TargetPath, generatorConfig.RelativePath))

	// Create all required directories if they are missing

	r.FileHelper.MkDirAll(finalPath, 0755)

	r.Logger.Info("Generating files for '%s' on directory: %s", gen.GetName(), finalPath)

	for _, file := range fileCollection.GetFiles() {
		path := filepath.Clean(fmt.Sprintf("%s/%s", finalPath, file.RelativePath))
		dir := filepath.Dir(path)

		r.FileHelper.MkDirAll(dir, 0755)

		if file.SkipIfExists && r.FileHelper.FileExists(path) {
			r.Logger.Info("File '%s' already exists and its 'skip if exists' option is true. Skipping...", path)

			continue
		}

		var buf bytes.Buffer

		err = file.Template.Execute(&buf, file.TemplateData)

		r.HandleIfError(err, "Could not execute template for path: %s", path)

		contents := buf.String()

		if file.PreFileWriteFunc != nil {
			contents, err = file.PreFileWriteFunc(path, contents)

			r.HandleIfError(err, "There was an error while executing a pre file write func for path: %s", path)
		}

		f, err := os.Create(path)

		r.HandleIfError(err, "Could not create file: %s", path)

		_, err = f.WriteString(contents)

		r.HandleIfError(err, "Could not generate file: %s", path)
	}

	return generator.NewExecutedGenerator(gen, fileCollection, generatorContext)
}

func (r *Rapidito) HandleError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.Handle(err, msg, args...)
}

func (r *Rapidito) HandleIfError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.HandleIfError(err, msg, args...)
}

func (r *Rapidito) initialize() {
	r.RegisterGenerator(goginrestapi.NewGoGinRestApiGenerator(r.GolangManager))
}

func NewRapidito() *Rapidito {
	log := logger.NewLogger()
	errorHandler := errorhandler.NewErrorHandler(log)
	golangManager := golang.NewGolangManager(log, errorHandler)

	rapidito := &Rapidito{
		ErrorHandler:   errorHandler,
		FileHelper:     helper.NewFileHelper(errorHandler),
		TemplateHelper: helper.NewTemplateHelper(errorHandler),
		Logger:         log,
		Generators:     make(map[string]generator.Generator),
		GolangManager:  golangManager,
	}

	rapidito.initialize()

	return rapidito
}
