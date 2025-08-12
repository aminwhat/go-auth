package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HealthCheck struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedDate primitive.DateTime `bson:"createdDate" json:"createdDate"`
}
