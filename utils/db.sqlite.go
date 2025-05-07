package utils

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func ConnectSqlite(uri string) (*sqlx.DB, error) {
	return sqlx.Open("sqlite", uri)
}
