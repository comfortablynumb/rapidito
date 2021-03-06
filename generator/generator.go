package generator

type Generator interface {
	Generate(fileCollection *FileCollection, context *GeneratorContext, helper *GeneratorHelper) error
	PostGeneration(fileCollection *FileCollection, context *GeneratorContext, helper *GeneratorHelper) error
	GetName() string
}
