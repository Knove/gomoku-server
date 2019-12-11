package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"server/models"
	"server/services"
)

/*
GetAllUser 获取全部用户数据

*/
func GetAllUser(c *gin.Context) {

	log.WithFields(log.Fields{
		"action": "GetAllUser",
	}).Info("GetAllUser 接口调用")

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Error(err)
	}

	users, err := services.GetAll()
	if err != nil {
		log.Error(err)
	}
	response := models.NewAPIResponse(users, "Success", http.StatusOK)

	c.JSON(http.StatusOK, response)
}
