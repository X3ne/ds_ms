package main

import (
	"fmt"
	app "github.com/X3ne/ds_ms/gateway"
	"github.com/X3ne/ds_ms/gateway/config"
	"github.com/X3ne/ds_ms/gateway/docs"
)

// @title           GoCord Gateway
// @version         1.0
// @description     This is the gateway specification for the GoCord project.

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8079
// @BasePath  /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	app.Start(cfg)
}
