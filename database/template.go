package database

import "github.com/google/uuid"

type Template struct {
	UUID        TemplateID
	OwnerUUID   UserID
	Name        string
	Description string
	Preview     []byte
	Template    string
	Parameters  []Parameter
}

type TemplateID uuid.UUID
