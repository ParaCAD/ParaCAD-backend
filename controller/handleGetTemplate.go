package controller

import (
	"errors"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/julienschmidt/httprouter"
)

func (c *Controller) HandleGetTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.HandleErr(r, w, http.StatusNotImplemented, errors.New("not implemented"))
}
