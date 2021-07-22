package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type Book struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Author string `json:"author" form:"author"`
}
