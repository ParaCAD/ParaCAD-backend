package database

type Database interface {
	GetUserByUUID(UserID) (User, error)
	GetUserByUsername(string) (User, error)
	GetUserByEmail(string) (User, error)

	GetTemplateByUUID(TemplateID) (Template, error)
	SearchTemplates(SearchParameters) ([]Template, error)
}
