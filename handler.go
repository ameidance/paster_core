package main

import (
	"context"

	"github.com/ameidance/paster_core/model/dto/kitex_gen/paster/core"
	"github.com/ameidance/paster_core/service"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

// CoreImpl implements the last service interface defined in the IDL.
type CoreImpl struct{}

// GetPost implements the CoreImpl interface.
func (s *CoreImpl) GetPost(ctx context.Context, req *core.GetPostRequest) (resp *core.GetPostResponse, err error) {
	klog.Infof("[GetPost] req:%v", util.GetJsonString(req))
	resp = service.GetPost(ctx, req)
	klog.Infof("[GetPost] resp:%v", util.GetJsonString(resp))
	return
}

// SavePost implements the CoreImpl interface.
func (s *CoreImpl) SavePost(ctx context.Context, req *core.SavePostRequest) (resp *core.SavePostResponse, err error) {
	klog.Infof("[SavePost] req:%v", util.GetJsonString(req))
	resp = service.SavePost(ctx, req)
	klog.Infof("[SavePost] resp:%v", util.GetJsonString(resp))
	return
}

// DeletePost implements the CoreImpl interface.
func (s *CoreImpl) DeletePost(ctx context.Context, req *core.DeletePostRequest) (resp *core.DeletePostResponse, err error) {
	klog.Infof("[DeletePost] req:%v", util.GetJsonString(req))
	resp = service.DeletePost(ctx, req)
	klog.Infof("[DeletePost] resp:%v", util.GetJsonString(resp))
	return
}

// GetComments implements the CoreImpl interface.
func (s *CoreImpl) GetComments(ctx context.Context, req *core.GetCommentsRequest) (resp *core.GetCommentsResponse, err error) {
	klog.Infof("[GetComments] req:%v", util.GetJsonString(req))
	resp = service.GetComments(ctx, req)
	klog.Infof("[GetComments] resp:%v", util.GetJsonString(resp))
	return
}

// SaveComment implements the CoreImpl interface.
func (s *CoreImpl) SaveComment(ctx context.Context, req *core.SaveCommentRequest) (resp *core.SaveCommentResponse, err error) {
	klog.Infof("[SaveComment] req:%v", util.GetJsonString(req))
	resp = service.SaveComment(ctx, req)
	klog.Infof("[SaveComment] resp:%v", util.GetJsonString(resp))
	return
}

// Check implements the CoreImpl interface.
func (s *CoreImpl) Check(ctx context.Context, req *core.HealthCheckRequest) (resp *core.HealthCheckResponse, err error) {
	return &core.HealthCheckResponse{Status: core.ServingStatus_SERVING}, nil
}

func (s *CoreImpl) Watch(req *core.HealthCheckRequest, stream core.Core_WatchServer) (err error) {
	return stream.Close()
}
