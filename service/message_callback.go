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

func MessageCallback(ctx *gin.Context) {
	var req MessageCallbackReq
	defer func() {
		log.Printf("[QA] request=%+v", utils.ToJsonString(&req))
	}()

	reqPath := ctx.FullPath()
	err := ctx.Bind(&req)
	if err != nil {
		TemplateFailure(ctx, Err.NewQaError(Err.ParamsResolveErr))
		return
	}

	if err := utils.CheckHeaders(ctx); err != nil {
		TemplateFailure(ctx, Err.NewQaError(Err.InvalidParamErr, err.Error()))
		return
	}

	if len(utils.GetHeaderByName(ctx, consts.Header_StressTag)) > 0 {
		// 举例：sleep 1秒
		//time.Sleep(time.Duration(1) * time.Second)

		//sleep 随机 100ms ~ 1000ms（0.1s ~ 1s）
		num := utils.GenerateRandInt(100, 1000)
		time.Sleep(time.Duration(num) * time.Millisecond)
	}

	msgJson := req.Msg
	body := &msgJson
	switch gjson.Get(msgJson, "qa_command").String() {
	case "resp_with_error":
		TemplateFailure(ctx, Err.NewQaError(Err.QaCommandErr, "qa_command = resp_with_error"))
		return
	case "http_status_code_not_200":
		TemplateFailureWithHttpStatusCode(ctx, 500, Err.NewQaError(Err.QaCommandErr, "qa_command = http_status_code_not_200"))
		return
	case "panic":
		PanicControl()
	case "timeout":
		time.Sleep(time.Duration(5) * time.Second)
	case "big_json":
		body = utils.StringPtr(consts.TestDataBigJson)
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

type MessageCallbackReq struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type MessageCallbackResp struct {
	ErrNo   int      `json:"err_no"`
	ErrTips string   `json:"err_tips"`
	Data    *Data    `json:"data"`
	QaExtra *QaExtra `json:"qa_extra"`
}
