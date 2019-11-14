/**
 * @Desc: Request
 * @Author: Knove
 * @createTime: 2019/11/7 23:18
 * @Email: knove@qq.com
 */

package models

/*
Request 通用请求数据格式

*/
type Request struct {
	TraceID  string      `json:"traceID"`        // 消息的唯一追踪Id
	InfoType string      `json:"infoType"`       // 消息类型
	Data     interface{} `json:"data,omitempty"` // 数据 json
	UserID   string      `json:"userID,omitempty"`
}
