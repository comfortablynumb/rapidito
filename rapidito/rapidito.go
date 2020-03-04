package rapidito

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/comfortablynumb/rapidito/configuration"
	"github.com/comfortablynumb/rapidito/errorhandler"
	"github.com/comfortablynumb/rapidito/generator"
	"github.com/comfortablynumb/rapidito/helper"
	"github.com/comfortablynumb/rapidito/logger"
)

// Rapidito This is the rapidito instance.
type Rapidito struct {
	ErrorHandler *errorhandler.ErrorHandler
	FileHelper   *helper.FileHelper
	Logger       *logger.Logger
	Generators   map[string]generator.Generator
}

func (r *Rapidito) Generate(configFile string) error {
	config := r.parseConfig(configFile)
	generatorHelper := generator.NewGeneratorHelper(r.ErrorHandler, r.FileHelper, r.Logger)

	for _, generatorConfig := range config.Generators {
		r.runGenerator(config, generatorConfig, generatorHelper)
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

func (r *Rapidito) runGenerator(globalConfig *configuration.Config, generatorConfig configuration.GeneratorConfig, generatorHelper *generator.GeneratorHelper) {
	gen, found := r.Generators[generatorConfig.Type]

	if !found {
		r.HandleError(errors.New("Unknown generator type"), "Generator type is invalid. Please check that the value '%s' is correct", generatorConfig.Type)
	}

	generatorContext := generator.NewGeneratorContext(r.ErrorHandler, generatorConfig.Options)
	fileCollection := generator.NewFileCollection()

	err := gen.Generate(fileCollection, generatorContext, generatorHelper)

	r.HandleIfError(err, "Generator of type '%s' returned an error.", generatorConfig.Type)

	finalPath := filepath.Clean(fmt.Sprintf("%s/%s", globalConfig.TargetPath, generatorConfig.RelativePath))

	// Create all required directories if they are missing

	r.FileHelper.MkDirAll(finalPath, 0755)

	for _, file := range fileCollection.GetFiles() {
		path := filepath.Clean(fmt.Sprintf("%s/%s", finalPath, file.RelativePath))
		dir := filepath.Dir(path)

		r.FileHelper.MkDirAll(dir, 0755)

		f, err := os.Create(path)

		r.HandleIfError(err, "Could not create file: %s", path)

		err = file.Template.Execute(f, file.TemplateData)

		r.HandleIfError(err, "Could not generate file: %s", path)
	}

}

func (r *Rapidito) HandleError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.Handle(err, msg, args...)
}

func (r *Rapidito) HandleIfError(err error, msg string, args ...interface{}) {
	r.ErrorHandler.HandleIfError(err, msg, args...)
}

func NewRapidito() *Rapidito {
	log := logger.NewLogger()
	errorHandler := errorhandler.NewErrorHandler(log)

	return &Rapidito{
		ErrorHandler: errorHandler,
		FileHelper:   helper.NewFileHelper(errorHandler),
		Logger:       log,
		Generators:   make(map[string]generator.Generator),
	}
}