package models

import (
	"github.com/google/uuid"
	"github.com/tnyidea/go-sample-userdata/types"
	"log"
	"testing"
)

const TestUserId = "cefe108b-bf45-4133-9aa8-560aa2cdd681"

func TestFindAllUsers(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	users, err := db.FindAllUsers()

	log.Println(users)
}

func TestCount(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	count, err := db.Count()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if count != 500 {
		log.Println("error: Expected count == 500")
		t.FailNow()
	}
}

func TestFindUserByUUID(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err := db.FindUserByUUID(TestUserId)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(&user)
}

func TestCreateUserWithUUID(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	uuidString := uuid.NewString()
	newUser := types.User{
		Id:          uuidString,
		FirstName:   "John",
		LastName:    "Smith",
		CompanyName: "Smith and Associates",
		Address:     "123 Main Street",
		City:        "Springfield",
		County:      "Springfield",
		State:       "DE",
		Zip:         "12345",
		Phone1:      "555-555-5555",
		Phone2:      "555-555-5556",
		Email:       "john.smith@email.com",
		Web:         "http://johnsmith.com",
	}

	createdUser, err := db.CreateUser(newUser)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !createdUser.IsEqual(newUser) {
		log.Println("error: createdUser != newUser")
		log.Println("createdUser: ", &createdUser)
		log.Println("newUser: ", &newUser)
		t.FailNow()
	}

	user, err := db.FindUserByUUID(uuidString)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if !user.IsEqual(newUser) {
		log.Println("error: user != newUser")
		t.FailNow()
	}

	log.Println(&user)
}

func TestUpdateUser(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	modifyUser, err := db.FindUserByUUID(TestUserId)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	modifyUser.FirstName = "Jane"
	updatedUser, err := db.UpdateUser(modifyUser)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if !updatedUser.IsEqual(modifyUser) {
		log.Println("error: updatedUser != modifyUser")
		t.FailNow()
	}

	user, err := db.FindUserByUUID(TestUserId)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if !user.IsEqual(modifyUser) {
		log.Println("error: Update failed. db.FindUserByUUID() != modifyUser")
	}

	log.Println(&user)
}

func TestDeleteAllUsers(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.DeleteAllUsers()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	count, err := db.Count()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if count != 0 {
		log.Println("error: Expected count == 0")
	}

	users, err := db.FindAllUsers()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if users != nil {
		log.Println("error: Expected db.FindAllUsers() to return nil slice")
		t.FailNow()
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err := db.FindUserByUUID(TestUserId)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	deletedUser, err := db.DeleteUserByUUID(TestUserId)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !deletedUser.IsEqual(user) {
		log.Println("error: deletedUser != user")
		t.FailNow()
	}

	user, err = db.FindUserByUUID(TestUserId)
	if err == nil {
		log.Println("error: db.DeleteUserByUUID() failed: Expected err result from db.FindUserByUUID()")
		t.FailNow()
	}
	if !user.IsEqual(types.User{}) {
		log.Println("error: user != types.User{}")
		t.FailNow()
	}
}
