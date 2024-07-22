package main

import (
	"log"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/utils"
)

func main() {
	cfg := utils.MustLoadConfig()

	// TODO: create database driver

	con := controller.New()

	api := api.New(con, cfg.Port)
	err := api.Serve()
	log.Fatal(err)
}
