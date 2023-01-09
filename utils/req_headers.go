package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pipiguanli/douyincloud_mock/consts"
	Err "github.com/pipiguanli/douyincloud_mock/errors"
)

func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func CheckHeaders(ctx *gin.Context) (err error) {
	if GetHeaderByName(ctx, consts.CommonHead_DataFormat) != consts.CommonHead_DataFormatJSON {
		return Err.NewQaError(Err.InvalidParamErr, fmt.Sprintf("请求头非法 %s", consts.CommonHead_DataFormat))
	}
	if GetHeaderByName(ctx, consts.CommonHead_Charset) != consts.CommonHead_Charset_UTF8 {
		return Err.NewQaError(Err.InvalidParamErr, fmt.Sprintf("请求头非法 %s", consts.CommonHead_Charset))
	}

	return nil
}
