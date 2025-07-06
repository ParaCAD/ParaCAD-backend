package sqldb

import (
	"fmt"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/google/uuid"
)

func (db *SQLDB) SearchTemplates(searchParameters database.SearchParameters) ([]database.SearchResult, error) {
	results := []database.SearchResult{}
	query := `
	SELECT t.uuid, t.name, t.preview, t.created, t.owner_uuid, u.username AS owner_name, u.deleted AS owner_deleted
	FROM templates t
		JOIN users u ON t.owner_uuid = u.uuid
	WHERE t.deleted IS NULL AND 
		(
			$1 = '' -- empty search query - return all
			OR 
			(t.name ILIKE '%%' || $1 || '%%') -- search by name
			OR
			($2 IS TRUE AND t.description ILIKE '%%' || $1 || '%%') -- search by description
		) 
	ORDER BY %s
	LIMIT $3 OFFSET $4;
	`

	orderByString := "t.created DESC"
	switch searchParameters.Sorting {
	case database.Newest:
		orderByString = "t.created DESC"
	case database.Oldest:
		orderByString = "t.created ASC"
	}

	query = fmt.Sprintf(query, orderByString)

	err := db.db.Select(&results, query,
		searchParameters.Query,
		searchParameters.SearchDescriptions,
		searchParameters.PageSize,
		searchParameters.PageNumber-1,
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (db *SQLDB) GetTemplatesByOwnerUUID(ownerUUID uuid.UUID, pageNumber int, pageSize int) ([]database.SearchResult, error) {
	results := []database.SearchResult{}

	query := `
	SELECT t.uuid, t.name, t.preview, t.created, t.owner_uuid, u.username AS owner_name
	FROM templates t
		JOIN users u ON t.owner_uuid = u.uuid
	WHERE t.deleted IS NULL AND t.owner_uuid = $1
	ORDER BY t.created DESC
	LIMIT $2 OFFSET $3;
	`
	err := db.db.Select(&results, query,
		ownerUUID,
		pageSize,
		pageNumber-1,
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}
