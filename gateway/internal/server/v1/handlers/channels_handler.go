package handlers

import (
	"connectrpc.com/connect"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	"github.com/X3ne/ds_ms/gateway/internal/responses"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChannelsHandler struct {
	server         *s.Server
	ChannelsClient channelsv1connect.ChannelsServiceClient
}

func NewChannelsHandler(server *s.Server, channelsClient channelsv1connect.ChannelsServiceClient) *ChannelsHandler {
	return &ChannelsHandler{
		server:         server,
		ChannelsClient: channelsClient,
	}
}

// GetChannel godoc
// @Summary Get channel by ID
// @Description Get channel by ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Channel ID"
// @Success 200 {object} channelsv1.Channel
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{id} [get]
func (h *ChannelsHandler) GetChannel(c echo.Context) error {
	channelID := c.Param("id")

	channel, err := h.ChannelsClient.GetById(c.Request().Context(), &connect.Request[channelsv1.GetByIdRequest]{
		Msg: &channelsv1.GetByIdRequest{
			Id: channelID,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channel.Msg)
}

// ModifyChannel godoc
// @Summary Modify channel with given ID
// @Description Modify channel with given ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Channel ID"
// @Param request body channelsv1.UpdateRequest true "Channel data"
// @Success 200 {object} channelsv1.Channel
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{id} [patch]
func (h *ChannelsHandler) ModifyChannel(c echo.Context) error {
	modifyRequest := new(channelsv1.UpdateRequest)

	if err := c.Bind(modifyRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	channel, err := h.ChannelsClient.Update(c.Request().Context(), &connect.Request[channelsv1.UpdateRequest]{
		Msg: &channelsv1.UpdateRequest{
			Id:   c.Param("id"),
			Name: modifyRequest.Name,
			Icon: modifyRequest.Icon,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channel.Msg)
}
