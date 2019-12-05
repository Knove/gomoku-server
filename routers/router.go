/**
 * @Desc: Router
 * @Author: Knove
 * @createTime: 2019/11/8 18:04
 * @Email: knove@qq.com
 */

package routers

import (
	"server/controllers"
	"server/models"
	"server/services"

	"github.com/astaxie/beego"
)

func init() {

	// websocket Router and Register
	beego.Router("/user", &controllers.UsersController{})
	beego.Router("/ws", &controllers.WebsocketController{})

	models.Register("gomoku", &services.Gomoku{})
}
