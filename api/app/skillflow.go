package app

import "github.com/nixoncode/skillflow/core"

var _ core.App = (*SkillFlowApp)(nil)

type SkillFlowApp struct {
}

func New() *SkillFlowApp {
	return &SkillFlowApp{}
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
