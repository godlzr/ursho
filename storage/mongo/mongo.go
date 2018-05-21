package mongo

import (
        "gopkg.in/mgo.v2"
				"gopkg.in/mgo.v2/bson"
				"github.com/Ziyang2go/ursho/storage"			
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
	mongoId := bson.ObjectIdHex(code)
	c := m.db.DB("clients").C("shortener")

	var item storage.Item
	c.FindId(mongoId).One(&item)	
	return item.URL, nil
}

func (m *mongo) Save(url string) (string, error) {
	item := &storage.Item{ bson.NewObjectId(), url, false, 0}
	c := m.db.DB("clients").C("shortener")
	err := c.Insert(item)

	if err != nil {
		return "", err
	}
	
	objectIdString := item.ID.Hex()

	return objectIdString, nil
}

func (m *mongo) LoadInfo(code string) (*storage.Item, error) { 
	mongoId := bson.ObjectIdHex(code)
	c := m.db.DB("clients").C("shortener")
	var item storage.Item
	c.FindId(mongoId).One(&item)	
	return &item, nil
}

type mongo struct{ db *mgo.Session }
