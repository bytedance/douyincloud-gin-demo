package errors

import (
	"fmt"
	"strings"
)

type QaError struct {
	ErrNo   int64  `json:"err_no"`
	ErrTips string `json:"err_tips"`
}

func (e *QaError) Error() string {
	return fmt.Sprintf("errNo:%d, errTips:%s", e.ErrNo, e.ErrTips)
}

func (e *QaError) GetError() error {
	return fmt.Errorf(e.ErrTips)
}

func NewQaError(code int64, msg ...string) *QaError {
	if code == 0 {
		return &QaError{
			ErrNo:   code,
			ErrTips: "success",
		}
	}
	return &QaError{
		ErrNo:   code,
		ErrTips: fmt.Sprintf(QaErrorMap[code], strings.Join(msg, "|")),
	}
}

const (
	SuccessCode = 0

	// 参数错误 10000-10999
	InvalidParamErr  = 10000
	ParamsResolveErr = 10002

	// 内部系统错误（db redis tcc rpc ... 错误） 13000-13999
	SystemErr = 13000

	// 控制型错误
	QaCommandErr = 77777
	TccContolErr = 88888

	// sdk错误
	SdkErr = 99999
)

var (
	QaErrorMap = map[int64]string{
		SuccessCode:      "success",
		InvalidParamErr:  "参数错误",
		ParamsResolveErr: "参数解析异常，请注意参数格式",
		SystemErr:        "系统错误，请重试",
		TccContolErr:     "通过TCC的配置项来控制本次响应返回异常", // （注意：并非是TCC本身出现异常）
		SdkErr:           "调用sdk方法返回了错误",
	}
)