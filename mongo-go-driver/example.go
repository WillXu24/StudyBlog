package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

/*****************************************初始化数据库****************************************************/
// 数据库参数
const (
	DatabaseURI  = "mongodb://localhost:27017"
	DatabaseName = "example"
)

// 声明全局用户集合
var UserColl *mongo.Collection

func init() {
	// 连接数据库
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DatabaseURI))
	if err != nil {
		log.Fatal(err)
	}
	// 测试连通性
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// 打开/创建数据库
	db := client.Database(DatabaseName)
	// 创建集合
	UserColl = db.Collection("user")
}

/*****************************************基本CRUD操作****************************************************/
// 用户结构体
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username" `
	Password  string             `bson:"password" `
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// 新增一个记录
func UserInsertOne(msg User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(context.Background(), bson.M{
		"username":   msg.Username,
		"password":   msg.Password,
		"created_at": time.Now(),
		"updated_at": time.Now()})
	return res, err
}

// 新增多个记录（不太好用
func UserInsertMany(msg []User) (*mongo.InsertManyResult, error) {
	res, err := UserColl.InsertMany(context.Background(), []interface{}{
		bson.M{
			"username":   msg[0].Username,
			"password":   msg[0].Password,
			"created_at": time.Now(),
			"updated_at": time.Now()},
		bson.M{
			"username":   msg[1].Username,
			"password":   msg[1].Password,
			"created_at": time.Now(),
			"updated_at": time.Now()}})
	return res, err
}

// 查找一个记录
func UserFindOne(filter bson.M) (User, error) {
	var msg User
	err := UserColl.FindOne(context.Background(), filter).Decode(&msg)
	return msg, err
}

// 查找多个记录
func UserFindMany(filter bson.M, options *options.FindOptions) ([]User, error) {
	ctx := context.Background()
	cursor, err := UserColl.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var res []User
	for cursor.Next(ctx) {
		var temp User
		if err := cursor.Decode(&temp); err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}

// 更新一个记录
func UserUpdateOne(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(context.Background(), filter, update, options)
	return res, err
}

// 更新多个记录
func UserUpdateMany(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateMany(context.Background(), filter, update, options)
	return res, err
}

// 删除一个记录
func UserDeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteOne(context.Background(), filter)
	return res, err
}

// 删除多个记录
func UserDeleteMany(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteMany(context.Background(), filter)
	return res, err
}

/*****************************************调用CRUD操作****************************************************/

func main() {
	// 新增一个用户
	user := User{Username: "userA", Password: "123456"}
	if res, err := UserInsertOne(user); err != nil {
		log.Println(err)
	} else {
		log.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
	}
	// 新增多个用户
	users := []User{{Username: "userB", Password: "123456"}, {Username: "userC", Password: "123456"}}
	if res, err := UserInsertMany(users); err != nil {
		log.Println(err)
	} else {
		log.Printf("inserted ids: %v\n", res.InsertedIDs)
	}
	// 查找一个用户
	filter := bson.M{"username": "userA"}
	if res, err := UserFindOne(filter); err != nil {
		log.Println(err)
	} else {
		log.Printf("user: %+v\n", res)
	}
	// 查找多个用户
	filter = bson.M{"password": "123456"}
	if res, err := UserFindMany(filter, nil); err != nil {
		log.Println(err)
	} else {
		log.Printf("user: %+v\n", res)
	}
	// 更新一个用户
	filter = bson.M{"username": "userA"}
	update := bson.M{"username": "userA++"}
	if res, err := UserUpdateOne(filter, bson.M{"$set":update}, nil); err != nil {
		log.Println(err)
	} else {
		log.Printf("modified count: %d\n", res.ModifiedCount)
	}
	// 更新多个用户
	filter = bson.M{"password": "123456"}
	update = bson.M{"password": "654321"}
	if res, err := UserUpdateMany(filter, bson.M{"$set":update}, nil); err != nil {
		log.Println(err)
	} else {
		log.Printf("modified count: %d\n", res.ModifiedCount)
	}
	// 删除一个用户
	filter = bson.M{"username": "userA++"}
	if res, err := UserDeleteOne(filter); err != nil {
		log.Println(err)
	} else {
		log.Printf("delete count: %d\n", res.DeletedCount)
	}
	// 删除多个用户
	filter = bson.M{"password": "654321"}
	if res, err := UserDeleteMany(filter); err != nil {
		log.Println(err)
	} else {
		log.Printf("delete count: %d\n", res.DeletedCount)
	}
}
