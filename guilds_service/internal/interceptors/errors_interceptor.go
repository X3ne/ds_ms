package interceptors

import (
	"context"
	"log"

	api_errors "github.com/X3ne/ds_ms/guilds_service/internal/errors"
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
				return connect.NewError(connect.CodeNotFound, api_errors.ErrGuildNotFound)
			case gorm.ErrInvalidDB:
				return connect.NewError(connect.CodeInternal, api_errors.ErrInternalServer)
			case gorm.ErrDuplicatedKey:
				return connect.NewError(connect.CodeAlreadyExists, api_errors.ErrGuildAlreadyExists)
			case err, gorm.ErrInvalidTransaction:
				return connect.NewError(connect.CodeInternal, api_errors.ErrInternalServer)
			default:
				return connect.NewError(connect.CodeUnknown, api_errors.ErrInternalServer)
		}
	}
}

func NewErrorInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
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
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
