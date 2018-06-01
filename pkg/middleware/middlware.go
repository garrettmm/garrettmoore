package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Register(server *echo.Echo) {
	server.Use(middleware.Recover())
	server.Use(middleware.LoggerWithConfig(loggerConfig()))
}

func loggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}
}