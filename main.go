package main

import (
	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	client.InitDB()

	srv := pastercoreservice.NewServer(new(PasterCoreServiceImpl))
	if err := srv.Run(); err != nil {
		klog.Error(err)
		panic(err)
	}
}
