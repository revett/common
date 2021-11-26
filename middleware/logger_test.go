package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/revett/common/middleware"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestLoggerUsingZerolog(t *testing.T) { // nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		status   int
		contains []string
	}{
		"StatusOK": {
			status: http.StatusOK,
			contains: []string{
				`"level":"info"`,
				`"remote_ip":"192.0.2.1"`,
				`"host":"example.com"`,
				`"method":"GET"`,
				`"uri":"/"`,
				`"status":200`,
				`"bytes_in":0`,
				`"bytes_out":4`,
				`"message":"GET / 200"`,
			},
		},
		"StatusInternalServerError": {
			status: http.StatusInternalServerError,
			contains: []string{
				`"level":"error"`,
				`"remote_ip":"192.0.2.1"`,
				`"host":"example.com"`,
				`"method":"GET"`,
				`"uri":"/"`,
				`"status":500`,
				`"bytes_in":0`,
				`"bytes_out":4`,
				`"message":"GET / 500"`,
			},
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			var buf bytes.Buffer
			zerologger := zerolog.New(&buf)

			loggerMiddleware := middleware.LoggerUsingZerolog(zerologger)

			handler := loggerMiddleware(
				func(c echo.Context) error {
					return c.String(tc.status, "test")
				},
			)

			err := handler(ctx)

			require.NoError(t, err)
			for _, s := range tc.contains {
				require.Contains(t, buf.String(), s)
			}
		})
	}
}
