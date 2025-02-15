package sqldb

import (
	"fmt"
	"strconv"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

type templateParameter struct {
	UUID         uuid.UUID `db:"uuid"`
	Name         string    `db:"name"`
	Type         string    `db:"type"`
	DisplayName  string    `db:"display_name"`
	DefaultValue string    `db:"default_value"`
}

type templateParameterConstraint struct {
	UUID            uuid.UUID `db:"uuid"`
	ConstraintType  string    `db:"constraint_type_name"`
	ConstraintValue string    `db:"constraint_value"`
}

func (db *SQLDB) getTemplateParameters(templateUUID uuid.UUID) ([]dbparameter.Parameter, error) {
	var dbParameters []templateParameter
	parametersQuery := `
	SELECT uuid, name, type, display_name, default_value
	FROM template_parameters
	WHERE template_uuid = $1
	`
	err := db.db.Select(&dbParameters, parametersQuery, templateUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get parameters: %w", err)
	}
	var parameters []dbparameter.Parameter
	for _, dbParameter := range dbParameters {
		var parameter dbparameter.Parameter
		switch dbParameter.Type {
		case "string":
			minLength, maxLength, err := db.getParameterConstraintsString(dbParameter.UUID)
			if err != nil {
				return nil, fmt.Errorf("failed to get constraints for parameter %s: %w", dbParameter.UUID, err)
			}
			parameter = dbparameter.StringParameter{
				Name:         dbParameter.Name,
				DisplayName:  dbParameter.DisplayName,
				DefaultValue: dbParameter.DefaultValue,
				MinLength:    minLength,
				MaxLength:    maxLength,
			}
		case "int":
			value, err := strconv.Atoi(dbParameter.DefaultValue)
			if err != nil {
				return nil, fmt.Errorf("failed to parse default value %s as int: %w", dbParameter.DefaultValue, err)
			}
			minValue, maxValue, err := db.getParameterConstraintsInt(dbParameter.UUID)
			if err != nil {
				return nil, fmt.Errorf("failed to get constraints for parameter %s: %w", dbParameter.UUID, err)
			}
			parameter = dbparameter.IntParameter{
				Name:         dbParameter.Name,
				DisplayName:  dbParameter.DisplayName,
				DefaultValue: value,
				MinValue:     minValue,
				MaxValue:     maxValue,
			}
		case "float":
			value, err := strconv.ParseFloat(dbParameter.DefaultValue, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse default value %s as float: %w", dbParameter.DefaultValue, err)
			}
			minValue, maxValue, step, err := db.getParameterConstraintsFloat(dbParameter.UUID)
			if err != nil {
				return nil, fmt.Errorf("failed to get constraints for parameter %s: %w", dbParameter.UUID, err)
			}
			parameter = dbparameter.FloatParameter{
				Name:         dbParameter.Name,
				DisplayName:  dbParameter.DisplayName,
				DefaultValue: value,
				MinValue:     minValue,
				MaxValue:     maxValue,
				Step:         step,
			}
		case "bool":
			value, err := strconv.ParseBool(dbParameter.DefaultValue)
			if err != nil {
				return nil, fmt.Errorf("failed to parse default value %s as bool: %w", dbParameter.DefaultValue, err)
			}
			parameter = dbparameter.BoolParameter{
				Name:         dbParameter.Name,
				DisplayName:  dbParameter.DisplayName,
				DefaultValue: value,
			}
		}

		if parameter == nil {
			return nil, fmt.Errorf("unknown parameter type %s: %s", dbParameter.UUID, dbParameter.Type)
		}
		parameters = append(parameters, parameter)
	}
	return parameters, nil
}

func (db *SQLDB) getParameterConstraints(parameterUUID uuid.UUID) ([]templateParameterConstraint, error) {
	var constraints []templateParameterConstraint
	constraintsQuery := `
	SELECT c.uuid, ct.constraint_type_name, c.constraint_value
	FROM template_parameters_constraints c
		JOIN parameter_constraint_types ct ON c.constraint_type_id = ct.constraint_type_id
	WHERE c.template_parameter_uuid = $1
	`
	err := db.db.Select(&constraints, constraintsQuery, parameterUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get constraints: %w", err)
	}
	return constraints, nil
}

func (db *SQLDB) getParameterConstraintsString(parameterUUID uuid.UUID) (int, int, error) {
	constraints, err := db.getParameterConstraints(parameterUUID)
	if err != nil {
		return 0, 0, err
	}

	var minLength, maxLength int
	for _, constraint := range constraints {
		switch constraint.ConstraintType {
		case "min_length":
			minLength, err = strconv.Atoi(constraint.ConstraintValue)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse %s min_length %s as int: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		case "max_length":
			maxLength, err = strconv.Atoi(constraint.ConstraintValue)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse %s max_length %s as int: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		}
	}
	return minLength, maxLength, nil
}

func (db *SQLDB) getParameterConstraintsInt(parameterUUID uuid.UUID) (int, int, error) {
	constraints, err := db.getParameterConstraints(parameterUUID)
	if err != nil {
		return 0, 0, err
	}

	var minValue, maxValue int
	for _, constraint := range constraints {
		switch constraint.ConstraintType {
		case "min_value":
			minValue, err = strconv.Atoi(constraint.ConstraintValue)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse %s min_value %s as int: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		case "max_value":
			maxValue, err = strconv.Atoi(constraint.ConstraintValue)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse %s max_value %s as int: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		}
	}
	return minValue, maxValue, nil
}

func (db *SQLDB) getParameterConstraintsFloat(parameterUUID uuid.UUID) (float64, float64, float64, error) {
	constraints, err := db.getParameterConstraints(parameterUUID)
	if err != nil {
		return 0, 0, 0, err
	}

	var minValue, maxValue, step float64
	for _, constraint := range constraints {
		switch constraint.ConstraintType {
		case "min_value":
			minValue, err = strconv.ParseFloat(constraint.ConstraintValue, 64)
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to parse %s min_value %s as float: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		case "max_value":
			maxValue, err = strconv.ParseFloat(constraint.ConstraintValue, 64)
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to parse %s max_value %s as float: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		case "step":
			step, err = strconv.ParseFloat(constraint.ConstraintValue, 64)
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to parse %s step %s as float: %w", constraint.UUID, constraint.ConstraintValue, err)
			}
		}
	}
	return minValue, maxValue, step, nil
}
