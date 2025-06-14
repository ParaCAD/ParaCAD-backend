package api

import (
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
	"github.com/julienschmidt/httprouter"
)

func (a *API) middlewareOpen(h httprouter.Handle) httprouter.Handle {
	// might add some rate limiting here
	return logging.Middleware(h)
}

func (a *API) middlewareAuth(h httprouter.Handle) httprouter.Handle {
	h = a.auth.Middleware(h)
	h = logging.Middleware(h)
	return h
}

func (a *API) addRoutes() *API {
	a.router.Handle(http.MethodGet, "/", a.middlewareOpen(a.controller.HandleRoot))

	// TEMPLATE
	a.router.Handle(http.MethodPost, "/search", a.middlewareOpen(a.controller.HandleSearch))
	a.router.Handle(http.MethodGet, "/template/:UUID", a.middlewareOpen(a.controller.HandleGetTemplate))
	a.router.Handle(http.MethodGet, "/template/:UUID/content", a.middlewareOpen(a.controller.HandleGetTemplateContent))
	a.router.Handle(http.MethodPost, "/template/:UUID/model", a.middlewareOpen(a.controller.HandleGenerateModel))

	a.router.Handle(http.MethodPost, "/template", a.middlewareAuth(a.controller.HandleCreateTemplate))
	a.router.Handle(http.MethodDelete, "/template/:UUID", a.middlewareAuth(a.controller.HandleDeleteTemplate))
	a.router.Handle(http.MethodPatch, "/template/:UUID/mark", a.middlewareAuth(a.controller.HandleMarkTemplate))

	a.router.Handle(http.MethodGet, "/image/:FILENAME", a.middlewareOpen(a.controller.HandleGetImage))

	// USER
	a.router.Handle(http.MethodGet, "/user/:UUID", a.middlewareOpen(a.controller.HandleGetUser))

	a.router.Handle(http.MethodPost, "/user/:UUID", a.middlewareAuth(a.controller.HandleEditUser))
	a.router.Handle(http.MethodDelete, "/user/:UUID", a.middlewareAuth(a.controller.HandleDeleteUser))

	// AUTH
	a.router.Handle(http.MethodPost, "/register", a.middlewareOpen(a.controller.HandleRegister))
	a.router.Handle(http.MethodPost, "/login", a.middlewareOpen(a.controller.HandleLogin))

	a.router.HandleOPTIONS = true
	a.router.HandleMethodNotAllowed = true
	return a
}
