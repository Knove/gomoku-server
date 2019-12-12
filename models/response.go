/**
 * @Desc: Reponse
 * @Author: Knove
 * @createTime: 2019/11/7 23:18
 * @Email: knove@qq.com
 */

package models

import (
	"server/common"
)

/*
Response 通用返回数据格式

*/
type Response struct {
	TraceID  string      `json:"traceId"`        // 消息的唯一追踪Id
	InfoType string      `json:"infoType"`       // 消息类型
	Data     interface{} `json:"data,omitempty"` // 数据 json
}

/*
Data 通用返回数据格式

*/
type Data struct {
	Code uint64      `json:"code"` // code 值
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据 json
}

/*
NewResponse 创建 websocket 返回数据实例

*/
func NewResponse(traceID string, infoType string, data interface{}, msg string, code uint64) *Response {
	backData := Data{Code: code, Msg: msg, Data: data}
	return &Response{TraceID: traceID, InfoType: infoType, Data: backData}
}

/*
NewAPIResponse 创建标准接口实例

*/
func NewAPIResponse(data interface{}, msg string, code uint64) *Data {
	backData := Data{Code: code, Msg: msg, Data: data}
	return &backData
}

/*
ErrorResponse 创建错误返回接口实例

*/
func ErrorResponse(code uint64) *Data {
	backData := Data{Code: code, Msg: common.GetErrorMessage(code, ""), Data: nil}
	return &backData
}
