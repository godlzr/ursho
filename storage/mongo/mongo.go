package mongo

import (
				"gopkg.in/mgo.v2"
				"time"
				"gopkg.in/mgo.v2/bson"
				"github.com/Ziyang2go/ursho/storage"
				"github.com/Ziyang2go/ursho/base62"
				"math/rand"
)

type Person struct {
	Name string
	Phone string
}

func New(host, port, dbName string) (storage.Service, error) {
	db, err :=  mgo.Dial(host + ":" + port)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &mongo{db}, nil
}

func (m *mongo) Close() error {
	m.db.Close()
	return nil;
}

func (m *mongo) Load(code string) (string, error) {
	// mongoId := bson.ObjectIdHex(code)
	id, _ := base62.Decode(code)
	c := m.db.DB("clients").C("shortener")

	 var item storage.Item
	 c.Find(bson.M{"uuid": id}).One(&item)
	return item.URL, nil
}

func (m *mongo) Save(url string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	uuid := rand.Int63();
	item := &storage.Item{ bson.NewObjectId(), uuid, url, false, 0}
	c := m.db.DB("clients").C("shortener")
	err := c.Insert(item)

	if err != nil {
		return "", err
	}

	return base62.Encode(uuid), nil
}

func (m *mongo) LoadInfo(code string) (*storage.Item, error) {
	id, _ := base62.Decode(code)
	c := m.db.DB("clients").C("shortener")
	var item storage.Item
	c.Find(bson.M{"uuid": id}).One(&item)
	return &item, nil
}

type mongo struct{ db *mgo.Session }
