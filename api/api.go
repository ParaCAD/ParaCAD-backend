package api

import (
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	port       string
	router     *httprouter.Router
	controller *controller.Controller
}

func New(controller *controller.Controller, port string) *API {
	return &API{
		port:       port,
		router:     httprouter.New(),
		controller: controller,
	}
}
