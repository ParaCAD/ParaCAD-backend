package dbparameter

import "strconv"

type IntParameter struct {
	Name         string
	DisplayName  string
	DefaultValue int
	MinValue     int
	MaxValue     int
}

func (p IntParameter) GetType() ParameterType {
	return ParameterTypeInt
}

func (p IntParameter) GetName() string {
	return p.Name
}

func (p IntParameter) GetDisplayName() string {
	return p.DisplayName
}

func (p IntParameter) String() string {
	return strconv.Itoa(p.DefaultValue)
}
