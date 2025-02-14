package dbparameter

// TODO: consider allowing constrains containing other parameters (make constrain an interface)

type Parameter interface {
	GetType() ParameterType
	GetName() string
	GetDisplayName() string
	String() string
	VerifyValue(string) error
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

func (p parameterType) String() string {
	return string(p)
}

func (p parameterType) ParameterType() parameterType {
	return p
}
