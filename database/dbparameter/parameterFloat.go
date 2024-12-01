package dbparameter

import (
	"fmt"
	"strconv"
)

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

func (p FloatParameter) VerifyValue(value string) error {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	if val < p.MinValue || val > p.MaxValue {
		return fmt.Errorf("value of %s: %f out of range (%f, %f)", p.Name, val, p.MinValue, p.MaxValue)
	}
	return nil
}
