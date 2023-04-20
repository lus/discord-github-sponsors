package main

import (
	"github.com/lus/discord-github-sponsors/internal/config"
	"github.com/lus/discord-github-sponsors/internal/meta"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	// Set up the logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	if !meta.IsProdEnvironment() {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
		log.Warn().Msg("This distribution was compiled for development mode and is thus not meant to be run in production!")
	}

	// Load the configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load the configuration.")
	}

	// Adjust the log level
	if !meta.IsProdEnvironment() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		level, err := zerolog.ParseLevel(cfg.LogLevel)
		if err != nil {
			log.Warn().Msg("An invalid log level was configured. Falling back to 'info'.")
			level = zerolog.InfoLevel
		}
		zerolog.SetGlobalLevel(level)
	}

	// Wait for an interrupt signal
	log.Info().Msg("The application has been started. use Ctrl+C to shut it down.")
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	<-shutdownChan
}
