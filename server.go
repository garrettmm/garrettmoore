package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(LoggerConfig()))

	// Routes
	e.GET("/ping", pong)

	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!\n")
}

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}
}
