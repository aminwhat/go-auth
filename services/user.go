package services

import (
	"go-auth/dtos"
	"go-auth/models"
	"go-auth/repositories"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUser(userId string) (dtos.GetCurrentUserResponse, error)
	GetAllUsers() ([]models.User, error)
	GetAllUsersWithPagination(request dtos.GetAllUsersRequest) (dtos.GetAllUsersResponse, error)
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

func (u *userService) GetAllUsersWithPagination(request dtos.GetAllUsersRequest) (dtos.GetAllUsersResponse, error) {
	// Set default values and validate
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PageSize < 1 {
		request.PageSize = 10
	}
	if request.PageSize > 100 {
		request.PageSize = 100
	}

	// Build filter
	filter := bson.M{}
	if request.Phone != "" {
		filter["phoneNumber"] = bson.M{"$regex": request.Phone, "$options": "i"}
	}

	// Get paginated users
	users, totalCount, err := u.userRepo.FindAllWithPagination(filter, request.Page, request.PageSize)
	if err != nil {
		return dtos.GetAllUsersResponse{Succeed: false, Message: err.Error()}, err
	}

	// Initialize empty slice if users is nil
	if users == nil {
		users = []models.User{}
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalCount) / float64(request.PageSize)))
	if totalPages == 0 && totalCount > 0 {
		totalPages = 1
	}

	return dtos.GetAllUsersResponse{
		Succeed:    true,
		Message:    "Success",
		Users:      users,
		TotalCount: totalCount,
		Page:       request.Page,
		PageSize:   request.PageSize,
		TotalPages: totalPages,
	}, nil
}
