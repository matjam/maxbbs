package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Username    string
	Password    string
	Email       string
	AccessLevel int
}

type DBSession struct {
	db *gorm.DB
}

func Connect() (*DBSession, error) {
	var err error
	session := DBSession{}
	session.db, err = gorm.Open(sqlite.Open("maxbbs.db"), &gorm.Config{})
	if err != nil {

	}

	err = session.db.AutoMigrate(&User{})

	return &session, err
}
