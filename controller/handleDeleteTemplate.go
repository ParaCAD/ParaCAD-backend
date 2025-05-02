package controller

import (
	"errors"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (c *Controller) HandleDeleteTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	loggedInUserUUID, loggedInUserRole, err := auth.GetUserIDAndRoleFromRequest(r)
	if err != nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, err)
		return
	}

	templateMeta, err := c.db.GetTemplateMetaByUUID(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, err)
		return
	}
	if templateMeta == nil {
		utils.HandleErr(r, w, http.StatusNotFound, errors.New("template not found"))
		return
	}

	if templateMeta.OwnerUUID != loggedInUserUUID && loggedInUserRole != auth.RoleAdmin {
		utils.HandleErr(r, w, http.StatusForbidden, errors.New("user is not the owner of the template"))
		return
	}

	err = c.db.DeleteTemplate(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
