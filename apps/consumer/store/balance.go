package store

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Balance struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	CustomerId string             `bson:"customerId,omitempty"`
	Balance float64            `bson:"balance,omitempty"`
}