package service

import (
	"github.com/gin-gonic/gin"
	Err "github.com/pipiguanli/douyincloud_mock/errors"
	"github.com/pipiguanli/douyincloud_mock/utils"
	"log"
)

func Ping(ctx *gin.Context) {
	reqPath := ctx.FullPath()

	code := int64(0)
	resp := &TemplateResp{
		ErrNo:   code,
		ErrTips: Err.QaErrorMap[code],
		QaExtra: &QaExtra{
			QaPath: &reqPath,
		},
	}
	ctx.JSON(200, resp)
}

func PanicControl() {
	log.Printf("[QA] 请求中解析到 qa_command:panic 指令，服务随后会 panic, 本次请求不会返回响应与响应日志")
	panic("qa_command:panic")
}

func TemplateFailure(ctx *gin.Context, err *Err.QaError) {
	reqPath := ctx.FullPath()

	resp := &TemplateResp{
		ErrNo:   err.ErrNo,
		ErrTips: err.ErrTips,
		QaExtra: &QaExtra{
			QaPath: &reqPath,
		},
	}
	httpStatusCode := 200
	log.Printf("[QA] response=%+v, httpStatusCode=%+v, err=%+v", utils.ToJsonString(resp), httpStatusCode, err.Error())
	ctx.JSON(httpStatusCode, resp)
}

func TemplateFailureWithHttpStatusCode(ctx *gin.Context, httpStatusCode int, err *Err.QaError) {
	reqPath := ctx.FullPath()

	resp := &TemplateResp{
		ErrNo:   err.ErrNo,
		ErrTips: err.ErrTips,
		QaExtra: &QaExtra{
			QaPath: &reqPath,
		},
	}
	log.Printf("[QA] response=%+v, httpStatusCode=%+v, err=%+v", utils.ToJsonString(resp), httpStatusCode, err.Error())
	ctx.JSON(httpStatusCode, resp)
}

type TemplateResp struct {
	ErrNo   int64       `json:"err_no"`
	ErrTips string      `json:"err_tips"`
	Data    interface{} `json:"data"`
	QaExtra *QaExtra    `json:"qa_extra"`
}

type QaExtra struct {
	QaPath *string `json:"qa_path"`
}
