package middleware

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// LoggerUsingZerolog is an labstack/echo middlewware which takes the fields
// from the normal log output and instead outputs them using a rs/zerolog
// logger.
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
func (z zerologWriter) Write(b []byte) (int, error) {
	var f fields
	err := json.Unmarshal(b, &f)
	if err != nil {
		return 0, err
	}

	z.logger.Info().
		Time("time", f.Time).
		Str("id", f.ID).
		Str("remote_ip", f.RemoteIP).
		Str("host", f.Host).
		Str("method", f.Method).
		Str("uri", f.URI).
		Str("user_agent", f.UserAgent).
		Int("status", f.Status).
		Str("error", f.Error).
		Int("latency", f.Latency).
		Str("latency_human", f.LatencyHuman).
		Int("bytes_in", f.BytesIn).
		Int("bytes_out", f.BytesOut).
		Send()

	return len(b), nil
}

type fields struct {
	Time         time.Time `json:"time"`
	ID           string    `json:"id"`
	RemoteIP     string    `json:"remote_ip"`
	Host         string    `json:"host"`
	Method       string    `json:"method"`
	URI          string    `json:"uri"`
	UserAgent    string    `json:"user_agent"`
	Status       int       `json:"status"`
	Error        string    `json:"error"`
	Latency      int       `json:"latency"`
	LatencyHuman string    `json:"latency_human"`
	BytesIn      int       `json:"bytes_in"`
	BytesOut     int       `json:"bytes_out"`
}
