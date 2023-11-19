package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// LoggerUsingZerolog is a labstack/echo middlewware which takes the fields from
// the normal log output and instead outputs them using a rs/zerolog logger.
func LoggerUsingZerolog(l zerolog.Logger) echo.MiddlewareFunc {
	c := middleware.DefaultLoggerConfig
	c.Output = zerologWriter{
		logger: l,
	}
	return middleware.LoggerWithConfig(c)
}

type zerologWriter struct {
	logger zerolog.Logger
}

// Write implements the io.Writer interface.
func (z zerologWriter) Write(b []byte) (int, error) { //nolint:varnamelen
	var f fields //nolint:varnamelen
	if err := json.Unmarshal(b, &f); err != nil {
		return 0, fmt.Errorf("failed to unmarshal json log data: %w", err)
	}

	event := z.logger.Info()
	if f.Error != "" || f.Status >= http.StatusBadRequest {
		event = z.logger.Error()
	}

	if f.Error != "" {
		event.Str("error", f.Error)
	}

	if f.ID != "" {
		event.Str("id", f.ID)
	}

	event.Time("time", f.Time).
		Str("remote_ip", f.RemoteIP).
		Str("host", f.Host).
		Str("method", f.Method).
		Str("uri", f.URI).
		Str("user_agent", f.UserAgent).
		Int("status", f.Status).
		Int("latency", f.Latency).
		Str("latency_human", f.LatencyHuman).
		Int("bytes_in", f.BytesIn).
		Int("bytes_out", f.BytesOut)

	m := fmt.Sprintf("%s %s %d", f.Method, f.URI, f.Status)
	event.Msg(m)

	return len(b), nil
}

type fields struct {
	Time         time.Time `json:"time"`
	ID           string    `json:"id"`
	RemoteIP     string    `json:"remote_ip"` //nolint:tagliatelle
	Host         string    `json:"host"`
	Method       string    `json:"method"`
	URI          string    `json:"uri"`
	UserAgent    string    `json:"user_agent"` //nolint:tagliatelle
	Status       int       `json:"status"`
	Error        string    `json:"error"`
	Latency      int       `json:"latency"`
	LatencyHuman string    `json:"latency_human"` //nolint:tagliatelle
	BytesIn      int       `json:"bytes_in"`      //nolint:tagliatelle
	BytesOut     int       `json:"bytes_out"`     //nolint:tagliatelle
}
