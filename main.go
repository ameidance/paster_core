package main

import (
	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/bytedance/gopkg/util/logger"
)

func main() {
	client.InitDB()

	srv := pastercoreservice.NewServer(new(PasterCoreServiceImpl))
	if err := srv.Run(); err != nil {
		logger.Error(err)
		panic(err)
	}
}
