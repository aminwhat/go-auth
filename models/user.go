package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	OtpCode     int                `bson:"otpCode" json:"otpCode"`
}
