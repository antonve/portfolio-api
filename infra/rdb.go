package infra

import (
	"context"

	"github.com/jmoiron/sqlx"
	// Postgres driver that's used to connect to the db
	_ "github.com/lib/pq"
)

// RDB is a relational database connection pool
type RDB = sqlx.DB

// NewRDB creates a new relational database connection pool
func NewRDB(URL string, maxIdleConns, maxOpenConns int) (*RDB, error) {
	db, err := sqlx.Open("postgres", URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	return db, nil
}

type Querier interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
}
