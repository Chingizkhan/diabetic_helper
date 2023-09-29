package postgres

import (
	"context"
	"database/sql"
	"diabetHelperTelegramBot/diabetHelper/storage"
	"diabetHelperTelegramBot/diabetHelper/storage/postgres/queryBuilder"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Db *sqlx.DB
	*sqlx.Tx
}

func New() *Storage {
	return &Storage{
		Db: db,
		Tx: nil,
	}
}

var (
	ErrSLAlreadyExists = errors.New("sugar level with this id already exists")
)

func (s *Storage) Begin(ctx context.Context, options *sql.TxOptions) (err error) {
	s.Tx, err = s.Db.BeginTxx(ctx, options)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	return nil
}

func (s *Storage) Commit() error {
	if s.Tx == nil {
		return errors.New("can't commit when transaction was not began")
	}
	err := s.Tx.Commit()
	if err != nil {
		return fmt.Errorf("can't commit: %w", err)
	}
	return nil
}

func (s *Storage) Add(ctx context.Context, level *storage.SugarLevel) error {
	q := `
		INSERT INTO sugar_levels (user_id, value, created_at, updated_at)
		VALUES (:user_id, :value, :created_at, :updated_at);
	`

	//exists, err := s.Exists(ctx, level.Id)
	//if err != nil {
	//	return err
	//}
	//if exists {
	//	return ErrSLAlreadyExists
	//}

	_, err := s.Tx.NamedExecContext(ctx, q, level)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Exists(ctx context.Context, userId string) (has bool, err error) {
	q := `SELECT EXISTS(SELECT true FROM sugar_levels WHERE user_id = $1);`

	err = s.Tx.GetContext(ctx, &has, q, userId)
	if err != nil {
		return false, fmt.Errorf("can't check if sugar_level already exists: %w", err)
	}
	return has, nil
}

func (s *Storage) Update(ctx context.Context, level storage.SugarLevel) error {
	return nil
}

func (s *Storage) Find(ctx context.Context, dto storage.SugarLevelFindDTO) ([]storage.SugarLevel, error) {
	var levels []storage.SugarLevel

	qb := queryBuilder.NewQueryBuilder(
		"sugar_levels",
		"*",
		"sl",
	)
	err := fillQueryForFind(qb, dto)
	if err != nil {
		return nil, fmt.Errorf("can't 'fill query for find': %w", err)
	}

	err = qb.OrderBy("sl.created_at", true)
	if err != nil {
		return nil, fmt.Errorf("can't 'order by created_at': %w", err)
	}

	qb.Limit(dto.Limit)
	qb.Offset(dto.Offset)

	err = s.Db.SelectContext(ctx, &levels, qb.GetQuery(), qb.GetParameters()...)
	switch err {
	case nil, sql.ErrNoRows:
		return levels, nil
	case context.DeadlineExceeded:
		return nil, fmt.Errorf("can't make request to db: %w", err)
	default:
		return nil, fmt.Errorf("can't make request to db: %w", err)
	}
}

func (s *Storage) GetCountAllForFind(ctx context.Context, dto storage.SugarLevelFindDTO) (uint32, error) {
	var count uint32

	qb := queryBuilder.NewQueryBuilder(
		"sugar_levels",
		"count(*)",
		"sl",
	)
	err := fillQueryForFind(qb, dto)
	if err != nil {
		return 0, fmt.Errorf("can't 'fill query for find' count: %w", err)
	}

	err = s.Db.GetContext(ctx, &count, qb.GetQuery(), qb.GetParameters()...)
	switch err {
	case nil, sql.ErrNoRows:
		return count, nil
	case context.DeadlineExceeded:
		return 0, fmt.Errorf("can't make request to db: %w", err)
	default:
		return 0, fmt.Errorf("can't make request to db: %w", err)
	}
}

func fillQueryForFind(qb *queryBuilder.QueryBuilder, dto storage.SugarLevelFindDTO) error {
	if dto.UserId != 0 {
		_ = qb.WhereEqual("sl.user_id", dto.UserId, queryBuilder.And)
	}

	if dto.Value != "" {
		_ = qb.WhereEqual("sl.value", dto.Value, queryBuilder.And)
	}

	if !dto.CreatedAtStart.IsZero() {
		_ = qb.WhereGreaterOrEqual("sl.created_at", dto.CreatedAtStart, queryBuilder.And)
	}

	if !dto.CreatedAtFinish.IsZero() {
		_ = qb.WhereLessOrEqual("sl.created_at", dto.CreatedAtFinish, queryBuilder.And)
	}
	return nil
}
