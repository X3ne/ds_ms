package handlers

import (
	"connectrpc.com/connect"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	"github.com/X3ne/ds_ms/gateway/internal/responses"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
// @Summary Get Channel
// @Description Get channel by ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Success 200 {object} channelsv1.Channel
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id} [get]
func (h *ChannelsHandler) GetChannel(c echo.Context) error {
	channelID := c.Param("channel.id")

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
// @Summary Modify Channel
// @Description Modify channel by ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param request body channelsv1.UpdateRequest true "Channel data"
// @Success 200 {object} channelsv1.Channel
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id} [patch]
func (h *ChannelsHandler) ModifyChannel(c echo.Context) error {
	modifyRequest := new(channelsv1.UpdateRequest)

	if err := c.Bind(modifyRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	channel, err := h.ChannelsClient.Update(c.Request().Context(), &connect.Request[channelsv1.UpdateRequest]{
		Msg: &channelsv1.UpdateRequest{
			Id:   c.Param("channel.id"),
			Name: modifyRequest.Name,
			Icon: modifyRequest.Icon,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channel.Msg)
}

// DeleteChannel godoc
// @Summary Delete Channel
// @Description Delete the channel by ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Success 200 {object} channelsv1.DeleteResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id} [delete]
func (h *ChannelsHandler) DeleteChannel(c echo.Context) error {
	deleteResponse, err := h.ChannelsClient.Delete(c.Request().Context(), &connect.Request[channelsv1.DeleteRequest]{
		Msg: &channelsv1.DeleteRequest{
			Id: c.Param("channel.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, deleteResponse.Msg)
}

// GetChannelMessages godoc
// @Summary Get Channel Messages
// @Description Get messages for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param limit query int false "Limit the number of messages returned"
// @Param around query string false "Get messages around a specific message ID"
// @Param before query string false "Get messages before a specific message ID"
// @Param after query string false "Get messages after a specific message ID"
// @Success 200 {object} channelsv1.GetChannelMessagesResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages [get]
func (h *ChannelsHandler) GetChannelMessages(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 50
	}
	messagesResponse, err := h.ChannelsClient.GetChannelMessages(c.Request().Context(), &connect.Request[channelsv1.GetChannelMessagesRequest]{
		Msg: &channelsv1.GetChannelMessagesRequest{
			ChannelId: c.Param("channel.id"),
			Limit:     int32(limit),
			Around:    c.QueryParam("around"),
			Before:    c.QueryParam("before"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, messagesResponse.Msg)
}

// GetChannelMessage godoc
// @Summary Get message
// @Description Get message by id for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param message.id path string true "Message ID"
// @Success 200 {object} channelsv1.GetChannelMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages/{message.id} [get]
func (h *ChannelsHandler) GetChannelMessage(c echo.Context) error {
	messageResponse, err := h.ChannelsClient.GetChannelMessage(c.Request().Context(), &connect.Request[channelsv1.GetChannelMessageRequest]{
		Msg: &channelsv1.GetChannelMessageRequest{
			ChannelId: c.Param("channel.id"),
			MessageId: c.Param("message.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, messageResponse.Msg)
}

// CreateMessage godoc
// @Summary Create Message
// @Description Create a message for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param request body channelsv1.CreateMessageRequest true "Message data"
// @Success 200 {object} channelsv1.CreateMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages [post]
func (h *ChannelsHandler) CreateMessage(c echo.Context) error {
	createRequest := new(channelsv1.CreateMessageRequest)

	if err := c.Bind(createRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	message, err := h.ChannelsClient.CreateMessage(c.Request().Context(), &connect.Request[channelsv1.CreateMessageRequest]{
		Msg: &channelsv1.CreateMessageRequest{
			ChannelId: c.Param("channel.id"),
			AuthorId:  "1769332944884731904", // TODO: change this to the actual author ID
			Content:   createRequest.Content,
			Embeds:    createRequest.Embeds,
			Nonce:     createRequest.Nonce,
			Type:      channelsv1.MessageType_DEFAULT,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, message.Msg)
}

// EditMessage godoc
// @Summary Edit Message
// @Description Edit a message for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param message.id path string true "Message ID"
// @Param request body channelsv1.UpdateMessageRequest true "Message data"
// @Success 200 {object} channelsv1.UpdateMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages/{message.id} [patch]
func (h *ChannelsHandler) EditMessage(c echo.Context) error {
	editRequest := new(channelsv1.UpdateMessageRequest)

	if err := c.Bind(editRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	message, err := h.ChannelsClient.UpdateMessage(c.Request().Context(), &connect.Request[channelsv1.UpdateMessageRequest]{
		Msg: &channelsv1.UpdateMessageRequest{
			ChannelId: c.Param("channel.id"),
			MessageId: c.Param("message.id"),
			Content:   editRequest.Content,
			Embeds:    editRequest.Embeds,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, message.Msg)
}

// DeleteMessage godoc
// @Summary Delete Message
// @Description Delete a message for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param message.id path string true "Message ID"
// @Success 200 {object} channelsv1.DeleteMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages/{message.id} [delete]
func (h *ChannelsHandler) DeleteMessage(c echo.Context) error {
	deleteResponse, err := h.ChannelsClient.DeleteMessage(c.Request().Context(), &connect.Request[channelsv1.DeleteMessageRequest]{
		Msg: &channelsv1.DeleteMessageRequest{
			ChannelId: c.Param("channel.id"),
			MessageId: c.Param("message.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, deleteResponse.Msg)
}

// BulkDeleteMessages godoc
// @Summary Bulk Delete sMessages
// @Description Bulk delete messages for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param request body channelsv1.BulkDeleteMessagesRequest true "Message IDs array"
// @Success 200 {object} channelsv1.BulkDeleteMessagesResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/messages/bulk-delete [post]
func (h *ChannelsHandler) BulkDeleteMessages(c echo.Context) error {
	bulkDeleteRequest := new(channelsv1.BulkDeleteMessagesRequest)

	if err := c.Bind(bulkDeleteRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	deleteResponse, err := h.ChannelsClient.BulkDeleteMessages(c.Request().Context(), &connect.Request[channelsv1.BulkDeleteMessagesRequest]{
		Msg: &channelsv1.BulkDeleteMessagesRequest{
			ChannelId: c.Param("channel.id"),
			Messages:  bulkDeleteRequest.Messages,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, deleteResponse.Msg)
}

// EditChannelPermissions godoc
// @Summary Edit Channel Permissions
// @Description Edit channel permissions for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param overwrite.id path string true "ID of a user or role to overwrite permissions for"
// @Param request body channelsv1.EditChannelPermissionsRequest true "Permissions data"
// @Success 200 {object} channelsv1.EditChannelPermissionsResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/permissions/{overwrite.id} [put]
func (h *ChannelsHandler) EditChannelPermissions(c echo.Context) error {
	permissionsRequest := new(channelsv1.EditChannelPermissionsRequest)

	if err := c.Bind(permissionsRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	permissionsResponse, err := h.ChannelsClient.EditChannelPermissions(c.Request().Context(), &connect.Request[channelsv1.EditChannelPermissionsRequest]{
		Msg: &channelsv1.EditChannelPermissionsRequest{
			ChannelId:   c.Param("channel.id"),
			OverwriteId: c.Param("overwrite.id"),
			Allow:       permissionsRequest.Allow,
			Deny:        permissionsRequest.Deny,
			Type:        permissionsRequest.Type,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, permissionsResponse.Msg)
}

// DeleteChannelPermission godoc
// @Summary Delete Channel Permission
// @Description Delete channel permissions for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param overwrite.id path string true "ID of a user or role to overwrite permissions for"
// @Success 200 {object} channelsv1.DeleteChannelPermissionResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/permissions/{overwrite.id} [delete]
func (h *ChannelsHandler) DeleteChannelPermission(c echo.Context) error {
	deleteResponse, err := h.ChannelsClient.DeleteChannelPermission(c.Request().Context(), &connect.Request[channelsv1.DeleteChannelPermissionRequest]{
		Msg: &channelsv1.DeleteChannelPermissionRequest{
			ChannelId:   c.Param("channel.id"),
			OverwriteId: c.Param("overwrite.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, deleteResponse.Msg)
}

// TriggerTypingIndicator godoc
// @Summary Trigger Typing Indicator
// @Description Trigger the typing indicator for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Success 200 {object} channelsv1.TriggerTypingIndicatorResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/typing [post]
func (h *ChannelsHandler) TriggerTypingIndicator(c echo.Context) error {
	typingResponse, err := h.ChannelsClient.TriggerTypingIndicator(c.Request().Context(), &connect.Request[channelsv1.TriggerTypingIndicatorRequest]{
		Msg: &channelsv1.TriggerTypingIndicatorRequest{
			ChannelId: c.Param("channel.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, typingResponse.Msg)
}

// GetPinnedMessages godoc
// @Summary Get Pinned Messages
// @Description Get pinned messages for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Success 200 {object} channelsv1.GetPinnedMessagesResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/pins [get]
func (h *ChannelsHandler) GetPinnedMessages(c echo.Context) error {
	pinnedResponse, err := h.ChannelsClient.GetPinnedMessages(c.Request().Context(), &connect.Request[channelsv1.GetPinnedMessagesRequest]{
		Msg: &channelsv1.GetPinnedMessagesRequest{
			ChannelId: c.Param("channel.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, pinnedResponse.Msg)
}

// AddPinnedMessage godoc
// @Summary Add Pinned Message
// @Description Add a pinned message for the channel associated with the given ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param message.id path string true "Message ID"
// @Success 200 {object} channelsv1.AddPinnedMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/pins/{message.id} [put]
func (h *ChannelsHandler) AddPinnedMessage(c echo.Context) error {
	pinnedResponse, err := h.ChannelsClient.AddPinnedMessage(c.Request().Context(), &connect.Request[channelsv1.AddPinnedMessageRequest]{
		Msg: &channelsv1.AddPinnedMessageRequest{
			ChannelId: c.Param("channel.id"),
			MessageId: c.Param("message.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, pinnedResponse.Msg)
}

// DeletePinnedMessage godoc
// @Summary Delete Pinned Message
// @Description Delete a pinned message for the given channel ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param message.id path string true "Message ID"
// @Success 200 {object} channelsv1.DeletePinnedMessageResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/pins/{message.id} [delete]
func (h *ChannelsHandler) DeletePinnedMessage(c echo.Context) error {
	pinnedResponse, err := h.ChannelsClient.DeletePinnedMessage(c.Request().Context(), &connect.Request[channelsv1.DeletePinnedMessageRequest]{
		Msg: &channelsv1.DeletePinnedMessageRequest{
			ChannelId: c.Param("channel.id"),
			MessageId: c.Param("message.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, pinnedResponse.Msg)
}

// GroupDMAddRecipient godoc
// @Summary Group DM Add Recipient
// @Description Add a recipient to the group DM associated with the given ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param user.id path string true "User ID"
// @Param request body channelsv1.GroupDMAddRecipientRequest true "Recipient data"
// @Success 200 {object} channelsv1.GroupDMAddRecipientResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/recipients/{user.id} [put]
func (h *ChannelsHandler) GroupDMAddRecipient(c echo.Context) error {
	addRecipientRequest := new(channelsv1.GroupDMAddRecipientRequest)

	if err := c.Bind(addRecipientRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	addRecipientResponse, err := h.ChannelsClient.GroupDMAddRecipient(c.Request().Context(), &connect.Request[channelsv1.GroupDMAddRecipientRequest]{
		Msg: &channelsv1.GroupDMAddRecipientRequest{
			ChannelId:   c.Param("channel.id"),
			UserId:      c.Param("user.id"),
			Nick:        addRecipientRequest.Nick,
			AccessToken: addRecipientRequest.AccessToken,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, addRecipientResponse.Msg)
}

// GroupDMRemoveRecipient godoc
// @Summary Group DM Remove Recipient
// @Description Remove a recipient from the group DM associated with the given ID
// @Tags Channels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param channel.id path string true "Channel ID"
// @Param user.id path string true "User ID"
// @Success 200 {object} channelsv1.GroupDMRemoveRecipientResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/channels/{channel.id}/recipients/{user.id} [delete]
func (h *ChannelsHandler) GroupDMRemoveRecipient(c echo.Context) error {
	removeRecipientResponse, err := h.ChannelsClient.GroupDMRemoveRecipient(c.Request().Context(), &connect.Request[channelsv1.GroupDMRemoveRecipientRequest]{
		Msg: &channelsv1.GroupDMRemoveRecipientRequest{
			ChannelId: c.Param("channel.id"),
			UserId:    c.Param("user.id"),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, removeRecipientResponse.Msg)
}
