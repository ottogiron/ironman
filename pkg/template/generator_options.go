package template

import (
	"io"

	"github.com/ironman-project/ironman/pkg/template/engine"
)

//GeneratorOption represents a generatorOption setter
type GeneratorOption func(*generator)

//SetGeneratorOutput sets the generator logger
func SetGeneratorOutput(out io.Writer) GeneratorOption {
	return func(generator *generator) {
		generator.out = out
	}
}

//SetGeneratorEngine sets the generator template engine
func SetGeneratorEngine(engine engine.Factory) GeneratorOption {
	return func(generator *generator) {
		generator.engine = engine
	}
}

//SetGeneratorIgnoreList sets the generator file ignore lists
func SetGeneratorIgnoreList(ignoreList []string) GeneratorOption {
	return func(generator *generator) {
		generator.ignore = ignoreList
	}
}
