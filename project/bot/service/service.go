package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
)

func GetConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if conn != nil {
			_ = conn.Close()
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.Code(http.StatusServiceUnavailable), "Service at "+addr+" is unavailable")
		}
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable), err.Error())
	}

	return conn, nil
}
