package api

import (
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/julienschmidt/httprouter"
)

type Api struct {
	port       string
	router     *httprouter.Router
	controller *controller.Controller
}

func New(controller *controller.Controller, port string) *Api {
	return &Api{
		port:       port,
		router:     httprouter.New(),
		controller: controller,
	}
}
