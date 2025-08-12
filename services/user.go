package services

import (
	"go-auth/models"
	"go-auth/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUser(userId string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetUser(userId string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	return u.userRepo.Find(bson.M{"_id": objID})
}
