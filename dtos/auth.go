package dtos

type AuthSignupRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=11,max=11" example:"09123456789"`
}

type AuthSignupResponse struct {
	Succeed bool   `example:"true"`
	Message string `example:"string"`
}

type AuthLoginRequest struct {
}

type AuthSignupConfirmOtpRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=11,max=11" example:"09123456789"`
	OtpCode     string `json:"otpCode" example:"1234"`
}

type AuthTokenResponse struct {
	Succeed bool   `example:"true"`
	Message string `example:"Succeed"`
	Token   string `example:"the token will be here"`
}
