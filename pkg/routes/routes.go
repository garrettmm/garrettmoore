package routes

import (
	"github.com/labstack/echo"

	"github.com/garrettmm/garrettmoore/pkg/context"
)

func Register(server *echo.Echo) {
	// Ping Pong
	server.GET("/ping", context.Pong)
	server.GET("/pong", context.Ping)
}
