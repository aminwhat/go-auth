package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type HealthCheckRepository interface {
	CheckTheHealth() (bool, error)
}

type healthChekRepository struct {
	collection *mongo.Collection
}

func NewHealthCheckRepository(db *mongo.Database) HealthCheckRepository {
	return &healthChekRepository{
		collection: db.Collection("health_checks"),
	}
}

func (h *healthChekRepository) CheckTheHealth() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := h.collection.InsertOne(ctx, map[string]interface{}{
		"createdDate": time.Now(),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
