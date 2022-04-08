package user

import (
	"Settings/models"
	"context"
	"github.com/go-kit/kit/log"
	"os"
)

type UserService interface {
	AddUser(ctx context.Context, user *models.User) (models.Status, models.User, error)
	UpdateUser(ctx context.Context, userId string, user models.User) (models.Status, error)
	GetUsers(ctx context.Context, filters ...models.Filter) ([]models.User, error)
}

type userService struct{
	check string
}

func NewUserService() UserService { return &userService{} }

func (w *userService) GetUsers(_ context.Context, filters ...models.Filter) ([]models.User, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	doc := models.User{
		Name:   "J.K. Rowling",
		Gender: "male",
		Status: "active",
	}
	return []models.User{doc}, nil
}
func (w *userService) AddUser(ctx context.Context, user *models.User) (models.Status, models.User, error) {
	err := logger.Log("adding user", "check this")
	if err != nil {
		return "", models.User{}, err
	}
	return models.Success, *user, nil
}

func (w *userService) UpdateUser(ctx context.Context, userId string, user models.User) (models.Status, error) {
	return models.Success, nil
}
var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}