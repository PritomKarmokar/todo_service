package routes

import (
	"github.com/labstack/echo/v4"
	"todo_service/cmd/handlers"
)

var RegisterRoutes = func(e *echo.Echo) {
	basePrefix := e.Group("")

	basePrefix.GET("/", handlers.HelloWorld)
}
