package main

import (
	"fmt"
	"myapp/project/config"
	"myapp/project/middlewares"
	"myapp/project/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	middlewares.LogMiddlewares(e)
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
