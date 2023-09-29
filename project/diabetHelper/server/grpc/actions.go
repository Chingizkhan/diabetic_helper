package grpc

import (
	"context"
	slAdd "diabetHelperTelegramBot/diabetHelper/handler/sugar_level/add"
	slFind "diabetHelperTelegramBot/diabetHelper/handler/sugar_level/find"
	slUpdate "diabetHelperTelegramBot/diabetHelper/handler/sugar_level/update"
	userAdd "diabetHelperTelegramBot/diabetHelper/handler/user/add"
	userFind "diabetHelperTelegramBot/diabetHelper/handler/user/find"
	userUpdate "diabetHelperTelegramBot/diabetHelper/handler/user/update"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
)

func (s Server) AddSL(ctx context.Context, request *pb.AddSLRequest) (*pb.SugarLevel, error) {
	return slAdd.Handle(ctx, request)
}

func (s Server) UpdateSL(ctx context.Context, request *pb.UpdateSLRequest) (*pb.SugarLevel, error) {
	return slUpdate.Handle(ctx, request)
}

func (s Server) FindSL(ctx context.Context, request *pb.FindSLRequest) (*pb.SugarLevels, error) {
	return slFind.Handle(ctx, request)
}

func (s Server) AddUser(ctx context.Context, request *pb.AddUserRequest) (*pb.User, error) {
	return userAdd.Handle(ctx, request)
}

func (s Server) FindUser(ctx context.Context, request *pb.FindUserRequest) (*pb.Users, error) {
	return userFind.Handle(ctx, request)
}

func (s Server) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	return userUpdate.Handle(ctx, request)
}
