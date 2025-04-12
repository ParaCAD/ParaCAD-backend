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
	SELECT uuid, username, email, password, role, deleted, created, last_login
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

func (db *SQLDB) IsUsernameOrEmailUsed(username, email string) (bool, error) {
	query := `
	SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2
	`
	var count int
	err := db.db.Get(&count, query, username, email)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (db *SQLDB) CreateUser(user database.User) error {
	query := `
	INSERT INTO users (uuid, username, email, password, role, created)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := db.db.Exec(query, user.UUID, user.Username, user.Email, user.Password, user.Role, user.Created)
	if err != nil {
		return err
	}

	return nil
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
