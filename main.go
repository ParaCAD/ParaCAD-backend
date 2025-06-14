package main

import (
	"log"
	"log/slog"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/database/sqldb"
	"github.com/ParaCAD/ParaCAD-backend/fsstore"
	"github.com/ParaCAD/ParaCAD-backend/generator/cachinggenerator"
	"github.com/ParaCAD/ParaCAD-backend/generator/openscadgenerator"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
)

func main() {
	cfg := utils.MustLoadConfig()

	logging.Init(slog.LevelDebug)

	auth := auth.New(cfg.JWTSecret, 15)

	sqlDB, err := sqldb.New(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Init()
	defer sqlDB.Close()

	imageStore, err := fsstore.New(cfg.ImageStorageDir)
	if err != nil {
		log.Fatal(err)
	}

	generator := cachinggenerator.NewCachingGenerator(openscadgenerator.NewOpenSCADGenerator(), sqlDB)

	con := controller.New(auth, sqlDB, imageStore, generator)

	createExampleUsersAndTemplates(con)

	api := api.New(cfg.Port, auth, con)
	err = api.Serve()
	log.Fatal(err)
}
