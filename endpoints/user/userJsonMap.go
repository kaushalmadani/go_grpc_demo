package user

import (
	models "Settings/models"
)

type AddUserRequest struct {
	User *models.User `json:"user"`
}
type AddUserResponse struct {
	Status models.Status `json:"status"`
	User   models.User   `json:"user"`
	Error  string        `json:"error"`
}
