package config

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartServer(e *echo.Echo) {
	go func() {
		if err := e.Start(":" + viper.GetString("PORT")); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err.Error() + " shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Gracefully shutting down the server for interrupt signal or deployment")
	if err := e.Shutdown(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
}
