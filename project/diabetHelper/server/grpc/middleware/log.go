package middleware

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
)

func Logging(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		reqData, _ := json.Marshal(req)
		log.Println(reqData)
		// todo: golemService.RegisterGrpcRequestWithError(string(reqData), err)
	}

	return resp, err
}
