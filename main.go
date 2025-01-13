package main

import (
	"github.com/labstack/echo/v4"
	"todo_service/cmd/config"
	"todo_service/cmd/routes"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	config.LoggerConfig()
	config.ConnectDB()
	routes.RegisterRoutes(e)
	config.StartServer(e)
}
