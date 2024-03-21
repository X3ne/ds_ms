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

	"connectrpc.com/connect"
)

type ChannelsServer struct {
	Repository   *repositories.ChannelsRepository
	GuildsClient guildsv1connect.GuildsServiceClient
	UsersClient  usersv1connect.UsersServiceClient
}

func createChannelResponse(channel *models.Channel) (retChannel *channelsv1.Channel) {
	retChannel = &channelsv1.Channel{
		Id:   channel.ID,
		Name: channel.Name,
	}

	if channel.Icon.Valid {
		retChannel.Icon = channel.Icon.String
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

	if req.Msg.GuildId != 0 {
		if _, err := s.GuildsClient.GetById(ctx, &connect.Request[guildsv1.GetByIdRequest]{
			Msg: &guildsv1.GetByIdRequest{
				Id: req.Msg.GuildId,
			},
		}); err != nil {
			return nil, err
		}
	}

	if req.Msg.OwnerId != 0 {
		if _, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
			Msg: &usersv1.GetByIdRequest{
				Id: req.Msg.OwnerId,
			},
		}); err != nil {
			return nil, err
		}
	}

	if req.Msg.ParentId != 0 {
		if _, err := s.GetById(ctx, &connect.Request[channelsv1.GetByIdRequest]{
			Msg: &channelsv1.GetByIdRequest{
				Id: req.Msg.ParentId,
			},
		}); err != nil {
			return nil, err
		}
	}

	newChannel := &models.Channel{
		Name: req.Msg.Name,
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
