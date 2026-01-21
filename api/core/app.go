package core

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type App interface {
	Bootstrap() error
	Start() error
	Execute() error
	Shutdown() error

	DB() *sqlx.DB
	Log() *zerolog.Logger
}
