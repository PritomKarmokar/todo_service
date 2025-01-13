package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"todo_service/cmd/config"
	"todo_service/cmd/routes"
)

func main() {
	fmt.Println("Hello the first todo service")
	e := echo.New()

	config.LoadEnv()
	config.LoggerConfig()
	routes.RegisterRoutes(e)
	config.StartServer(e)
}
