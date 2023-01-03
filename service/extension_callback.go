package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pipiguanli/douyincloud_mock/consts"
	Err "github.com/pipiguanli/douyincloud_mock/errors"
	"github.com/pipiguanli/douyincloud_mock/utils"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

func ExtensionCallback(ctx *gin.Context) {
	var req ExtensionCallbackReq

	reqPath := ctx.FullPath()

	err := ctx.Bind(&req)
	if err != nil {
		TemplateFailure(ctx, Err.NewQaError(Err.ParamsResolveErr))
		return
	}

	msgJson := req.Msg
	body := &msgJson
	switch gjson.Get(msgJson, "qa_command").String() {
	case "resp_with_error":
		TemplateFailure(ctx, Err.NewQaError(Err.QaCommandErr))
		return
	case "http_status_code_not_200":
		TemplateFailureWithHttpStatusCode(ctx, 500, Err.NewQaError(Err.QaCommandErr))
		return
	case "panic":
		PanicControl()
	case "timeout":
		time.Sleep(time.Duration(5) * time.Second)
	case "big_json":
		body = &consts.TestDataBigJson
	}

	resp := &ExtensionCallbackResp{
		ErrNo:   0,
		ErrTips: "success",
		Data: &Data{
			Body: body,
		},
		QaExtra: &QaExtra{
			QaPath: &reqPath,
		},
	}

	httpStatusCode := 200
	log.Printf("[QA] response=%+v, httpStatusCode=%+v", utils.ToJsonString(resp), httpStatusCode) // 只有正常返回才打上日志，其他异常返回都没打日志，以后再改吧，要么改 demo，要么改日志中间件
	ctx.JSON(httpStatusCode, resp)
}

type ExtensionCallbackReq struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type ExtensionCallbackResp struct {
	ErrNo   int      `json:"err_no"`
	ErrTips string   `json:"err_tips"`
	Data    *Data    `json:"data"`
	QaExtra *QaExtra `json:"qa_extra"`
}

type Data struct {
	Body *string `json:"body"`
}
