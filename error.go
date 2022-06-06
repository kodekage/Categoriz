package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type httpError struct {
	Code    int    `json:"code"`
	Key     string `json:"error"`
	Message string `json:"message"`
}

func newHTTPError(code int, key string, message string) *httpError {
	return &httpError{
		Code:    code,
		Key:     key,
		Message: message,
	}
}

func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

func customHttpErrorHandler(err error, c echo.Context) {
	var (
		code    = http.StatusInternalServerError
		key     = "ServerError"
		message string
	)

	if he, ok := err.(*httpError); ok {
		code = he.Code
		key = he.Key
		message = he.Message
	} else {
		message = http.StatusText(code)
	}

	err = c.JSON(code, newHTTPError(code, key, message))
	c.Logger().Error(message)
}
