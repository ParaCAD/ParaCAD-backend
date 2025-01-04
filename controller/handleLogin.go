package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Controller) HandleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginRequest := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}
	loginRequest.Username = strings.TrimSpace(loginRequest.Username)

	user, err := c.db.GetUserSecurityByUsername(loginRequest.Username)
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, err)
		return
	}
	if user == nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	if user.Deleted != nil {
		utils.HandleErr(r, w, http.StatusUnauthorized, fmt.Errorf("this account has been deactivated on %s", user.Deleted.Format(time.RFC3339)))
		return
	}

	token, err := c.auth.CreateToken(user.UUID.String(), auth.AuthRole(user.Role))
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	err = c.db.SetUserLastLogin(user.UUID, time.Now())
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	response := LoginResponse{
		Token: token,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
}
