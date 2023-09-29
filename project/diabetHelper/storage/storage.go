package storage

import (
	"context"
	"time"
)

type SugarLevelStorage interface {
	Add(ctx context.Context, sl SugarLevel) error
	Update(ctx context.Context, sl SugarLevel) error
	Find(ctx context.Context, dto SugarLevelFindDTO) ([]SugarLevel, error)
	Exists(ctx context.Context) (bool, error)
}

type UserStorage interface {
	Add(ctx context.Context, u User) error
	Update(ctx context.Context, u User) error
	Find(ctx context.Context, u UserFindDTO) ([]User, error)
	Exists(ctx context.Context) (bool, error)
}

type User struct {
	Id           string `db:"id"`
	Username     string `db:"user_name"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	LanguageCode string `db:"language_code"`
}

type SugarLevel struct {
	UserId    int64     `db:"user_id"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewSugarLevel(value string, userId int64) SugarLevel {
	return SugarLevel{
		UserId:    userId,
		Value:     value,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func NewUser(id, username, firstname, lastname, languageCode string) User {
	return User{
		Id:           id,
		Username:     username,
		FirstName:    firstname,
		LastName:     lastname,
		LanguageCode: languageCode,
	}
}

type SugarLevelFindDTO struct {
	UserId          int64
	Value           string
	CreatedAtStart  time.Time
	CreatedAtFinish time.Time
	Limit, Offset   uint32
}

type UserFindDTO struct {
	Id       string
	Username string
}
