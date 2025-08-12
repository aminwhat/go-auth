package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthRegister struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	OtpCode     string             `bson:"otpCode" json:"otpCode"`
	Trys        int                `bson:"trys" json:"trys"`
	UpdatedDate primitive.DateTime `bson:"updatedDate" json:"updatedDate"`
	CreatedDate primitive.DateTime `bson:"createdDate" json:"createdDate"`
}
