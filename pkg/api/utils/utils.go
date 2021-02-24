package utils

import (
	"github.com/juju/errors"
	"net/http"
)

// APIHTTPError defines api default error output
type APIHTTPError struct {
	Error string `json:"error"`
}

type ReplyCommon struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ErrCodeOK       = 200000 // 成功
	ErrCodeNotFound = 404000 // 没找到
	ErrCodeInternal = 500000 // 内部错误
)

func MakeReply(data interface{}, err error) (int, interface{}) {
	if err == nil {
		return http.StatusOK, ReplyCommon{
			Code: ErrCodeOK,
			Data: data,
		}
	} else if errors.IsNotFound(err) {
		return http.StatusOK, ReplyCommon{
			Code: ErrCodeNotFound,
			Msg:  err.Error(),
		}
	} else {
		return http.StatusInternalServerError, ReplyCommon{
			Code: ErrCodeInternal,
			Msg:  err.Error(),
		}
	}
}
