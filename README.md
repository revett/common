# common

[![Builds](https://img.shields.io/github/checks-status/revett/common/main?label=build&style=flat-square)](https://github.com/revett/common/actions?query=branch%3Amain)
[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/revett/common)
[![Go Report Card](https://goreportcard.com/badge/github.com/revett/common?style=flat-square)](https://goreportcard.com/report/github.com/revett/common)
[![License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/revett/common/blob/main/LICENSE)

Shared common code for my Go projects.

## `log`

```go
package main

import (
	"context"

	commonlog "github.com/revett/common/log"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = commonlog.New()
	ctx := context.Background()

	// log.Info().Msg("...")
	// log.Fatal().Err(err).Send()
}
```
