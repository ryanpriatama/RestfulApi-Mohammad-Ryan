package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

//get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

//get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user := users[id-1]
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success added data",
		"user":    user,
	})

}

//delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user1 := users[:id-1]
	user2 := users[id:]
	userDeleted := users[id-1]
	user1 = append(user1, user2...)
	users = user1
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleted data",
		"user":    userDeleted,
	})
}

//update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user, userSelected User
	user = User{}
	c.Bind(&user)
	userSelected = users[id-1]
	if user.Name != "" {
		userSelected.Name = user.Name
	}
	if user.Email != "" {
		userSelected.Email = user.Email
	}
	if user.Password != "" {
		userSelected.Password = user.Password
	}
	users[id-1] = userSelected
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    userSelected,
	})
}

//create new user
func CreateUserController(c echo.Context) error {
	//binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func main() {
	e := echo.New()
	//routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
