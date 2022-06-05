package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Key     string `json:"error"`
	Message string `json:"message"`
}

func newHTTPError(code int, key string, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Key:     key,
		Message: message,
	}
}

func (e *HttpError) Error() string {
	return e.Key + ": " + e.Message
}

func customHttpErrorHandler(err error, c echo.Context) {
	var (
		code    = http.StatusInternalServerError
		key     = "ServerError"
		message string
	)

	if he, ok := err.(*HttpError); ok {
		code = he.Code
		key = he.Key
		message = he.Message
	} else {
		message = http.StatusText(code)
	}

	err = c.JSON(code, newHTTPError(code, key, message))
	c.Logger().Error(err)
}
