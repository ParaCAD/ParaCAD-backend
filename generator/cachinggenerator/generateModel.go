package cachinggenerator

import (
	"slices"
	"strings"

	"github.com/ParaCAD/ParaCAD-backend/generator"
)

func (g CachingGenerator) GenerateModel(template generator.FilledTemplate) ([]byte, error) {
	cacheKey := createCacheKey(template)
	model, err := g.db.CacheGetModel(cacheKey)
	if err == nil && model != nil {
		return model, nil
	}
	if err != nil {
		return nil, err
	}

	model, err = g.generator.GenerateModel(template)
	if err != nil {
		return nil, err
	}

	go g.db.CacheSetModel(cacheKey, model)

	return model, nil
}

func createCacheKey(template generator.FilledTemplate) string {
	key := strings.Builder{}
	key.WriteString(template.UUID.String())
	key.WriteString("-")
	slices.SortFunc(template.Params, func(a, b generator.Parameter) int {
		return strings.Compare(a.Key, b.Key)
	})
	for i, param := range template.Params {
		key.WriteString(param.Key)
		key.WriteString("=")
		key.WriteString(param.Value)
		if i < len(template.Params)-1 {
			key.WriteString("&")
		}
	}
	return key.String()
}
