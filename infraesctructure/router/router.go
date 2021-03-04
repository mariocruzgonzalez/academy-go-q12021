package router

import (
	"academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/csvJokes", func(context echo.Context) error { return controller.GetAllCSVJokes(context) })

	e.GET("/isAlive", func(context echo.Context) error { return  context.JSON(http.StatusOK, "it is alive")})

	return e
}
