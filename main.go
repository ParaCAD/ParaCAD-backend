package main

import (
	"log"
	"log/slog"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/database/sqldb"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
)

func main() {
	cfg := utils.MustLoadConfig()

	logging.Init(slog.LevelDebug)

	auth := auth.New(cfg.JWTSecret, 15)

	// TODO: a real create database driver
	// db := dummydb.New()

	sqlDB, err := sqldb.New(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Init()
	defer sqlDB.Close()

	con := controller.New(auth, sqlDB)

	api := api.New(cfg.Port, auth, con)
	err = api.Serve()
	log.Fatal(err)
}
