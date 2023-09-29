package diabet_helper

import (
	"context"
	grpc "diabetHelperTelegramBot/bot/service"
	"diabetHelperTelegramBot/proto/diabetHelper"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (c ClientGrpc) AddSL(ctx context.Context, req *diabetHelper.AddSLRequest) (*diabetHelper.SugarLevel, error) {
	clientConn, err := grpc.GetConn(ctx, c.addr)
	defer func() {
		if clientConn != nil {
			_ = clientConn.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := diabetHelper.NewDiabetHelperClient(clientConn)
	res, err := client.AddSL(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable),
			fmt.Sprintf("%s:%s context deadline exceeded", errPrefix, "AddSL"))
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c ClientGrpc) FindSL(ctx context.Context, req *diabetHelper.FindSLRequest) (*diabetHelper.SugarLevels, error) {
	clientConn, err := grpc.GetConn(ctx, c.addr)
	defer func() {
		if clientConn != nil {
			_ = clientConn.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := diabetHelper.NewDiabetHelperClient(clientConn)
	res, err := client.FindSL(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable),
			fmt.Sprintf("%s:%s context deadline exceeded", errPrefix, "FindSL"))
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
