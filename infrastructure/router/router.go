package router

import (
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Initializes the endpoints of the pokemon API.
func NewRouter(c controller.AppController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons/:id", func(context echo.Context) error { return c.Pokemon.ActionGetByID(context) })
	e.GET("/pokemons", func(context echo.Context) error { return c.Pokemon.ActionGetAll(context) })
	e.POST("/pokemons/:id", func(context echo.Context) error { return c.Pokemon.ActionPostByID(context) })

	return e
}
