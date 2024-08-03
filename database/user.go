package database

import "github.com/google/uuid"

type User struct {
	UUID     UserID
	Username string
	Password string
	Email    string
	Role     role
	Deleted  bool
}

type UserID uuid.UUID

type role string

const (
	RoleAdmin role = "admin"
	RoleUser  role = "user"
)

type Role interface {
	Role() role
}

func (r role) Role() role {
	return r
}
