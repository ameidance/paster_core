package frame

import (
	"io/ioutil"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
	"gopkg.in/yaml.v3"
)

var (
	Address net.Addr
)

type kitexConf struct {
	Address string `yaml:"Address"`
}

func (*kitexConf) Network() string {
	return "tcp"
}

func (m *kitexConf) String() string {
	return m.Address
}

func init() {
	conf, err := getKitexConfig()
	if conf == nil || err != nil {
		panic(err)
	}
	Address = conf
}

func getKitexConfig() (*kitexConf, error) {
	conf := new(kitexConf)
	file, err := ioutil.ReadFile(utils.GetConfFile())
	if err != nil {
		klog.Errorf("[getKitexConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getKitexConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
