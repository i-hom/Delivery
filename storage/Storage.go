package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type IStorage interface {
	Connect(string, string) error
	Close() error
	Find(string, bson.M, interface{}) error
	FindAll(string, bson.M, interface{}) error
	Insert(string, interface{}) error
	Update(string, bson.M, interface{}) error
	Delete(string, bson.M) error
}

type Storage struct {
	db *mongo.Database
	conn *mongo.Client
}

func (s *Storage) Connect(dbUrl, dbName string) error {
	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)
	s.conn, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}
	s.db = s.conn.Database(dbName)
	return nil
}

func (s *Storage) Close() error {
	return s.conn.Disconnect(context.TODO())
}

func (s *Storage) Find(collectionName string, filter bson.M, result interface{}) error {
	return s.db.Collection(collectionName).FindOne(context.TODO(), filter).Decode(result)
}

func (s *Storage) FindAll(collectionName string, filter bson.M, result interface{}) error {
	curr, _ := s.db.Collection(collectionName).Find(context.TODO(), filter)
	return curr.All(context.TODO(), result)
}

func (s *Storage) Insert(collectionName string, document interface{}) error {
	_, err := s.db.Collection(collectionName).InsertOne(context.TODO(), document)
	return err
}

func (s *Storage) Update(collectionName string, filter bson.M, update interface{}) error {
	_, err := s.db.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	return err
}

func (s *Storage) Delete(collectionName string, filter bson.M) error {
	_, err := s.db.Collection(collectionName).DeleteOne(context.TODO(), filter)
	return err
}