package role

import (
	"Settings/models"
	"context"
	"github.com/go-kit/kit/log"
	"os"
)

type RoleService interface {
	AddRole(ctx context.Context, role *models.Role) (models.Status, models.Role, error)
}

type roleService struct{
}

func NewRoleService() RoleService { return &roleService{} }

func (w *roleService) AddRole(ctx context.Context, role *models.Role) (models.Status, models.Role, error) {
	err := logger.Log("adding role", "check this")
	if err != nil {
		return "", models.Role{}, err
	}
	return models.Success, *role, nil
}
var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}