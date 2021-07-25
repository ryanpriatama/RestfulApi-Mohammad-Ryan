package models

import (
	"gorm.io/gorm"
)

//User struct and gorm.Model for adding updated_at,created_at,and deleted_ad
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

//Book struct and gorm.Model for adding updated_at,created_at,and deleted_ad
type Book struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Author string `json:"author" form:"author"`
}
