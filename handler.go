package main

import (
    "context"

    "github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core"
    "github.com/ameidance/paster_core/service"
)

// PasterCoreServiceImpl implements the last service interface defined in the IDL.
type PasterCoreServiceImpl struct{}

// GetPost implements the PasterCoreServiceImpl interface.
func (s *PasterCoreServiceImpl) GetPost(ctx context.Context, req *core.GetPostRequest) (resp *core.GetPostResponse, err error) {
    return service.GetPost(ctx, req), nil
}

// SavePost implements the PasterCoreServiceImpl interface.
func (s *PasterCoreServiceImpl) SavePost(ctx context.Context, req *core.SavePostRequest) (resp *core.SavePostResponse, err error) {
    return service.SavePost(ctx, req), nil
}

// DeletePost implements the PasterCoreServiceImpl interface.
func (s *PasterCoreServiceImpl) DeletePost(ctx context.Context, req *core.DeletePostRequest) (resp *core.DeletePostResponse, err error) {
    return service.DeletePost(ctx, req), nil
}

// GetComments implements the PasterCoreServiceImpl interface.
func (s *PasterCoreServiceImpl) GetComments(ctx context.Context, req *core.GetCommentsRequest) (resp *core.GetCommentsResponse, err error) {
    return service.GetComments(ctx, req), nil
}

// SaveComment implements the PasterCoreServiceImpl interface.
func (s *PasterCoreServiceImpl) SaveComment(ctx context.Context, req *core.SaveCommentRequest) (resp *core.SaveCommentResponse, err error) {
    return service.SaveComment(ctx, req), nil
}
