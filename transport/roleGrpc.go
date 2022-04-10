package transport

import (
	"Settings/endpoints/role"
	"Settings/models"
	"Settings/output"
	"context"
	"fmt"
	_ "github.com/go-kit/kit/transport/grpc"
	transport "github.com/go-kit/kit/transport/grpc"
)

type roleGrpcServer struct {
	addRole transport.Handler
}

func (r roleGrpcServer) AddRole(ctx context.Context, request *output.AddRoleRequest) (*output.AddRoleResponse, error) {
	_, rep, err := r.addRole.ServeGRPC(ctx, request)
	if err != nil {
		fmt.Print("error while serving grpc")
		return nil, err
	}
	return rep.(*output.AddRoleResponse), nil
}

func NewRoleServer(ep role.RoleSet) output.RoleServiceServer {
	return &roleGrpcServer{
		addRole: transport.NewServer(
			ep.AddRoleEndpoint,
			decodeGRPCAddRoleRequest,
			decodeGRPCAddRoleResponse,
		),
	}
}
func decodeGRPCAddRoleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*output.AddRoleRequest)
	roleObj := models.Role{
		Name:   req.Role.Name,
		Status: req.Role.Status,
	}
	return role.AddRoleRequest{Role: &roleObj}, nil
}
func decodeGRPCAddRoleResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(role.AddRoleResponse)
	return &output.AddRoleResponse{
		Status: string(reply.Status),
		Role: &output.Role{
			Name:   reply.Role.Name,
			Status: reply.Role.Status,
		},
		Error: reply.Error,
	}, nil
}
