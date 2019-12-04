package common

import (
	"github.com/astaxie/beego"

	"github.com/jinzhu/gorm"
	// Register some standard stuff
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
Connect 数据库连接操作库

*/
type Connect struct {
}

var db *gorm.DB
var errDb error
var c Conf

/*
GetInstance 初始化数据库

*/
func GetInstance() (instance *Connect) {
	instance = &Connect{}
	return
}

/*
Init 初始化数据库

*/
func (connect *Connect) Init() (issucc bool) {

	conf := c.GetConf()
	connConf := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ")/" + conf.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, errDb = gorm.Open("mysql", connConf)

	if errDb != nil {
		beego.Error(errDb)
		return false
	}
	return true
}

/*
GetDB 获取数据库实例

*/
func (connect *Connect) GetDB() (dbCon *gorm.DB) {
	return db
}
