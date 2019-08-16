package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	storage "ursho/storage"
	base62 "ursho/base62"
	"math/rand"
	"fmt"
	"context"
	"log"
)

type Person struct {
	Name string
	Phone string
}

func New(host, port, dbName string) (storage.Service, error) {
	fmt.Println("Connecting Mongo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// db, err :=  mgo.Dial(host + ":" + port)
	if err != nil {
		return nil, err
	}
	//err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("Connected Mongo")
	return &mongodb{client}, nil
}

func (m *mongodb) Close() error {
	err := m.client.Disconnect(context.Background())
	//m.db.Close()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (m *mongodb) Load(code string) (string, error) {
	// mongoId := bson.ObjectIdHex(code)
	id, _ := base62.Decode(code)
	collection := m.client.Database("clients").Collection("shortener")
	//c := m.db.DB("clients").C("shortener")

	var item storage.Item
	err := collection.FindOne(context.Background(), bson.M{"uuid": id}).Decode(&item)
	if err != nil { log.Fatal(err) }
	//c.Find(bson.M{"uuid": id}).One(&item)
	return item.URL, nil
}

func (m *mongodb) Save(url string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	uuid := rand.Int63();
	item := &storage.Item{ primitive.NewObjectID(), uuid, url, false, 0}
	collection := m.client.Database("clients").Collection("shortener")
	//c := m.db.DB("clients").C("shortener")
	_, err := collection.InsertOne(context.Background(), item)
	//err := c.Insert(item)

	if err != nil {
		return "", err
	}

	return base62.Encode(uuid), nil
}

func (m *mongodb) LoadInfo(code string) (*storage.Item, error) {
	id, _ := base62.Decode(code)
	collection := m.client.Database("clients").Collection("shortener")
	//c := m.db.DB("clients").C("shortener")
	var item storage.Item
	err := collection.FindOne(context.Background(), bson.M{"uuid": id}).Decode(&item)
	//c.Find(bson.M{"uuid": id}).One(&item)

	if err != nil { log.Fatal(err) }

	return &item, nil
}

type mongodb struct{ client *mongo.Client }
