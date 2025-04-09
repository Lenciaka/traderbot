package utils

import (
	"github.com/jmoiron/sqlx"
)

func ConnectSqlite(uri string) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", uri)
}
