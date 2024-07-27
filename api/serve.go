package api

import "net/http"

func (a *API) Serve() error {
	a.addRoutes()
	return http.ListenAndServe(":"+a.port, a.router)
}
