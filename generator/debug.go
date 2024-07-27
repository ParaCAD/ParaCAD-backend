package generator

import "github.com/google/uuid"

var cubeScad string = `
cube(10, true);
`

var CubeTemplate FilledTemplate = FilledTemplate{
	UUID:     uuid.UUID{},
	Template: []byte(cubeScad),
	Params:   []Parameter{},
}
