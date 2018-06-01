package app

import (
	"os"

	"github.com/labstack/echo"

	"github.com/garrettmm/garrettmoore/pkg/middleware"
	"github.com/garrettmm/garrettmoore/pkg/routes"
)

func Init() {
	Server := echo.New()

	routes.Register(Server)
	middleware.Register(Server)

	Server.Logger.Fatal(Server.Start(getEnv("PORT", ":1323")))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return ":" + value
	}
	return fallback
}