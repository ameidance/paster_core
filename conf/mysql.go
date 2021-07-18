package conf

import (
	"io/ioutil"

	"github.com/bytedance/gopkg/util/logger"
	"gopkg.in/yaml.v3"
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

func GetDBConfig() (*DBConf, error) {
	conf := new(DBConf)
	file, err := ioutil.ReadFile(_DB_CONF_PATH)
	if err != nil {
		logger.Errorf("[GetDBConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		logger.Errorf("[GetDBConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
