package auth

import "fmt"

type AuthRole string

const (
	RoleUser  AuthRole = "user"
	RoleAdmin AuthRole = "admin"
)

func GetRole(role string) (AuthRole, error) {
	switch role {
	case "user":
		return RoleUser, nil
	case "admin":
		return RoleAdmin, nil
	default:
		return "", fmt.Errorf("invalid role %s", role)
	}
}
