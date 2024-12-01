package dbparameter

import (
	"fmt"
	"strconv"
)

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

func (p IntParameter) VerifyValue(value string) error {
	val, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if val < p.MinValue || val > p.MaxValue {
		return fmt.Errorf("value of %s: %d out of range (%d, %d)", p.Name, val, p.MinValue, p.MaxValue)
	}
	return nil
}
