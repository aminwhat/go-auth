package dtos

import "go-auth/models"

type GetCurrentUserResponse struct {
	Succeed bool        `json:"succeed" example:"true"`
	Message string      `json:"message" example:"Succeed"`
	User    models.User `json:"user"`
}
