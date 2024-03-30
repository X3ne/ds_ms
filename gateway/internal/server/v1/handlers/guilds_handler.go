package handlers

import (
	"connectrpc.com/connect"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	guildsv1 "github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1"
	"github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1/guildsv1connect"
	"github.com/X3ne/ds_ms/gateway/internal/responses"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GuildsHandler struct {
	server         *s.Server
	ChannelsClient channelsv1connect.ChannelsServiceClient
	GuildsClient   guildsv1connect.GuildsServiceClient
}

func NewGuildsHandler(server *s.Server, channelsClient channelsv1connect.ChannelsServiceClient, guildsClient guildsv1connect.GuildsServiceClient) *GuildsHandler {
	return &GuildsHandler{
		server:         server,
		ChannelsClient: channelsClient,
		GuildsClient:   guildsClient,
	}
}

// CreateGuild godoc
// @Summary Create a new guild
// @Description Create a new guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild body guildsv1.CreateRequest true "Guild object"
// @Success 200 {object} guildsv1.CreateResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds [post]
func (h *GuildsHandler) CreateGuild(c echo.Context) error {
	guildRequest := new(guildsv1.CreateRequest)

	if err := c.Bind(guildRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	guild, err := h.GuildsClient.Create(c.Request().Context(), &connect.Request[guildsv1.CreateRequest]{
		Msg: &guildsv1.CreateRequest{
			Name:    guildRequest.Name,
			Icon:    guildRequest.Icon,
			OwnerId: "1769332944884731904", // TODO: change this to the actual user id
		},
	})
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return responses.Response(c, http.StatusOK, guild.Msg)
}
