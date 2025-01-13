package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}
