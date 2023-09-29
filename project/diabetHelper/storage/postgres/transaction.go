package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Transaction struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewTransaction(ctx context.Context, options *sql.TxOptions) (*Transaction, error) {
	tr := &Transaction{}

	tr.db = GetDbConnection()
	var err error

	tr.tx, err = tr.db.BeginTxx(ctx, options)
	if err != nil {
		if tr.tx != nil {
			_ = tr.tx.Rollback()
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return tr, nil
}

func (tr *Transaction) RollbackIfError(err *error) {
	if (*err) != nil {
		_ = tr.tx.Rollback()
	}
}

func (tr *Transaction) PersistNamedCtx(ctx context.Context, query string, entity interface{}) error {
	_, err := tr.tx.NamedExecContext(ctx, query, entity)
	if err != nil {
		return status.Error(
			codes.Code(500),
			fmt.Sprintf("transaction PersistNamedCtx: entity=%v, err=%v, q=%s", entity, err, query),
		)
	}

	return nil
}

func (tr *Transaction) PersistExecContext(ctx context.Context, query string) error {
	_, err := tr.tx.ExecContext(ctx, query)
	if err != nil {
		return status.Error(
			codes.Code(500),
			fmt.Sprintf("transaction PersistExecContext: err=%v, q=%s", err, query),
		)
	}

	return nil
}

func (tr *Transaction) Rollback() error {
	if err := tr.tx.Rollback(); err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (tr *Transaction) Commit() error {
	if err := tr.tx.Commit(); err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func RollbackIfError(tr *Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
