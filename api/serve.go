package api

import "net/http"

func (a *Api) Serve() error {
	a.addRoutes()
	return http.ListenAndServe(":"+a.port, a.router)
}
