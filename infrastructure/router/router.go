package router

import (
	"academy_bootcamp/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/items/:id", func(context echo.Context) error { return c.Item.GetItemById(context) })
	// e.POST("/items", func(context echo.Context) error { return c.Item.CreateItem(context) })

	return e
}
