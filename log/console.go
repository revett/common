package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New creates a new zerolog console logger, by default it uses RFC3339 as the
// time format.
func New(opts ...func(*zerolog.ConsoleWriter)) zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	writer.FormatTimestamp = func(i interface{}) string {
		return time.Now().Format(writer.TimeFormat)
	}

	for _, o := range opts {
		o(&writer)
	}

	return log.Output(writer)
}

// NewFuncOption is a functional option for the New function.
type NewFuncOption func(*zerolog.ConsoleWriter)

// WithNoColor is a functional option for the New function that disables
// colorized output for the logger.
func WithNoColor() NewFuncOption {
	return func(cw *zerolog.ConsoleWriter) {
		cw.NoColor = true
	}
}

// WithPrefix is a functional option for the New function that allows a prefix
// to be added to the logger.
func WithPrefix(p string) NewFuncOption {
	return func(cw *zerolog.ConsoleWriter) {
		cw.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s %s", p, i)
		}
	}
}

// WithTimeFormat is a functional option for the New function that allows a
// custom time format to be set for the logger.
func WithTimeFormat(t string) NewFuncOption {
	return func(cw *zerolog.ConsoleWriter) {
		cw.TimeFormat = t
	}
}
