package infrastructure

import (
	"emoshu_practice/interfaces/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

var Router *echo.Echo

// @title                      emoshu practice
// @version                    1.0
// @license.name               kenshin kawai
// @host                       localhost:3000
// @BasePath                   /
func init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	memberController := controller.NewMemberController(InitDB())

	e.GET("/api/member/:id", memberController.Show)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
