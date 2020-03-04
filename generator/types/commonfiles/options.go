package commonfiles

// Structs

type CommonFilesOptions struct {
	Name         string `yaml:"name" mapstructure:"name"`
	FriendlyName string `yaml:"friendly_name" mapstructure:"friendly_name"`
	Description  string `yaml:"description" mapstructure:"description"`
	LicenseType  string `yaml:"license_type" mapstructure:"license_type"`
}

// Static functions

func NewCommonFilesOptions() *CommonFilesOptions {
	return &CommonFilesOptions{
		Name:         "project-name",
		FriendlyName: "Project Name (Friendly Name)",
		Description:  "Some description for this project.",
	}
}
