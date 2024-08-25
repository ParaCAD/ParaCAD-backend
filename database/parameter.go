package database

import "strconv"

// TODO: consider allowing constrains containing other parameters (make constrain an interface)
// TODO: consider making just one parameter with Type field and value of type T

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
