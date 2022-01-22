package database

import (
	"context"
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
	if insertResult.InsertedID == nil || err != nil {
		log.Printf("Invalid insertion %s, InsertedId: %d", err, insertResult.InsertedID)
		return false, err
	}
	return true, nil
}
func UpdateOne(collection string, filter interface{}, updatedDocument interface{}) (isUpdated bool, err error) {
	updateResult, err := db.Collection(collection).UpdateOne(context.Background(), filter, updatedDocument)
	if err != nil || updateResult.MatchedCount == 0 {
		log.Printf("Update err: %s, updatedCount: %d ", err, updateResult.UpsertedCount)
		return false, err
	}
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
	if err != nil || deleteResult.DeletedCount == 0 {
		log.Printf("Deletion err: %s, deletetedCount: %d ", err, deleteResult.DeletedCount)
		return false, err
	}
	return true, nil
}
