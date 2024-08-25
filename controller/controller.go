package controller

import "github.com/ParaCAD/ParaCAD-backend/database"

type Controller struct {
	db database.Database
}

func New(db database.Database) *Controller {
	return &Controller{
		db: db,
	}
}
