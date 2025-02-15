package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/google/uuid"
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
	return nil, nil
}

func (db *SQLDB) GetTemplateMetaByUUID(templateUUID uuid.UUID) (*database.TemplateMeta, error) {
	return nil, nil
}

func (db *SQLDB) DeleteTemplate(templateUUID uuid.UUID) error {
	return nil
}

func (db *SQLDB) SearchTemplates(searchParameters database.SearchParameters) ([]database.Template, error) {
	return nil, nil
}

func (db *SQLDB) SetTemplateMarked(templateUUID uuid.UUID, marked bool) error {
	return nil
}
