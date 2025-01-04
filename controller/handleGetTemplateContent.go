package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type GetTemplateContentResponse []byte

func (c *Controller) HandleGetTemplateContent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	templateContent, err := c.db.GetTemplateContentByUUID(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusNotFound, err)
		return
	}

	sb := strings.Builder{}
	for _, param := range templateContent.Parameters {
		sb.WriteString(param.GetName())
		sb.WriteString(" = ")
		sb.WriteString(param.String())
		sb.WriteString(";\n")
	}
	sb.WriteString(templateContent.Template)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.scad\"", templateContent.Name))
	w.Header().Set("Content-Length", string(len(sb.String())))
	w.Write([]byte(sb.String()))
}
