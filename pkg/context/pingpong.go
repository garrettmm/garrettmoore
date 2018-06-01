package context

import (
	"net/http"
	"github.com/labstack/echo"
)

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!\n")
}

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Ping!\n")
}