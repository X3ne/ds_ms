package gateway

import (
	"github.com/X3ne/ds_ms/gateway/config"
	"github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/X3ne/ds_ms/gateway/internal/server/v1/routes"
	"github.com/labstack/echo/v4"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	//app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins:     []string{"http://localhost:5500"},
	//	AllowCredentials: true,
	//}))

	app.Echo.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	routes.ConfigureV1Routes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
