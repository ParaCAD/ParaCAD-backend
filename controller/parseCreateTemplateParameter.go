package controller

import (
	"fmt"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
)

func parseCreateTemplateParameter(parameter CreateTemplateRequestParameter) (dbparameter.Parameter, error) {
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
