// THIS IS ONLY FOR SWAGGER USAGE PURPOSES, NOT TO USE IN THE CODE BASE

package controllers_swagger

type AuthSignupBadResponse struct {
	Succeed bool   `example:"false"`
	Message string `example:"string"`
}

type AuthSignupConfirmOtpBadResponse struct {
	Succeed bool   `example:"false"`
	Message string `example:"string"`
	Token   string `example:""`
}
