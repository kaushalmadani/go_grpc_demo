package user

import (
	"Settings/models"
	"Settings/output"
	"context"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService UserService
}

func (ctlr *userServiceController) AddUser(ctx context.Context, request *output.AddUserRequest) (*output.AddUserResponse, error) {
	status, user, err := ctlr.userService.AddUser(ctx, &models.User{
		Name:   request.User.Name,
		Gender: request.User.Gender,
		Status: request.User.Status,
	})
	if err != nil {
		return &output.AddUserResponse{
			Status: string(status),
			User:   nil,
			Error:  err.Error(),
		}, err
	}
	resp := &output.AddUserResponse{
		Status: string(status),
		User: &output.User{
			Name:   user.Name,
			Gender: user.Gender,
			Status: user.Status,
		},
		Error: "",
	}
	return resp, nil
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService UserService) output.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

