package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CtxKey string

const UserIDKey CtxKey = "userID"
const RoleKey CtxKey = "role"

func (a *Auth) Middleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no token provided"))
			return
		}

		userID, role, err := a.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)
		*r = *r.WithContext(ctx)

		h(w, r, p)
	}
}

func GetUserIDAndRoleFromRequest(r *http.Request) (string, AuthRole, error) {
	ctx := r.Context()
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", "", errors.New("no user ID in context")
	}
	role, ok := ctx.Value(RoleKey).(AuthRole)
	if !ok {
		return "", "", errors.New("no role in context")
	}
	return userID, role, nil
}
