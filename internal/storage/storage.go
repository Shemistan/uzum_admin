package storage

import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"

func NewDatabase(connStr string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", connStr)
}
