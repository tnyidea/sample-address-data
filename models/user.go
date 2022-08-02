package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tnyidea/go-sample-userdata/types"
)

func (d *DB) Count() (int, error) {
	users, err := d.FindAllUsers()
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func (d *DB) CreateUser(user types.User) (types.User, error) {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	err := tx.Insert("user", user)
	if err != nil {
		return types.User{}, err
	}

	tx.Commit()

	return user, nil
}

func (d *DB) FindAllUsers() ([]types.User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	result, err := tx.Get("user", "id")
	if err != nil {
		return nil, err
	}

	var users []types.User
	for v := result.Next(); v != nil; v = result.Next() {
		user := v.(types.User)
		users = append(users, user)
	}
	tx.Commit()

	return users, nil
}

func (d *DB) FindUserByUUID(uuidString string) (types.User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	// id is unique, so we are sure to get only one result for uuidString
	v, err := tx.First("user", "id", uuidString)
	if err != nil {
		return types.User{}, err
	}
	tx.Commit()

	if v == nil {
		return types.User{}, ErrorNotFound
	}

	return v.(types.User), nil
}

func (d *DB) UpdateUser(v types.User) (types.User, error) {
	if v.Id == "" {
		return types.User{}, errors.New("invalid Id")
	}

	_, err := d.FindUserByUUID(v.Id)
	if err != nil {
		return types.User{}, err
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

func (d *DB) DeleteUserByUUID(uuidString string) (types.User, error) {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	user, err := d.FindUserByUUID(uuidString)
	if err != nil {
		return types.User{}, err
	}

	err = tx.Delete("user", types.User{
		Id: uuidString,
	})
	if err != nil {
		return types.User{}, err
	}
	tx.Commit()

	return user, err
}
