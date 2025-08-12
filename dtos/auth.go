package dtos

type AuthSignupRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=11,max=11"`
}

type AuthSignupResponse struct {
	Succeed bool
	Message string
}

type AuthLoginRequest struct {
}

type AuthSignupConfirmOtpRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=11,max=11"`
	OtpCode     string `json:"otpCode"`
}

type AuthTokenResponse struct {
	Succeed bool
	Message string
	Token   string
}
