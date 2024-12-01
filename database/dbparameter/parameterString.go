package dbparameter

import "fmt"

type StringParameter struct {
	Name         string
	DisplayName  string
	DefaultValue string
	MinLength    int
	MaxLength    int
}

func (p StringParameter) GetType() ParameterType {
	return ParameterTypeString
}

func (p StringParameter) GetName() string {
	return p.Name
}

func (p StringParameter) GetDisplayName() string {
	return p.DisplayName
}

func (p StringParameter) String() string {
	return p.DefaultValue
}

func (p StringParameter) VerifyValue(value string) error {
	if len(value) < p.MinLength || len(value) > p.MaxLength {
		return fmt.Errorf("length of %s: %s (%d) out of range (%d, %d)", p.Name, value, len(value), p.MinLength, p.MaxLength)
	}
	return nil
}
