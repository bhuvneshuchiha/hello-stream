package main

import (
	"github.com/bhuvneshuchiha/hello-stream/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", routes.RootCall)

	e.Logger.Fatal(e.Start(":8080"))
}
