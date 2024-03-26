package api_errors

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrChannelNotFound = connect.NewError(connect.CodeNotFound, errors.New("channel not found"))
	ErrMessageNotFound = connect.NewError(connect.CodeNotFound, errors.New("message not found"))
	ErrForbidden       = connect.NewError(connect.CodePermissionDenied, errors.New("forbidden"))
	ErrInternalServer  = connect.NewError(connect.CodeInternal, errors.New("internal server error"))
)
