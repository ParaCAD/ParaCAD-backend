package dummydb

import (
	"fmt"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
)

var dummyTemplateID uuid.UUID = uuid.Nil
var dummyTemplateTemplate string = `
cube([10,width,10],false);
`

func (db *DummyDB) getDummyTemplate() database.Template {
	return database.Template{
		UUID:        dummyTemplateID,
		OwnerUUID:   dummyUserID,
		Name:        "Test cube",
		Description: "Simple cube for testing",
		Preview:     nil,
		Template:    dummyTemplateTemplate,
		Parameters: []dbparameter.Parameter{
			dbparameter.IntParameter{
				Name:         "width",
				DisplayName:  "Width of the cube",
				DefaultValue: 20,
				MinValue:     10,
				MaxValue:     30,
			},
		},
	}
}

func (db *DummyDB) GetTemplateByUUID(templateID uuid.UUID) (database.Template, error) {
	if templateID == dummyTemplateID {
		return db.getDummyTemplate(), nil
	}
	return database.Template{}, fmt.Errorf("template %v not found", templateID)
}

func (db *DummyDB) SearchTemplates(searchParameters database.SearchParameters) ([]database.Template, error) {
	return []database.Template{db.getDummyTemplate()}, nil
}
