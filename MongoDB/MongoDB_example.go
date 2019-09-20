package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var ctx = context.Background() // create a new context
var UserColl *mongo.Collection // create a new collection

type User struct { //user type
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username" `
	Psw       string             `bson:"password" `
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// connect database
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// check connection
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	UserColl = client.Database("[db name]").Collection("[coll name]")
}

func ExampleMongo() {
	// insert one user
	NewUser := User{
		Username: "userA",
		Psw:      "123",
	}
	if res, err := InsertOneUser(NewUser); err != nil {
		return
	} else {
		fmt.Printf("new post created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	}

	// insert many users
	NewUsers := make([]User, 1)
	NewUsers = append(NewUsers, User{
		Username: "userB",
		Psw:      "123",
	}, User{
		Username: "userC",
		Psw:      "123",
	})
	if res, err := InsertManyUsers(NewUsers); err != nil {
		return
	} else {
		fmt.Printf("inserted ids: %v\n", res.InsertedIDs)
	}

	// delete one user
	// you can design your own filter
	id, err := primitive.ObjectIDFromHex("[user id]")
	if err != nil {
		return
	}
	filterDeleteOne := bson.M{"_id": id}
	if res, err := DeleteOneUser(filterDeleteOne); err != nil {
		return
	} else {
		fmt.Printf("deleted count: %d\n", res.DeletedCount)
	}

	// delete many users
	// delete documents created older than 2 days
	filterDeleteMany := bson.M{"created_at": bson.M{
		"$lt": time.Now().Add(-2 * 24 * time.Hour),
	}}
	if res, err := DeleteManyUsers(filterDeleteMany); err != nil {
		return
	} else {
		fmt.Printf("deleted count: %d\n", res.DeletedCount)
	}

	// find one user
	id, err = primitive.ObjectIDFromHex("[user id]")
	if err != nil {
		return
	}
	filterFindOne := bson.M{"_id": id}
	if res, err := FindOneUser(filterFindOne); err != nil {
		return
	} else {
		fmt.Printf("post: %+v\n", res)
	}

	// find many user
	filterFindMany := bson.M{"[key]": "[value]"}
	if res, err := FindManyUsers(filterFindMany); err != nil {
		return
	} else {
		fmt.Printf("post: %+v\n", res)
	}

	// update one user
	filter := bson.M{"[key]": "[value]"}
	update := bson.M{"[key]": "[value]"}
	if res, err := UpdateOneUser(filter, update); err != nil {
		return
	} else {
		fmt.Printf("modified count: %d\n", res.ModifiedCount)
	}

	// update many users
	if res, err := UpdateManyUsers(filter, update); err != nil {
		return
	} else {
		fmt.Printf("modified count: %d\n", res.ModifiedCount)
	}
}

// insert one
func InsertOneUser(user User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(ctx, bson.M{
		"username":   user.Username,
		"psw":        user.Psw,
		"created_at": time.Now(),
		"updated_at": time.Now()})
	return res, err
}

// insert many
func InsertManyUsers(users []User) (*mongo.InsertManyResult, error) {
	res, err := UserColl.InsertMany(ctx, []interface{}{
		bson.M{
			"username":   users[1].Username,
			"psw":        users[1].Psw,
			"created_at": time.Now(),
			"updated_at": time.Now()},
		bson.M{
			"username":   users[2].Username,
			"psw":        users[2].Psw,
			"created_at": time.Now(),
			"updated_at": time.Now()}})
	return res, err
}

// delete one
func DeleteOneUser(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteOne(ctx, filter)
	return res, err
}

// delete many
func DeleteManyUsers(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteMany(ctx, filter)
	return res, err
}

// find one
func FindOneUser(filter bson.M) (User, error) {
	Msg := User{}
	result := UserColl.FindOne(ctx, filter)
	err := result.Decode(&Msg)
	return Msg, err
}

// find many
func FindManyUsers(filter bson.M) ([]User, error) {
	cursor, err := UserColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []User
	for cursor.Next(ctx) {
		var p User
		// decode the document into given type
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

// update one
func UpdateOneUser(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(ctx, filter, update)
	return res, err
}

// update many
func UpdateManyUsers(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateMany(ctx, filter, update)
	return res, err
}
