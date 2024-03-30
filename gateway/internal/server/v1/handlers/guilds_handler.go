package handlers

import (
	"connectrpc.com/connect"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	guildsv1 "github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1"
	"github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1/guildsv1connect"
	"github.com/X3ne/ds_ms/gateway/internal/responses"
	s "github.com/X3ne/ds_ms/gateway/internal/server"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type GuildsHandler struct {
	server         *s.Server
	ChannelsClient channelsv1connect.ChannelsServiceClient
	GuildsClient   guildsv1connect.GuildsServiceClient
	Channel        channelsv1.Channel
}

type GuildChannel struct {
	Id            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Type          channelsv1.ChannelType `json:"type,omitempty"`
	GuildId       string                 `json:"guild_id,omitempty"`
	Position      int32                  `json:"position,omitempty"`
	Topic         string                 `json:"topic,omitempty"`
	UserLimit     int32                  `json:"user_limit,omitempty"`
	ParentId      string                 `json:"parent_id,omitempty"`
	Permissions   string                 `json:"permissions,omitempty"`
	LastMessageId string                 `json:"last_message_id,omitempty"`
	IsNsfw        bool                   `json:"is_nsfw,omitempty"`
	IsVoice       bool                   `json:"is_voice,omitempty"`
	CreatedAt     int64                  `json:"created_at,omitempty"`
	UpdatedAt     int64                  `json:"updated_at,omitempty"`
} // TODO maybe rewrite the channels service to differentiate between guild and dm channels

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
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, guild.Msg)
}

// GetGuild godoc
// @Summary Get a guild by id
// @Description Get a guild by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.GetByIdResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id} [get]
func (h *GuildsHandler) GetGuild(c echo.Context) error {
	guildId := c.Param("guild.id")

	guild, err := h.GuildsClient.GetById(c.Request().Context(), &connect.Request[guildsv1.GetByIdRequest]{
		Msg: &guildsv1.GetByIdRequest{
			Id: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, guild.Msg)
}

// ModifyGuild godoc
// @Summary Modify a guild by id
// @Description Modify a guild by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param guild body guildsv1.UpdateRequest true "Guild object"
// @Success 200 {object} guildsv1.UpdateResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id} [patch]
func (h *GuildsHandler) ModifyGuild(c echo.Context) error {
	guildId := c.Param("guild.id")
	guildRequest := new(guildsv1.UpdateRequest)

	if err := c.Bind(guildRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	guild, err := h.GuildsClient.Update(c.Request().Context(), &connect.Request[guildsv1.UpdateRequest]{
		Msg: &guildsv1.UpdateRequest{
			Id:          guildId,
			Name:        guildRequest.Name,
			Icon:        guildRequest.Icon,
			OwnerId:     guildRequest.OwnerId,
			Banner:      guildRequest.Banner,
			Description: guildRequest.Description,
			Splash:      guildRequest.Splash,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, guild.Msg)
}

// DeleteGuild godoc
// @Summary Delete a guild by id
// @Description Delete a guild by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.DeleteResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id} [delete]
func (h *GuildsHandler) DeleteGuild(c echo.Context) error {
	guildId := c.Param("guild.id")

	guild, err := h.GuildsClient.Delete(c.Request().Context(), &connect.Request[guildsv1.DeleteRequest]{
		Msg: &guildsv1.DeleteRequest{
			Id: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, guild.Msg)
}

// GetGuildChannels godoc
// @Summary Get a guild's channels by id
// @Description Get a guild's channels by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.GetGuildChannelsResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/channels [get]
func (h *GuildsHandler) GetGuildChannels(c echo.Context) error {
	guildId := c.Param("guild.id")

	channels, err := h.GuildsClient.GetGuildChannels(c.Request().Context(), &connect.Request[guildsv1.GetGuildChannelsRequest]{
		Msg: &guildsv1.GetGuildChannelsRequest{
			GuildId: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channels.Msg)
}

// CreateGuildChannel godoc
// @Summary Create a new guild channel
// @Description Create a new guild channel
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param channel body guildsv1.CreateGuildChannelRequest true "Channel object"
// @Success 200 {object} guildsv1.CreateGuildChannelResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/channels [post]
func (h *GuildsHandler) CreateGuildChannel(c echo.Context) error {
	guildId := c.Param("guild.id")
	channelRequest := new(guildsv1.CreateGuildChannelRequest)

	if err := c.Bind(channelRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	channel, err := h.GuildsClient.CreateGuildChannel(c.Request().Context(), &connect.Request[guildsv1.CreateGuildChannelRequest]{
		Msg: &guildsv1.CreateGuildChannelRequest{
			GuildId:   guildId,
			Name:      channelRequest.Name,
			Type:      channelRequest.Type,
			UserLimit: channelRequest.UserLimit,
			ParentId:  channelRequest.ParentId,
			Position:  channelRequest.Position,
			IsNsfw:    channelRequest.IsNsfw,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channel.Msg)
}

// ModifyGuildChannelPositions godoc
// @Summary Modify a guild's channel positions
// @Description Modify a guild's channel positions
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param channels body guildsv1.ModifyGuildChannelPositionsRequest true "Channel positions object"
// @Success 200 {object} guildsv1.ModifyGuildChannelPositionsResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/channels [patch]
func (h *GuildsHandler) ModifyGuildChannelPositions(c echo.Context) error {
	guildId := c.Param("guild.id")
	channelsRequest := new(guildsv1.ModifyGuildChannelPositionsRequest)

	if err := c.Bind(channelsRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	channels, err := h.GuildsClient.ModifyGuildChannelPositions(c.Request().Context(), &connect.Request[guildsv1.ModifyGuildChannelPositionsRequest]{
		Msg: &guildsv1.ModifyGuildChannelPositionsRequest{
			GuildId:   guildId,
			Positions: channelsRequest.Positions,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, channels.Msg)
}

// GetGuildMember godoc
// @Summary Get a guild member by id
// @Description Get a guild member by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Success 200 {object} guildsv1.GetGuildMemberResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/{user.id} [get]
func (h *GuildsHandler) GetGuildMember(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")

	member, err := h.GuildsClient.GetGuildMember(c.Request().Context(), &connect.Request[guildsv1.GetGuildMemberRequest]{
		Msg: &guildsv1.GetGuildMemberRequest{
			GuildId: guildId,
			UserId:  userId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, member.Msg)
}

// ListGuildMembers godoc
// @Summary List a guild's members
// @Description List a guild's members
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.ListGuildMembersResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members [get]
func (h *GuildsHandler) ListGuildMembers(c echo.Context) error {
	guildId := c.Param("guild.id")

	members, err := h.GuildsClient.ListGuildMembers(c.Request().Context(), &connect.Request[guildsv1.ListGuildMembersRequest]{
		Msg: &guildsv1.ListGuildMembersRequest{
			GuildId: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, members.Msg)
}

// SearchGuildMembers godoc
// @Summary Search a guild's members
// @Description Search a guild's members
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param query query string true "Query string to match username(s) and nickname(s) against"
// @Param limit query int false "Limit"
// @Success 200 {object} guildsv1.SearchGuildMembersResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/search [get]
func (h *GuildsHandler) SearchGuildMembers(c echo.Context) error {
	guildId := c.Param("guild.id")
	query := c.QueryParam("query")

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	members, err := h.GuildsClient.SearchGuildMembers(c.Request().Context(), &connect.Request[guildsv1.SearchGuildMembersRequest]{
		Msg: &guildsv1.SearchGuildMembersRequest{
			GuildId: guildId,
			Query:   query,
			Limit:   int32(limit),
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, members.Msg)
}

// AddGuildMember godoc
// @Summary Add a member to a guild
// @Description Add a member to a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Param member body guildsv1.AddGuildMemberRequest true "Member object"
// @Success 200 {object} guildsv1.AddGuildMemberResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/{user.id} [put]
func (h *GuildsHandler) AddGuildMember(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")
	memberRequest := new(guildsv1.AddGuildMemberRequest)

	if err := c.Bind(memberRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	member, err := h.GuildsClient.AddGuildMember(c.Request().Context(), &connect.Request[guildsv1.AddGuildMemberRequest]{
		Msg: &guildsv1.AddGuildMemberRequest{
			GuildId:     guildId,
			UserId:      userId,
			Roles:       memberRequest.Roles,
			AccessToken: memberRequest.AccessToken,
			Nick:        memberRequest.Nick,
			Mute:        memberRequest.Mute,
			Deaf:        memberRequest.Deaf,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, member.Msg)
}

// ModifyCurrentMember godoc
// @Summary Modify the current member of a guild
// @Description Modify the current member of a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param member body guildsv1.ModifyCurrentMemberRequest true "Member object"
// @Success 200 {object} guildsv1.ModifyCurrentMemberResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/@me [patch]
