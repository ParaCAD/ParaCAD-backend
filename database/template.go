package database

import (
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

// Database offers templates with different amount of fields, depending on the use case

type PageTemplate struct {
	UUID        uuid.UUID
	Name        string
	Description string
	PreviewURL  string
	Parameters  []dbparameter.Parameter

	OwnerUUID uuid.UUID
	OwnerName string
}

type ContentTemplate struct {
	UUID       uuid.UUID
	Name       string
	Template   string
	Parameters []dbparameter.Parameter
}

type Template struct {
	UUID        uuid.UUID
	OwnerUUID   uuid.UUID
	Name        string
	Description string
	Preview     []byte
	Template    string
	Parameters  []dbparameter.Parameter
}
