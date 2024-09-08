package database

import (
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

type Template struct {
	UUID        uuid.UUID
	OwnerUUID   uuid.UUID
	Name        string
	Description string
	Preview     []byte
	Template    string
	Parameters  []dbparameter.Parameter
}
