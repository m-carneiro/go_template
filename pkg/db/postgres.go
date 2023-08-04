package db

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	ConnString string
	Pool       *pgxpool.Pool
}

func (dbp *Pool) Connect() error {
	pool, err := pgxpool.Connect(context.Background(), dbp.ConnString)
	if err != nil {
		return err
	}
	dbp.Pool = pool
	return nil
}

func (dbp *Pool) GetConnection() (conn *pgxpool.Conn, err error) {
	return dbp.Pool.Acquire(context.TODO())
}

func (dbp *Pool) NewTx(conn *pgxpool.Conn) (tx pgx.Tx, err error) {
	return conn.BeginTx(context.Background(), pgx.TxOptions{})
}

func (dbp *Pool) CommitOrRollback(tx pgx.Tx, err *error) (er error) {
	if *err != nil {
		er = tx.Rollback(context.TODO())
	} else {
		er = tx.Commit(context.TODO())
	}
	return
}

func (dbp *Pool) Exec(tx pgx.Tx, sql string, args ...interface{}) (tag pgconn.CommandTag, err error) {
	command, err := tx.Exec(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}

	return command, nil
}

func (dbp *Pool) Select(tx pgx.Tx, store interface{}, sql string, args ...interface{}) (err error) {
	var er error

	if args == nil {
		er = pgxscan.Select(context.Background(), tx, store, sql)
	} else {
		er = pgxscan.Select(context.Background(), tx, store, sql, args...)
	}

	if er != nil {
		fmt.Println("error on pgxscan.Select", err)
		return er
	}
	return nil
}
