package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID  `db:"uuid"`
	Username    string     `db:"username"`
	Email       string     `db:"email"`
	Password    []byte     `db:"password"`
	Description string     `db:"description"`
	Role        role       `db:"role"`
	Deleted     *time.Time `db:"deleted"`
	Created     time.Time  `db:"created"`
	LastLogin   *time.Time `db:"last_login"`
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
