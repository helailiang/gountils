/*
@Created by : 2019/12/11 14:39
@Author : 何来亮
@Descripition :
*/
package app

import (
	"fiveCity/logger"
	"fiveCity/pkg/pub"
	"fiveCity/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type Gin struct {
	C *gin.Context
}

type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode string, data interface{}) {
	res := Response{
		Code: errCode,
		Msg:  utils.GetErrMsg(errCode),
		Data: data,
	}
	logger.Log.Infof("返回值:%+v\n", res)
	g.C.JSON(httpCode, res)
	return
}

// NewError example
func (g *Gin) NewError(status int, err error) {
	res := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	logger.Log.Infof("返回值:%+v\n", res)
	g.C.JSON(status, res)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"msg" example:"status bad request"`
}

type JsonTime time.Time


//MarshalJSON jsonTime序列化调用的方法
func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	//当返回时间为空时，需特殊处理
	if time.Time(jsonTime).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(jsonTime).Format(pub.LocalDateTimeFormat) + `"`), nil
}
func (jsonTime JsonTime) UnmarshalJSON(bin []byte) error {
	now, err := time.ParseInLocation(`"`+pub.LocalDateTimeFormat+`"`, string(bin), time.Local)
	jsonTime = JsonTime(now)
	return err

	return nil
}
