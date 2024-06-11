package repository

import (
	"context"
	"time"

	"logo-lab-demo/domain"
	"logo-lab-demo/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type messageRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageRepository(db mongo.Database, collection string) domain.MessageRepository {
	return &messageRepository{
		database:   db,
		collection: collection,
	}
}

func (mr *messageRepository) Create(c context.Context, message domain.CreateMessage) error {
	collection := mr.database.Collection(mr.collection)

	message.CreatedAt = time.Now().Unix()
	_, err := collection.InsertOne(c, message)

	return err
}

func (mr *messageRepository) Fetch(c context.Context, page int) ([]domain.MessageDB, error) {
	collection := mr.database.Collection(mr.collection)

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var messages []domain.MessageDB

	err = cursor.All(c, &messages)
	if messages == nil {
		return []domain.MessageDB{}, err
	}

	return messages, err
}
