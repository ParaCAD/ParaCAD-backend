package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *Controller) HandleRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("ParaCAD-backend"))
}
