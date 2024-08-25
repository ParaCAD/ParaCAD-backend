package database

import "github.com/google/uuid"

type Database interface {
	GetUserByUUID(uuid.UUID) (User, error)
	GetUserByUsername(string) (User, error)
	GetUserByEmail(string) (User, error)

	GetTemplateByUUID(uuid.UUID) (Template, error)
	SearchTemplates(SearchParameters) ([]Template, error)
}
