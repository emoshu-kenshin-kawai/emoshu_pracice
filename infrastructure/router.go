package infrastructure

import (
	"emoshu_practice/interfaces/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "emoshu_practice/docs"
)

var Router *echo.Echo

// @title                      emoshu practice
// @version                    1.0
// @license.name               kenshin kawai
// @host                       localhost:3000
// @BasePath                   /
func InitRouter() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	dbHandler := InitDB()
	memberController := controller.NewMemberController(dbHandler)

	e.GET("/api/members", memberController.Index)
	e.GET("/api/members/:id", memberController.Show)
	e.POST("/api/members", memberController.Create)
	e.PUT("/api/members/:id", memberController.Update)
	e.DELETE("/api/members/:id", memberController.Delete)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
