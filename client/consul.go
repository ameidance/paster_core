package client

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/ameidance/paster_core/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	ConsulClient *api.Client
)

type _ConsulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

func getConsulConfig() (*_ConsulConf, error) {
	conf := new(_ConsulConf)
	file, err := ioutil.ReadFile(constant.CONSUL_CONF_PATH)
	if err != nil {
		klog.Errorf("[getConsulConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getConsulConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}

func InitConsul() {
	address, err := getConsulConfig()
	if address == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", address.Hostname, address.Port)
	ConsulClient, err = api.NewClient(config)
	if ConsulClient == nil || err != nil {
		panic(err)
	}
}

type ConsulRegistry struct {
	ServiceName string
	InstanceId  string
}

func NewConsulRegistry() *ConsulRegistry {
	return &ConsulRegistry{
		ServiceName: constant.SERVICE_NAME,
		InstanceId:  strconv.Itoa(rand.Int()),
	}
}

func (m *ConsulRegistry) Register(info *registry.Info) (err error) {
	if ConsulClient == nil {
		return nil
	}

	var host string
	var port int
	addr := strings.Split(info.Addr.String(), ":")
	if len(addr) != 2 {
		klog.Errorf("[ConsulRegistry -> Register] registry info addr split failed.")
		return fmt.Errorf("registry info addr split failed")
	}
	port, err = strconv.Atoi(addr[1])
	if err != nil {
		klog.Errorf("[ConsulRegistry -> Register] registry info addr port atoi failed. err:%v", err)
		return
	}
	if len(addr[0]) > 0 {
		host = addr[0]
	} else {
		// use local ip by default
		host = "localhost"
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = m.InstanceId
	registration.Name = m.ServiceName
	registration.Address = host
	registration.Port = port

	return ConsulClient.Agent().ServiceRegister(registration)
}

func (m *ConsulRegistry) Deregister(info *registry.Info) error {
	if ConsulClient == nil {
		return nil
	}
	return ConsulClient.Agent().ServiceDeregister(m.InstanceId)
}
