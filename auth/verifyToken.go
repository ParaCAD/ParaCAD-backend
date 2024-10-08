package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func (a *Auth) VerifyToken(tokenString string) (string, AuthRole, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return a.JWTSecret, nil })
	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("no claims in token")
	}

	userID, ok := claims[claimSUB].(string)
	if !ok {
		return "", "", errors.New("no username in token")
	}

	role, ok := claims[claimROLE].(string)
	if !ok {
		return "", "", errors.New("no role in token")
	}

	authRole, err := GetRole(role)
	if err != nil {
		return "", "", err
	}

	return userID, authRole, nil
}
