package api_errors

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrUserNotFound										= connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	ErrUserAlreadyExists							= connect.NewError(connect.CodeAlreadyExists, errors.New("user already exists"))
	ErrForbidden											= connect.NewError(connect.CodePermissionDenied, errors.New("forbidden"))

	ErrInternalServer									= connect.NewError(connect.CodeInternal, errors.New("internal server error"))
)
