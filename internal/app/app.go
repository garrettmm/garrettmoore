package app

import (
	"github.com/labstack/echo"

	"github.com/garrettmm/garrettmoore/internal/middleware"
	"github.com/garrettmm/garrettmoore/internal/routes"
)

func Init() {
	Server := echo.New()

	routes.Register(Server)
	middleware.Register(Server)

	Server.Logger.Fatal(Server.Start(":1323"))
}
