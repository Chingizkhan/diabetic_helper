package diabet_helper

import (
	"context"
	grpc "diabetHelperTelegramBot/bot/service"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (c ClientGrpc) AddSL(ctx context.Context, req *pb.AddSLRequest) (*pb.SugarLevel, error) {
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
	client := pb.NewDiabetHelperClient(clientConn)
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

func (c ClientGrpc) FindSL(ctx context.Context, req *pb.FindSLRequest) (*pb.SugarLevels, error) {
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
	client := pb.NewDiabetHelperClient(clientConn)
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

func (c ClientGrpc) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.User, error) {
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
	client := pb.NewDiabetHelperClient(clientConn)
	res, err := client.AddUser(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable),
			fmt.Sprintf("%s:%s context deadline exceeded", errPrefix, "AddUser"))
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c ClientGrpc) FindUser(ctx context.Context, req *pb.FindUserRequest) (*pb.Users, error) {
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
	client := pb.NewDiabetHelperClient(clientConn)
	res, err := client.FindUser(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable),
			fmt.Sprintf("%s:%s context deadline exceeded", errPrefix, "FindUser"))
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c ClientGrpc) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
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
	client := pb.NewDiabetHelperClient(clientConn)
	res, err := client.UpdateUser(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(http.StatusServiceUnavailable),
			fmt.Sprintf("%s:%s context deadline exceeded", errPrefix, "UpdateUser"))
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
