package router

import (
	"net/http"

	"github.com/mktbsh/golang-devcontainer/config"
	"github.com/mktbsh/golang-devcontainer/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	if config.Config.IsDev() {
		e.Use(middleware.Logger())
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})

	users := e.Group("/users")
	{
		users.GET("", controller.GetUser)
		users.POST("", controller.CreateUser)
	}

	if config.Config.IsDev() {
		debug := e.Group("/_debug")
		{
			debug.GET("/sqlite3.db", controller.GetSqliteFile)
		}
	}

	return e
}
