package client

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ameidance/paster_core/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBClient *gorm.DB
)

func InitDB() {
	conf, err := getDBConfig()
	if conf == nil || err != nil {
		panic(err)
	}
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Hostname + ":" + strconv.Itoa(conf.Port) +
		")/" + conf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	if DBClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		klog.Errorf("[InitDB] connect db failed. err:%v", err)
		panic(err)
	}

	migrator := DBClient.Migrator()
	if !migrator.HasTable("post") && !migrator.HasTable("comment") {
		klog.Info("[InitDB] migrating...")
		dbScript, err := getDBScript()
		if err != nil {
			panic(err)
		}
		sqls := strings.Split(dbScript, ";")
		for _, sql := range sqls {
			if sql = strings.Trim(sql, "\n"); len(sql) > 0 {
				if err = DBClient.Exec(sql).Error; err != nil {
					panic(err)
				}
			}
		}
	}
}

type dbConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

func getDBConfig() (*dbConf, error) {
	conf := new(dbConf)
	file, err := ioutil.ReadFile(constant.DB_CONF_PATH)
	if err != nil {
		klog.Errorf("[getDBConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getDBConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}

func getDBScript() (string, error) {
	file, err := ioutil.ReadFile(constant.DB_SCRIPT_PATH)
	if err != nil {
		klog.Errorf("[getDBScript] open file failed. err:%v", err)
		return "", err
	}
	return string(file), nil
}
