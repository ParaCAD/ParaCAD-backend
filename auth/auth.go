package auth

import "time"

type Auth struct {
	JWTSecret []byte
	Duration  time.Duration
}

func New(secret string, duration time.Duration) *Auth {
	return &Auth{
		JWTSecret: []byte(secret),
		Duration:  duration,
	}
}
