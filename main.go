package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/generator"
	"github.com/ParaCAD/ParaCAD-backend/utils"
)

func main() {
	_, err := generator.Generate(generator.FilledTemplate{})
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(0)

	cfg := utils.MustLoadConfig()

	// TODO: create database driver

	con := controller.New()

	api := api.New(con, cfg.Port)
	err = api.Serve()
	log.Fatal(err)
}
