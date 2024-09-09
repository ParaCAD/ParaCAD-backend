package dbparameter

// TODO: consider allowing constrains containing other parameters (make constrain an interface)

type Parameter interface {
	GetType() ParameterType
	GetName() string
	GetDisplayName() string
	String() string
}

type parameterType string

const (
	ParameterTypeString parameterType = "string"
	ParameterTypeInt    parameterType = "int"
	ParameterTypeFloat  parameterType = "float"
	ParameterTypeBool   parameterType = "bool"
)

type ParameterType interface {
	ParameterType() parameterType
}

func (p parameterType) ParameterType() parameterType {
	return p
}
