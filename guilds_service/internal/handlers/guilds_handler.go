package handlers

import (
	"context"
	"database/sql"

	usersv1 "github.com/X3ne/ds_ms/api/gen/users_service/users/v1"

	guildsv1 "github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1"
	"github.com/X3ne/ds_ms/api/gen/users_service/users/v1/usersv1connect"
	apiErrors "github.com/X3ne/ds_ms/guilds_service/internal/errors"
	"github.com/X3ne/ds_ms/guilds_service/internal/models"
	"github.com/X3ne/ds_ms/guilds_service/internal/repositories"
	"github.com/X3ne/ds_ms/guilds_service/internal/validator"

	"connectrpc.com/connect"
)

type GuildsServer struct {
	Repository *repositories.GuildRepository
	UserClient usersv1connect.UsersServiceClient
}

func createGuildResponse(guild *models.Guild) (retGuild *guildsv1.Guild) {
	retGuild = &guildsv1.Guild{
		Id:        guild.ID,
		Name:      guild.Name,
		OwnerId:   guild.OwnerID,
		CreatedAt: guild.CreatedAt.Unix(),
		UpdatedAt: guild.UpdatedAt.Unix(),
	}

	if guild.Icon.Valid {
		retGuild.Icon = guild.Icon.String
	}

	if guild.Splash.Valid {
		retGuild.Splash = guild.Splash.String
	}

	if guild.Banner.Valid {
		retGuild.Banner = guild.Banner.String
	}

	if guild.Description.Valid {
		retGuild.Description = guild.Description.String
	}

	return
}

func (s *GuildsServer) Create(ctx context.Context, req *connect.Request[guildsv1.CreateRequest]) (*connect.Response[guildsv1.CreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}
	_, err := s.UserClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.OwnerId,
		},
	})

	if err != nil {
		return nil, err
	}

	newGuild := &models.Guild{
		Name:    req.Msg.Name,
		OwnerID: req.Msg.OwnerId,
	}

	if req.Msg.Icon != "" {
		newGuild.Icon = sql.NullString{String: req.Msg.Icon, Valid: true}
	}

	err = s.Repository.CreateGuild(ctx, newGuild)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.CreateResponse{
		Guild: createGuildResponse(newGuild),
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) GetById(ctx context.Context, req *connect.Request[guildsv1.GetByIdRequest]) (*connect.Response[guildsv1.GetByIdResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	guild, err := s.Repository.GetGuildByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.GetByIdResponse{
		Guild: createGuildResponse(guild),
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) Update(ctx context.Context, req *connect.Request[guildsv1.UpdateRequest]) (*connect.Response[guildsv1.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	guild, err := s.Repository.GetGuildByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	} else if guild == nil {
		return nil, apiErrors.ErrGuildNotFound
	}

	newGuild := &models.Guild{
		ID:        req.Msg.Id,
		CreatedAt: guild.CreatedAt,
	}

	if req.Msg.Name != "" {
		newGuild.Name = req.Msg.Name
	}

	if req.Msg.Icon != "" {
		newGuild.Icon = sql.NullString{String: req.Msg.Icon, Valid: true}
	}

	if req.Msg.Splash != "" {
		newGuild.Splash = sql.NullString{String: req.Msg.Splash, Valid: true}
	}

	if req.Msg.Banner != "" {
		newGuild.Banner = sql.NullString{String: req.Msg.Banner, Valid: true}
	}

	if req.Msg.Description != "" {
		newGuild.Description = sql.NullString{String: req.Msg.Description, Valid: true}
	}

	if req.Msg.OwnerId != 0 {
		newGuild.OwnerID = req.Msg.OwnerId
	}

	err = s.Repository.UpdateGuild(ctx, newGuild)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.UpdateResponse{
		Guild: createGuildResponse(newGuild),
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) Delete(ctx context.Context, req *connect.Request[guildsv1.DeleteRequest]) (*connect.Response[guildsv1.DeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if user, err := s.Repository.GetGuildByID(ctx, req.Msg.Id); err != nil {
		return nil, err
	} else if user == nil {
		return nil, apiErrors.ErrGuildNotFound
	}

	err := s.Repository.DeleteGuild(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.DeleteResponse{
		Success: true,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}
