package common

import (
	"diabetHelperTelegramBot/diabetHelper/storage"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
)

func MapSLToPb(level storage.SugarLevel) *pb.SugarLevel {
	return &pb.SugarLevel{
		UserId:    level.UserId,
		Value:     level.Value,
		CreatedAt: level.CreatedAt.Unix(),
		UpdatedAt: level.UpdatedAt.Unix(),
	}
}

func MapUserToPb(u storage.User) *pb.User {
	return &pb.User{
		Id:           u.Id,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		LanguageCode: u.LanguageCode,
	}
}
