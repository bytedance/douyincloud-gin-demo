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
package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/pipiguanli/douyincloud_mock/component"
	"github.com/pipiguanli/douyincloud_mock/service"
	"log"
)

func main() {
	//component.InitComponents()
	r := gin.Default()
	r.Use(gin.Logger()) // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	//r.Use(gin.Recovery()) // Recovery middleware recovers from any panics and writes a 500 if there was one.

	r.POST("/api/douyincloud/dev/extension_callback", service.ExtensionCallback)
	r.POST("/api/douyincloud/prod/extension_callback", service.ExtensionCallback)
	r.POST("/api/douyincloud/dev/message_callback", service.MessageCallback)
	r.POST("/api/douyincloud/prod/message_callback", service.MessageCallback)

	err := r.Run(":8000")
	if err != nil {
		log.Println(err.Error())
	}
}
