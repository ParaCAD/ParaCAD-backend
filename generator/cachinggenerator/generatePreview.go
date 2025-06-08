package cachinggenerator

import (
	"github.com/ParaCAD/ParaCAD-backend/generator"
)

func (g CachingGenerator) GeneratePreview(template generator.FilledTemplate) ([]byte, error) {
	return g.generator.GeneratePreview(template)
}
