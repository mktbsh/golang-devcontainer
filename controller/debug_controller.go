package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mktbsh/golang-devcontainer/config"
)

func GetSqliteFile(c echo.Context) error {
	return c.File(config.Config.Database.Path)
}
