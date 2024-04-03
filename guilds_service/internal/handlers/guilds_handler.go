package handlers

import (
	"context"
	"database/sql"
	channelsv1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
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
	Repository     *repositories.GuildRepository
	UsersClient    usersv1connect.UsersServiceClient
	ChannelsClient channelsv1connect.ChannelsServiceClient
}

func createGuildResponse(guild *models.Guild) (retGuild *guildsv1.Guild) {
	retGuild = &guildsv1.Guild{
		Id:          guild.ID,
		Name:        guild.Name,
		OwnerId:     guild.OwnerID,
		Icon:        guild.Icon.String,
		Splash:      guild.Splash.String,
		Banner:      guild.Banner.String,
		Description: guild.Description.String,
		CreatedAt:   guild.CreatedAt.Unix(),
		UpdatedAt:   guild.UpdatedAt.Unix(),
	}

	return
}

func createGuildChannelResponse(channel *channelsv1.Channel) (retChannel *guildsv1.GuildChannel) {
	retChannel = &guildsv1.GuildChannel{
		Id:            channel.Id,
		GuildId:       channel.GuildId,
		Name:          channel.Name,
		Type:          guildsv1.GuildChannelType(channel.Type),
		IsNsfw:        channel.IsNsfw,
		IsVoice:       channel.IsVoice,
		Permissions:   channel.Permissions,
		Position:      channel.Position,
		UserLimit:     channel.UserLimit,
		LastMessageId: channel.LastMessageId,
		ParentId:      channel.ParentId,
		Topic:         channel.Topic,
		CreatedAt:     channel.CreatedAt,
		UpdatedAt:     channel.UpdatedAt,
	}

	return
}

func createGuildMemberResponse(member *models.GuildMember, s *GuildsServer) (retMember *guildsv1.GuildMember, err error) {
	user, err := s.UsersClient.GetById(context.Background(), &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: member.UserID,
		},
	})
	if err != nil {
		return nil, err
	}

	retMember = &guildsv1.GuildMember{
		User:     user.Msg.User,
		Nick:     member.Nick.String,
		JoinedAt: member.JoinedAt.Unix(),
		//Roles:    member.Roles,
		Avatar: member.Avatar.String,
		Mute:   member.Mute,
		Deaf:   member.Deaf,
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

	_, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
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

	if req.Msg.OwnerId != "" {
		_, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
			Msg: &usersv1.GetByIdRequest{
				Id: req.Msg.OwnerId,
			},
		})
		if err != nil {
			return nil, err
		}
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

	if guild, err := s.Repository.GetGuildByID(ctx, req.Msg.Id); err != nil {
		return nil, err
	} else if guild == nil {
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

func (s *GuildsServer) GetGuildChannels(ctx context.Context, req *connect.Request[guildsv1.GetGuildChannelsRequest]) (*connect.Response[guildsv1.GetGuildChannelsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	channels, err := s.ChannelsClient.GetGuildChannels(ctx, &connect.Request[channelsv1.GetGuildChannelsRequest]{
		Msg: &channelsv1.GetGuildChannelsRequest{
			GuildId: req.Msg.GuildId,
		},
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.GetGuildChannelsResponse{
		Channels: make([]*guildsv1.GuildChannel, 0, len(channels.Msg.Channels)),
	})

	for _, channel := range channels.Msg.Channels {
		res.Msg.Channels = append(res.Msg.Channels, createGuildChannelResponse(channel))
	}

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) CreateGuildChannel(ctx context.Context, req *connect.Request[guildsv1.CreateGuildChannelRequest]) (*connect.Response[guildsv1.CreateGuildChannelResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	channel, err := s.ChannelsClient.Create(ctx, &connect.Request[channelsv1.CreateRequest]{
		Msg: &channelsv1.CreateRequest{
			GuildId: req.Msg.GuildId,
			Name:    req.Msg.Name,
			Type:    channelsv1.ChannelType(req.Msg.Type), // TODO: check if the binding works
		},
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.CreateGuildChannelResponse{
		Channel: createGuildChannelResponse(channel.Msg.Channel),
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) ModifyGuildChannelPositions(ctx context.Context, req *connect.Request[guildsv1.ModifyGuildChannelPositionsRequest]) (*connect.Response[guildsv1.ModifyGuildChannelPositionsResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	// TODO: maybe update all positions in a single transaction
	for _, channel := range req.Msg.Positions {
		_, err := s.ChannelsClient.Update(ctx, &connect.Request[channelsv1.UpdateRequest]{
			Msg: &channelsv1.UpdateRequest{
				Id:       channel.Id,
				Position: channel.Position,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	res := connect.NewResponse(&guildsv1.ModifyGuildChannelPositionsResponse{
		Positions: make([]*channelsv1.ChannelPosition, 0),
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) GetGuildMember(ctx context.Context, req *connect.Request[guildsv1.GetGuildMemberRequest]) (*connect.Response[guildsv1.GetGuildMemberResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	member, err := s.Repository.GetGuildMember(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	}

	retMember, err := createGuildMemberResponse(member, s)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.GetGuildMemberResponse{
		Member: retMember,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) ListGuildMembers(ctx context.Context, req *connect.Request[guildsv1.ListGuildMembersRequest]) (*connect.Response[guildsv1.ListGuildMembersResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	members, err := s.Repository.GetGuildMembers(ctx, req.Msg.GuildId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.ListGuildMembersResponse{
		Members: make([]*guildsv1.GuildMember, 0, len(members)),
	})

	for _, member := range members {
		retMember, err := createGuildMemberResponse(member, s)
		if err != nil {
			return nil, err
		}

		res.Msg.Members = append(res.Msg.Members, retMember)
	}

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) SearchGuildMembers(ctx context.Context, req *connect.Request[guildsv1.SearchGuildMembersRequest]) (*connect.Response[guildsv1.SearchGuildMembersResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	members, err := s.Repository.SearchGuildMembers(ctx, req.Msg.GuildId, req.Msg.Query)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.SearchGuildMembersResponse{
		Members: make([]*guildsv1.GuildMember, 0, len(members)),
	})

	for _, member := range members {
		retMember, err := createGuildMemberResponse(member, s)
		if err != nil {
			return nil, err
		}

		res.Msg.Members = append(res.Msg.Members, retMember)
	}

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) AddGuildMember(ctx context.Context, req *connect.Request[guildsv1.AddGuildMemberRequest]) (*connect.Response[guildsv1.AddGuildMemberResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	_, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Repository.GetGuildByID(ctx, req.Msg.GuildId)
	if err != nil {
		return nil, err
	}

	member := &models.GuildMember{
		UserID:  req.Msg.UserId,
		GuildID: req.Msg.GuildId,
		Mute:    req.Msg.Mute,
		Deaf:    req.Msg.Deaf,
	}

	// TODO: check if roles exist and add them to the member

	if req.Msg.Nick != "" {
		member.Nick = sql.NullString{String: req.Msg.Nick, Valid: true}
	}

	//if req.Msg.Roles != nil {
	//	member.Roles = req.Msg.Roles
	//}

	err = s.Repository.AddGuildMember(ctx, member)
	if err != nil {
		return nil, err
	}

	retMember, err := createGuildMemberResponse(member, s)

	res := connect.NewResponse(&guildsv1.AddGuildMemberResponse{
		Member: retMember,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) ModifyGuildMember(ctx context.Context, req *connect.Request[guildsv1.ModifyGuildMemberRequest]) (*connect.Response[guildsv1.ModifyGuildMemberResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	member, err := s.Repository.GetGuildMember(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	} else if member == nil {
		return nil, apiErrors.ErrGuildMemberNotFound
	}

	newMember := &models.GuildMember{
		UserID:  req.Msg.UserId,
		GuildID: req.Msg.GuildId,
		Mute:    req.Msg.Mute,
		Deaf:    req.Msg.Deaf,
		//Roles:   req.Msg.Roles,
	}

	if req.Msg.Nick != "" {
		newMember.Nick = sql.NullString{String: req.Msg.Nick, Valid: true}
	}

	err = s.Repository.UpdateGuildMember(ctx, newMember)
	if err != nil {
		return nil, err
	}

	retMember, err := createGuildMemberResponse(newMember, s)

	res := connect.NewResponse(&guildsv1.ModifyGuildMemberResponse{
		Member: retMember,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) AddGuildMemberRole(ctx context.Context, req *connect.Request[guildsv1.AddGuildMemberRoleRequest]) (*connect.Response[guildsv1.AddGuildMemberRoleResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	member, err := s.Repository.GetGuildMember(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	} else if member == nil {
		return nil, apiErrors.ErrGuildMemberNotFound
	}

	//member.Roles = append(member.Roles, req.Msg.RoleId)

	err = s.Repository.UpdateGuildMember(ctx, member)
	if err != nil {
		return nil, err
	}

	retMember, err := createGuildMemberResponse(member, s)

	res := connect.NewResponse(&guildsv1.AddGuildMemberRoleResponse{
		Member: retMember,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) RemoveGuildMemberRole(ctx context.Context, req *connect.Request[guildsv1.RemoveGuildMemberRoleRequest]) (*connect.Response[guildsv1.RemoveGuildMemberRoleResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	member, err := s.Repository.GetGuildMember(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	} else if member == nil {
		return nil, apiErrors.ErrGuildMemberNotFound
	}

	//for i, role := range member.Roles {
	//	if role == req.Msg.RoleId {
	//		member.Roles = append(member.Roles[:i], member.Roles[i+1:]...)
	//		break
	//	}
	//}

	err = s.Repository.UpdateGuildMember(ctx, member)
	if err != nil {
		return nil, err
	}

	retMember, err := createGuildMemberResponse(member, s)

	res := connect.NewResponse(&guildsv1.RemoveGuildMemberRoleResponse{
		Member: retMember,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) RemoveGuildMember(ctx context.Context, req *connect.Request[guildsv1.RemoveGuildMemberRequest]) (*connect.Response[guildsv1.RemoveGuildMemberResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	err := s.Repository.DeleteGuildMember(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.RemoveGuildMemberResponse{
		Success: true,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) GetGuildBans(ctx context.Context, req *connect.Request[guildsv1.GetGuildBansRequest]) (*connect.Response[guildsv1.GetGuildBansResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	bans, err := s.Repository.GetGuildBans(ctx, req.Msg.GuildId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.GetGuildBansResponse{
		Bans: make([]*guildsv1.Ban, 0, len(bans)),
	})

	for _, ban := range bans {
		user, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
			Msg: &usersv1.GetByIdRequest{
				Id: ban.UserID,
			},
		})
		if err != nil {
			return nil, err
		}
		res.Msg.Bans = append(res.Msg.Bans, &guildsv1.Ban{
			User:   user.Msg.User,
			Reason: ban.Reason.String,
		})
	}

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) GetGuildBan(ctx context.Context, req *connect.Request[guildsv1.GetGuildBanRequest]) (*connect.Response[guildsv1.GetGuildBanResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	ban, err := s.Repository.GetGuildBan(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	}

	user, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: ban.UserID,
		},
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.GetGuildBanResponse{
		Ban: &guildsv1.Ban{
			User:   user.Msg.User,
			Reason: ban.Reason.String,
		},
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) CreateGuildBan(ctx context.Context, req *connect.Request[guildsv1.CreateGuildBanRequest]) (*connect.Response[guildsv1.CreateGuildBanResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	_, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: req.Msg.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Repository.GetGuildByID(ctx, req.Msg.GuildId)
	if err != nil {
		return nil, err
	}

	ban := &models.GuildBan{
		UserID:  req.Msg.UserId,
		GuildID: req.Msg.GuildId,
		Reason:  sql.NullString{String: req.Msg.Reason, Valid: true},
	}

	err = s.Repository.AddGuildBan(ctx, ban)
	if err != nil {
		return nil, err
	}

	user, err := s.UsersClient.GetById(ctx, &connect.Request[usersv1.GetByIdRequest]{
		Msg: &usersv1.GetByIdRequest{
			Id: ban.UserID,
		},
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.CreateGuildBanResponse{
		Ban: &guildsv1.Ban{
			User:   user.Msg.User,
			Reason: ban.Reason.String,
		},
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) RemoveGuildBan(ctx context.Context, req *connect.Request[guildsv1.RemoveGuildBanRequest]) (*connect.Response[guildsv1.RemoveGuildBanResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	err := s.Repository.DeleteGuildBan(ctx, req.Msg.GuildId, req.Msg.UserId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&guildsv1.RemoveGuildBanResponse{
		Success: true,
	})

	res.Header().Set("Guilds-Version", "v1")

	return res, nil
}

func (s *GuildsServer) GetGuildRoles(ctx context.Context, req *connect.Request[guildsv1.GetGuildRolesRequest]) (*connect.Response[guildsv1.GetGuildRolesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, apiErrors.ErrNotImplemented)
}

func (s *GuildsServer) CreateGuildRole(ctx context.Context, req *connect.Request[guildsv1.CreateGuildRoleRequest]) (*connect.Response[guildsv1.CreateGuildRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, apiErrors.ErrNotImplemented)
}

func (s *GuildsServer) ModifyGuildRolePositions(ctx context.Context, req *connect.Request[guildsv1.ModifyGuildRolePositionsRequest]) (*connect.Response[guildsv1.ModifyGuildRolePositionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, apiErrors.ErrNotImplemented)
}

func (s *GuildsServer) ModifyGuildRole(ctx context.Context, req *connect.Request[guildsv1.ModifyGuildRoleRequest]) (*connect.Response[guildsv1.ModifyGuildRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, apiErrors.ErrNotImplemented)
}

func (s *GuildsServer) DeleteGuildRole(ctx context.Context, req *connect.Request[guildsv1.DeleteGuildRoleRequest]) (*connect.Response[guildsv1.DeleteGuildRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, apiErrors.ErrNotImplemented)
}
