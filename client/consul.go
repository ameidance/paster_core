package client

import (
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_core/constant"
	"github.com/ameidance/paster_core/frame"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	_consulConf  *consulConf
	consulClient *api.Client
)

type consulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

func InitConsul() {
	var err error
	_consulConf, err = getConsulConfig()
	if _consulConf == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", _consulConf.Hostname, _consulConf.Port)
	consulClient, err = api.NewClient(config)
	if consulClient == nil || err != nil {
		panic(err)
	}
}

type ConsulRegistry struct {
}

func NewConsulRegistry() *ConsulRegistry {
	if _consulConf == nil {
		return nil
	}
	return &ConsulRegistry{}
}

func (m *ConsulRegistry) Register(info *registry.Info) (err error) {
	if consulClient == nil {
		return nil
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = frame.GetInstanceId()
	registration.Name = frame.GetServiceName()
	registration.Address = util.GetInternalIP()
	_, err = fmt.Sscanf(info.Addr.String(), "[::]:%v", &registration.Port)
	if err != nil {
		klog.Errorf("[ConsulRegistry -> Register] get registry info addr port failed. err:%v", err)
		return
	}

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.GRPC = fmt.Sprintf("%s:%d/%s", registration.Address, registration.Port, registration.Name)
	registration.Check.Timeout = "5s"
	registration.Check.Interval = "5s"
	registration.Check.DeregisterCriticalServiceAfter = "10s"

	klog.Infof("[ConsulRegistry -> Register] registering... instance id:%v", frame.GetInstanceId())
	return consulClient.Agent().ServiceRegister(registration)
}

func (m *ConsulRegistry) Deregister(info *registry.Info) error {
	if consulClient == nil {
		return nil
	}

	klog.Infof("[ConsulRegistry -> Register] deregistering... instance id:%v", frame.GetInstanceId())
	return consulClient.Agent().ServiceDeregister(frame.GetInstanceId())
}

func getConsulConfig() (*consulConf, error) {
	conf := new(consulConf)
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
