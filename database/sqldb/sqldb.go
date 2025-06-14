package sqldb

import (
	"fmt"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
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

func (db *SQLDB) Init() {
	db.db.MustExec(schema)
	db.initParameterConstraintTypes()
	db.deleteUsersAndTemplates()
	go db.cacheCleanerJob()
}

func (db *SQLDB) initParameterConstraintTypes() {
	_, err := db.db.Exec(`TRUNCATE TABLE parameter_constraint_types CASCADE`)
	if err != nil {
		panic(err)
	}
	for _, constraint := range dbparameter.ParameterConstraints {
		db.db.MustExec(`
		INSERT INTO parameter_constraint_types 
		(constraint_type_id, constraint_type_name) 
		VALUES 
		($1, $2);
		`, constraint.ID(), constraint.String())
	}
}

func (db *SQLDB) Close() error {
	return db.db.Close()
}

func (db *SQLDB) deleteUsersAndTemplates() {
	db.deleteAllUsers()
	db.deleteAllTemplates()
}

func (db *SQLDB) deleteAllUsers() {
	db.db.MustExec(`
	TRUNCATE TABLE users CASCADE
	`)
}

func (db *SQLDB) deleteAllTemplates() {
	db.db.MustExec(`
	TRUNCATE TABLE templates CASCADE
	`)
}
