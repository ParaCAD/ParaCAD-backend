package generator

import "github.com/google/uuid"

type FilledTemplate struct {
	UUID     uuid.UUID
	Template []byte
	Params   []Parameter
}

type Parameter struct {
	Name  string
	Key   string
	Value string
}
