package sqldb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/google/uuid"
)

func (db *SQLDB) GetUserByUUID(uuid uuid.UUID) (*database.User, error) {
	var user database.User
	query := `
	SELECT uuid, username, email, role, deleted 
		FROM users WHERE uuid = $1
	`
	err := db.db.Get(&user, query, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (db *SQLDB) GetUserByUsername(username string) (*database.User, error) {
	var user database.User
	query := `
	SELECT uuid, username, email, role, deleted 
		FROM users WHERE username = $1
	`
	err := db.db.Get(&user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (db *SQLDB) DeleteUser(uuid uuid.UUID) error {
	query := `
	UPDATE users SET deleted = 1 WHERE uuid = $2
	`
	res, err := db.db.Exec(query, uuid)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user with uuid %s found", uuid)
	}
	return nil
}

func (db *SQLDB) GetUserSecurityByUsername(username string) (*database.UserSecurity, error) {
	var user database.UserSecurity
	query := `
	SELECT uuid, username, email, password, role, deleted, created, last_login 
		FROM users WHERE username = $1
	`
	err := db.db.Get(&user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (db *SQLDB) SetUserLastLogin(uuid uuid.UUID, loginTime time.Time) error {
	query := `
	UPDATE users SET last_login = $1 WHERE uuid = $2
	`
	res, err := db.db.Exec(query, loginTime, uuid)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user with uuid %s found", uuid)
	}
	return nil
}
