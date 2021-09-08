package main

import (
	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/frame"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/paster/core/core"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	client.InitDB()
	client.InitConsul()

	srv := core.NewServer(new(CoreImpl), server.WithServiceAddr(frame.Address),
		server.WithServerBasicInfo(frame.EBI), server.WithRegistry(client.NewConsulRegistry()))
	if err := srv.Run(); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	} else {
		klog.Infof("[main] server stopped.")
	}
}
