package sqldb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (db *SQLDB) GetTemplateByUUID(templateUUID uuid.UUID) (*database.Template, error) {
	var template database.Template
	query := `
	SELECT uuid, owner_uuid, name, description, preview, template
	FROM templates
	WHERE uuid = $1
	`
	err := db.db.Get(&template, query, templateUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	template.Parameters, err = db.getTemplateParameters(templateUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get parameters: %w", err)
	}
	return &template, nil
}

func (db *SQLDB) GetTemplateWithOwnerByUUID(templateUUID uuid.UUID) (*database.TemplatePage, error) {
	var template database.TemplatePage
	query := `
	SELECT t.uuid, t.name, t.description, t.preview,
	 	t.owner_uuid, u.username AS owner_name
	FROM templates t	
		JOIN users u ON t.owner_uuid = u.uuid
	WHERE t.uuid = $1
	`
	err := db.db.Get(&template, query, templateUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	template.Parameters, err = db.getTemplateParameters(templateUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get parameters: %w", err)
	}
	return &template, nil
}

func (db *SQLDB) GetTemplateContentByUUID(templateUUID uuid.UUID) (*database.TemplateContent, error) {
	var template database.TemplateContent
	query := `
	SELECT uuid, name, template,
	FROM templates
	WHERE uuid = $1
	`
	err := db.db.Get(&template, query, templateUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	template.Parameters, err = db.getTemplateParameters(templateUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get parameters: %w", err)
	}
	return &template, nil
}

func (db *SQLDB) GetTemplateMetaByUUID(templateUUID uuid.UUID) (*database.TemplateMeta, error) {
	return nil, nil
}

func (db *SQLDB) CreateTemplate(template database.Template) error {
	tx, err := db.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
	INSERT INTO templates 
	(uuid, owner_uuid, name, description, preview, template, created)
	VALUES
	($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = tx.Exec(query,
		template.UUID,
		template.OwnerUUID,
		template.Name,
		template.Description,
		template.Preview,
		template.Template,
		time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert template: %w", err)
	}

	for _, parameter := range template.Parameters {
		dbParameter := templateParameter{
			UUID:         uuid.New(),
			Name:         parameter.GetName(),
			Type:         parameter.GetType().String(),
			DisplayName:  parameter.GetDisplayName(),
			DefaultValue: parameter.String(),
		}
		query = `
		INSERT INTO template_parameters
		(uuid, template_uuid, name, type, display_name, default_value)
		VALUES
		($1, $2, $3, $4, $5, $6)
		`
		_, err = tx.Exec(query,
			dbParameter.UUID,
			template.UUID,
			dbParameter.Name,
			dbParameter.Type,
			dbParameter.DisplayName,
			dbParameter.DefaultValue)
		if err != nil {
			return fmt.Errorf("failed to insert parameter: %w", err)
		}

		err = insertParameterConstraints(tx, parameter, dbParameter.UUID)
		if err != nil {
			return fmt.Errorf("failed to insert parameter constraint: %w", err)
		}

	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func insertParameterConstraints(tx *sqlx.Tx, parameter dbparameter.Parameter, parameterUUID uuid.UUID) error {
	query := `
	INSERT INTO template_parameters_constraints 
	(uuid, template_parameter_uuid, constraint_type_id, constraint_value)
	VALUES
	($1, $2, $3, $4)
	`
	switch parameter.GetType() {

	case dbparameter.ParameterTypeFloat:
		floatParameter, ok := parameter.(dbparameter.FloatParameter)
		if !ok {
			return fmt.Errorf("parameter is not a float parameter")
		}
		_, err := tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMinValue.ID(),
			floatParameter.MinValue)
		if err != nil {
			return fmt.Errorf("failed to insert min value constraint: %w", err)
		}
		_, err = tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMaxValue.ID(),
			floatParameter.MaxValue)
		if err != nil {
			return fmt.Errorf("failed to insert max value constraint: %w", err)
		}
		_, err = tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintStep.ID(),
			floatParameter.Step)
		if err != nil {
			return fmt.Errorf("failed to insert step constraint: %w", err)
		}
	case dbparameter.ParameterTypeInt:
		intParameter, ok := parameter.(dbparameter.IntParameter)
		if !ok {
			return fmt.Errorf("parameter is not an int parameter")
		}
		_, err := tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMinValue.ID(),
			intParameter.MinValue)
		if err != nil {
			return fmt.Errorf("failed to insert min value constraint: %w", err)
		}
		_, err = tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMaxValue.ID(),
			intParameter.MaxValue)
		if err != nil {
			return fmt.Errorf("failed to insert max value constraint: %w", err)
		}
	case dbparameter.ParameterTypeString:
		stringParameter, ok := parameter.(dbparameter.StringParameter)
		if !ok {
			return fmt.Errorf("parameter is not a string parameter")
		}
		_, err := tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMinLength.ID(),
			stringParameter.MinLength)
		if err != nil {
			return fmt.Errorf("failed to insert min length constraint: %w", err)
		}
		_, err = tx.Exec(query,
			uuid.New(),
			parameterUUID,
			dbparameter.ParameterConstraintMaxLength.ID(),
			stringParameter.MaxLength)
		if err != nil {
			return fmt.Errorf("failed to insert max length constraint: %w", err)
		}
	}
	return nil
}

func (db *SQLDB) DeleteTemplate(templateUUID uuid.UUID) error {
	return nil
}

func (db *SQLDB) SetTemplateMarked(templateUUID uuid.UUID, marked bool) error {
	return nil
}
