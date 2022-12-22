package service

import (
	"github.com/gin-gonic/gin"
	Err "github.com/pipiguanli/douyincloud_mock/errors"
	"github.com/tidwall/gjson"
	"time"
)

func MessageCallback(ctx *gin.Context) {
	var req MessageCallbackReq

	err := ctx.Bind(&req)
	if err != nil {
		TemplateFailure(ctx, Err.NewQaError(Err.ParamsResolveErr))
		return
	}
	msgJson := req.Msg
	switch gjson.Get(msgJson, "qa_command").String() {
	case "resp_with_error":
		TemplateFailure(ctx, Err.NewQaError(Err.QaCommandErr))
		return
	case "http_status_code_not_200":
		TemplateFailureWithHttpStatusCode(ctx, 500, Err.NewQaError(Err.QaCommandErr))
		return
	case "panic":
		panic("qa_command:panic")
	case "timeout":
		time.Sleep(time.Duration(5) * time.Second)
	}

	resp := &ExtensionCallbackResp{
		ErrNo:   0,
		ErrTips: "success",
	}
	ctx.JSON(200, resp)
}

type MessageCallbackReq struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type MessageCallbackResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}
