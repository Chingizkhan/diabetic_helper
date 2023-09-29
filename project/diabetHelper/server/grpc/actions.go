package grpc

import (
	"context"
	"diabetHelperTelegramBot/diabetHelper/handler/add_sugar_level"
	"diabetHelperTelegramBot/diabetHelper/handler/find_sugar_level"
	"diabetHelperTelegramBot/diabetHelper/handler/update_sugar_level"
	pb "diabetHelperTelegramBot/proto"
)

func (s Server) AddSL(ctx context.Context, request *pb.AddSLRequest) (*pb.SugarLevel, error) {
	return add_sugar_level.Handle(ctx, request)
}

func (s Server) UpdateSL(ctx context.Context, request *pb.UpdateSLRequest) (*pb.SugarLevel, error) {
	return update_sugar_level.Handle(ctx, request)
}

func (s Server) FindSL(ctx context.Context, request *pb.FindSLRequest) (*pb.SugarLevels, error) {
	return find_sugar_level.Handle(ctx, request)
}
