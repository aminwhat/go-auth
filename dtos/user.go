package dtos

import "go-auth/models"

type GetCurrentUserResponse struct {
	Succeed bool         `json:"succeed" example:"true"`
	Message string       `json:"message" example:"Succeed"`
	User    *models.User `json:"user"`
}

type GetAllUsersRequest struct {
	Page     int    `form:"page" json:"page" example:"1"`
	PageSize int    `form:"pageSize" json:"pageSize" example:"10"`
	Phone    string `form:"phone" json:"phone" example:"09123456789"`
}

type GetAllUsersResponse struct {
	Succeed    bool          `json:"succeed" example:"true"`
	Message    string        `json:"message" example:"Succeed"`
	Users      []models.User `json:"users"`
	TotalCount int64         `json:"totalCount" example:"100"`
	Page       int           `json:"page" example:"1"`
	PageSize   int           `json:"pageSize" example:"10"`
	TotalPages int           `json:"totalPages" example:"10"`
}
