package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var URI = ""

var DB *mongo.Client

func init()  {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	DB=client
}
func main() {
	UserColl := DB.Database("example").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := UserColl.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err!=nil{
		log.Println("[Error]",err)
		return
	}
	fmt.Println("[Success]",res.InsertedID)
}
