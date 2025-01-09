package repository

import (
	"context"

	"github.com/Magowtham/go_book_app/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB(db *mongo.Database) *MongoDB {
	return &MongoDB{
		db,
	}
}

func (d *MongoDB) CreateUser(user *types.User) error {
	_, err := d.db.Collection("users").InsertOne(context.Background(), user)
	return err
}
