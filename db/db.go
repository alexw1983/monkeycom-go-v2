package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/alexw1983/monkeycom-go-v2/config"
	"github.com/alexw1983/monkeycom-go-v2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	GetCharacter(userId string, slug string) models.Character
	GetCharacters(userId string) []models.Character
}

type MongoDbStorage struct {
	connectionString string
	dbName           string
}

func NewDB() (DB, error) {
	if config.Envs.DbConnectionString == "" || config.Envs.DbName == "" {
		return nil, errors.New("config is not set")
	}

	return &MongoDbStorage{connectionString: config.Envs.DbConnectionString, dbName: config.Envs.DbName}, nil
}

// const DB_NAME = "MONKEYCOM_GO_DB"
const COLLECTION_NAME = "Characters"

func (storage *MongoDbStorage) GetCharacter(userId string, slug string) models.Character {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(storage.connectionString))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database(storage.dbName).Collection(COLLECTION_NAME)

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "UserId", Value: userId}},
				bson.D{{Key: "Slug", Value: slug}},
			},
		},
	}

	var result models.Character
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Character{}
		} else {
			panic(err)
		}
	}

	return result
}

func (storage *MongoDbStorage) GetCharacters(userId string) []models.Character {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client,
		err := mongo.Connect(ctx, options.Client().ApplyURI(storage.connectionString))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database(storage.dbName).Collection(COLLECTION_NAME)
	cursor,
		err2 := coll.Find(ctx, bson.D{{Key: "UserId", Value: userId}})
	if err2 != nil {
		log.Fatal(err)
	}

	var results []models.Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}
