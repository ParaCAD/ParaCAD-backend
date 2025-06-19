package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type EditUserRequest struct {
	Description string `json:"description"`
}

func (c *Controller) HandleEditUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}
	// TODO check if user can edit

	request := EditUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	// TODO
	fmt.Printf("Received edit user request: user: %s description: %s", userUUID, request.Description)
}
