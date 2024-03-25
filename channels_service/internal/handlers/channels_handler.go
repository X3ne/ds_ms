package handlers

import (
	"context"
	"database/sql"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	guildsv1 "github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1"
	"github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1/guildsv1connect"
	usersv1 "github.com/X3ne/ds_ms/api/gen/users_service/users/v1"
	"github.com/X3ne/ds_ms/api/gen/users_service/users/v1/usersv1connect"
	apiErrors "github.com/X3ne/ds_ms/channels_service/internal/errors"
	"github.com/X3ne/ds_ms/channels_service/internal/models"
	"github.com/X3ne/ds_ms/channels_service/internal/repositories"
	"github.com/X3ne/ds_ms/channels_service/internal/validator"
	"regexp"

	"connectrpc.com/connect"
)

type ChannelsServer struct {
	Repository   *repositories.ChannelsRepository
	GuildsClient guildsv1connect.GuildsServiceClient
	UsersClient  usersv1connect.UsersServiceClient
}

func createChannelResponse(channel *models.Channel) (retChannel *channelsv1.Channel) {
	retChannel = &channelsv1.Channel{
		Id:            channel.ID,
		Name:          channel.Name,
		Type:          channel.Type,
		Icon:          channel.Icon.String,
		OwnerId:       channel.OwnerID,
		GuildId:       channel.GuildID,
		ParentId:      channel.ParentID,
		Position:      channel.Position,
		Topic:         channel.Topic.String,
		UserLimit:     channel.UserLimit,
		Recipients:    channel.Recipients,
		Permissions:   channel.Permissions,
		LastMessageId: channel.LastMessageID,
		IsNsfw:        channel.IsNSFW,
		IsVoice:       channel.IsVoice,
		CreatedAt:     channel.CreatedAt.Unix(),
		UpdatedAt:     channel.UpdatedAt.Unix(),
	}

	return
}

func createMessageResponse(message *models.Message) (retMessage *channelsv1.Message) {
	retMessage = &channelsv1.Message{
		Id:              message.ID,
		ChannelId:       message.ChannelID,
		Content:         message.Content,
		Type:            channelsv1.MessageType(message.Type),
		Nonce:           message.Nonce,
		Pinned:          message.Pinned,
		Timestamp:       message.CreatedAt.Unix(),
		EditedTimestamp: message.UpdatedAt.Time.Unix(),
	}

	return
}

func (s *ChannelsServer) Create(ctx context.Context, req *connect.Request[channelsv1.CreateRequest]) (*connect.Response[channelsv1.CreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	newChannel := &models.Channel{
		Name:      req.Msg.Name,
		Type:      req.Msg.Type,
		Position:  req.Msg.Position,
		UserLimit: req.Msg.UserLimit,
		IsNSFW:    req.Msg.IsNsfw,
		IsVoice:   req.Msg.IsVoice,
	}

	if req.Msg.GuildId != "" {
		if _, err := s.GuildsClient.GetById(ctx, &connect.Request[guildsv1.GetByIdRequest]{
			Msg: &guildsv1.GetByIdRequest{
				Id: req.Msg.GuildId,
			},
		}); err != nil {
			return nil, err
		}
		newChannel.GuildID = req.Msg.GuildId
	}

	if req.Msg.OwnerId != "" {
		if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
			Msg: &usersv1.GetByIdRequest{
				Id: req.Msg.OwnerId,
			},
		}); err != nil {
			return nil, err
		}
		newChannel.OwnerID = req.Msg.OwnerId
	}

	if req.Msg.ParentId != "" {
		if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ParentId); err != nil {
			return nil, err
		}
		newChannel.ParentID = req.Msg.ParentId
	}

	if req.Msg.Topic != "" {
		newChannel.Topic = sql.NullString{String: req.Msg.Topic, Valid: true}
	}

	if req.Msg.Recipients != nil {
		for _, recipient := range req.Msg.Recipients {
			if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
				Msg: &usersv1.GetByIdRequest{
					Id: recipient,
				},
			}); err == nil {
				newChannel.Recipients = append(newChannel.Recipients, recipient)
			}
		}
	}

	err := s.Repository.CreateChannel(ctx, newChannel)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.CreateResponse{
		Channel: createChannelResponse(newChannel),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GetById(ctx context.Context, req *connect.Request[channelsv1.GetByIdRequest]) (*connect.Response[channelsv1.GetByIdResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	channel, err := s.Repository.GetChannelByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GetByIdResponse{
		Channel: createChannelResponse(channel),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) Update(ctx context.Context, req *connect.Request[channelsv1.UpdateRequest]) (*connect.Response[channelsv1.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	channel, err := s.Repository.GetChannelByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	} else if channel == nil {
		return nil, apiErrors.ErrChannelNotFound
	}

	newChannel := &models.Channel{
		ID:        req.Msg.Id,
		CreatedAt: channel.CreatedAt,
	}

	if req.Msg.Name != "" {
		newChannel.Name = req.Msg.Name
	}

	if req.Msg.Icon != "" {
		newChannel.Icon = sql.NullString{String: req.Msg.Icon, Valid: true}
	}

	err = s.Repository.UpdateChannel(ctx, newChannel)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.UpdateResponse{
		Channel: createChannelResponse(newChannel),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) Delete(ctx context.Context, req *connect.Request[channelsv1.DeleteRequest]) (*connect.Response[channelsv1.DeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if channel, err := s.Repository.GetChannelByID(ctx, req.Msg.Id); err != nil {
		return nil, err
	} else if channel == nil {
		return nil, apiErrors.ErrChannelNotFound
	}

	err := s.Repository.DeleteChannel(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.DeleteResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GetChannelMessages(ctx context.Context, req *connect.Request[channelsv1.GetChannelMessagesRequest]) (*connect.Response[channelsv1.GetChannelMessagesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	messages, err := s.Repository.GetChannelMessages(ctx, req.Msg.ChannelId, repositories.SearchRequest{
		After:  req.Msg.After,
		Before: req.Msg.Before,
		Around: req.Msg.Around,
		Limit:  req.Msg.Limit,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GetChannelMessagesResponse{
		Messages: make([]*channelsv1.Message, 0, len(messages)),
	})

	for _, message := range messages {
		res.Msg.Messages = append(res.Msg.Messages, createMessageResponse(&message))
	}

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GetChannelMessage(ctx context.Context, req *connect.Request[channelsv1.GetChannelMessageRequest]) (*connect.Response[channelsv1.GetChannelMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	message, err := s.Repository.GetMessageByID(ctx, req.Msg.MessageId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GetChannelMessageResponse{
		Message: createMessageResponse(message),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) CreateMessage(ctx context.Context, req *connect.Request[channelsv1.CreateMessageRequest]) (*connect.Response[channelsv1.CreateMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.AuthorId,
		},
	}); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	newMessage := &models.Message{
		ChannelID:       req.Msg.ChannelId,
		AuthorID:        req.Msg.AuthorId,
		Content:         req.Msg.Content,
		Type:            channelsv1.MessageType(req.Msg.Type),
		Mentions:        make([]string, 0),
		MentionChannels: make([]string, 0),
		MentionRoles:    make([]string, 0),
		MentionEveryone: false,
		UpdatedAt:       sql.NullTime{},
	}

	if req.Msg.Content != "" {
		content := req.Msg.Content

		// check if message content contains <@ID> or <@&ID> or <@everyone> or <#ID>
		regs := []string{
			`<@(\d+)>`,
			`<@&(\d+)>`,
			`<@everyone>`,
			`<#(\d+)>`,
		}

		// TODO: move this to utils to check mentions in embeds

		for _, reg := range regs {
			re := regexp.MustCompile(reg)
			matches := re.FindAllStringSubmatch(content, -1)
			for _, match := range matches {
				if len(match) > 1 {
					switch reg {
					case `<@(\d+)>`:
						newMessage.Mentions = append(newMessage.Mentions, match[1])
					case `<@&(\d+)>`:
						newMessage.MentionRoles = append(newMessage.MentionRoles, match[1])
					case `<@everyone>`:
						// TODO: check if the author has permission to mention everyone
						newMessage.MentionEveryone = true
					case `<#(\d+)>`:
						newMessage.MentionChannels = append(newMessage.MentionChannels, match[1])
					}
				}
			}
		}
	}

	err := s.Repository.CreateMessage(ctx, newMessage)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.CreateMessageResponse{
		Message: createMessageResponse(newMessage),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}
