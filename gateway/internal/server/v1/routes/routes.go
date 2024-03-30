package routes

import (
	"fmt"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/X3ne/ds_ms/gateway/internal/server/v1/handlers"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
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

	channels.PUT("/:channel.id/permission/:overwrite.id", channelsHandler.EditChannelPermissions)
	channels.DELETE("/:channel.id/permission/:overwrite.id", channelsHandler.DeleteChannelPermission)

	channels.POST("/:channel.id/typing", channelsHandler.TriggerTypingIndicator)

	channels.GET("/:channel.id/pins", channelsHandler.GetPinnedMessages)
	channels.PUT("/:channel.id/pins/:message.id", channelsHandler.AddPinnedMessage)
	channels.DELETE("/:channel.id/pins/:message.id", channelsHandler.DeletePinnedMessage)

	channels.PUT("/:channel.id/recipients", channelsHandler.GroupDMAddRecipient)
	channels.DELETE("/:channel.id/recipients/:user.id", channelsHandler.GroupDMRemoveRecipient)

	// v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	v1.GET("/swagger/docs/*", echoSwagger.WrapHandler)

	doc := redoc.Redoc{
		Title:       "V1 Api doc",
		Description: "1.0.0",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/v1/docs",
	}

	server.Echo.Use(echoredoc.New(doc))

	fmt.Println("V1 routes configured")
}
