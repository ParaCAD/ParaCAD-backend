package cachinggenerator

import (
	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/generator"
)

type CachingGenerator struct {
	generator generator.Generator
	db        database.Database
}

func NewCachingGenerator(generator generator.Generator, db database.Database) CachingGenerator {
	return CachingGenerator{
		generator: generator,
		db:        db,
	}
}
