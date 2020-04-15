package model

type Model struct {
	Name   string                `yaml:"-"`
	Fields map[string]ModelField `yaml:"fields"`
}

func (m *Model) GetFields() []ModelField {
	fields := make([]ModelField, 0)

	for name, field := range m.Fields {
		field.Name = name

		fields = append(fields, field)
	}

	return fields
}

type ModelField struct {
	Name       string    `yaml:"-"`
	Type       ModelType `yaml:"type"`
	CustomType string    `yaml:"custom_type"`
}
