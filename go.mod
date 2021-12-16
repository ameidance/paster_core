module github.com/ameidance/paster_core

go 1.15

require (
	github.com/cloudwego/kitex v0.1.1
	github.com/hashicorp/consul/api v1.10.1
	github.com/json-iterator/go v1.1.12
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
