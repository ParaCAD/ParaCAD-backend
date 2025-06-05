package database

import (
	"time"

	"github.com/google/uuid"
)

// Database methods do not error when no data is found, but return nil instead

type Database interface {
	GetUserByUUID(uuid.UUID) (*User, error)
	GetUserByUsername(string) (*User, error)
	IsUsernameOrEmailUsed(string, string) (bool, error)
	CreateUser(User) error
	DeleteUser(uuid.UUID) error

	SetUserLastLogin(uuid.UUID, time.Time) error

	GetTemplateByUUID(uuid.UUID) (*Template, error)
	GetTemplateWithOwnerByUUID(uuid.UUID) (*TemplatePage, error)
	GetTemplateContentByUUID(uuid.UUID) (*TemplateContent, error)
	GetTemplateMetaByUUID(uuid.UUID) (*TemplateMeta, error)
	GetTemplatesByOwnerUUID(uuid.UUID, int, int) ([]SearchResult, error)
	CreateTemplate(Template) error
	DeleteTemplate(uuid.UUID) error

	SearchTemplates(SearchParameters) ([]SearchResult, error)

	SetTemplateMarked(uuid.UUID, bool) error

	CacheGetModel(string) ([]byte, error)
	CacheSetModel(string, []byte) error
}
