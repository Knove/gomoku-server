package services

import (
	"server/common"
	"server/models"

	"github.com/astaxie/beego"
)

/*
Gomoku 五子棋服务层

*/
type Gomoku struct {
}

/*
Ping 测试链接

*/
func (gomoku *Gomoku) Ping(client *models.Client, traceID string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	beego.Info("webSocket request ping 接口 =>", client.Addr, traceID, message)

	data = "pong"

	return
}
