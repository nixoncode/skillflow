package core

import "github.com/jmoiron/sqlx"

type App interface {
	Bootstrap() error
	Start() error
	Execute() error
	Shutdown() error

	DB() *sqlx.DB
}
