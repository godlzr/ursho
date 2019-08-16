package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Save(string) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*Item, error)
	Close() error
}

type Item struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	UUID int64 `json:"uuid"`
	URL     string `json:"url"`
	Visited bool   `json:"visited"`
	Count   int    `json:"count"`
}
