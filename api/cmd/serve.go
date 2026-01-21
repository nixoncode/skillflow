package cmd

import (
	"fmt"

	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/internal/server"
	"github.com/spf13/cobra"
)

func NewServeCommand(app core.App) *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the SkillFlow API server",
		RunE: func(cmd *cobra.Command, args []string) error {

			srv := server.NewServer(app)

			addr := fmt.Sprintf("%s:%d", app.Config().Server.Host, app.Config().Server.Port)

			app.Log().Info().Msgf("Server configuration: Host=%s, Port=%d", app.Config().Server.Host, app.Config().Server.Port)
			if err := srv.Start(addr); err != nil {
				app.Log().Error().Err(err).Msg("Failed to start server")
				return err
			}

			return nil
		},
	}

	return serveCmd
}
