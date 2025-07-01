package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/validators"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UUID uuid.UUID `json:"uuid"`
}

func (c *Controller) HandleRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	request := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = validators.Username(request.Username)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = validators.Email(request.Email)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	exists, err := c.db.IsUsernameOrEmailUsed(request.Username, request.Email)
	if err != nil {
		utils.HandleErr(r, w, http.StatusConflict, err)
		return
	}

	if exists {
		utils.HandleErr(r, w, http.StatusConflict, errors.New("username or email already in use"))
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	user := database.User{
		UUID:      uuid.New(),
		Username:  request.Username,
		Email:     request.Email,
		Password:  passwordHash,
		Role:      database.RoleUser,
		Deleted:   nil,
		Created:   time.Now(),
		LastLogin: nil,
	}

	err = c.db.CreateUser(user)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := RegisterResponse{
		UUID: user.UUID,
	}
	json.NewEncoder(w).Encode(response)
}
