package httperr

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Error(code int, message string, err error) *echo.HTTPError {
	return &echo.HTTPError{
		Code:     code,
		Message:  message,
		Internal: err,
	}
}

func Message(code int, message string) *echo.HTTPError {
	return Error(code, message, errors.New(message))
}

func Forbidden(err error) *echo.HTTPError {
	return Error(http.StatusForbidden, "forbidden", err)
}

func Unauthorized(err error) *echo.HTTPError {
	return Error(http.StatusUnauthorized, "user is not authorized to use this resource", err)
}

func NotFound(err error) *echo.HTTPError {
	return Error(http.StatusNotFound, "not found", err)
}

func BadRequest(err error) *echo.HTTPError {
	return Error(http.StatusBadRequest, "bad request", err)
}

func InternalServerError(err error) *echo.HTTPError {
	return Error(http.StatusBadRequest, "internal server error", err)
}

func ServerErrorOrNotFound(err error) *echo.HTTPError {
	if errors.Is(err, sql.ErrNoRows) {
		return Error(http.StatusNotFound, "not found", err)
	}
	return InternalServerError(err)
}

func MalformedID(err error) *echo.HTTPError {
	return Error(http.StatusBadRequest, "malformed id", err)
}

func MalformedBody(err error) *echo.HTTPError {
	return Error(http.StatusBadRequest, "malformed request body", err)
}

func MalformedDate(err error) *echo.HTTPError {
	return Error(http.StatusBadRequest, "malformed date - use RFC3339 format", err)
}

func MissingQueryParameter(param string) *echo.HTTPError {
	msg := "missing query parameter " + param
	return Error(http.StatusBadRequest, msg, errors.New(msg))
}

func MatchRouteParam(param string) *echo.HTTPError {
	msg := "object " + param + " does not match route param"
	return Error(http.StatusBadRequest, msg, errors.New(msg))
}
