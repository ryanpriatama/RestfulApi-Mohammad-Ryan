package routes

import (
	"myapp/project/constants"
	"myapp/project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//CRUD method on users without jwt
	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUseridController)
	e.POST("/login", controllers.LoginUsersControllers)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	//CRUD method on books without jwt
	e.GET("/books", controllers.GetBookControllers)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books/:id", controllers.GetBookidController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	//CRUD method on users with jwt
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUseridController)
	r.GET("/users", controllers.GetUserControllers)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
}
