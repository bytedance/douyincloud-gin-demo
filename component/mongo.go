/*
Copyright (year) Bytedance Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package component

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"os"
	"time"
)

var (
	// mongo地址
	mongoAddr = ""
	// mongo用户名
	mongoUserName = ""
	// mongo密码
	mongoPassWord = ""
)

const collectionName = "demo"

type mongoComponent struct {
	client   *mongo.Client
	dataBase string
}

type model struct {
	Key   string `bson:"key"`   //类型
	Value string `bson:"value"` //值
}

func (m *mongoComponent) GetName(ctx context.Context, key string) (name string, err error) {

	coll := m.client.Database(m.dataBase).Collection(collectionName)
	doc := &model{}

	filter := bson.M{"key": key}
	result := coll.FindOne(ctx, filter)

	if err := result.Decode(doc); err != nil {
		return "", err
	}
	return doc.Value, err
}

func (m *mongoComponent) SetName(ctx context.Context, key string, name string) error {

	coll := m.client.Database(m.dataBase).Collection(collectionName)

	filter := bson.M{"key": key}
	update := bson.M{"$set": model{Key: key, Value: name}}
	_, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

//NewMongoComponent 新建一个mongodbComponent，其实现了HelloWorldComponent接口
func NewMongoComponent() *mongoComponent {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tmp, err := url.Parse(mongoAddr)
	if err != nil {
		panic("mongo addr parse error")
	}
	authSource := tmp.Query().Get("authSource")
	credential := options.Credential{
		AuthSource: authSource,
		Username:   mongoUserName,
		Password:   mongoPassWord,
	}
	mongoUrl := fmt.Sprintf("mongodb://%s", mongoAddr)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetAuth(credential))
	if err != nil {
		fmt.Printf("mongoClient init error. err %s\n", err)
		panic("mongo connect error")
	}

	dataBase := "demo"
	doc := &model{
		Key:   "name",
		Value: Mongo,
	}
	_, err = client.Database(dataBase).Collection(collectionName).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("mongoClient init error. err %s\n", err)
		panic("mongo init error")
	}
	return &mongoComponent{client, dataBase}
}

//init 项目启动时，会从环境变量中获取mongodb的地址，用户名和密码
func init() {
	mongoAddr = os.Getenv("MONGO_ADDRESS")
	mongoUserName = os.Getenv("MONGO_USERNAME")
	mongoPassWord = os.Getenv("MONGO_PASSWORD")
}
