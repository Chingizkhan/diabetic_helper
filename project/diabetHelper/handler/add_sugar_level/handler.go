package add_sugar_level

import (
	"context"
	"database/sql"
	"diabetHelperTelegramBot/diabetHelper/storage"
	"diabetHelperTelegramBot/diabetHelper/storage/postgres"
	pb "diabetHelperTelegramBot/proto"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func Handle(ctx context.Context, req *pb.AddSLRequest) (*pb.SugarLevel, error) {
	dto := &DTO{}
	if err := dto.validate(req); err != nil {
		return nil, status.Error(codes.Code(http.StatusBadRequest), err.Error())
	}

	sl := storage.NewSugarLevel(dto.value, dto.userId)

	repo := postgres.New()

	err := repo.Begin(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	err = repo.Add(ctx, &sl)
	if errors.Is(err, postgres.ErrSLAlreadyExists) {
		// todo: return existing sugar level
		return &pb.SugarLevel{
			Value: "11",
		}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusBadRequest), err.Error())
	}

	err = repo.Commit()
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	return &pb.SugarLevel{
		Value: "22",
	}, nil
}
