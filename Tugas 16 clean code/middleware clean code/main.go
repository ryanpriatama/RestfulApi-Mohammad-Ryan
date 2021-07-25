package main

import (
	"fmt"
	"myapp/project/config"
	"myapp/project/middlewares"
	"myapp/project/routes"

	"github.com/labstack/echo"
)

func main() {
	//to use echo framework
	e := echo.New()

	//connect to database
	config.InitDB()

	//conect to port
	config.InitPort()

	//add log middleware to e
	middlewares.LogMiddlewares(e)
	routes.New(e)
	//connect to port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
