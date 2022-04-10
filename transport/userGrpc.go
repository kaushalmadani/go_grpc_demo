package transport

import (
	"Settings/endpoints/user"
	"Settings/models"
	_ "Settings/models"
	"Settings/output"
	"context"
	"fmt"
	_ "github.com/go-kit/kit/transport/grpc"
	transport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	addUser transport.Handler
}

func (g grpcServer) AddUser(ctx context.Context, request *output.AddUserRequest) (*output.AddUserResponse, error) {
	_, rep, err := g.addUser.ServeGRPC(ctx, request)
	if err != nil {
		fmt.Print("error while serving grpc")
		return nil, err
	}
	return rep.(*output.AddUserResponse), nil
}

func NewUserServer(ep user.Set) output.UserServiceServer {
	return &grpcServer{
		addUser: transport.NewServer(
			ep.AddUserEndpoint,
			decodeGRPCAddUserRequest,
			decodeGRPCAddUserResponse,
		),
	}
}
func decodeGRPCAddUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*output.AddUserRequest)
	userObj := models.User{
		Name:   req.User.Name,
		Gender: req.User.Gender,
		Status: req.User.Status,
	}
	return user.AddUserRequest{User: &userObj}, nil
}
func decodeGRPCAddUserResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(user.AddUserResponse)
	return &output.AddUserResponse{
		Status: string(reply.Status),
		User: &output.User{
			Name:   reply.User.Name,
			Gender: reply.User.Gender,
			Status: reply.User.Status,
		},
		Error: reply.Error,
	}, nil
}
