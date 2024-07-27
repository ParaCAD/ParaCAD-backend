package api

import (
	"net/http"
)

func (a *API) addRoutes() *API {
	// TODO: add logging middleware
	// TODO: add auth middleware
	a.router.Handle(http.MethodGet, "/", a.controller.HandleRoot)

	a.router.HandleOPTIONS = true
	return a
}
