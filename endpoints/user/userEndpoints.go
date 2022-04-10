package user

import (
	"Settings/models"
	"Settings/packages/user"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
)

type Set struct {
	AddUserEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc user.UserService) Set {
	return Set{
		AddUserEndpoint: MakeAddUserEndpoint(svc),
	}
}
func MakeAddUserEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		status, userResponse, err := svc.AddUser(ctx, &models.User{
			Name:   req.User.Name,
			Gender: req.User.Gender,
			Status: req.User.Status,
		})
		fmt.Println("now checking this")
		if err != nil {
			return AddUserResponse{status, userResponse, err.Error()}, nil
		}
		return AddUserResponse{status, userResponse, ""}, nil
	}
}
func (s *Set) AddUser(ctx context.Context, user *models.User) (string, models.User, error) {
	fmt.Print("Coming here")
	resp, err := s.AddUserEndpoint(ctx, AddUserRequest{User: user})
	if err != nil {
		return string(models.Failed), *user, err
	}
	adResp := resp.(AddUserResponse)
	if adResp.Error != "" {
		return string(models.Failed), *user, errors.New(adResp.Error)
	}
	return string(adResp.Status), adResp.User, errors.New(adResp.Error)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
