package role

import (
	models "Settings/models"
)

type AddRoleRequest struct {
	Role *models.Role `json:"user"`
}
type AddRoleResponse struct {
	Status models.Status `json:"status"`
	Role   models.Role   `json:"user"`
	Error  string        `json:"error"`
}
