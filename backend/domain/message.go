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
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `bson:"text"`
	SentByName string             `bson:"sentByName"`
	SendByID   string             `bson:"sentByID"`
	CreatedAt  int64              `bson:"createdAt"`
}

type CreateMessage struct {
	Text       string `bson:"text"`
	SentByName string `bson:"sentByName"`
	SentByID   string `bson:"sentByID"`
	CreatedAt  int64  `bson:"createdAt"`
}

type MessageResponse struct {
	Messages []MessageDB `bson:"messages"`
	HasMore  bool        `bson:"hasMore"`
}

type MessageRepository interface {
	Create(c context.Context, message CreateMessage) error
	Fetch(c context.Context, page int) ([]MessageDB, bool, error)
}

type MessageService interface {
	GetMessages(c context.Context, page int) (*MessageResponse, error)
	InsertMessage(c context.Context, message CreateMessage) error
}
