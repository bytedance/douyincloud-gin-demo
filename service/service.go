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
package service

import (
	"douyincloud-gin-demo/component"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	target := ctx.Query("target")
	if target == "" {
		Failure(ctx, fmt.Errorf("param invalid"))
		return
	}
	fmt.Printf("target= %s\n", target)
	hello, err := component.GetComponent(target)
	if err != nil {
		Failure(ctx, fmt.Errorf("param invalid"))
		return
	}

	name, err := hello.GetName(ctx, "name")
	if err != nil {
		Failure(ctx, err)
		return
	}
	Success(ctx, name)
}

func SetName(ctx *gin.Context) {
	var req SetNameReq
	err := ctx.Bind(&req)
	if err != nil {
		Failure(ctx, err)
		return
	}
	hello, err := component.GetComponent(req.Target)
	if err != nil {
		Failure(ctx, fmt.Errorf("param invalid"))
		return
	}
	err = hello.SetName(ctx, "name", req.Name)
	if err != nil {
		Failure(ctx, err)
		return
	}
	Success(ctx, "")
}

func Failure(ctx *gin.Context, err error) {
	resp := &Resp{
		ErrNo:  -1,
		ErrMsg: err.Error(),
	}
	ctx.JSON(200, resp)
}

func Success(ctx *gin.Context, data string) {
	resp := &Resp{
		ErrNo:  0,
		ErrMsg: "success",
		Data:   data,
	}
	ctx.JSON(200, resp)
}

type HelloResp struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Data   string `json:"data"`
}

type SetNameReq struct {
	Target string `json:"target"`
	Name   string `json:"name"`
}

type Resp struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}
