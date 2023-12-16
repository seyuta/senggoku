package bootstrap

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitZerolog init zerolog configuration
func InitZerolog() {
	write := io.Writer(consoleWriter())
	log.Logger = zerolog.New(write).With().Timestamp().Logger()

	log.Info().
		Bool("consoleLogging", true).
		Bool("jsonLogOutput", false).
		Msg("zerolog configured")
}

// consoleWriter configure zerolog writer
func consoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
}
