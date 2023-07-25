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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOpenID(ctx *gin.Context) {
	openID := ctx.GetHeader("X-Tt-OPENID")
	if openID == "" {
		Failure(ctx, fmt.Errorf("openID is empty"))
		return
	}
	Success(ctx, openID)
}

func TextAntidirt(ctx *gin.Context) {
	var textAntidirtReq TextAntidirtReq
	err := ctx.Bind(&textAntidirtReq)
	if err != nil {
		log.Printf("params bind error. err %s", err)
		Failure(ctx, err)
		return
	}

	url := "http://developer.toutiao.com/api/v2/tags/text/antidirt"
	input := AntiInput{
		Tasks: []Task{
			{
				Content: textAntidirtReq.Content,
			},
		},
	}
	body, _ := json.Marshal(input)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("http new request error. err %s", err)
		Failure(ctx, err)
		return
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("call developer.toutiao.com error. err %s", err)
		Failure(ctx, err)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("call developer.toutiao.com error. err %s", err)
		return
	}
	Success(ctx, string(respBody))
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

type TextAntidirtReq struct {
	Content string `json:"content"`
}

type AntiInput struct {
	Tasks []Task `json:"tasks"`
}
type Task struct {
	Content string `json:"content"`
}

type Resp struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}
