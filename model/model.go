package model

import "sort"

type Model struct {
	Name       string                `yaml:"-"`
	Fields     map[string]ModelField `yaml:"fields"`
	PrimaryKey []string              `yaml:"primary_key"`
}

func (m *Model) GetFields() []ModelField {
	fields := make([]ModelField, 0)
	fieldNames := make([]string, 0)

	for name, _ := range m.Fields {
		fieldNames = append(fieldNames, name)
	}

	sort.Strings(fieldNames)

	for _, name := range fieldNames {
		field := m.Fields[name]

		field.Name = name

		fields = append(fields, field)
	}

	return fields
}

func (m *Model) GetField(name string) *ModelField {
	field, found := m.Fields[name]

	if !found {
		return nil
	}

	return &field
}

type ModelField struct {
	Name           string    `yaml:"-"`
	Type           ModelType `yaml:"type"`
	CustomType     string    `yaml:"custom_type"`
	HideOnSearch   bool      `yaml:"hide_on_search"`
	HideOnCreate   bool      `yaml:"hide_on_create"`
	HideOnUpdate   bool      `yaml:"hide_on_update"`
	HideOnResponse bool      `yaml:"hide_on_response"`
}
