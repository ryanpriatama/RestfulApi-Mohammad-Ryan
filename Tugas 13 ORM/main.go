package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func GetUsersController(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	var user User
	if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes added data",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user User
	if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	if tx := DB.Delete(&user); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user User
	if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	c.Bind(&user)
	if tx := DB.Save(&user); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func main() {
	connectionString := "root:ryan81@tcp(localhost:3306)/latihan_orm?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{})
	fmt.Println(DB)

	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.Logger.Fatal(e.Start(":8000"))
}
