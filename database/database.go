package database

import (
	"time"

	"github.com/google/uuid"
)

type Database interface {
	GetUserByUUID(uuid.UUID) (User, error)
	GetUserByUsername(string) (User, error)

	GetUserSecurityByUsername(string) (UserSecurity, error)
	GetUserSecurityByEmail(string) (UserSecurity, error)
	SetUserLastLogin(uuid.UUID, time.Time) error

	GetTemplateByUUID(uuid.UUID) (Template, error)
	GetTemplateWithOwnerByUUID(uuid.UUID) (TemplateWithOwner, error)
	SearchTemplates(SearchParameters) ([]Template, error)

	SetTemplateMarked(uuid.UUID, bool) error
}
