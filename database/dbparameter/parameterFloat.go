package dbparameter

import "strconv"

type FloatParameter struct {
	Name         string
	DisplayName  string
	DefaultValue float64
	MinValue     float64
	MaxValue     float64
	Step         float64
}

func (p FloatParameter) GetType() ParameterType {
	return ParameterTypeFloat
}

func (p FloatParameter) GetName() string {
	return p.Name
}

func (p FloatParameter) GetDisplayName() string {
	return p.DisplayName
}

func (p FloatParameter) String() string {
	return strconv.FormatFloat(p.DefaultValue, 'f', -1, 64)
}
