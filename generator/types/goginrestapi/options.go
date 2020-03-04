package goginrestapi

// Structs

type RestOptions struct {
	Name         string   `yaml:"name" mapstructure:"name"`
	FriendlyName string   `yaml:"friendly_name" mapstructure:"friendly_name"`
	Description  string   `yaml:"description" mapstructure:"description"`
	Actions      []string `yaml:"actions" mapstructure:"actions"`
}

// Static functions

func NewRestOptions() *RestOptions {
	return &RestOptions{
		Name:         "rest-api-name",
		FriendlyName: "REST API Name",
		Description:  "REST API description",
	}
}
