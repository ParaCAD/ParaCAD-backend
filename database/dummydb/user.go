package dummydb

import (
	"fmt"
	"time"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/google/uuid"
)

var dummyUserID uuid.UUID = uuid.Nil
var dummyUsername string = "Dummy User"
var dummyEmail string = "test@test.com"

func (db *DummyDB) getDummyUser() *database.User {
	return &database.User{
		UUID:     dummyUserID,
		Username: dummyUsername,
		Email:    dummyEmail,
		Role:     database.RoleUser,
		Deleted:  nil,
	}
}

func (db *DummyDB) GetUserByUUID(userID uuid.UUID) (*database.User, error) {
	if userID == dummyUserID {
		return db.getDummyUser(), nil
	}
	return nil, nil
}

func (db *DummyDB) GetUserByUsername(username string) (*database.User, error) {
	if username == dummyUsername {
		return db.getDummyUser(), nil
	}
	return nil, nil
}

func (db *DummyDB) GetUserByEmail(email string) (*database.User, error) {
	if email == dummyEmail {
		return db.getDummyUser(), nil
	}
	return nil, nil
}

func (db *DummyDB) DeleteUser(userID uuid.UUID) error {
	if userID == dummyUserID {
		return nil
	}
	return fmt.Errorf("user %v not found", userID)
}

func (db *DummyDB) SetUserLastLogin(userID uuid.UUID, lastLoginTime time.Time) error {
	if userID == dummyUserID {
		return nil
	}
	return fmt.Errorf("user %v not found", userID)
}
