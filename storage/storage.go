package storage

import (
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	Save(string) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*Item, error)
	Close() error
}

type Item struct {
	ID   bson.ObjectId `json:"_id" bson:"_id"`
	UUID int64 `json:"uuid"`
	URL     string `json:"url"`
	Visited bool   `json:"visited"`
	Count   int    `json:"count"`
}
