package routes

import (
	"fmt"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/X3ne/ds_ms/gateway/internal/server/v1/handlers"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureV1Routes(server *s.Server) {
	channelsClient := channelsv1connect.NewChannelsServiceClient(http.DefaultClient, "http://127.0.0.1:8082")

	channelsHandler := handlers.NewChannelsHandler(server, channelsClient)

	v1 := server.Echo.Group("/v1")

	channels := v1.Group("/channels")
	channels.GET("/:channel.id", channelsHandler.GetChannel)
	channels.PATCH("/:channel.id", channelsHandler.ModifyChannel)
	channels.DELETE("/:channel.id", channelsHandler.DeleteChannel)

	channels.GET("/:channel.id/messages", channelsHandler.GetChannelMessages)
	channels.GET("/:channel.id/messages/:message.id", channelsHandler.GetChannelMessage)
	channels.POST("/:channel.id/messages", channelsHandler.CreateMessage)
	channels.PATCH("/:channel.id/messages/:message.id", channelsHandler.EditMessage)
	channels.DELETE("/:channel.id/messages/:message.id", channelsHandler.DeleteMessage)
	channels.POST("/:channel.id/messages/bulk-delete", channelsHandler.BulkDeleteMessages)

	// v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	v1.GET("/docs/*", echoSwagger.WrapHandler)

	fmt.Println("V1 routes configured")
}
