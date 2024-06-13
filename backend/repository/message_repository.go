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

func (mr *messageRepository) Fetch(c context.Context, page int) ([]domain.MessageDB, bool, error) {
	collection := mr.database.Collection(mr.collection)

	opts := options.Find().SetSkip(int64((page - 1) * domain.PageLimit)).SetLimit(domain.PageLimit)
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, false, err
	}

	hasMore := true
	total, err := collection.CountDocuments(c, bson.D{})
	if int64(page*domain.PageLimit) >= total {
		hasMore = false
	}

	if err != nil {
		return nil, false, err
	}

	var messages []domain.MessageDB

	err = cursor.All(c, &messages)
	if messages == nil {
		return []domain.MessageDB{}, false, err
	}

	return messages, hasMore, err
}
