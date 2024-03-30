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
// @Summary Create Guild
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
// @Summary Get Guild
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
// @Summary Modify Guild
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
// @Summary Delete Guild
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
// @Summary Get Guild Channels
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
// @Summary Create Guild Channel
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
// @Summary Modify Guild Channel Positions
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
// @Summary Get Guild Member
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
// @Summary List Guild Members
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
// @Summary Search Guild Members
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
// @Summary Add Guild Member
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
// @Summary Modify Current Member
// @Description Modify the current member of a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param member body guildsv1.ModifyGuildMemberRequest true "Member object"
// @Success 200 {object} guildsv1.ModifyGuildMemberResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/me [patch]
func (h *GuildsHandler) ModifyCurrentMember(c echo.Context) error {
	guildId := c.Param("guild.id")
	memberRequest := new(guildsv1.ModifyGuildMemberRequest)

	if err := c.Bind(memberRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// TODO: get the user id from the token

	member, err := h.GuildsClient.ModifyGuildMember(c.Request().Context(), &connect.Request[guildsv1.ModifyGuildMemberRequest]{
		Msg: &guildsv1.ModifyGuildMemberRequest{
			GuildId: guildId,
			UserId:  "1769332944884731904", // TODO: change this to the actual user id
			Nick:    memberRequest.Nick,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, member.Msg)
}

// AddGuildMemberRole godoc
// @Summary Add Guild Member Role
// @Description Add a role to a member of a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Param role.id path string true "Role ID"
// @Success 200 {object} guildsv1.AddGuildMemberRoleResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/{user.id}/roles/{role.id} [put]
func (h *GuildsHandler) AddGuildMemberRole(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")
	roleId := c.Param("role.id")

	_, err := h.GuildsClient.AddGuildMemberRole(c.Request().Context(), &connect.Request[guildsv1.AddGuildMemberRoleRequest]{
		Msg: &guildsv1.AddGuildMemberRoleRequest{
			GuildId: guildId,
			UserId:  userId,
			RoleId:  roleId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, nil)
}

// RemoveGuildMemberRole godoc
// @Summary Remove Guild Member Role
// @Description Remove a role from a member of a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Param role.id path string true "Role ID"
// @Success 200 {object} guildsv1.RemoveGuildMemberRoleResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/{user.id}/roles/{role.id} [delete]
func (h *GuildsHandler) RemoveGuildMemberRole(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")
	roleId := c.Param("role.id")

	_, err := h.GuildsClient.RemoveGuildMemberRole(c.Request().Context(), &connect.Request[guildsv1.RemoveGuildMemberRoleRequest]{
		Msg: &guildsv1.RemoveGuildMemberRoleRequest{
			GuildId: guildId,
			UserId:  userId,
			RoleId:  roleId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, nil)
}

// RemoveGuildMember godoc
// @Summary Remove Guild Member
// @Description Remove a member from a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Success 200 {object} guildsv1.RemoveGuildMemberResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/members/{user.id} [delete]
func (h *GuildsHandler) RemoveGuildMember(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")

	_, err := h.GuildsClient.RemoveGuildMember(c.Request().Context(), &connect.Request[guildsv1.RemoveGuildMemberRequest]{
		Msg: &guildsv1.RemoveGuildMemberRequest{
			GuildId: guildId,
			UserId:  userId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, nil)
}

// GetGuildBans godoc
// @Summary Get Guild Bans
// @Description Get users banned from a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.GetGuildBansResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/bans [get]
func (h *GuildsHandler) GetGuildBans(c echo.Context) error {
	guildId := c.Param("guild.id")

	bans, err := h.GuildsClient.GetGuildBans(c.Request().Context(), &connect.Request[guildsv1.GetGuildBansRequest]{
		Msg: &guildsv1.GetGuildBansRequest{
			GuildId: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, bans.Msg)
}

// GetGuildBan godoc
// @Summary Get Guild Ban
// @Description Get a user banned from a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Success 200 {object} guildsv1.GetGuildBanResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/bans/{user.id} [get]
func (h *GuildsHandler) GetGuildBan(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")

	ban, err := h.GuildsClient.GetGuildBan(c.Request().Context(), &connect.Request[guildsv1.GetGuildBanRequest]{
		Msg: &guildsv1.GetGuildBanRequest{
			GuildId: guildId,
			UserId:  userId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, ban.Msg)
}

// CreateGuildBan godoc
// @Summary Create Guild Ban
// @Description Ban a user from a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Param ban body guildsv1.CreateGuildBanRequest true "Ban object"
// @Success 200 {object} guildsv1.CreateGuildBanResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/bans/{user.id} [post]
func (h *GuildsHandler) CreateGuildBan(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")
	banRequest := new(guildsv1.CreateGuildBanRequest)

	if err := c.Bind(banRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	ban, err := h.GuildsClient.CreateGuildBan(c.Request().Context(), &connect.Request[guildsv1.CreateGuildBanRequest]{
		Msg: &guildsv1.CreateGuildBanRequest{
			GuildId: guildId,
			UserId:  userId,
			Reason:  banRequest.Reason,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, ban.Msg)
}

// RemoveGuildBan godoc
// @Summary Remove Guild Ban
// @Description Remove a user's ban from a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param user.id path string true "User ID"
// @Success 200 {object} guildsv1.RemoveGuildBanResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/bans/{user.id} [delete]
func (h *GuildsHandler) RemoveGuildBan(c echo.Context) error {
	guildId := c.Param("guild.id")
	userId := c.Param("user.id")

	_, err := h.GuildsClient.RemoveGuildBan(c.Request().Context(), &connect.Request[guildsv1.RemoveGuildBanRequest]{
		Msg: &guildsv1.RemoveGuildBanRequest{
			GuildId: guildId,
			UserId:  userId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, nil)
}

// GetGuildRoles godoc
// @Summary Get Guild Roles
// @Description Get a guild's roles
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Success 200 {object} guildsv1.GetGuildRolesResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/roles [get]
func (h *GuildsHandler) GetGuildRoles(c echo.Context) error {
	guildId := c.Param("guild.id")

	roles, err := h.GuildsClient.GetGuildRoles(c.Request().Context(), &connect.Request[guildsv1.GetGuildRolesRequest]{
		Msg: &guildsv1.GetGuildRolesRequest{
			GuildId: guildId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, roles.Msg)
}

// CreateGuildRole godoc
// @Summary Create Guild Role
// @Description Create a new role for a guild
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param role body guildsv1.CreateGuildRoleRequest true "Role object"
// @Success 200 {object} guildsv1.CreateGuildRoleResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/roles [post]
func (h *GuildsHandler) CreateGuildRole(c echo.Context) error {
	guildId := c.Param("guild.id")
	roleRequest := new(guildsv1.CreateGuildRoleRequest)

	if err := c.Bind(roleRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	role, err := h.GuildsClient.CreateGuildRole(c.Request().Context(), &connect.Request[guildsv1.CreateGuildRoleRequest]{
		Msg: &guildsv1.CreateGuildRoleRequest{
			GuildId:     guildId,
			Name:        roleRequest.Name,
			Color:       roleRequest.Color,
			Hoist:       roleRequest.Hoist,
			Permissions: roleRequest.Permissions,
			Mentionable: roleRequest.Mentionable,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, role.Msg)
}

// ModifyGuildRolePositions godoc
// @Summary Modify Guild Role Positions
// @Description Modify a guild's role positions
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param roles body guildsv1.ModifyGuildRolePositionsRequest true "Role positions object"
// @Success 200 {object} guildsv1.ModifyGuildRolePositionsResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/roles [patch]
func (h *GuildsHandler) ModifyGuildRolePositions(c echo.Context) error {
	guildId := c.Param("guild.id")
	rolesRequest := new(guildsv1.ModifyGuildRolePositionsRequest)

	if err := c.Bind(rolesRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	roles, err := h.GuildsClient.ModifyGuildRolePositions(c.Request().Context(), &connect.Request[guildsv1.ModifyGuildRolePositionsRequest]{
		Msg: &guildsv1.ModifyGuildRolePositionsRequest{
			GuildId:   guildId,
			Positions: rolesRequest.Positions,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, roles.Msg)
}

// ModifyGuildRole godoc
// @Summary Modify Guild Role
// @Description Modify a guild's role by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param role.id path string true "Role ID"
// @Param role body guildsv1.ModifyGuildRoleRequest true "Role object"
// @Success 200 {object} guildsv1.ModifyGuildRoleResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/roles/{role.id} [patch]
func (h *GuildsHandler) ModifyGuildRole(c echo.Context) error {
	guildId := c.Param("guild.id")
	roleId := c.Param("role.id")
	roleRequest := new(guildsv1.ModifyGuildRoleRequest)

	if err := c.Bind(roleRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	role, err := h.GuildsClient.ModifyGuildRole(c.Request().Context(), &connect.Request[guildsv1.ModifyGuildRoleRequest]{
		Msg: &guildsv1.ModifyGuildRoleRequest{
			GuildId:     guildId,
			RoleId:      roleId,
			Name:        roleRequest.Name,
			Color:       roleRequest.Color,
			Hoist:       roleRequest.Hoist,
			Permissions: roleRequest.Permissions,
			Mentionable: roleRequest.Mentionable,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, role.Msg)
}

// DeleteGuildRole godoc
// @Summary Delete Guild Role
// @Description Delete a guild's role by id
// @Tags Guilds
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param guild.id path string true "Guild ID"
// @Param role.id path string true "Role ID"
// @Success 200 {object} guildsv1.DeleteGuildRoleResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/guilds/{guild.id}/roles/{role.id} [delete]
func (h *GuildsHandler) DeleteGuildRole(c echo.Context) error {
	guildId := c.Param("guild.id")
	roleId := c.Param("role.id")

	_, err := h.GuildsClient.DeleteGuildRole(c.Request().Context(), &connect.Request[guildsv1.DeleteGuildRoleRequest]{
		Msg: &guildsv1.DeleteGuildRoleRequest{
			GuildId: guildId,
			RoleId:  roleId,
		},
	})
	if err != nil {
		return responses.ConnectErrorResponse(c, err)
	}

	return responses.Response(c, http.StatusOK, nil)
}
