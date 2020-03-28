package generator

// Structs

type ExecutedGenerator struct {
	Generator        Generator
	FileCollection   *FileCollection
	GeneratorContext *GeneratorContext
}

// Static functions

func NewExecutedGenerator(generator Generator, fileCollection *FileCollection, context *GeneratorContext) *ExecutedGenerator {
	return &ExecutedGenerator{
		Generator:        generator,
		FileCollection:   fileCollection,
		GeneratorContext: context,
	}
}
