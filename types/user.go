package types

import (
	"encoding/json"
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
