package database

import "github.com/google/uuid"

type Template struct {
	UUID        uuid.UUID
	OwnerUUID   uuid.UUID
	Name        string
	Description string
	Preview     []byte
	Template    string
	Parameters  []Parameter
}
