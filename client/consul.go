package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"

	"github.com/ameidance/paster_core/constant"
	"github.com/ameidance/paster_core/frame"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	consulCheckCounter int64
	consulConf         *_ConsulConf
	consulClient       *api.Client
)

type _ConsulConf struct {
	Hostname  string `yaml:"hostname"`
	Port      int    `yaml:"port"`
	CheckPort int    `yaml:"check_port"`
}

func InitConsul() {
	var err error
	consulConf, err = getConsulConfig()
	if consulConf == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", consulConf.Hostname, consulConf.Port)
	consulClient, err = api.NewClient(config)
	if consulClient == nil || err != nil {
		panic(err)
	}
	go healthCheckHandler(consulConf.CheckPort)
}

type ConsulRegistry struct {
}

func NewConsulRegistry() *ConsulRegistry {
	if consulConf == nil {
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
	_, err = fmt.Sscanf(info.Addr.String(), ":%v", &registration.Port)
	if err != nil {
		klog.Errorf("[ConsulRegistry -> Register] get registry info addr port failed. err:%v", err)
		return
	}

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/health", registration.Address, consulConf.CheckPort)
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
		atomic.AddInt64(&consulCheckCounter, 1)
		//klog.Debugf("[healthCheckHandler] counter:%v", atomic.LoadInt64(&consulCheckCounter))
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		klog.Errorf("[healthCheckHandler] serve failed. err:%v", err)
	}
}
