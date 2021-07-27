package database

import (
	"myapp/project/config"
	"myapp/project/middlewares"
	"myapp/project/models"
)

//get all user info from database
func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//get 1 specified user info from database
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

//create new user info from database
func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//delete user info by id from database
func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if tx := config.DB.Delete(&user, "id=?", id).Error; tx != nil {
		return nil, tx
	}
	return user, nil
}

//update user info from database
func UpdateUser(user models.User) (interface{}, error) {
	if tx := config.DB.Save(&user).Error; tx != nil {
		return nil, tx
	}
	return user, nil
}

//get 1 specified user with User struct output
func GetUpdateUser(id int) models.User {
	var user models.User
	config.DB.Find(&user, "id=?", id)
	return user
}

//login user with matching data from database
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

//get all book info from database
func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

//get 1 specified book info by id from database
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

//create book info from database
func CreateBook(book models.Book) (interface{}, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

//delete book by id from database
func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if tx := config.DB.Delete(&book, "id=?", id).Error; tx != nil {
		return nil, tx
	}
	return book, nil
}

//update book from database
func UpdateBook(book models.Book) (interface{}, error) {
	if tx := config.DB.Save(&book).Error; tx != nil {
		return nil, tx
	}
	return book, nil
}

//get book info by id with Book struct output
func GetUpdateBook(id int) models.Book {
	var book models.Book
	config.DB.Find(&book, "id=?", id)
	return book
}
