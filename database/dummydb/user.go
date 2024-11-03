package dummydb

import (
	"fmt"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var dummyUserID uuid.UUID = uuid.Nil
var dummyUsername string = "Dummy User"
var dummyEmail string = "test@test.com"
var dummyPassword string = "password"

func (db *DummyDB) getDummyUser() database.User {
	return database.User{
		UUID:     dummyUserID,
		Username: dummyUsername,
		Email:    dummyEmail,
		Role:     database.RoleUser,
		Deleted:  false,
	}
}

func (db *DummyDB) GetUserByUUID(userID uuid.UUID) (database.User, error) {
	if userID == dummyUserID {
		return db.getDummyUser(), nil
	}
	return database.User{}, fmt.Errorf("user %v not found", userID)
}

func (db *DummyDB) GetUserByUsername(username string) (database.User, error) {
	if username == dummyUsername {
		return db.getDummyUser(), nil
	}
	return database.User{}, fmt.Errorf("user %v not found", username)
}

func (db *DummyDB) GetUserByEmail(email string) (database.User, error) {
	if email == dummyEmail {
		return db.getDummyUser(), nil
	}
	return database.User{}, fmt.Errorf("user %v not found", email)
}

func (db *DummyDB) GetUserSecurityByUsername(username string) (database.UserSecurity, error) {
	if username == dummyEmail {
		u := db.getDummyUser()
		password, _ := bcrypt.GenerateFromPassword([]byte(dummyPassword), bcrypt.DefaultCost)
		return database.UserSecurity{
			Username: u.Username,
			Email:    u.Email,
			Password: password,
			Role:     u.Role,
			Deleted:  u.Deleted,
		}, nil
	}
	return database.UserSecurity{}, fmt.Errorf("user %v not found", username)
}

func (db *DummyDB) GetUserSecurityByEmail(email string) (database.UserSecurity, error) {
	if email == dummyEmail {
		u := db.getDummyUser()
		password, _ := bcrypt.GenerateFromPassword([]byte(dummyPassword), bcrypt.DefaultCost)
		return database.UserSecurity{
			Username: u.Username,
			Email:    u.Email,
			Password: password,
			Role:     u.Role,
			Deleted:  u.Deleted,
		}, nil
	}
	return database.UserSecurity{}, fmt.Errorf("user %v not found", email)
}

func (db *DummyDB) SetUserLastLogin(userID uuid.UUID, lastLoginTime time.Time) error {
	if userID == dummyUserID {
		return nil
	}
	return fmt.Errorf("user %v not found", userID)
}
