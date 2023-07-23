package models

import "gorm/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

func MigrateUser() {
	db.Database.AutoMigrate(&User{})
}
