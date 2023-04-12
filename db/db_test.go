package db

import (
	"database/sql"
	"gorm_study/entity"
	"testing"
	"time"
)

func TestGetDB(t *testing.T) {
	_, err := GetDB()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreate(t *testing.T) {
	email := "test@123.com"
	birthDay := time.Date(1992, 02, 24, 0, 0, 0, 0, time.UTC)
	user := &entity.User{
		Name:         "zhennlli",
		Email:        &email,
		Age:          30,
		Birthday:     &birthDay,
		MemberNumber: sql.NullString{},
		ActivedAt:    sql.NullTime{},
	}
	err := Create(user)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	println(user.ID)
}
