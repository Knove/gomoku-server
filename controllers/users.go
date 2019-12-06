package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"server/models"
)

/*
GetAllUser 获取全部用户数据

*/
func GetAllUser(c *gin.Context) {

	log.WithFields(log.Fields{
		"action": "GetAllUser",
	}).Info("GetAllUser 接口调用")

	var user models.Users

	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}

	response := models.NewAPIResponse(user, "Success", http.StatusOK)

	c.JSON(http.StatusOK, response)
}
