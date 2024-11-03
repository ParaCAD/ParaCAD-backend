package main

import (
	"log"
	"log/slog"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/database/dummydb"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
)

func main() {
	cfg := utils.MustLoadConfig()

	logging.Init(slog.LevelDebug)

	auth := auth.New(cfg.JWTSecret, 15)

	// TODO: a real create database driver
	db := dummydb.New()

	con := controller.New(auth, db)

	api := api.New(cfg.Port, auth, con)
	err := api.Serve()
	log.Fatal(err)
}
