package cmd

import (
	"database/sql"
	"fmt"

	"github.com/nixoncode/skillflow/internal/database"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func NewMigrateCommand(logger *zerolog.Logger, db *sql.DB, isDebug bool) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Database migration commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			create, _ := cmd.Flags().GetString("create")

			if create != "" {
				if !isDebug {
					logger.Warn().Msg("Migration creation is only allowed in debug mode")
					return fmt.Errorf("migration creation is only allowed in debug mode")
				}
				if err := database.CreateMigration(create); err != nil {
					logger.Error().Err(err).Msg("Failed to create migration")
					return err
				}
				logger.Info().Msgf("Successfully created migration: %s", create)
				return nil
			}

			switch {
			case getBool(cmd, "rollback"):
				return database.RollbackLastMigration(db)
			case getBool(cmd, "refresh"):
				return database.RefreshMigrations(db)
			case getBool(cmd, "status"):
				return database.MigrationStatus(db)
			case getBool(cmd, "reset"):
				return database.ResetMigrations(db)
			default: // up as default action
				return database.RunMigrations(db)
			}
		},
	}
	migrateCmd.Flags().Bool("rollback", false, "Rollback the last migration")
	migrateCmd.Flags().Bool("refresh", false, "Refresh all migrations")
	migrateCmd.Flags().Bool("status", false, "Show migration status")
	migrateCmd.Flags().StringP("create", "c", "", "Create a new migration with the given name e.g  create_users_table")
	migrateCmd.Flags().Bool("reset", false, "Reset all migrations")

	return migrateCmd

}

func getBool(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}
