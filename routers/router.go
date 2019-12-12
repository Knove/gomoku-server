/**
 * @Desc: Router
 * @Author: Knove
 * @createTime: 2019/11/8 18:04
 * @Email: knove@qq.com
 */

package routers

import (
	"server/common"
	"server/controllers"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
)

/*
Init 初始化路由

*/
func Init() {
	router := gin.Default()

	/* websocket */

	// websocket Router and Register
	router.GET("/ws", controllers.Websocket)

	models.Register("gomoku", &services.Gomoku{})

	// middleware
	router.Use(common.CORSMiddleware())
	router.Use(common.JWTMiddleware())

	/* api Router */

	// user
	router.POST("/user/login", controllers.Login)           // 登录
	router.POST("/user/getAllUser", controllers.GetAllUser) // 获取全部用户

	router.Run(":7777")
}
