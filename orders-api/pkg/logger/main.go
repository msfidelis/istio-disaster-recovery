package logger

import (
	"github.com/rs/zerolog"
	"os"
)

func Instance() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return logger
}
