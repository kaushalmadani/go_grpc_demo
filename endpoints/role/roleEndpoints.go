package role

import (
	"Settings/models"
	"Settings/packages/role"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
)

type RoleSet struct {
	AddRoleEndpoint endpoint.Endpoint
}

func NewRoleEndpointSet(svc role.RoleService) RoleSet {
	return RoleSet{
		AddRoleEndpoint: MakeAddRoleEndpoint(svc),
	}
}
func MakeAddRoleEndpoint(svc role.RoleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRoleRequest)
		status, userResponse, err := svc.AddRole(ctx, &models.Role{
			Name:   req.Role.Name,
			Status: req.Role.Status,
		})
		if err != nil {
			return AddRoleResponse{status, userResponse, err.Error()}, nil
		}
		return AddRoleResponse{status, userResponse, ""}, nil
	}
}
func (s *RoleSet) AddRole(ctx context.Context, role *models.Role) (string, models.Role, error) {
	fmt.Print("Coming here")
	resp, err := s.AddRoleEndpoint(ctx, AddRoleRequest{Role: role})
	if err != nil {
		return string(models.Failed), *role, err
	}
	adResp := resp.(AddRoleResponse)
	if adResp.Error != "" {
		return string(models.Failed), *role, errors.New(adResp.Error)
	}
	return string(adResp.Status), adResp.Role, errors.New(adResp.Error)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
