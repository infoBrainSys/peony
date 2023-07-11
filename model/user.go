package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UUID     string
	Email    string
	Password string
	Phone    int
}

type UserEmail struct {
	Email string
}
