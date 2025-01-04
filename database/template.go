package database

import (
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

// Database offers templates with different amount of fields, depending on the use case

type TemplatePage struct {
	UUID        uuid.UUID
	Name        string
	Description string
	PreviewURL  string
	Parameters  []dbparameter.Parameter

	OwnerUUID uuid.UUID
	OwnerName string
}

type TemplateContent struct {
	UUID       uuid.UUID
	Name       string
	Template   string
	Parameters []dbparameter.Parameter
}

type TemplateMeta struct {
	UUID      uuid.UUID
	Name      string
	OwnerUUID uuid.UUID
}

// TODO: remove

type Template struct {
	UUID        uuid.UUID
	OwnerUUID   uuid.UUID
	Name        string
	Description string
	Preview     []byte
	Template    string
	Parameters  []dbparameter.Parameter
}
