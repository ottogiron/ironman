package template

import (
	"io"
	"log"

	"github.com/ironman-project/ironman/template/engine"
)

//GeneratorOption represents a generatorOption setter
type GeneratorOption func(*generator)

//SetGeneratorOutput sets the generator logger
func SetGeneratorOutput(output io.Writer) GeneratorOption {
	return func(generator *generator) {
		generator.logger = log.New(output, "", 0)
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
