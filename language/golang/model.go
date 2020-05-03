package golang

// Structs

type GolangModel struct {
	Name        string
	Filename    string
	StructName  string
	BuilderName string
	Fields      map[string]GolangModelField
	PrimaryKey  []string
}

type GolangModelField struct {
	Name             string
	StructFieldName  string
	BuilderFieldName string
	Type             GolangType
	CustomType       string
	HideOnSearch     bool
	HideOnCreate     bool
	HideOnUpdate     bool
	HideOnResponse   bool
}

func (m *GolangModel) GetField(name string) *GolangModelField {
	field, found := m.Fields[name]

	if !found {
		return nil
	}

	return &field
}
