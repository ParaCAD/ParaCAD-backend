package sqldb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLDB struct {
	db *sqlx.DB
}

func New(host, user, password, dbName string) (*SQLDB, error) {
	dbURL := fmt.Sprintf("postgresql://%s@%s/%s?connect_timeout=10&password=%s&sslmode=disable", user, host, dbName, password)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return &SQLDB{
		db: db,
	}, nil
}

func (s *SQLDB) Init() {
	s.db.MustExec(schema)
}

func (s *SQLDB) Close() error {
	return s.db.Close()
}
