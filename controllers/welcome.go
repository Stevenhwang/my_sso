package controllers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, world")
}
