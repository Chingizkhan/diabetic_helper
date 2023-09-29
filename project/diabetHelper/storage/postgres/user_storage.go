package postgres

import (
	"context"
	"database/sql"
	"diabetHelperTelegramBot/diabetHelper/storage"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	Db *sqlx.DB
	tr *sqlx.Tx
}

func NewUserStorage() UserStorage {
	return UserStorage{Db: GetDbConnection()}
}

func (s *UserStorage) Begin(ctx context.Context, options *sql.TxOptions) (err error) {
	s.tr, err = s.Db.BeginTxx(ctx, options)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	return nil
}

func (s *UserStorage) Commit() error {
	if s.tr == nil {
		return errors.New("can't commit when transaction was not began")
	}
	err := s.tr.Commit()
	if err != nil {
		return fmt.Errorf("can't commit: %w", err)
	}
	return nil
}

func (s *UserStorage) Add(ctx context.Context, u *storage.User) (err error) {
	defer s.RollbackIfError(&err)

	q := `
		INSERT INTO users (id, username, first_name, last_name, language_code)
		VALUES (:id, :username, :first_name, :last_name, :language_code);
	`

	exists, err := s.Exists(ctx, u.Id)
	if err != nil {
		return err
	}
	if exists {
		return storage.ErrUserAlreadyExists
	}

	_, err = s.tr.NamedExecContext(ctx, q, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) Update(ctx context.Context, u storage.User) error {
	panic("not implemented")
}

func (s *UserStorage) Find(ctx context.Context, u storage.UserFindDTO) ([]storage.User, error) {
	panic("not implemented")
}

func (s *UserStorage) Exists(ctx context.Context, id int64) (has bool, err error) {
	q := `SELECT EXISTS(SELECT true FROM users WHERE id = $1);`

	err = s.tr.GetContext(ctx, &has, q, id)
	if err != nil {
		return false, fmt.Errorf("can't check if user already exists: %w", err)
	}
	return has, nil
}

func (s *UserStorage) RollbackIfError(err *error) {
	if (*err) != nil {
		_ = s.tr.Rollback()
	}
}
