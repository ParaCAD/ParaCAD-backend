package sqldb

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
	db.initParameterConstraintTypes(
		"min_value",
		"max_value",
		"step",
		"min_length",
		"max_length",
	)
	db.createTestData()
}

func (db *SQLDB) initParameterConstraintTypes(constraints ...string) {
	for i, constraint := range constraints {
		existingConstraint := ""
		err := db.db.Get(&existingConstraint, `SELECT constraint_type_name FROM parameter_constraint_types WHERE constraint_type_name = $1`, constraint)
		if err == nil {
			continue
		}
		db.db.MustExec(`
		INSERT INTO parameter_constraint_types (constraint_type_id, constraint_type_name) VALUES ($1, $2);
	`, i, constraint)
	}
}

func (db *SQLDB) Close() error {
	return db.db.Close()
}

// Test data

func (db *SQLDB) createTestData() {
	db.createTestUser()
	db.createTestTemplate()
}

func (db *SQLDB) createTestUser() {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	db.db.MustExec(`
		INSERT INTO users 
		(uuid, username, email, password, role, created)
		VALUES
		('00000000-0000-0000-0000-000000000000', 
		'Dummy User', 'test@test.com', $1, 'user', $2)
	`, password, time.Now())
}

func (db *SQLDB) createTestTemplate() {
	db.db.MustExec(`
		INSERT INTO templates
		(uuid, owner_uuid, name, description, preview, template)
		VALUES
		('00000000-0000-0000-0000-000000000000',
		'00000000-0000-0000-0000-000000000000',
		'Test cube', 'Simple cube for testing', NULL, 'template')
	`)
	db.db.MustExec(`
		INSERT INTO template_parameters
		(uuid, template_uuid, name, type, display_name, default_value)
		VALUES
		('00000000-0000-0000-0000-000000000000',
		'00000000-0000-0000-0000-000000000000', 
		'width', 'int', 'Width of the cube', '20')
	`)
	db.db.MustExec(`
		INSERT INTO template_parameters_constraints
		(uuid, template_parameter_uuid, constraint_type_id, constraint_value)
		VALUES
		('00000000-0000-0000-0000-000000000000',
		'00000000-0000-0000-0000-000000000000', 0, '10')
	`)
	db.db.MustExec(`
		INSERT INTO template_parameters_constraints
		(uuid, template_parameter_uuid, constraint_type_id, constraint_value)
		VALUES
		('00000000-0000-0000-0000-000000000001',
		'00000000-0000-0000-0000-000000000000', 1, '30')
	`)
}
