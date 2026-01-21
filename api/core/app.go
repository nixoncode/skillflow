package core

type App interface {
	Bootstrap() error
	Start() error
	Execute() error
	Shutdown() error
}
