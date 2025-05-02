package controller

import (
	"errors"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (c *Controller) HandleDeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	loggedInUserUUID, loggedInUserRole, err := auth.GetUserIDAndRoleFromRequest(r)
	if err != nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, err)
		return
	}

	if loggedInUserUUID != userUUID && loggedInUserRole != auth.RoleAdmin {
		utils.HandleErr(r, w, http.StatusForbidden, errors.New("can not delete another user"))
		return
	}

	err = c.db.DeleteUser(userUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
