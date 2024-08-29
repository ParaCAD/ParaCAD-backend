package auth

import "time"

type Auth struct {
	JWTSecret []byte
	Duration  time.Duration
}

type AuthRole string

const (
	RoleUser  AuthRole = "user"
	RoleAdmin AuthRole = "admin"
)

func Init(secret string, durationMinutes int) *Auth {
	return &Auth{
		JWTSecret: []byte(secret),
		Duration:  time.Duration(durationMinutes) * time.Minute,
	}
}
