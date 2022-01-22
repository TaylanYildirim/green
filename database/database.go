package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"green/config"
	"green/models/maze"
	"log"
	"time"
)

var db *mongo.Database

func Init() error {
	dbUri := config.Get().Servers.DB.Uri
	clientOptions := options.Client()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(20)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri), clientOptions)

	if err != nil {
		return err
	}
	db = client.Database(config.Get().Servers.DB.Name)
	return nil
}

func InsertOne(collection string, newDocument interface{}) (isInserted bool, err error) {
	insertResult, err := db.Collection(collection).InsertOne(context.Background(), newDocument)
	fmt.Println(insertResult, err)
	return true, nil
}
func FindById(collection string, filter interface{}, results *maze.Maze) error {
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(results)
	if nil != err {
		return err
	}
	return nil
}

func DeleteOne(collection string, filter interface{}) (bool, error) {
	deleteResult, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		log.Printf("Invalid maze id")
		return false, err
	}
	if nil != err {
		return false, err
	}
	return true, nil
}
