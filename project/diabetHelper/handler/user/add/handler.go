package add

import (
	"context"
	"database/sql"
	"diabetHelperTelegramBot/diabetHelper/handler/common"
	"diabetHelperTelegramBot/diabetHelper/storage"
	"diabetHelperTelegramBot/diabetHelper/storage/postgres"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func Handle(ctx context.Context, req *pb.AddUserRequest) (*pb.User, error) {
	dto := DTO{}
	if err := dto.validate(req); err != nil {
		return nil, status.Error(codes.Code(http.StatusBadRequest), err.Error())
	}

	// create user
	user := storage.NewUser(dto.id, dto.username, dto.firstName, dto.lastName, dto.languageCode)

	// save user
	repo := postgres.NewUserStorage()
	err := repo.Begin(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}
	err = repo.Add(ctx, &user)
	if errors.Is(err, storage.ErrUserAlreadyExists) {
		return common.MapUserToPb(user), nil
	}
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	err = repo.Commit()
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	return common.MapUserToPb(user), nil
}
