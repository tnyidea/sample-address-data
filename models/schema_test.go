package models

import (
	"log"
	"testing"
)

func TestNewUserDatabase(t *testing.T) {
	_, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}
