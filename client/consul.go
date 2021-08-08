package client

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync/atomic"

	"github.com/ameidance/paster_core/constant"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	ConsulCheckCounter int64
	ConsulConf         *_ConsulConf
	ConsulClient       *api.Client
)

type _ConsulConf struct {
	Hostname  string `yaml:"hostname"`
	Port      int    `yaml:"port"`
	CheckPort int    `yaml:"check_port"`
}

func InitConsul() {
	var err error
	ConsulConf, err = getConsulConfig()
	if ConsulConf == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", ConsulConf.Hostname, ConsulConf.Port)
	ConsulClient, err = api.NewClient(config)
	if ConsulClient == nil || err != nil {
		panic(err)
	}
	go healthCheckHandler(ConsulConf.CheckPort)
}

type ConsulRegistry struct {
	ServiceName string
	InstanceId  string
}

func NewConsulRegistry() *ConsulRegistry {
	if ConsulConf == nil {
		return nil
	}
	return &ConsulRegistry{
		ServiceName: constant.SERVICE_NAME,
		InstanceId:  fmt.Sprintf("%s-%d", constant.SERVICE_NAME, rand.Int()),
	}
}

func (m *ConsulRegistry) Register(info *registry.Info) (err error) {
	if ConsulClient == nil {
		return nil
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = m.InstanceId
	registration.Name = m.ServiceName
	registration.Address = util.GetInternalIP()
	_, err = fmt.Sscanf(info.Addr.String(), ":%v", &registration.Port)
	if err != nil {
		klog.Errorf("[ConsulRegistry -> Register] get registry info addr port failed. err:%v", err)
		return
	}

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/health", registration.Address, ConsulConf.CheckPort)
	registration.Check.Timeout = "5s"
	registration.Check.Interval = "5s"
	registration.Check.DeregisterCriticalServiceAfter = "10s"

	klog.Infof("[ConsulRegistry -> Register] registering... instance id:%v", m.InstanceId)
	return ConsulClient.Agent().ServiceRegister(registration)
}

func (m *ConsulRegistry) Deregister(info *registry.Info) error {
	if ConsulClient == nil {
		return nil
	}

	klog.Infof("[ConsulRegistry -> Register] deregistering... instance id:%v", m.InstanceId)
	return ConsulClient.Agent().ServiceDeregister(m.InstanceId)
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

func healthCheckHandler(port int) {
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		atomic.AddInt64(&ConsulCheckCounter, 1)
		klog.Infof("[healthCheckHandler] counter:%v", atomic.LoadInt64(&ConsulCheckCounter))
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		klog.Errorf("[healthCheckHandler] serve failed. err:%v", err)
	}
}
