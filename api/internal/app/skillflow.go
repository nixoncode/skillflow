package app

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nixoncode/skillflow/config"
	"github.com/nixoncode/skillflow/core"
	"github.com/spf13/cobra"
)

var _ core.App = (*SkillFlowApp)(nil)

type SkillFlowApp struct {
	rootCmd *cobra.Command
	config  *config.Config
	db      *sqlx.DB
}

func New() *SkillFlowApp {
	app := &SkillFlowApp{
		config: config.LoadConfig(),
	}

	app.rootCmd = &cobra.Command{
		Use:     app.config.App.Name,
		Short:   "SkillFlow is a learning platform API",
		Version: app.config.App.Version,
	}

	return app
}

func (app *SkillFlowApp) Bootstrap() error {

	if err := app.initDB(); err != nil {
		return err
	}

	return nil
}

func (app *SkillFlowApp) Start() error {
	return nil
}

func (app *SkillFlowApp) Execute() error {
	return nil
}

func (app *SkillFlowApp) Shutdown() error {
	return nil
}

func (app *SkillFlowApp) DB() *sqlx.DB {
	return app.db
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
