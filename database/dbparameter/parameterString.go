package dbparameter

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
