package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMessage = "messages"
	PageLimit         = 10
)

type MessageDB struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Text       string             `bson:"text,omitempty"`
	SentByName string             `bson:"sentByName,omitempty"`
	SendByID   string             `bson:"sentByID,omitempty"`
	CreatedAt  int64              `bson:"createdAt,omitempty"`
}

type CreateMessage struct {
	Text       string `bson:"text,omitempty"`
	SentByName string `bson:"sentByName,omitempty"`
	SentByID   string `bson:"sentByID,omitempty"`
	CreatedAt  int64  `bson:"createdAt,omitempty"`
}

type MessageRepository interface {
	Create(c context.Context, message CreateMessage) error
	Fetch(c context.Context, page int) ([]MessageDB, error)
}

type MessageService interface {
	GetMessages(c context.Context, page int) ([]MessageDB, error)
	InsertMessage(c context.Context, message CreateMessage) error
}
