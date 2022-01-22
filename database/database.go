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

/*clientOptions := options.Client().
ApplyURI("mongodb+srv://<username>:<password>@cluster0.f3wv0.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
log.Fatal(err)
}
*/
func Init() error {
	dbUri := config.Get().Servers.DB.Uri
	clientOptions := options.Client()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(20)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri), clientOptions)
	/*	databases, _ := client.ListDatabaseNames(ctx, bson.M{})
		log.Printf("%v", databases)
		collections, _ := client.Database("green").ListCollectionNames(ctx, bson.M{})
		log.Printf("%v", collections)*/
	if err != nil {
		return err
	}
	db = client.Database(config.Get().Servers.DB.Name)
	return nil
}

func Aggregate(collection string, filter interface{}, results interface{}) error {
	cursor, err := db.Collection(collection).Aggregate(context.TODO(), filter)
	if nil != err {
		log.Printf("DB Aggregate error %v\n", err)
		return err
	}

	err = cursor.All(context.TODO(), results)
	if nil != err {
		log.Printf("DB Aggregate error %v\n", err)
		return err
	}
	return nil

}

func FindById(collection string, filter interface{}, results *maze.MongoMaze) error {
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(results)
	if nil != err {
		return err
	}
	return nil

}
