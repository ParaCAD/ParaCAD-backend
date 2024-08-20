package api

import (
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
	"github.com/julienschmidt/httprouter"
)

func middleware(h httprouter.Handle) httprouter.Handle {
	return logging.Middleware(h)
}

func (a *API) addRoutes() *API {
	// TODO: add auth middleware
	a.router.Handle(http.MethodGet, "/", middleware(a.controller.HandleRoot))
	a.router.Handle(http.MethodGet, "/template/:uuid", middleware(a.controller.HandleGetTemplate))

	// TODO: get template list endpoint
	// TODO: generate model endpoint
	// TODO: create template endpoint (auth)
	// TODO: delete template endpoint (auth)
	// TODO: get user profile endpoint
	// TODO: create user (register) endpoint
	// TODO: login endpoint
	// TODO: edit user endpoint (auth)
	// TODO: delete user endpoint (auth)
	a.router.HandleOPTIONS = true
	a.router.HandleMethodNotAllowed = true
	return a
}
