package model

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// DBTX includes all methods shared by sqlx.DB and sqlx.Tx, allowing
// either type to be used interchangeably.
// https://github.com/jmoiron/sqlx/pull/809
type DBTX interface {
	sqlx.Ext
	sqlx.ExecerContext
	sqlx.PreparerContext
	sqlx.QueryerContext
	sqlx.Preparer

	GetContext(context.Context, interface{}, string, ...interface{}) error
	SelectContext(context.Context, interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	MustExecContext(context.Context, string, ...interface{}) sql.Result
	PreparexContext(context.Context, string) (*sqlx.Stmt, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Select(interface{}, string, ...interface{}) error
	QueryRow(string, ...interface{}) *sql.Row
	PrepareNamedContext(context.Context, string) (*sqlx.NamedStmt, error)
	PrepareNamed(string) (*sqlx.NamedStmt, error)
	Preparex(string) (*sqlx.Stmt, error)
	NamedExec(string, interface{}) (sql.Result, error)
	NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
	MustExec(string, ...interface{}) sql.Result
	NamedQuery(string, interface{}) (*sqlx.Rows, error)
}

type DBRows interface {
	Close() error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...interface{}) error
	SliceScan() ([]interface{}, error)
	MapScan(dest map[string]interface{}) error
	StructScan(dest interface{}) error
}

type Tx interface {
	Commit() error
	Rollback() error
}

var _ DBTX = (*sqlx.DB)(nil)
var _ DBTX = (*sqlx.Tx)(nil)
var _ DBRows = (*sqlx.Rows)(nil)
var _ Tx = (*sqlx.Tx)(nil)

var sqlIn = sqlx.In

type Database struct {
	*sqlx.DB
}

func (db *Database) Queries() *Queries {
	return &Queries{db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx *sqlx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

func TxDo(rollback func() error) {
	err := rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		log.Print(err.Error())
	}
}

func NewDatabase(cfg *config.DBConfig) *Database {
	db, err := sqlx.Connect("pgx", cfg.ConnStr())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err.Error())
	}
	if db == nil {
		log.Panicf("database is nil")
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 30)

	return &Database{db}
}
