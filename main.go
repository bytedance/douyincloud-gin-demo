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
	"douyincloud-gin-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/get_open_id", service.GetOpenID)
	r.POST("/api/text/antidirt", service.TextAntidirt)

	r.Run(":8000")
}
