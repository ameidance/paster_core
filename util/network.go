package util

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
)

func GetInternalIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		klog.Errorf("[GetInternalIP] get network address failed. err:%v", err)
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	klog.Errorf("[GetInternalIP] get network address failed.")
	return
}
