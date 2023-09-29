package grpc

import (
	"context"
	"diabetHelperTelegramBot/diabetHelper/config"
	"diabetHelperTelegramBot/diabetHelper/server/grpc/middleware"
	pb "diabetHelperTelegramBot/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GraceFulInit(ctx context.Context) {
	listener, grpcServer := registerServer()

	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()
}

func registerServer() (net.Listener, *grpc.Server) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+config.Get().Port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.Logging))

	pb.RegisterDiabetHelperServer(grpcServer, &Server{})

	return listener, grpcServer
}
