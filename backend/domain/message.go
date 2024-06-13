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
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Text       string             `json:"text" bson:"text"`
	SentByName string             `json:"sentByName" bson:"sentByName"`
	SendByID   string             `json:"sentByID" bson:"sentByID"`
	CreatedAt  int64              `json:"createdAt" bson:"createdAt"`
}

type CreateMessage struct {
	Text       string `bson:"text"`
	SentByName string `bson:"sentByName"`
	SentByID   string `bson:"sentByID"`
	CreatedAt  int64  `bson:"createdAt"`
}

type MessageResponse struct {
	Messages []MessageDB `json:"messages"`
	HasMore  bool        `json:"hasMore"`
}

type MessageRepository interface {
	Create(c context.Context, message CreateMessage) error
	Fetch(c context.Context, page int) ([]MessageDB, bool, error)
}

type MessageService interface {
	GetMessages(c context.Context, page int) (*MessageResponse, error)
	InsertMessage(c context.Context, message CreateMessage) error
}
