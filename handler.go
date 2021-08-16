package main

import (
	"context"

	"github.com/ameidance/paster_core/model/dto/kitex_gen/core"
	"github.com/ameidance/paster_core/service"
)

// PasterCoreImpl implements the last service interface defined in the IDL.
type PasterCoreImpl struct{}

// GetPost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) GetPost(ctx context.Context, req *core.GetPostRequest) (resp *core.GetPostResponse, err error) {
	return service.GetPost(ctx, req), nil
}

// SavePost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) SavePost(ctx context.Context, req *core.SavePostRequest) (resp *core.SavePostResponse, err error) {
	return service.SavePost(ctx, req), nil
}

// DeletePost implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) DeletePost(ctx context.Context, req *core.DeletePostRequest) (resp *core.DeletePostResponse, err error) {
	return service.DeletePost(ctx, req), nil
}

// GetComments implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) GetComments(ctx context.Context, req *core.GetCommentsRequest) (resp *core.GetCommentsResponse, err error) {
	return service.GetComments(ctx, req), nil
}

// SaveComment implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) SaveComment(ctx context.Context, req *core.SaveCommentRequest) (resp *core.SaveCommentResponse, err error) {
	return service.SaveComment(ctx, req), nil
}

// Check implements the PasterCoreImpl interface.
func (s *PasterCoreImpl) Check(ctx context.Context, req *core.HealthCheckRequest) (resp *core.HealthCheckResponse, err error) {
	return &core.HealthCheckResponse{Status: core.HealthCheckResponse_SERVING}, nil
}

func (s *PasterCoreImpl) Watch(req *core.HealthCheckRequest, stream core.PasterCore_WatchServer) (err error) {
	return stream.Close()
}
