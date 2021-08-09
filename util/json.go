package util

import (
	"github.com/cloudwego/kitex/pkg/klog"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetJsonMap(obj interface{}) map[string]interface{} {
	data, err := json.Marshal(obj)
	if err != nil {
		klog.Errorf("[GetJsonMap] marshal failed. err:%v", err)
		return nil
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		klog.Errorf("[GetJsonMap] unmarshal failed. err:%v", err)
		return nil
	}
	return m
}

func GetJsonString(obj interface{}) string {
	str, err := json.MarshalToString(obj)
	if err != nil {
		klog.Errorf("[GetJsonString] marshal failed. err:%v", err)
		return ""
	}
	return str
}
