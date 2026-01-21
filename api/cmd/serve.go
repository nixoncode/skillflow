package cmd

import (
	"fmt"

	"github.com/nixoncode/skillflow/core"
	"github.com/spf13/cobra"
)

func NewServeCommand(app core.App) *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the SkillFlow API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implementation for starting the server goes here
			fmt.Println(app.Config().App.Name)
			return nil
		},
	}

	return serveCmd
}
