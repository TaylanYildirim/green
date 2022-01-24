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

const (
	InsertSuccessMessage string = "Maze successfully inserted."
	InsertFailMessage           = "Maze couldn't be inserted. MazeId is not."
	UpdateSuccessMessage        = "Maze successfully updated."
	UpdateFailMessage           = "Maze couldn't be updated. MazeId couldn't found."
	DeleteSuccessMessage        = "Maze successfully deleted."
	DeleteFailMessage           = "Maze couldn't be deleted. MazeId couldn't found."
)

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

func InsertOne(collection string, newDocument interface{}) (string, error) {
	insertResult, err := db.Collection(collection).InsertOne(context.Background(), newDocument)
	if err != nil {
		log.Printf("Invalid insertion %s, InsertedId: %d", err, insertResult.InsertedID)
		return "", err
	} else if insertResult.InsertedID == nil {
		return InsertFailMessage, nil
	}
	return InsertSuccessMessage, nil
}
func UpdateOne(collection string, filter interface{}, updatedDocument interface{}) (string, error) {
	updateResult, err := db.Collection(collection).UpdateOne(context.Background(), filter, updatedDocument)
	if err != nil {
		log.Printf("Update err: %s, updatedCount: %d ", err, updateResult.UpsertedCount)
		return "", err
	} else if updateResult.MatchedCount == 0 {
		return UpdateFailMessage, nil
	}
	return UpdateSuccessMessage, nil
}
func FindById(collection string, filter interface{}, results *maze.Maze) error {
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(results)
	if nil != err {
		return err
	}
	return nil
}

func DeleteOne(collection string, filter interface{}) (string, error) {
	deleteResult, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Deletion err: %s, deletetedCount: %d ", err, deleteResult.DeletedCount)
		return "", err
	} else if deleteResult.DeletedCount == 0 {
		return DeleteFailMessage, nil
	}
	return DeleteSuccessMessage, nil
}
