package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type con struct {
	Host          string
	Port          string
	User          string
	Pass          string
	AuthSource    string
	AuthMechanism string
}

type mgo struct {
	database   string
	collection string
}

const uri = "mongodb://tit:tit@127.0.0.1:27017/"

var client *mongo.Client

func init() {
	if err := conn(); err != nil {
		fmt.Println("connect to mongodb failed! ", err)
	}
}

// connect mongodb
func conn() error {
	// https://github.com/rosspatil/Golang-Mongo-Driver-Example/blob/master/mongo-exmple.go
	// credential := options.Credential{
	//   AuthMechanism: "SCRAM-SHA-256",
	// 	AuthSource: "tit",
	// 	Username:   "<username>",
	// 	Password:   "<password>",
	// }
	// client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential).SetMaxPoolSize(20))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // connect pool
	if err != nil {
		return err
	}

	return client.Ping(ctx, readpref.Primary())
}

func New(database, collection string) *mgo {
	return &mgo{
		database, collection,
	}
}

// 查询单个
func (m *mgo) FindOne(key string, value interface{}) *mongo.SingleResult {
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	//collection.
	filter := bson.D{{key, value}}
	return collection.FindOne(context.TODO(), filter)
}

// 查询多个
func (m *mgo) FindMany(mp map[string]any) (*mongo.Cursor, error) {
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()

	var filter bson.D
	for k, v := range mp {
		filter = append(filter, bson.E{k, v})
	}
	if len(filter) == 0 {
		filter = bson.D{{"foo", nil}}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return cursor, err

	// var bs []bson.M

	// if err = cursor.All(context.TODO(), &bs); err != nil {
	// 	return nil, err
	// }

	// return bs, nil
}

//插入单个
func (m *mgo) InsertOne(value interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(m.database).Collection(m.collection)
	return collection.InsertOne(context.TODO(), value)
}

//插入多个
func (m *mgo) InsertMany(values []interface{}) (*mongo.InsertManyResult, error) {
	collection := client.Database(m.database).Collection(m.collection)
	return collection.InsertMany(context.TODO(), values)
}

//查询集合里有多少数据
func (m *mgo) CollectionCount() (int64, error) {
	collection := client.Database(m.database).Collection(m.collection)
	//name := collection.Name()
	return collection.EstimatedDocumentCount(context.TODO())
}

//按选项查询集合 Skip 跳过 Limit 读取数量 sort 1 ，-1 . 1 为最初时间读取 ， -1 为最新时间读取
func (m *mgo) CollectionDocuments(Skip, Limit int64, sort int) (*mongo.Cursor, error) {
	collection := client.Database(m.database).Collection(m.collection)
	SORT := bson.D{{"_id", sort}} //filter := bson.D{{key,value}}
	filter := bson.D{{}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	//findOptions.SetLimit(i)
	return collection.Find(context.Background(), filter, findOptions)
}

//获取集合创建时间和编号
func (m *mgo) ParsingId(result string) (time.Time, uint64) {
	temp1 := result[:8]
	timestamp, _ := strconv.ParseInt(temp1, 16, 64)
	dateTime := time.Unix(timestamp, 0) //这是截获情报时间 时间格式 2019-04-24 09:23:39 +0800 CST
	temp2 := result[18:]
	count, _ := strconv.ParseUint(temp2, 16, 64) //截获情报的编号
	return dateTime, count
}

//删除和查询
func (m *mgo) DeleteAndFind(key string, value interface{}) (int64, *mongo.SingleResult, error) {
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	DeleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	return DeleteResult.DeletedCount, singleResult, err
}

//删除
func (m *mgo) Delete(key string, value interface{}) (int64, error) {
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	return count.DeletedCount, err
}

//删除多个
func (m *mgo) DeleteMany(key string, value interface{}) (int64, error) {
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{{key, value}}

	count, err := collection.DeleteMany(context.TODO(), filter)
	return count.DeletedCount, err
}
