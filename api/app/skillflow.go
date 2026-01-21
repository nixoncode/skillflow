package app

import (
	"github.com/nixoncode/skillflow/core"
	"github.com/spf13/cobra"
)

var _ core.App = (*SkillFlowApp)(nil)

type SkillFlowApp struct {
	rootCmd *cobra.Command
}

func New() *SkillFlowApp {
	return &SkillFlowApp{
		rootCmd: &cobra.Command{
			Use:     "skillflow",
			Short:   "SkillFlow is a learning platform API",
			Version: "1.0.0",
		},
	}
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
