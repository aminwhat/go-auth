package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id" example:"689b9bbf5800ec55229e240b"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" example:"09123456789"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate" example:"2025-08-12T19:53:35.685Z"`
}
