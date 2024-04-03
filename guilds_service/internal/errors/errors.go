package api_errors

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrGuildNotFound        = connect.NewError(connect.CodeNotFound, errors.New("guild not found"))
	ErrGuildAlreadyExists   = connect.NewError(connect.CodeAlreadyExists, errors.New("guild already exists"))
	ErrUserNotFound         = connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	ErrForbidden            = connect.NewError(connect.CodePermissionDenied, errors.New("forbidden"))
	ErrGuildMemberNotFound  = connect.NewError(connect.CodeAlreadyExists, errors.New("guild member not found"))
	ErrChannelNotFound      = connect.NewError(connect.CodeNotFound, errors.New("channel not found"))
	ErrChannelAlreadyExists = connect.NewError(connect.CodeAlreadyExists, errors.New("channel already exists"))
	ErrMessageNotFound      = connect.NewError(connect.CodeNotFound, errors.New("message not found"))
	ErrInternalServer       = connect.NewError(connect.CodeInternal, errors.New("internal server error"))
	ErrNotImplemented       = connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
)
