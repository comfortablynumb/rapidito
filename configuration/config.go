package configuration

import "github.com/comfortablynumb/rapidito/model"

// Structs

type Config struct {
	Version    string                 `yaml:"file_version"`
	TargetPath string                 `yaml:"target_path"`
	Generators []GeneratorConfig      `yaml:"generators"`
	Models     map[string]model.Model `yaml:"models"`
}

func (c *Config) GetModels() []model.Model {
	models := make([]model.Model, 0)

	for name, myModel := range c.Models {
		myModel.Name = name

		models = append(models, myModel)
	}

	return models
}

type GeneratorConfig struct {
	Type         string                 `yaml:"type"`
	RelativePath string                 `yaml:"relative_path"`
	Options      map[string]interface{} `yaml:"options"`
}
