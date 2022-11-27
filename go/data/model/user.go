package model

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"reflect"
)

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Phone1      string `json:"phone1"`
	Phone2      string `json:"phone2"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

func (v *User) String() string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func (v *User) IsEqual(user User) bool {
	return reflect.DeepEqual(*v, user)
}

func (d *DB) Count() (int, error) {
	users, err := d.FindAllUsers()
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func (d *DB) CreateUser(user User) (User, error) {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	err := tx.Insert("user", user)
	if err != nil {
		return User{}, err
	}

	tx.Commit()

	return user, nil
}

func (d *DB) FindAllUsers() ([]User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	result, err := tx.Get("user", "id")
	if err != nil {
		return nil, err
	}

	var users []User
	for v := result.Next(); v != nil; v = result.Next() {
		user := v.(User)
		users = append(users, user)
	}
	tx.Commit()

	return users, nil
}

func (d *DB) FindUserByUUID(uuidString string) (User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	// id is unique, so we are sure to get only one result for uuidString
	v, err := tx.First("user", "id", uuidString)
	if err != nil {
		return User{}, err
	}
	tx.Commit()

	if v == nil {
		return User{}, ErrorNotFound
	}

	return v.(User), nil
}

func (d *DB) UpdateUser(v User) (User, error) {
	if v.Id == "" {
		return User{}, errors.New("invalid Id")
	}

	_, err := d.FindUserByUUID(v.Id)
	if err != nil {
		return User{}, err
	}

	return d.CreateUser(v)
}

func (d *DB) DeleteAllUsers() error {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	_, err := tx.DeleteAll("user", "id")
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (d *DB) DeleteUserByUUID(uuidString string) (User, error) {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	user, err := d.FindUserByUUID(uuidString)
	if err != nil {
		return User{}, err
	}

	err = tx.Delete("user", User{
		Id: uuidString,
	})
	if err != nil {
		return User{}, err
	}
	tx.Commit()

	return user, err
}
