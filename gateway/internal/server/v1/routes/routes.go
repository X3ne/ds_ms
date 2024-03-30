package routes

import (
	"fmt"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	"github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1/guildsv1connect"
	"github.com/X3ne/ds_ms/api/gen/users_service/users/v1/usersv1connect"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/X3ne/ds_ms/gateway/internal/server/v1/handlers"
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

type Clients struct {
	ChannelsClient channelsv1connect.ChannelsServiceClient
	UsersClient    usersv1connect.UsersServiceClient
	GuildsClient   guildsv1connect.GuildsServiceClient
}

func configureChannelsRoutes(server *s.Server, group *echo.Group, clients *Clients) {
	channelsHandler := handlers.NewChannelsHandler(server, clients.ChannelsClient)

	channels := group.Group("/channels")

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

	fmt.Println("Channels routes configured")
}

func configureGuildsRoutes(server *s.Server, group *echo.Group, clients *Clients) {
	guildsHandler := handlers.NewGuildsHandler(server, clients.ChannelsClient, clients.GuildsClient)

	guilds := group.Group("/guilds")

	guilds.POST("", guildsHandler.CreateGuild)
	guilds.GET("/:guild.id", guildsHandler.GetGuild)
	guilds.PATCH("/:guild.id", guildsHandler.ModifyGuild)
	guilds.DELETE("/:guild.id", guildsHandler.DeleteGuild)

	guilds.GET("/:guild.id/channels", guildsHandler.GetGuildChannels)
	guilds.POST("/:guild.id/channels", guildsHandler.CreateGuildChannel)

	guilds.GET("/:guild.id/members/:user.id", guildsHandler.GetGuildMember)
	guilds.GET("/:guild.id/members", guildsHandler.ListGuildMembers)
	guilds.GET("/:guild.id/members/search", guildsHandler.SearchGuildMembers)
}

func ConfigureV1Routes(server *s.Server) {
	v1 := server.Echo.Group("/v1")

	clients := &Clients{
		ChannelsClient: channelsv1connect.NewChannelsServiceClient(http.DefaultClient, "http://127.0.0.1:8082"),
		UsersClient:    usersv1connect.NewUsersServiceClient(http.DefaultClient, "http://127.0.0.1:8081"),
		GuildsClient:   guildsv1connect.NewGuildsServiceClient(http.DefaultClient, "http://127.0.0.1:8083"),
	}

	configureChannelsRoutes(server, v1, clients)
	configureGuildsRoutes(server, v1, clients)

	// v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	doc := redoc.Redoc{
		Title:       "V1 Api docs",
		Description: "1.0.0",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/v1/docs",
	}

	v1.Static("/docs", "./docs")

	server.Echo.Use(echoredoc.New(doc))

	fmt.Println("V1 routes configured")
}
