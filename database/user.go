package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID
	Username string
	Email    string
	Role     role
	Deleted  bool
}

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

type UserSecurity struct {
	UUID      uuid.UUID
	Username  string
	Email     string
	Password  []byte
	Role      role
	Deleted   bool
	Created   time.Time
	LastLogin time.Time
}
