package configuration

// Structs

type Config struct {
	Version    string            `yaml:"file_version"`
	TargetPath string            `yaml:"target_path"`
	Generators []GeneratorConfig `yaml:"generators"`
}

type GeneratorConfig struct {
	Type         string                 `yaml:"type"`
	RelativePath string                 `yaml:"relative_path"`
	Options      map[string]interface{} `yaml:"options"`
}
