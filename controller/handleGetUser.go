package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type GetUserResponse struct {
	UserUUID string `json:"user_uuid"`
	UserName string `json:"user_name"`

	Templates []TemplatePreview `json:"templates"`
}

func (c *Controller) HandleGetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	user, err := c.db.GetUserByUUID(userUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, err)
		return
	}

	if user == nil {
		utils.HandleErr(r, w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	templates, err := c.db.GetTemplatesByOwnerUUID(user.UUID, 1, 100) // TODO: pagination
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, err)
		return
	}

	response := GetUserResponse{
		UserUUID:  user.UUID.String(),
		UserName:  user.Username,
		Templates: []TemplatePreview{},
	}

	for _, template := range templates {
		response.Templates = append(response.Templates, searchResponseToTemplatePreview(template))
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
}
