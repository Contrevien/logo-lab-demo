package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Name      primitive.ObjectID `bson:"name,omitempty"`
	CreatedAt uint32             `bson:"created_at,omitempty"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}

type UserService interface {
	GetProfileByEmail(c context.Context, email string) (*User, error)
}
