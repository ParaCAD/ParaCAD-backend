package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
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

func GetUserIDAndRoleFromRequest(r *http.Request) (uuid.UUID, AuthRole, error) {
	ctx := r.Context()
	userUUIDStr, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return uuid.UUID{}, "", errors.New("no user ID in context")
	}
	userUUID, err := uuid.Parse(userUUIDStr)
	if err != nil {
		return uuid.UUID{}, "", err
	}

	role, ok := ctx.Value(RoleKey).(AuthRole)
	if !ok {
		return uuid.UUID{}, "", errors.New("no role in context")
	}
	return userUUID, role, nil
}
