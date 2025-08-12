package controllers_swagger

type GetCurrentUserBadResponse struct {
	Succeed bool   `json:"succeed" example:"false"`
	Message string `json:"message" example:"string"`
	User    string `json:"user" example:"null"`
}
