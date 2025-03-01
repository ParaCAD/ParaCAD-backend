package sqldb

import "github.com/ParaCAD/ParaCAD-backend/database"

func (db *SQLDB) SearchTemplates(searchParameters database.SearchParameters) ([]database.SearchResult, error) {
	results := []database.SearchResult{}

	query := `
	SELECT t.uuid, t.name, t.preview, t.owner_uuid, u.username AS owner_name
	FROM templates t
		JOIN users u ON t.owner_uuid = u.uuid
	WHERE 1=1 AND 
		(
			$1 = '' -- empty search query - return all
			OR 
			(t.name ILIKE '%' || $1 || '%') -- search by name
			OR
			($2 IS TRUE AND t.description ILIKE '%' || $1 || '%') -- search by description
		) 
	ORDER BY $3
	LIMIT $4 OFFSET $5;
	`

	orderByString := "t.created DESC"
	switch searchParameters.Sorting {
	case database.Newest:
		orderByString = "t.created DESC"
	case database.Oldest:
		orderByString = "t.created ASC"
	}

	err := db.db.Select(&results, query,
		searchParameters.Query,
		searchParameters.SearchDescriptions,
		orderByString,
		searchParameters.PageSize,
		searchParameters.PageNumber-1,
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}
