package frame

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ameidance/paster_core/constant"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	EBI *rpcinfo.EndpointBasicInfo
)

func init() {
	EBI = new(rpcinfo.EndpointBasicInfo)
	EBI.Tags = make(map[string]string)

	rand.Seed(time.Now().Unix())
	SetServiceName(constant.SERVICE_NAME)
	SetInstanceId(fmt.Sprintf("%s_%d", constant.SERVICE_NAME, rand.Int()))
}

func SetServiceName(name string) {
	if EBI == nil {
		return
	}
	EBI.ServiceName = name
}

func SetInstanceId(id string) {
	if EBI == nil || EBI.Tags == nil {
		return
	}
	EBI.Tags["id"] = id
}

func GetServiceName() string {
	if EBI == nil {
		return ""
	}
	return EBI.ServiceName
}

func GetInstanceId() string {
	if EBI == nil || EBI.Tags == nil {
		return ""
	}
	id, _ := EBI.Tags["id"]
	return id
}
