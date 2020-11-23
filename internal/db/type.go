package db

import (
	"context"
	"database/sql"
)

// Executor can perform SQL queries with context
type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Transactor can commit and rollback, on top of being able to execute
// context-aware queries.
type Transactor interface {
	Commit() error
	Rollback() error

	Executor
}

// Beginner allows creation of context aware transactions with options.
type Beginner interface {
	BeginTx(context.Context, *sql.TxOptions) (Transactor, error)
}

// BeginnerExecutor can context-aware perform SQL queries and
// create context-aware transactions with options
type BeginnerExecutor interface {
	Beginner
	Executor
}

// DB is an implementation of Beginner and Executor
type DB struct {
	*sql.DB
}

// BeginTx begins a transaction with the database in receiver and returns a Transactor
func (d *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transactor, error) {
	tx, err := d.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
