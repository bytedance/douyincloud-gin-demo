package service

import (
	"github.com/gin-gonic/gin"
	common_errors "github.com/pipiguanli/douyincloud_mock/errors"
)

func TemplateFailure(ctx *gin.Context, err *common_errors.QaError) {
	resp := &TemplateResp{
		ErrNo:   err.ErrNo,
		ErrTips: err.ErrTips,
	}
	ctx.JSON(200, resp)
}

func TemplateFailureWithHttpStatusCode(ctx *gin.Context, httpStatusCode int, err *common_errors.QaError) {
	resp := &TemplateResp{
		ErrNo:   err.ErrNo,
		ErrTips: err.ErrTips,
	}
	ctx.JSON(httpStatusCode, resp)
}

type TemplateResp struct {
	ErrNo   int64       `json:"err_no"`
	ErrTips string      `json:"err_tips"`
	Data    interface{} `json:"data"`
}
