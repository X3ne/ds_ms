package responses

import (
	"connectrpc.com/connect"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func getConnectError(err error) (*connect.Error, bool) {
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		return nil, false
	}

	return connectErr, true
}

func connectErrorToHTTPStatus(err *connect.Error) int {
	switch err.Code() {
	case connect.CodeNotFound:
		return http.StatusNotFound
	case connect.CodeInvalidArgument:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func Response(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, Error{
		Code:    status,
		Message: err.Error(),
	})
}

func ConnectErrorResponse(c echo.Context, err error) error {
	if connectErr, ok := getConnectError(err); ok {
		code := connectErrorToHTTPStatus(connectErr)
		return c.JSON(code, Error{
			Code:    code,
			Message: connectErr.Message(),
		})
	}

	return ErrorResponse(c, http.StatusInternalServerError, err)
}

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, Success{
		Success: true,
		Data:    data,
	})
}
