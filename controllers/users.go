package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
GetAllUser 获取全部用户数据

*/
func GetAllUser(c *gin.Context) {

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
