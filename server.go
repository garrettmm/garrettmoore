package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(LoggerConfig()))

	// Static
	e.File("/favicon.ico", "public/images/favicon.ico")

	// Routes
	e.GET("/", index)
	e.GET("/ping", pong)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":1323"
	}

	e.Logger.Fatal(e.Start(port))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!\n")
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!\n")
}

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}
}
