package sqldb

import (
	"fmt"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/google/uuid"
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
	db.initParameterConstraintTypes()
	db.createTestData()
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

// Test data

func (db *SQLDB) createTestData() {
	db.createTestUser()
	db.createTestTemplate()
}

func (db *SQLDB) createTestUser() {
	db.db.MustExec(`
	TRUNCATE TABLE users CASCADE
	`)
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
	TRUNCATE TABLE templates CASCADE
	`)
	err := db.CreateTemplate(
		database.Template{
			UUID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			OwnerUUID:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Name:        "Test Template",
			Description: "This is a test template",
			Preview:     nil,
			Template:    exampleTemplateCube,
			Parameters: []dbparameter.Parameter{
				dbparameter.IntParameter{
					Name:         "width",
					DisplayName:  "Width of the cube",
					DefaultValue: 20,
					MinValue:     10,
					MaxValue:     30,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	time.Sleep(500 * time.Millisecond)

	err = db.CreateTemplate(
		database.Template{
			UUID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			OwnerUUID:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Name:        "Lorem Ipsum",
			Description: "Lorem Ipsum Dolor Sit",
			Preview:     nil,
			Template:    exampleTemplateCube,
			Parameters: []dbparameter.Parameter{
				dbparameter.IntParameter{
					Name:         "width",
					DisplayName:  "Width of the cube",
					DefaultValue: 30,
					MinValue:     10,
					MaxValue:     90,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	time.Sleep(500 * time.Millisecond)

	err = db.CreateTemplate(
		database.Template{
			UUID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			OwnerUUID:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Name:        "Box with sliding lid",
			Description: "Simple box with sliding lid. Lid is not attached to the box, allowing for easy access to the contents. Set parameters, generate box, check 'Generate lid', generate lid. All dimensions are in mm.",
			Preview:     nil,
			Template:    exampleTemplateBox,
			Parameters: []dbparameter.Parameter{
				dbparameter.IntParameter{
					Name:         "content_length",
					DisplayName:  "Content length",
					DefaultValue: 139,
					MinValue:     15,
					MaxValue:     200,
				},
				dbparameter.IntParameter{
					Name:         "content_width",
					DisplayName:  "Content width",
					DefaultValue: 70,
					MinValue:     30,
					MaxValue:     100,
				},
				dbparameter.IntParameter{
					Name:         "content_height",
					DisplayName:  "Content height",
					DefaultValue: 15,
					MinValue:     10,
					MaxValue:     100,
				},
				dbparameter.FloatParameter{
					Name:         "wall_thickness",
					DisplayName:  "Wall thickness",
					DefaultValue: 3,
					MinValue:     2,
					MaxValue:     10,
				},
				dbparameter.FloatParameter{
					Name:         "lid_thickness",
					DisplayName:  "Lid thickness",
					DefaultValue: 1.5,
					MinValue:     0.4,
					MaxValue:     5,
				},
				dbparameter.FloatParameter{
					Name:         "clearance",
					DisplayName:  "Clearance",
					DefaultValue: 0.2,
					MinValue:     0,
					MaxValue:     1,
				},
				dbparameter.BoolParameter{
					Name:         "lid",
					DisplayName:  "Generate lid",
					DefaultValue: false,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}
}
