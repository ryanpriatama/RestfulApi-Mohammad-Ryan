package controllers

import (
	"myapp/project/lib/database"
	"myapp/project/models"
	"net/http"

	"strconv"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUseridController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user, err := database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	userAdd, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    userAdd,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user, err1 := database.DeleteUser(id)
	if err1 != nil {
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
	user := database.GetUpdateUser(id)
	c.Bind(&user)
	userUpdate, err1 := database.UpdateUser(user)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    userUpdate,
	})
}

func LoginUsersControllers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login",
		"users":  users,
	})
}

func GetBookControllers(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   books,
	})
}

func GetBookidController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	book, err := database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	bookAdd, err := database.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    bookAdd,
	})
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	book, err1 := database.DeleteBook(id)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	book := database.GetUpdateBook(id)
	c.Bind(&book)
	bookUpdate, err1 := database.UpdateBook(book)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    bookUpdate,
	})
}
