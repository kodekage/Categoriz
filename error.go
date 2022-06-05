package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func newHTTPError(code int, key string, message string) *httpError {
	return &httpError{
		code:    code,
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
		code = he.code
		key = he.Key
		message = he.Message
	} else {
		message = http.StatusText(code)
	}

	if c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
			//} else if config.Debug {
			//	message = err.Error()
		} else {
			err := c.JSON(code, newHTTPError(code, key, message))

			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
