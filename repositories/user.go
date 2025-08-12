package repositories

import (
	"context"
	"go-auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Find(filter interface{}) (*models.User, error)
	Create(user models.User) (models.User, error)
	ExistsByPhoneNumber(phoneNumber string) (bool, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) Find(filter interface{}) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user *models.User
	r.collection.FindOne(ctx, filter).Decode(&user)

	if user != nil {
		return user, nil
	}

	return nil, nil
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *userRepository) ExistsByPhoneNumber(phoneNumber string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user *models.User

	filter := bson.M{"phoneNumber": phoneNumber}

	r.collection.FindOne(ctx, filter).Decode(&user)

	if user != nil {
		return true, nil
	}

	return false, nil
}
