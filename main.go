package main

import (
	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/frame"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	client.InitDB()
	client.InitConsul()

	srv := pastercoreservice.NewServer(new(PasterCoreServiceImpl), server.WithServerBasicInfo(frame.EBI),
		server.WithRegistry(client.NewConsulRegistry()), server.WithMiddleware(frame.LogMiddleware))
	if err := srv.Run(); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	} else {
		klog.Infof("[main] server stopped.")
	}
}
