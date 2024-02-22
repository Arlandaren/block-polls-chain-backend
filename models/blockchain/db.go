package blockchain

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func InitDB() {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        panic(err)
    }
    // defer client.Disconnect(context.Background())

    err = client.Ping(context.Background(), nil)
    if err != nil {
        panic(err)
    }

    database := client.Database("blockchain")
    Collection = database.Collection("blocks")
}

func GetLastBlock() *Block{
    opts := options.FindOne().SetSort(bson.D{{"Index", -1}})

    var lastblock Block
    if err := Collection.FindOne(context.Background(), bson.D{}, opts).Decode(&lastblock); err != nil {
        fmt.Println("error while finding the last block:", err)
    }

    fmt.Println("Last inserted document:", lastblock)
    return &lastblock

}