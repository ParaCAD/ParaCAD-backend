package dbparameter

type Parameter interface {
	GetType() ParameterType
	GetName() string
	GetDisplayName() string
	String() string
	VerifyValue(string) error
}

type ParameterType string

const (
	ParameterTypeString ParameterType = "string"
	ParameterTypeInt    ParameterType = "int"
	ParameterTypeFloat  ParameterType = "float"
	ParameterTypeBool   ParameterType = "bool"
)

func (p ParameterType) String() string {
	return string(p)
}

func (p ParameterType) ID() int {
	switch p {
	case ParameterTypeString:
		return 0
	case ParameterTypeInt:
		return 1
	case ParameterTypeFloat:
		return 2
	case ParameterTypeBool:
		return 3
	default:
		return -1
	}
}

type ParameterConstraintType string

const (
	ParameterConstraintMinLength ParameterConstraintType = "min_length"
	ParameterConstraintMaxLength ParameterConstraintType = "max_length"
	ParameterConstraintMinValue  ParameterConstraintType = "min_value"
	ParameterConstraintMaxValue  ParameterConstraintType = "max_value"
	ParameterConstraintStep      ParameterConstraintType = "step"
)

var ParameterConstraints = []ParameterConstraintType{
	ParameterConstraintMinLength,
	ParameterConstraintMaxLength,
	ParameterConstraintMinValue,
	ParameterConstraintMaxValue,
	ParameterConstraintStep,
}

func (p ParameterConstraintType) String() string {
	return string(p)
}

func (p ParameterConstraintType) ID() int {
	switch p {
	case ParameterConstraintMinLength:
		return 0
	case ParameterConstraintMaxLength:
		return 1
	case ParameterConstraintMinValue:
		return 2
	case ParameterConstraintMaxValue:
		return 3
	case ParameterConstraintStep:
		return 4
	default:
		return -1
	}
}
