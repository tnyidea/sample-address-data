package models

import (
	"encoding/json"
	"github.com/hashicorp/go-memdb"
	"github.com/tnyidea/go-sample-userdata/types"
	"github.com/tnyidea/typeutils"
	"io/ioutil"
)

type DB struct {
	memDB *memdb.MemDB
}

func NewUserDatabase(filename ...string) (DB, error) {
	var filenameString string
	if filename != nil {
		filenameString = filename[0]
	}

	b, err := ioutil.ReadFile(typeutils.StringDefault(filenameString, "./us-500.json"))
	if err != nil {
		return DB{}, err
	}

	var users []types.User
	err = json.Unmarshal(b, &users)
	if err != nil {
		return DB{}, err
	}

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return DB{}, err
	}

	tx := db.Txn(true)
	for _, user := range users {
		err := tx.Insert("user", user)
		if err != nil {
			return DB{}, err
		}
	}
	tx.Commit()

	return DB{
		memDB: db,
	}, nil
}

func (d *DB) Close() error {
	// Really a no-op, but here for completeness
	return nil
}
