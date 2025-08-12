package services

import (
	"go-auth/dtos"
	"go-auth/repositories"
)

type AuthService interface {
	Signup(model dtos.AuthSignupRequest) (dtos.AuthSignupResponse, error)
	SignupConfirmOtp(model dtos.AuthSignupConfirmOtpRequest) (dtos.AuthTokenResponse, error)
	Login(model dtos.AuthLoginRequest) (dtos.AuthTokenResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Signup(model dtos.AuthSignupRequest) (dtos.AuthSignupResponse, error) {
	panic("unimplemented")
}

func (s *authService) SignupConfirmOtp(model dtos.AuthSignupConfirmOtpRequest) (dtos.AuthTokenResponse, error) {
	panic("unimplemented")
}

func (s *authService) Login(model dtos.AuthLoginRequest) (dtos.AuthTokenResponse, error) {
	panic("unimplemented")
}
