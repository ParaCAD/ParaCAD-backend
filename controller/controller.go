package controller

import (
	"fmt"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/database"
)

type Controller struct {
	db database.Database
}

func New(db database.Database) *Controller {
	return &Controller{
		db: db,
	}
}

const notImplementedString = "%s %s is not implemented yet (try again later).\nFor now, have this racoon --> 🦝"

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, notImplementedString, r.Method, r.URL.Path)
}
