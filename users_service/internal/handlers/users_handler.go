package handlers

import (
	"context"
	"database/sql"

	usersv1 "github.com/X3ne/ds_ms/api/gen/users_service/users/v1"
	apiErrors "github.com/X3ne/ds_ms/users_service/internal/errors"
	"github.com/X3ne/ds_ms/users_service/internal/models"
	"github.com/X3ne/ds_ms/users_service/internal/repositories"
	"github.com/X3ne/ds_ms/users_service/internal/validator"

	"connectrpc.com/connect"
)

type UsersServer struct {
	Repository *repositories.UserRepository
}

func createUserResponse(user *models.User) (retUser *usersv1.User) {
	retUser = &usersv1.User{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}

	return
}

func (s *UsersServer) Create(ctx context.Context, req *connect.Request[usersv1.CreateRequest]) (*connect.Response[usersv1.CreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username: req.Msg.Username,
		Email:    req.Msg.Email,
		Password: sql.NullString{
			String: req.Msg.Password,
			Valid:  true,
		},
	}

	err := s.Repository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.CreateResponse{
		User: createUserResponse(newUser),
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (s *UsersServer) GetById(ctx context.Context, req *connect.Request[usersv1.GetByIdRequest]) (*connect.Response[usersv1.GetByIdResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	user, err := s.Repository.GetUserByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.GetByIdResponse{
		User: createUserResponse(user),
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (s *UsersServer) GetByEmail(ctx context.Context, req *connect.Request[usersv1.GetByEmailRequest]) (*connect.Response[usersv1.GetByEmailResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	user, err := s.Repository.GetUserByEmail(ctx, req.Msg.Email)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.GetByEmailResponse{
		User: createUserResponse(user),
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (s *UsersServer) Update(ctx context.Context, req *connect.Request[usersv1.UpdateRequest]) (*connect.Response[usersv1.UpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	user, err := s.Repository.GetUserByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, apiErrors.ErrUserNotFound
	}

	newUser := &models.User{
		ID:       req.Msg.Id,
		Username: req.Msg.Username,
		Email:    req.Msg.Email,
		Password: sql.NullString{
			String: req.Msg.Password,
			Valid:  true,
		},
		CreatedAt: user.CreatedAt,
	}

	err = s.Repository.UpdateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.UpdateResponse{
		User: createUserResponse(newUser),
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}

func (s *UsersServer) Delete(ctx context.Context, req *connect.Request[usersv1.DeleteRequest]) (*connect.Response[usersv1.DeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := validator.Validate(req.Msg); err != nil {
		return nil, err
	}

	if user, err := s.Repository.GetUserByID(ctx, req.Msg.Id); err != nil {
		return nil, err
	} else if user == nil {
		return nil, apiErrors.ErrUserNotFound
	}

	err := s.Repository.DeleteUser(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&usersv1.DeleteResponse{
		Success: true,
	})

	res.Header().Set("Users-Version", "v1")

	return res, nil
}
