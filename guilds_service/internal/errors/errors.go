package api_errors

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrGuildNotFound									= connect.NewError(connect.CodeNotFound, errors.New("guild not found"))
	ErrGuildAlreadyExists							= connect.NewError(connect.CodeAlreadyExists, errors.New("guild already exists"))
	ErrUserNotFound										= connect.NewError(connect.CodeNotFound, errors.New("user not found"))

	ErrForbidden											= connect.NewError(connect.CodePermissionDenied, errors.New("forbidden"))

	ErrInternalServer									= connect.NewError(connect.CodeInternal, errors.New("internal server error"))
)
