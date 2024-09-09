package dbparameter

import "strconv"

type BoolParameter struct {
	Name         string
	DisplayName  string
	DefaultValue bool
}

func (p BoolParameter) GetType() ParameterType {
	return ParameterTypeBool
}

func (p BoolParameter) GetName() string {
	return p.Name
}

func (p BoolParameter) GetDisplayName() string {
	return p.DisplayName
}

func (p BoolParameter) GetValue() interface{} {
	return p.DefaultValue
}

func (p BoolParameter) String() string {
	return strconv.FormatBool(p.DefaultValue)
}
