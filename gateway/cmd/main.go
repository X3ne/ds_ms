package main

import (
	"fmt"
	app "github.com/X3ne/ds_ms/gateway"
	"github.com/X3ne/ds_ms/gateway/config"
	"github.com/X3ne/ds_ms/gateway/docs"
)

// @title Torrents API
// @version 1.0.0
// @description v1.0.0 Torrents API

// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	app.Start(cfg)
}
