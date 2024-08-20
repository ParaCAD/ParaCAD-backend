package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ParaCAD/ParaCAD-backend/api"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dummydb"
	"github.com/ParaCAD/ParaCAD-backend/generator"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
	"github.com/google/uuid"
)

func main() {
	cfg := utils.MustLoadConfig()

	logging.Init(slog.LevelDebug)

	// TODO: a real create database driver
	db := dummydb.New()
	template, _ := db.GetTemplateByUUID(database.TemplateID(uuid.Nil))

	filledTemplate := generator.FilledTemplate{
		UUID:     uuid.UUID(template.UUID),
		Template: []byte(template.Template),
		Params: []generator.Parameter{
			{
				Name:  template.Parameters[0].GetDisplayName(),
				Key:   template.Parameters[0].GetName(),
				Value: template.Parameters[0].String(),
			},
		},
	}

	generated, _ := generator.Generate(filledTemplate)
	os.WriteFile(template.Name+".stl", generated, 0644)

	con := controller.New()

	api := api.New(con, cfg.Port)
	err := api.Serve()
	log.Fatal(err)
}
