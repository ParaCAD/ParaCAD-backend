package auth

import "time"

type Auth struct {
	JWTSecret []byte
	Duration  time.Duration
}

func New(secret string, durationMinutes int) *Auth {
	return &Auth{
		JWTSecret: []byte(secret),
		Duration:  time.Duration(durationMinutes) * time.Minute,
	}
}
