package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Status is a generic status endpoint which returns a 200 response and "ok" as
// the response body.
func Status(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "ok")
}
