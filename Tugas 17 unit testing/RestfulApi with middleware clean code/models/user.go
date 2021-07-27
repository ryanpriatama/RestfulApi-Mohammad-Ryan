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

type BookModel interface {
	Get() []Book
	Insert(Book) error
}

type BookModelPost interface {
	Post() []Book
	Insert(Book) error
}

type BookModelDelete interface {
	Delete() []Book
	Insert(Book) error
}

type BookModelUpdate interface {
	Update() []Book
	Insert(Book) error
}

type MockBookModel struct {
	books []Book
}

func (m *MockBookModel) Get() []Book {
	return m.books
}

func (m *MockBookModel) Post() []Book {
	return m.books
}

func (m *MockBookModel) Delete() []Book {
	return m.books[1:]
}

func (m *MockBookModel) Update() []Book {
	return m.books[1:]
}

func (m *MockBookModel) Insert(book Book) error {
	m.books = append(m.books, book)
	return nil
}

func NewMockBookModel() *MockBookModel {
	return &MockBookModel{
		books: []Book{},
	}
}

type UserResponse struct {
	Message string
	Data    []User
}
