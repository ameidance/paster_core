package main

import (
	"context"

	"github.com/ameidance/paster_core/model/dto/kitex_gen/core"
	"github.com/ameidance/paster_core/service"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

// PasterCoreImpl implements the last service interface defined in the IDL.
type PasterCoreImpl struct{}

// GetPost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) GetPost(ctx context.Context, req *core.GetPostRequest) (resp *core.GetPostResponse, err error) {
	klog.Infof("[GetPost] req:%v", util.GetJsonString(req))
	resp = service.GetPost(ctx, req)
	klog.Infof("[GetPost] resp:%v", util.GetJsonString(resp))
	return
}

// SavePost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) SavePost(ctx context.Context, req *core.SavePostRequest) (resp *core.SavePostResponse, err error) {
	klog.Infof("[SavePost] req:%v", util.GetJsonString(req))
	resp = service.SavePost(ctx, req)
	klog.Infof("[SavePost] resp:%v", util.GetJsonString(resp))
	return
}

// DeletePost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) DeletePost(ctx context.Context, req *core.DeletePostRequest) (resp *core.DeletePostResponse, err error) {
	klog.Infof("[DeletePost] req:%v", util.GetJsonString(req))
	resp = service.DeletePost(ctx, req)
	klog.Infof("[DeletePost] resp:%v", util.GetJsonString(resp))
	return
}

// GetComments implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) GetComments(ctx context.Context, req *core.GetCommentsRequest) (resp *core.GetCommentsResponse, err error) {
	klog.Infof("[GetComments] req:%v", util.GetJsonString(req))
	resp = service.GetComments(ctx, req)
	klog.Infof("[GetComments] resp:%v", util.GetJsonString(resp))
	return
}

// SaveComment implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) SaveComment(ctx context.Context, req *core.SaveCommentRequest) (resp *core.SaveCommentResponse, err error) {
	klog.Infof("[SaveComment] req:%v", util.GetJsonString(req))
	resp = service.SaveComment(ctx, req)
	klog.Infof("[SaveComment] resp:%v", util.GetJsonString(resp))
	return
}

// Check implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) Check(ctx context.Context, req *core.HealthCheckRequest) (resp *core.HealthCheckResponse, err error) {
	return &core.HealthCheckResponse{Status: core.ServingStatus_SERVING}, nil
}

func (s *PasterCoreImpl) Watch(req *core.HealthCheckRequest, stream core.PasterCore_WatchServer) (err error) {
	return stream.Close()
}
