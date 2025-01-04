package database

import (
	"time"

	"github.com/google/uuid"
)

// Database methods do not error when no data is found, but return nil instead

type Database interface {
	GetUserByUUID(uuid.UUID) (*User, error)
	GetUserByUsername(string) (*User, error)
	DeleteUser(uuid.UUID) error

	GetUserSecurityByUsername(string) (*UserSecurity, error)
	SetUserLastLogin(uuid.UUID, time.Time) error

	GetTemplateByUUID(uuid.UUID) (*Template, error)
	GetTemplateWithOwnerByUUID(uuid.UUID) (*TemplatePage, error)
	GetTemplateContentByUUID(uuid.UUID) (*TemplateContent, error)
	GetTemplateMetaByUUID(uuid.UUID) (*TemplateMeta, error)
	DeleteTemplate(uuid.UUID) error

	SearchTemplates(SearchParameters) ([]Template, error)

	SetTemplateMarked(uuid.UUID, bool) error
}
