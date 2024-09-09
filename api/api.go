package api

import (
	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	port       string
	router     *httprouter.Router
	auth       *auth.Auth
	controller *controller.Controller
}

func New(port string, auth *auth.Auth, controller *controller.Controller) *API {
	return &API{
		port:       port,
		router:     httprouter.New(),
		auth:       auth,
		controller: controller,
	}
}
