package common

import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"

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

	db, errDb = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pwd, conf.Host, conf.Dbname))

	if errDb != nil {
		log.Error(errDb)
		return false
	}
	
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return true
}

/*
GetDB 获取数据库实例

*/
func (connect *Connect) GetDB() (dbCon *gorm.DB) {
	return db
}

/*
CloseDB 关闭数据库

*/
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateDate"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateDate"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateDate", time.Now().Unix())
	}
}