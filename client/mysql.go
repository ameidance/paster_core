package client

import (
	"strconv"

	"github.com/ameidance/paster_core/conf"
	"github.com/bytedance/gopkg/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBClient *gorm.DB
)

func InitDB() {
	dbConf, err := conf.GetDBConfig()
	if dbConf == nil || err != nil {
		panic(err)
	}
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Hostname + ":" + strconv.Itoa(dbConf.Port) +
		")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	if DBClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		logger.Errorf("[InitDB] connect db failed. err:%v", err)
		panic(err)
	}
}
