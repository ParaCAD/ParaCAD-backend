package database

import (
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

// Database offers templates with different amount of fields, depending on the use case

type TemplatePage struct {
	UUID        uuid.UUID `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	PreviewURL  *string   `db:"preview"`
	Parameters  []dbparameter.Parameter

	OwnerUUID    uuid.UUID  `db:"owner_uuid"`
	OwnerName    string     `db:"owner_name"`
	OwnerDeleted *time.Time `db:"owner_deleted"`
}

type TemplateContent struct {
	UUID       uuid.UUID `db:"uuid"`
	Name       string    `db:"name"`
	Template   string    `db:"template"`
	Parameters []dbparameter.Parameter
}

type TemplateMeta struct {
	UUID      uuid.UUID `db:"uuid"`
	Name      string    `db:"name"`
	OwnerUUID uuid.UUID `db:"owner_uuid"`
}

// TODO: remove

type Template struct {
	UUID        uuid.UUID `db:"uuid"`
	OwnerUUID   uuid.UUID `db:"owner_uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Preview     *string   `db:"preview"`
	Template    string    `db:"template"`
	Parameters  []dbparameter.Parameter
}
