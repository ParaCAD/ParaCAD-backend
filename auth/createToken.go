package auth

import "github.com/golang-jwt/jwt"

func (a *Auth) CreateToken(userID string, username string, role AuthRole) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimISS:   "ParaCAD",
		claimSUB:   userID,
		claimAUD:   "ParaCAD",
		claimEXP:   jwt.TimeFunc().Add(a.Duration).Unix(),
		claimIAT:   jwt.TimeFunc().Unix(),
		claimROLE:  role,
		claimLOGIN: username,
	})
	token, err := claims.SignedString(a.JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
