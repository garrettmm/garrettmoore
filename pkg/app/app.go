package app

import (
	"github.com/labstack/echo"

	"github.com/garrettmm/garrettmoore/pkg/middleware"
	"github.com/garrettmm/garrettmoore/pkg/routes"
)

func Init() {
	Server := echo.New()

	routes.Register(Server)
	middleware.Register(Server)

	Server.Logger.Fatal(Server.Start(":1323"))
}
