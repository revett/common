package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Status is a generic status endpoint which returns a 200 response and "ok" as
// the response body.
func Status(ctx echo.Context) error {
	if err := ctx.String(http.StatusOK, "ok"); err != nil {
		return fmt.Errorf("failed to send ok response: %w", err)
	}

	return nil
}
