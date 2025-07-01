package controller

import (
	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/fsstore"
	"github.com/ParaCAD/ParaCAD-backend/generator"
)

type Controller struct {
	auth       *auth.Auth
	db         database.Database
	imageStore *fsstore.FSStore
	generator  generator.Generator
}

func New(auth *auth.Auth, db database.Database, imageStore *fsstore.FSStore, generator generator.Generator) *Controller {
	return &Controller{
		auth:       auth,
		db:         db,
		imageStore: imageStore,
		generator:  generator,
	}
}
