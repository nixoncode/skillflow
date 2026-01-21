package main

import "github.com/nixoncode/skillflow/internal/app"

func main() {
	sf := app.New()

	if err := sf.Bootstrap(); err != nil {
		sf.Log().Fatal().Err(err).Msg("Failed to bootstrap the application")
	}

	if err := sf.Start(); err != nil {
		sf.Log().Fatal().Err(err).Msg("Failed to start the application")
	}
}
