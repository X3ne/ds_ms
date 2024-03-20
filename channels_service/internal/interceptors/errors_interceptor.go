package interceptors

import (
	"context"
	"log"

	apiErrors "github.com/X3ne/ds_ms/channels_service/internal/errors"
	"github.com/bufbuild/protovalidate-go"
	"gorm.io/gorm"

	"connectrpc.com/connect"
)

func handleErrors(ctx context.Context, err error) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	switch e := err.(type) {
	case *protovalidate.ValidationError:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case *connect.Error:
		return err
	default:
		switch e {
		case gorm.ErrRecordNotFound:
			return connect.NewError(connect.CodeNotFound, apiErrors.ErrChannelNotFound)
		case gorm.ErrInvalidDB:
			return connect.NewError(connect.CodeInternal, apiErrors.ErrInternalServer)
		case gorm.ErrDuplicatedKey:
			return connect.NewError(connect.CodeAlreadyExists, apiErrors.ErrInternalServer)
		case err, gorm.ErrInvalidTransaction:
			return connect.NewError(connect.CodeInternal, apiErrors.ErrInternalServer)
		default:
			return connect.NewError(connect.CodeUnknown, apiErrors.ErrInternalServer)
		}
	}
}

func NewErrorInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if err := ctx.Err(); err != nil {
				return nil, err
			}

			res, err := next(ctx, req)
			if err != nil {
				log.Println(err)
				return nil, handleErrors(ctx, err)
			}

			return res, nil
		}
	}
	return interceptor
}
