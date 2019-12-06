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
	// middleware
	router.Use(common.CORSMiddleware())

	/* api Router */

	// user
	router.POST("/user/getAllUser", controllers.GetAllUser)

	/* websocket */

	// websocket Router and Register
	router.GET("/ws", controllers.Websocket)

	models.Register("gomoku", &services.Gomoku{})

	router.Run(":7777")
}
