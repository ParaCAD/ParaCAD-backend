package database

type SearchParameters struct {
	Query              string
	SearchDescriptions bool
	Sorting            Sorting
	PageNumber         int
	PageSize           int
}

type Sorting string

const (
	Newest Sorting = "newest"
	Oldest Sorting = "oldest"
)

type SearchResult struct {
	UUID      string `db:"uuid"`
	Name      string `db:"name"`
	Preview   string `db:"preview"`
	OwnerUUID string `db:"owner_uuid"`
	OwnerName string `db:"owner_name"`
}

func ToSorting(s string) Sorting {
	switch s {
	case "newest":
		return Newest
	case "oldest":
		return Oldest
	default:
		return Newest
	}
}
