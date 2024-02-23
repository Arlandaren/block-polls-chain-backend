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

func GetLastBlock() *Block {
	opts := options.FindOne().SetSort(bson.D{{"Index", -1}})

	var lastblock Block
	if err := Collection.FindOne(context.Background(), bson.D{}, opts).Decode(&lastblock); err != nil {
		fmt.Println("error while finding the last block:", err)
	}
	return &lastblock
}

func InsertBlockIntoDb(block *Block) error {
	data := bson.M{
		"Index":        block.Index,
		"PreviousHash": block.PreviousHash,
		"Timestamp":    block.Timestamp,
		"Hash":         block.Hash,
		"Data":         block.Data,
		"Owner":        block.Owner,
	}
	_, err := Collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func FindByHash(hash string) (*Block, error) {
    var block Block
	if err := Collection.FindOne(context.Background(), bson.M{"Hash": hash}).Decode(&block); err != nil {
		return nil, err
	}
    return &block, nil
}
