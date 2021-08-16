package frame

import (
	"context"

	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = LogMiddleware

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		method := rpcinfo.GetRPCInfo(ctx).Invocation().MethodName()
		klog.Debugf("[LogMiddleware] rpc method:%v, request:%v", method, util.GetJsonString(req))
		if err = next(ctx, req, resp); err != nil {
			return
		}
		klog.Debugf("[LogMiddleware] rpc method:%v, response:%v", method, util.GetJsonString(resp))
		return nil
	}
}
