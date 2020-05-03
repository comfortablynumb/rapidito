package sql

import "github.com/comfortablynumb/rapidito/model"

// Interfaces

type Vendor interface {
	GenerateUpMigrationsForModels(models []model.Model) string
	GenerateDownMigrationsForModels(models []model.Model) string
}
