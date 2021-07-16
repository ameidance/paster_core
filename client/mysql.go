package client

import (
    "io/ioutil"
    "strconv"

    "github.com/bytedance/gopkg/util/logger"
    "gopkg.in/yaml.v3"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

const (
    _DB_CONF_PATH = "conf/mysql.yml"
)

type DBConf struct {
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Hostname string `yaml:"hostname"`
    Port     int    `yaml:"port"`
    Name     string `yaml:"name"`
}

var (
    DBClient *gorm.DB
)

func InitDB() {
    conf, err := getDBConf(_DB_CONF_PATH)
    if conf == nil || err != nil {
        panic(err)
    }
    // https://github.com/go-sql-driver/mysql#dsn-data-source-name
    dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Hostname + ":" + strconv.Itoa(conf.Port) +
        ")/" + conf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

    if DBClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
        logger.Errorf("[InitDB] connect db failed. err:%v", err)
        panic(err)
    }
}

func getDBConf(filePath string) (*DBConf, error) {
    conf := new(DBConf)
    file, err := ioutil.ReadFile(filePath)
    if err != nil {
        logger.Errorf("[getDBConf] open file failed. err:%v", err)
        return nil, err
    }
    if err = yaml.Unmarshal(file, conf); err != nil {
        logger.Errorf("[getDBConf] unmarshal file failed. err:%v", err)
        return nil, err
    }
    return conf, nil
}
