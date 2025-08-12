package services

import (
	"go-auth/dtos"
	"go-auth/models"
	"go-auth/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUser(userId string) (dtos.GetCurrentUserResponse, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetUser(userId string) (dtos.GetCurrentUserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return dtos.GetCurrentUserResponse{Succeed: false, Message: err.Error()}, err
	}
	user, err := u.userRepo.Find(bson.M{"_id": objID})
	if err != nil {
		return dtos.GetCurrentUserResponse{Succeed: false, Message: err.Error()}, err
	}

	if user != nil {
		return dtos.GetCurrentUserResponse{Succeed: true, Message: "Succeed", User: user}, nil
	}

	return dtos.GetCurrentUserResponse{Succeed: false, Message: "User Not Found"}, nil
}

func (u *userService) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}
