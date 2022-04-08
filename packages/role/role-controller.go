package role

import (
	"Settings/models"
	"Settings/output"
	"context"
)

type roleServiceController struct {
	roleService RoleService
}
func (ctr *roleServiceController) AddRole(ctx context.Context, request *output.AddRoleRequest) (*output.AddRoleResponse, error) {
	status, role, err := ctr.roleService.AddRole(ctx, &models.Role{
		Name:   request.Role.Name,
		Status: request.Role.Status,
	})
	if err != nil {
		return &output.AddRoleResponse{
			Status: string(status),
			Error:  err.Error(),
			Role: nil,
		}, err
	}
	resp := &output.AddRoleResponse{
		Status: string(status),
		Role: &output.Role{
			Name:   role.Name,
			Status: role.Status,
		},
		Error: "",
	}
	return resp, nil
}

func NewRoleServiceController(roleService RoleService) output.RoleServiceServer {
	return &roleServiceController{
		roleService: roleService,
	}
}
