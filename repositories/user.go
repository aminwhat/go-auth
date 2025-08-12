package repositories

import (
	"context"
	"fmt"
	"go-auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Find(filter interface{}) (*models.User, error)
	Create(user models.User) (*models.User, error)
	ExistsByPhoneNumber(phoneNumber string) (*models.User, error)
	FindAllWithPagination(filter interface{}, page int, pageSize int) ([]models.User, int64, error)
	CountDocuments(filter interface{}) (int64, error)
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

	fmt.Println("User with filter of " + fmt.Sprintf("%v", filter) + " not found")
	return nil, nil
}

func (r *userRepository) Create(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return &models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (r *userRepository) ExistsByPhoneNumber(phoneNumber string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user *models.User

	filter := bson.M{"phoneNumber": phoneNumber}

	r.collection.FindOne(ctx, filter).Decode(&user)

	if user != nil {
		return user, nil
	}

	return nil, nil
}

func (r *userRepository) FindAllWithPagination(filter interface{}, page int, pageSize int) ([]models.User, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)

	cursor, err := r.collection.Find(ctx, filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  bson.D{{"createdDate", -1}}, // Sort by creation date (newest first)
	})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	totalCount, err := r.CountDocuments(filter)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (r *userRepository) CountDocuments(filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
