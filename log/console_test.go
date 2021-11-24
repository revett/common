package log_test

import (
	"github.com/revett/common/log"
	zerolog "github.com/rs/zerolog/log"
)

func ExampleNew() {
	zerolog.Logger = log.New(
		log.WithTimeFormat(""), log.WithNoColor(),
	)
	zerolog.Info().Msg("example")
	// Output: INF example
}

func ExampleWithPrefix() {
	zerolog.Logger = log.New(
		log.WithTimeFormat(""), log.WithNoColor(), log.WithPrefix("service"),
	)
	zerolog.Info().Msg("example")
	// Output: INF service example
}
