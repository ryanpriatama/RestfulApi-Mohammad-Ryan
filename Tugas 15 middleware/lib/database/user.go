package database

import (
	"myapp/project/config"
	"myapp/project/middlewares"
	"myapp/project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User
	var count int64
	if err1 := config.DB.Model(&user).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if tx := config.DB.Delete(&user, "id=?", id).Error; tx != nil {
		return nil, tx
	}
	return user, nil
}

func UpdateUser(user models.User) (interface{}, error) {
	if tx := config.DB.Save(&user).Error; tx != nil {
		return nil, tx
	}
	return user, nil
}

func GetUpdateUser(id int) models.User {
	var user models.User
	config.DB.Find(&user, "id=?", id)
	return user
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var book models.Book
	var count int64
	if err1 := config.DB.Model(&book).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if tx := config.DB.Delete(&book, "id=?", id).Error; tx != nil {
		return nil, tx
	}
	return book, nil
}

func UpdateBook(book models.Book) (interface{}, error) {
	if tx := config.DB.Save(&book).Error; tx != nil {
		return nil, tx
	}
	return book, nil
}

func GetUpdateBook(id int) models.Book {
	var book models.Book
	config.DB.Find(&book, "id=?", id)
	return book
}
