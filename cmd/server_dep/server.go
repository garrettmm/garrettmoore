package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
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
	e.File("/", "index.html")

	// Routes
	e.GET("/", index)
	e.GET("/ping", pong)

	e.Logger.Fatal(e.Start(getEnv("PORT", ":1323")))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!\n")
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!\n")
}

func getEnv(key, fallback string) string {
	log.Println("Looking for environment variable: " + key)
	if value, ok := os.LookupEnv(key); ok {
		log.Println("Found: " + value)
		return ":" + value
	}
	log.Println("Not Found, using fallback")
	return fallback
}

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}
}
