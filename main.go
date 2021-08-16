package main

import (
	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/frame"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/core/pastercore"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	client.InitDB()
	client.InitConsul()

	srv := pastercore.NewServer(new(PasterCoreImpl), server.WithServerBasicInfo(frame.EBI), server.WithRegistry(client.NewConsulRegistry()))
	if err := srv.Run(); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	} else {
		klog.Infof("[main] server stopped.")
	}
}
