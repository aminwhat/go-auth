package services

import "go-auth/repositories"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthService interface {
	Signup() User
	Login() User
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Signup() User {
	panic("unimplemented")
}

func (s *authService) Login() User {
	panic("unimplemented")
}
