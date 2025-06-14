package controller

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
)

func parseAndValidateCreateTemplateParameter(parameter CreateTemplateRequestParameter) (dbparameter.Parameter, error) {
	if parameter.ParameterName == "" {
		return nil, fmt.Errorf("parameter name must not be empty")
	}
	if parameter.ParameterDisplayName == "" {
		return nil, fmt.Errorf("parameter display name must not be empty")
	}

	var param dbparameter.Parameter
	var err error
	switch parameter.ParameterType {
	case string(dbparameter.ParameterTypeFloat):
		param, err = parseFloatParameter(parameter)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float parameter %s: %w", parameter.ParameterDisplayName, err)
		}
	case string(dbparameter.ParameterTypeInt):
		param, err = parseIntParameter(parameter)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int parameter %s: %w", parameter.ParameterDisplayName, err)
		}
	case string(dbparameter.ParameterTypeString):
		param, err = parseStringParameter(parameter)
		if err != nil {
			return nil, fmt.Errorf("failed to parse string parameter %s: %w", parameter.ParameterDisplayName, err)
		}
	case string(dbparameter.ParameterTypeBool):
		param, err = parseBoolParameter(parameter)
		if err != nil {
			return nil, fmt.Errorf("failed to parse bool parameter %s: %w", parameter.ParameterDisplayName, err)
		}
	default:
		return nil, fmt.Errorf("unknown parameter type: %s", parameter.ParameterType)
	}

	return param, nil
}

func parseFloatParameter(parameter CreateTemplateRequestParameter) (dbparameter.FloatParameter, error) {
	param := dbparameter.FloatParameter{
		Name:        parameter.ParameterName,
		DisplayName: parameter.ParameterDisplayName,
	}

	defaultValue, err := strconv.ParseFloat(parameter.ParameterDefaultValue, 64)
	if err != nil {
		return dbparameter.FloatParameter{}, fmt.Errorf("invalid default value: %w", err)
	}
	param.DefaultValue = defaultValue

	idx := slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMinValue)
	})
	if idx == -1 {
		return dbparameter.FloatParameter{}, errors.New("parameter must have a minimum value constraint")
	}
	minValue, err := strconv.ParseFloat(parameter.ParameterConstraints[idx].Value, 64)
	if err != nil {
		return dbparameter.FloatParameter{}, fmt.Errorf("invalid minimum value: %w", err)
	}
	if minValue > defaultValue {
		return dbparameter.FloatParameter{}, errors.New("minimum value must not be greater than the default value")
	}
	param.MinValue = minValue

	idx = slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMaxValue)
	})
	if idx == -1 {
		return dbparameter.FloatParameter{}, errors.New("parameter must have a maximum value constraint")
	}
	maxValue, err := strconv.ParseFloat(parameter.ParameterConstraints[idx].Value, 64)
	if err != nil {
		return dbparameter.FloatParameter{}, fmt.Errorf("invalid maximum value for parameter: %w", err)
	}
	if maxValue < defaultValue {
		return dbparameter.FloatParameter{}, errors.New("maximum value must not be less than the default value")
	}
	param.MaxValue = maxValue

	if minValue >= maxValue {
		return dbparameter.FloatParameter{}, errors.New("minimum value must be less than maximum value")
	}

	idx = slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintStep)
	})
	if idx == -1 {
		return dbparameter.FloatParameter{}, errors.New("parameter must have a step constraint")
	}
	step, err := strconv.ParseFloat(parameter.ParameterConstraints[idx].Value, 64)
	if err != nil {
		return dbparameter.FloatParameter{}, fmt.Errorf("invalid step value for parameter: %w", err)
	}
	if step <= 0 {
		return dbparameter.FloatParameter{}, errors.New("step value must be greater than 0")
	}
	if (maxValue-minValue)/step < 1 {
		return dbparameter.FloatParameter{}, errors.New("step value must allow at least one valid value between min and max")
	}
	param.Step = step

	return param, nil
}

func parseIntParameter(parameter CreateTemplateRequestParameter) (dbparameter.IntParameter, error) {
	param := dbparameter.IntParameter{
		Name:        parameter.ParameterName,
		DisplayName: parameter.ParameterDisplayName,
	}

	defaultValue, err := strconv.Atoi(parameter.ParameterDefaultValue)
	if err != nil {
		return dbparameter.IntParameter{}, fmt.Errorf("invalid default value: %w", err)
	}
	param.DefaultValue = defaultValue

	idx := slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMinValue)
	})
	if idx == -1 {
		return dbparameter.IntParameter{}, errors.New("parameter must have a minimum value constraint")
	}
	minValue, err := strconv.Atoi(parameter.ParameterConstraints[idx].Value)
	if err != nil {
		return dbparameter.IntParameter{}, fmt.Errorf("invalid minimum value: %w", err)
	}
	if minValue > defaultValue {
		return dbparameter.IntParameter{}, errors.New("minimum value must not be greater than the default value")
	}
	param.MinValue = minValue

	idx = slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMaxValue)
	})
	if idx == -1 {
		return dbparameter.IntParameter{}, errors.New("parameter must have a maximum value constraint")
	}
	maxValue, err := strconv.Atoi(parameter.ParameterConstraints[idx].Value)
	if err != nil {
		return dbparameter.IntParameter{}, fmt.Errorf("invalid maximum value for parameter: %w", err)
	}
	if maxValue < defaultValue {
		return dbparameter.IntParameter{}, errors.New("maximum value must not be less than the default value")
	}
	param.MaxValue = maxValue

	if minValue >= maxValue {
		return dbparameter.IntParameter{}, errors.New("minimum value must be less than maximum value")
	}

	return param, nil
}

func parseStringParameter(parameter CreateTemplateRequestParameter) (dbparameter.StringParameter, error) {
	param := dbparameter.StringParameter{
		Name:        parameter.ParameterName,
		DisplayName: parameter.ParameterDisplayName,
	}

	param.DefaultValue = parameter.ParameterDefaultValue

	idx := slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMinLength)
	})
	if idx == -1 {
		return dbparameter.StringParameter{}, errors.New("parameter must have a minimum length constraint")
	}
	minLength, err := strconv.Atoi(parameter.ParameterConstraints[idx].Value)
	if err != nil {
		return dbparameter.StringParameter{}, fmt.Errorf("invalid minimum length: %w", err)
	}
	if minLength > len(param.DefaultValue) {
		return dbparameter.StringParameter{}, errors.New("minimum length must not be greater than the default value length")
	}
	param.MinLength = minLength

	idx = slices.IndexFunc(parameter.ParameterConstraints, func(c CreateTemplateRequestParameterConstraint) bool {
		return c.Type == string(dbparameter.ParameterConstraintMaxLength)
	})
	if idx == -1 {
		return dbparameter.StringParameter{}, errors.New("parameter must have a maximum length constraint")
	}
	maxLength, err := strconv.Atoi(parameter.ParameterConstraints[idx].Value)
	if err != nil {
		return dbparameter.StringParameter{}, fmt.Errorf("invalid maximum length for parameter: %w", err)
	}
	if maxLength < len(param.DefaultValue) {
		return dbparameter.StringParameter{}, errors.New("maximum length must not be less than the default value length")
	}
	param.MaxLength = maxLength

	if minLength >= maxLength {
		return dbparameter.StringParameter{}, errors.New("minimum length must be less than maximum length")
	}

	return param, nil
}

func parseBoolParameter(parameter CreateTemplateRequestParameter) (dbparameter.BoolParameter, error) {
	param := dbparameter.BoolParameter{
		Name:        parameter.ParameterName,
		DisplayName: parameter.ParameterDisplayName,
	}

	var err error
	param.DefaultValue, err = strconv.ParseBool(parameter.ParameterDefaultValue)
	if err != nil {
		return dbparameter.BoolParameter{}, fmt.Errorf("invalid default value: %w", err)
	}

	return param, nil
}
