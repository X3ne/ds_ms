package handlers

import (
	"connectrpc.com/connect"
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
	"reflect"
	"regexp"
)

type ChannelsServer struct {
	Repository   *repositories.ChannelsRepository
	GuildsClient guildsv1connect.GuildsServiceClient
	UsersClient  usersv1connect.UsersServiceClient
}

func mergeProtobufChannelToGORM(protoChannel interface{}, gormChannel *models.Channel) {
	protoValue := reflect.ValueOf(protoChannel).Elem()
	gormValue := reflect.ValueOf(gormChannel).Elem()

	for i := 0; i < protoValue.NumField(); i++ {
		fieldName := protoValue.Type().Field(i).Name

		if fieldName == "ID" {
			continue
		}

		protoField := protoValue.Field(i)
		gormField := gormValue.FieldByName(fieldName)

		if gormField.IsValid() && gormField.CanSet() && !isEmptyValue(protoField) {
			gormField.Set(protoField)
		}
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	}
	return false
}

func createChannelResponse(channel *models.Channel) (retChannel *channelsv1.Channel) {
	retChannel = &channelsv1.Channel{
		Id:            channel.ID,
		Name:          channel.Name,
		Type:          channel.Type,
		Icon:          channel.Icon.String,
		OwnerId:       channel.OwnerID.String,
		GuildId:       channel.GuildID.String,
		ParentId:      channel.ParentID.String,
		Position:      channel.Position.Int32,
		Topic:         channel.Topic.String,
		UserLimit:     channel.UserLimit.Int32,
		Recipients:    channel.Recipients,
		Permissions:   channel.Permissions.String,
		LastMessageId: channel.LastMessageID.String,
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
		Content:         message.Content.String,
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
		Position:  sql.NullInt32{Int32: req.Msg.Position, Valid: true},
		UserLimit: sql.NullInt32{Int32: req.Msg.UserLimit, Valid: true},
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
		newChannel.GuildID = sql.NullString{String: req.Msg.GuildId, Valid: true}
	}

	if req.Msg.OwnerId != "" {
		if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
			Msg: &usersv1.GetByIdRequest{
				Id: req.Msg.OwnerId,
			},
		}); err != nil {
			return nil, err
		}
		newChannel.OwnerID = sql.NullString{String: req.Msg.OwnerId, Valid: true}
	}

	if req.Msg.ParentId != "" {
		if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ParentId); err != nil {
			return nil, err
		}
		newChannel.ParentID = sql.NullString{String: req.Msg.ParentId, Valid: true}
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

	mergeProtobufChannelToGORM(req.Msg, channel)

	err = s.Repository.UpdateChannel(ctx, channel)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.UpdateResponse{
		Channel: createChannelResponse(channel),
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

func (s *ChannelsServer) GetGuildChannels(ctx context.Context, req *connect.Request[channelsv1.GetGuildChannelsRequest]) (*connect.Response[channelsv1.GetGuildChannelsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	channels, err := s.Repository.GetGuildChannels(ctx, req.Msg.GuildId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GetGuildChannelsResponse{
		Channels: make([]*channelsv1.Channel, 0, len(channels)),
	})

	for _, channel := range channels {
		res.Msg.Channels = append(res.Msg.Channels, createChannelResponse(&channel))
	}

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
		Content:         sql.NullString{String: req.Msg.Content, Valid: true},
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
						// TODO: check if the role is mentionable
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

func (s *ChannelsServer) UpdateMessage(ctx context.Context, req *connect.Request[channelsv1.UpdateMessageRequest]) (*connect.Response[channelsv1.UpdateMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	message, err := s.Repository.GetMessageByID(ctx, req.Msg.MessageId)
	if err != nil {
		return nil, err
	} else if message == nil {
		return nil, apiErrors.ErrMessageNotFound
	}

	//newMessage, err := mergeStruct(message, req.Msg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.Repository.UpdateMessage(ctx, newMessage.(*models.Message))
	//if err != nil {
	//	return nil, err
	//}
	//
	res := connect.NewResponse(&channelsv1.UpdateMessageResponse{
		Message: createMessageResponse(message),
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) DeleteMessage(ctx context.Context, req *connect.Request[channelsv1.DeleteMessageRequest]) (*connect.Response[channelsv1.DeleteMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	// TODO: check if this code is needed
	if message, err := s.Repository.GetMessageByID(ctx, req.Msg.MessageId); err != nil {
		return nil, err
	} else if message == nil {
		return nil, apiErrors.ErrMessageNotFound
	}

	err := s.Repository.DeleteMessage(ctx, req.Msg.MessageId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.DeleteMessageResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) BulkDeleteMessages(ctx context.Context, req *connect.Request[channelsv1.BulkDeleteMessagesRequest]) (*connect.Response[channelsv1.BulkDeleteMessagesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	err := s.Repository.DeleteMessages(ctx, req.Msg.ChannelId, req.Msg.Messages)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.BulkDeleteMessagesResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) EditChannelPermissions(ctx context.Context, req *connect.Request[channelsv1.EditChannelPermissionsRequest]) (*connect.Response[channelsv1.EditChannelPermissionsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	// TODO: define new permissions here

	res := connect.NewResponse(&channelsv1.EditChannelPermissionsResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) DeleteChannelPermission(ctx context.Context, req *connect.Request[channelsv1.DeleteChannelPermissionRequest]) (*connect.Response[channelsv1.DeleteChannelPermissionResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	// TODO: define new permissions here

	res := connect.NewResponse(&channelsv1.DeleteChannelPermissionResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) TriggerTypingIndicator(ctx context.Context, req *connect.Request[channelsv1.TriggerTypingIndicatorRequest]) (*connect.Response[channelsv1.TriggerTypingIndicatorResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	// TODO: Send typing indicator to the channel through the gateway_api API

	res := connect.NewResponse(&channelsv1.TriggerTypingIndicatorResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GetPinnedMessages(ctx context.Context, req *connect.Request[channelsv1.GetPinnedMessagesRequest]) (*connect.Response[channelsv1.GetPinnedMessagesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	messages, err := s.Repository.GetPinnedMessages(ctx, req.Msg.ChannelId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GetPinnedMessagesResponse{
		Messages: make([]*channelsv1.Message, 0, len(messages)),
	})

	for _, message := range messages {
		res.Msg.Messages = append(res.Msg.Messages, createMessageResponse(&message))
	}

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) AddPinnedMessage(ctx context.Context, req *connect.Request[channelsv1.AddPinnedMessageRequest]) (*connect.Response[channelsv1.AddPinnedMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	if err := s.Repository.PinMessage(ctx, req.Msg.MessageId); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.AddPinnedMessageResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) DeletePinnedMessage(ctx context.Context, req *connect.Request[channelsv1.DeletePinnedMessageRequest]) (*connect.Response[channelsv1.DeletePinnedMessageResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	if err := s.Repository.UnpinMessage(ctx, req.Msg.MessageId); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.DeletePinnedMessageResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GroupDMAddRecipient(ctx context.Context, req *connect.Request[channelsv1.GroupDMAddRecipientRequest]) (*connect.Response[channelsv1.GroupDMAddRecipientResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.UserId,
		},
	}); err != nil {
		return nil, err
	}

	//TODO: verify access_token and add user to the recipients

	res := connect.NewResponse(&channelsv1.GroupDMAddRecipientResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}

func (s *ChannelsServer) GroupDMRemoveRecipient(ctx context.Context, req *connect.Request[channelsv1.GroupDMRemoveRecipientRequest]) (*connect.Response[channelsv1.GroupDMRemoveRecipientResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if _, err := s.Repository.GetChannelByID(ctx, req.Msg.ChannelId); err != nil {
		return nil, err
	}

	if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.UserId,
		},
	}); err != nil {
		return nil, err
	}

	if err := s.Repository.RemoveGroupDMRecipient(ctx, req.Msg.ChannelId, req.Msg.UserId); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&channelsv1.GroupDMRemoveRecipientResponse{
		Success: true,
	})

	res.Header().Set("Channels-Version", "v1")

	return res, nil
}
