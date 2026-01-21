package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nixoncode/skillflow/cmd"
	"github.com/nixoncode/skillflow/config"
	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/pkg/logs"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var _ core.App = (*SkillFlowApp)(nil)

type SkillFlowApp struct {
	rootCmd *cobra.Command
	config  *config.Config
	db      *sqlx.DB
	logger  *zerolog.Logger
}

func New() *SkillFlowApp {
	app := &SkillFlowApp{
		config: config.LoadConfig(),
	}

	app.rootCmd = &cobra.Command{
		Use:     strings.ToLower(app.config.App.Name),
		Short:   "SkillFlow is a learning platform API",
		Version: app.config.App.Version,
	}

	return app
}

func (app *SkillFlowApp) Bootstrap() error {

	app.logger = logs.SetupLogger(app.config.App.LogLevel, app.config.App.IsDebug)

	if err := app.initDB(); err != nil {
		return err
	}

	return nil
}

func (app *SkillFlowApp) Start() error {

	app.logger.Info().Msgf("Starting %s version %s", app.config.App.Name, app.config.App.Version)

	app.rootCmd.AddCommand(cmd.NewServeCommand(app))
	app.rootCmd.AddCommand(cmd.NewMigrateCommand(app.logger, app.db.DB, app.config.App.IsDebug))

	return app.Execute()
}

func (app *SkillFlowApp) Execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	defer stop()

	errChan := make(chan error, 1)

	go func() {
		if err := app.rootCmd.ExecuteContext(ctx); err != nil {
			errChan <- fmt.Errorf("cobra command failed: %w", err)
		}
		errChan <- nil
	}()

	select {
	case <-ctx.Done():
		app.logger.Info().Msg("Shutdown down gracefully...")
		return app.Shutdown()
	case err := <-errChan:
		if err != nil {
			return err
		}
	}

	return nil

}

func (app *SkillFlowApp) Shutdown() error {
	if app.db != nil {
		if err := app.db.Close(); err != nil {
			app.logger.Error().Err(err).Msg("Failed to close database connection")
			return err
		}
		app.logger.Info().Msg("Database connection closed")
	}
	app.logger.Info().Msgf("%s has been shut down", app.config.App.Name)
	return nil
}

func (app *SkillFlowApp) DB() *sqlx.DB {
	return app.db
}

func (app *SkillFlowApp) Log() *zerolog.Logger {
	return app.logger
}

func (app *SkillFlowApp) Config() *config.Config {
	return app.config
}

func (app *SkillFlowApp) initDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		app.config.DB.User,
		app.config.DB.Password,
		app.config.DB.Host,
		app.config.DB.Port,
		app.config.DB.Name,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	app.db = db
	return nil
}
