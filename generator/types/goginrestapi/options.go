package goginrestapi

// Structs

type GoGinRestApiOptions struct {
	Name         string   `yaml:"name" mapstructure:"name"`
	Package      Package  `yaml:"package" mapstructure:"package"`
	FriendlyName string   `yaml:"friendly_name" mapstructure:"friendly_name"`
	Description  string   `yaml:"description" mapstructure:"description"`
	LicenseType  string   `yaml:"license_type" mapstructure:"license_type"`
	Version      string   `yaml:"version" mapstructure:"version"`
	Actions      []string `yaml:"actions" mapstructure:"actions"`
}

type Package struct {
	Name string `yaml:"name" mapstructure:"name"`
}

// Static functions

func NewGoGinRestApiOptions() *GoGinRestApiOptions {
	return &GoGinRestApiOptions{
		Name:         "rest-api-name",
		FriendlyName: "REST API Name",
		Description:  "REST API description",
		LicenseType:  "MIT",
		Version:      "1.0",
	}
}
