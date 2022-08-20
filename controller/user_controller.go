package controller

import (
	"net/http"

	"github.com/mktbsh/golang-devcontainer/model"
	"github.com/mktbsh/golang-devcontainer/repository"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {

	r := repository.NewUserRepository()

	return c.JSON(http.StatusOK, r.FindAll())
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")

	user := model.NewUser(name)

	result := repository.NewUserRepository().Create(user)

	if result {
		return c.NoContent(http.StatusNoContent)
	}

	return c.NoContent(http.StatusInternalServerError)
}
