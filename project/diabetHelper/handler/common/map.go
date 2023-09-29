package common

import (
	"diabetHelperTelegramBot/diabetHelper/storage"
	pb "diabetHelperTelegramBot/proto"
)

func MapSLToPb(level storage.SugarLevel) *pb.SugarLevel {
	return &pb.SugarLevel{
		UserId:    level.UserId,
		Value:     level.Value,
		CreatedAt: level.CreatedAt.Unix(),
		UpdatedAt: level.UpdatedAt.Unix(),
	}
}
