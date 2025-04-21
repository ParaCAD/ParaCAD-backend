package controller

import (
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/julienschmidt/httprouter"
)

func (c *Controller) HandleGetImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	imageName := p.ByName("FILENAME")
	imageData, err := c.imageStore.GetFile(imageName)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
	if imageData == nil {
		utils.HandleErr(r, w, http.StatusNotFound, nil)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageData)
}
