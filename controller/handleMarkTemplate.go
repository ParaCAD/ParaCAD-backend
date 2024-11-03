package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type MarkTemplateRequest struct {
	Marked bool `json:"marked"`
}

func (c *Controller) HandleMarkTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, role, err := auth.GetUserIDAndRoleFromRequest(r)
	if err != nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, err)
		return
	}

	if role != auth.RoleAdmin {
		utils.HandleErr(r, w, http.StatusForbidden, errors.New("only admins can mark templates"))
		return
	}

	templateUUIDStr := p.ByName("UUID")
	templateUUID, err := uuid.Parse(templateUUIDStr)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	request := MarkTemplateRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = c.db.SetTemplateMarked(templateUUID, request.Marked)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
