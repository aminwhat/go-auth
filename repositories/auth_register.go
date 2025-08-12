package repositories

import (
	"context"
	"go-auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRegisterRepository interface {
	Create(authRegister models.AuthRegister) error
	Update(authRegister models.AuthRegister) error
	Exists(phoneNumber string, otpCode int) (bool, error)
	ExistsByPhoneNumber(phoneNumber string) (*models.AuthRegister, error)
	ExistsByOtpCode(otpCode int) (bool, error)
}

type authRegisterRepository struct {
	collection *mongo.Collection
}

func NewAuthRegisterRepository(db *mongo.Database) AuthRegisterRepository {
	return &authRegisterRepository{
		collection: db.Collection("auth_registers"),
	}
}

func (a *authRegisterRepository) Create(authRegister models.AuthRegister) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.collection.InsertOne(ctx, authRegister)
	if err != nil {
		return err
	}

	return nil
}

func (a *authRegisterRepository) Exists(phoneNumber string, otpCode int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var auth_register *models.AuthRegister

	filter := bson.M{"phoneNumber": phoneNumber, "otpCode": otpCode}
	a.collection.FindOne(ctx, filter).Decode(&auth_register)

	if auth_register != nil {
		return true, nil
	}

	return false, nil
}

func (a *authRegisterRepository) ExistsByPhoneNumber(phoneNumber string) (*models.AuthRegister, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var model *models.AuthRegister

	filter := bson.M{"phoneNumber": phoneNumber}
	a.collection.FindOne(ctx, filter).Decode(&model)

	if model != nil {
		return model, nil
	}

	return nil, nil
}

func (a *authRegisterRepository) ExistsByOtpCode(otpCode int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var auth_register *models.AuthRegister

	filter := bson.M{"otpCode": otpCode}
	a.collection.FindOne(ctx, filter).Decode(&auth_register)

	if auth_register != nil {
		return true, nil
	}

	return false, nil
}

func (a *authRegisterRepository) Update(authRegister models.AuthRegister) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": authRegister.ID}
	update := bson.M{"$set": authRegister}

	_, err := a.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
