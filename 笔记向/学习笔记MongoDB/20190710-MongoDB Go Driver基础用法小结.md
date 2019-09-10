mongoDB的官方golang支持库，尚处于测试阶段。

以下为示例代码分析总结，实现最基本的CRUD操作，

可供初学者学习参考和引用。

注意：源码并不能直接运行！

<!--more-->

### 安装

推荐使用goland与go mod（go modules ）

官方网站：[**mongo-go-driver**](https://github.com/mongodb/mongo-go-driver), [**go modules**](https://github.com/golang/go/wiki/Modules).

在需要用到的项目中任意一个包import链接，然后右键进行sync操作，便可以导入包，注意设置代理

```go
//example
...
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
...
```

### 初始化

在本地数据库环境下进行，所以请确保已经开启本地的数据库服务，可通过终端运行`mongo`命令进行判断。

```go
var ctx = context.Background() 		// create a new context
var UserColl *mongo.Collection 		// create a new collection
type User struct { 					// user type for example
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username" `
	Psw       string             `bson:"password" `
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// connect database
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))  // change the URI when you need
	if err != nil {
		log.Fatal(err)
	}
	// check connection
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
    	//example coll
	UserColl = client.Database("[db name]").Collection("[coll name]")
}
```

### 增

#### 单个新增

```go
// controllers
func ExampleMongo() {
	// insert one user
	NewUser := User{
		Username: "userA",
		Psw:      "123",
	}
	if res, err := InsertOneUser(NewUser); err != nil {
		return
	} else {
		fmt.Println("new post created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	}
}
    
// models
func InsertOneUser(user User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(ctx, bson.M{
		"username":   user.Username,
		"psw":        user.Psw,
		"created_at": time.Now(),
		"updated_at": time.Now()})
	return res, err
}
```

#### 多个新增

```go
// controllers
func ExampleMongo()  {
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
}

// models
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
```



### 删

#### 单个删除

```go
// controllers
func ExampleMongo()  {
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
}

// models
func DeleteOneUser(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteOne(ctx, filter)
	return res, err
}
```



#### 多个删除

```go
// controllers
func ExampleMongo()  {
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
}

// models
func DeleteManyUsers(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteMany(ctx, filter)
	return res, err
}
```



### 查

#### 单个查找

```go
// controllers
func ExampleMongo()  {
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
}

// models
func FindOneUser(filter bson.M) (User, error) {
	Msg := User{}
	result := UserColl.FindOne(ctx, filter)
	err := result.Decode(&Msg)
	return Msg, err
}
```



#### 多个查找

```go
// controllers
func ExampleMongo()  {
	// find many user
	filterFindMany := bson.M{"[key]": "[value]"}
	if res, err := FindManyUsers(filterFindMany); err != nil {
		return
	} else {
		fmt.Printf("post: %+v\n", res)
	}
}

// models
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
```



### 改

#### 单个更改

```go
// controllers
func ExampleMongo()  {
	// update one user
	filter := bson.M{"[key]": "[value]"}
	update := bson.M{"[key]": "[value]"}
	if res, err := UpdateOneUser(filter, update); err != nil {
		return
	} else {
		fmt.Printf("modified count: %d\n", res.ModifiedCount)
	}
}

// models
func UpdateOneUser(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(ctx, filter, update)
	return res, err
}
```



#### 多个更改

```go
// controllers
func ExampleMongo()  {
	// update many users
    filter := bson.M{"[key]": "[value]"}
	update := bson.M{"[key]": "[value]"}
	if res, err := UpdateManyUsers(filter, update); err != nil {
		return
	} else {
		fmt.Printf("modified count: %d\n", res.ModifiedCount)
	}
}

// models
func UpdateManyUsers(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateMany(ctx, filter, update)
	return res, err
}
```

### 示例源码

Github: [MongoDB Go Driver](https://github.com/WillXu24/study-space/blob/master/MongoDB%20Go%20Driver/MongoDB.go)

### 参考文档

[Using the official MongoDB Go driver](https://vkt.sh/go-mongodb-driver-cookbook/)