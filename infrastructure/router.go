package infrastructure

import (
	"emoshu_practice/interfaces/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Router *echo.Echo

func init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	memberController := controller.NewMemberController(InitDB())

	e.GET("/api/member/:id", memberController.Show)
	e.Logger.Fatal(e.Start(":3000"))
}
