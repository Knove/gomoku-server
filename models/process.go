package models

import (
	"encoding/json"
	"reflect"
	"server/common"

	"github.com/astaxie/beego"
)

var (
	routers = make(map[string]interface{})
)

/*
ProcessData 数据管道

*/
func ProcessData(client *Client, message []byte) {

	defer func() {
		if r := recover(); r != nil {
			beego.Error("处理数据 stop", r)
			client.SendMsg([]byte("出现错误"))
		}
	}()

	request := &Request{}

	err := json.Unmarshal(message, request)
	if err != nil {
		beego.Error("处理数据 json Unmarshal", err)
		client.SendMsg([]byte("数据不合法"))

		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		beego.Error("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))

		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(requestData, &dataMap)
	if err != nil {
		beego.Error("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))
		return
	}

	// 唯一追踪 ID
	traceID := request.TraceID
	// 请求类型
	infoType := request.InfoType
	// 请求方法
	getType := dataMap["getType"]

	beego.Error(traceID, infoType, getType)

	var (
		code uint64
		msg  string
		data interface{}
	)

	// 这里采取路由 + 自动匹配方法模式控制
	if value, ok := getRouter(infoType); ok {

		immutable := reflect.ValueOf(value)
		ret := immutable.MethodByName(getType)
		// service 中是否存在方法
		if ret.IsValid() {
			reflectArray := ret.Call([]reflect.Value{reflect.ValueOf(client), reflect.ValueOf(traceID), reflect.ValueOf(requestData)})
			code = reflectArray[0].Uint()
			msg = reflectArray[1].String()
			data = reflectArray[2].Interface()
		} else {
			code = common.ParameterIllegal
			beego.Error("路由中方法不存在", client.Addr, "infoType:", infoType)
		}
	} else {
		code = common.ParameterIllegal
		beego.Error("路由不存在", client.Addr, "infoType:", infoType)
	}

	msg = common.GetErrorMessage(code, msg)

	response := NewResponse(traceID, infoType, data, msg, code)

	sendByte, err := json.Marshal(response)
	if err != nil {
		beego.Error("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(sendByte)

	beego.Info("Response send", client.Addr, "traceID", traceID, "code", code)

	return
}

/*
Register 路由注册

*/
func Register(key string, value interface{}) {
	routers[key] = value

	return
}

/*
getRouter 获取注册的路由

*/
func getRouter(key string) (value interface{}, ok bool) {

	value, ok = routers[key]

	return
}
