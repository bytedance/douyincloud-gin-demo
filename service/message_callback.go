package service

import (
	"github.com/gin-gonic/gin"
	Err "github.com/pipiguanli/douyincloud_mock/errors"
	"github.com/pipiguanli/douyincloud_mock/utils"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

func MessageCallback(ctx *gin.Context) {
	var req *MessageCallbackReq
	reqPath := ctx.FullPath()

	err := ctx.Bind(req)
	if err != nil {
		TemplateFailure(ctx, Err.NewQaError(Err.ParamsResolveErr))
		return
	}
	msgJson := req.Msg
	switch gjson.Get(msgJson, "qa_command").String() {
	case "resp_with_error":
		TemplateFailure(ctx, Err.NewQaError(Err.QaCommandErr, "qa_command = resp_with_error"))
		return
	case "http_status_code_not_200":
		TemplateFailureWithHttpStatusCode(ctx, 500, Err.NewQaError(Err.QaCommandErr, "qa_command = http_status_code_not_200"))
		return
	case "panic":
		panic("qa_command:panic")
	case "timeout":
		time.Sleep(time.Duration(5) * time.Second)
	}

	resp := &ExtensionCallbackResp{
		ErrNo:   0,
		ErrTips: "success",
		QaExtra: &QaExtra{
			QaPath: &reqPath,
		},
	}
	log.Printf("[QA] request=%+v", utils.ToJsonString(req))   // 只有正常返回才打上日志，其他异常返回都没打日志，以后再改吧，要么改 demo，要么改日志中间件
	log.Printf("[QA] response=%+v", utils.ToJsonString(resp)) // 只有正常返回才打上日志，其他异常返回都没打日志，以后再改吧，要么改 demo，要么改日志中间件

	ctx.JSON(200, resp)
}

type MessageCallbackReq struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type MessageCallbackResp struct {
	ErrNo   int      `json:"err_no"`
	ErrTips string   `json:"err_tips"`
	QaExtra *QaExtra `json:"qa_extra"`
}
