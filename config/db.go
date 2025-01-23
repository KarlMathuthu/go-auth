package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Replace with your MongoDB URI
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
    DB = client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    return client.Database("gin_auth_db").Collection(collectionName)
}