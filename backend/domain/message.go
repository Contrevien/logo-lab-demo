package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Text      string             `bson:"text,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	CreatedAt uint32             `bson:"created_at,omitempty"`
}
