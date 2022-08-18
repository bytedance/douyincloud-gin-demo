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
	"github.com/go-redis/redis/v8"
	"os"
)

var (
	//Redis地址
	redisAddr = ""
	//Redis用户名
	redisUserName = ""
	//Redis密码
	redisPassword = ""
)

type redisComponent struct {
	client *redis.Client
}

func (r *redisComponent) GetName(ctx context.Context, key string) (name string, err error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisComponent) SetName(ctx context.Context, key string, name string) error {
	_, err := r.client.Set(ctx, key, name, 0).Result()
	return err
}

//NewRedisComponent 初始化一个实现了HelloWorldComponent接口的RedisComponent
func NewRedisComponent() *redisComponent {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: redisUserName,
		Password: redisPassword,
		DB:       0, // use default DB
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Printf("redisClient init error. err %s", err)
		panic(fmt.Sprintf("redis init failed. err %s\n", err))
	}
	return &redisComponent{
		client: rdb,
	}
}

//init 项目启动时会从环境变量中获取
func init() {
	redisAddr = os.Getenv("REDIS_ADDRESS")
	redisUserName = os.Getenv("REDIS_USERNAME")
	redisPassword = os.Getenv("REDIS_PASSWORD")
}
