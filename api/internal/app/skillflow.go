package app

import (
	"github.com/nixoncode/skillflow/config"
	"github.com/nixoncode/skillflow/core"
	"github.com/spf13/cobra"
)

var _ core.App = (*SkillFlowApp)(nil)

type SkillFlowApp struct {
	rootCmd *cobra.Command
	config  *config.Config
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
