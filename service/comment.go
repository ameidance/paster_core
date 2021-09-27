package service

import (
	"context"

	"github.com/ameidance/paster_core/client"
	"github.com/ameidance/paster_core/constant"
	"github.com/ameidance/paster_core/model/dto/kitex_gen/paster/core"
	"github.com/ameidance/paster_core/model/po"
	"github.com/ameidance/paster_core/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetComments(ctx context.Context, req *core.GetCommentsRequest) *core.GetCommentsResponse {
	resp := new(core.GetCommentsResponse)
	util.FillBizResp(resp, constant.SUCCESS)

	postMgr := po.PostMgr(client.DBClient)
	postPO, err := postMgr.GetFromID(req.GetPostId())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	// check if exists
	if postPO.ID == 0 {
		util.FillBizResp(resp, constant.ERR_RECORD_NOT_FOUND)
		return resp
	}
	// validate
	if !postPO.ValidatePassword(req.GetPassword()) {
		util.FillBizResp(resp, constant.ERR_WRONG_PASSWORD)
		return resp
	}

	commentMgr := po.CommentMgr(client.DBClient)
	commentsPO, err := commentMgr.GetBatchFromPostID([]int64{req.GetPostId()})
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}
	// PO->DTO
	commentsDTO, err := po.Comments(commentsPO).ConvertToDTO(req.GetPassword())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	resp.Info = commentsDTO
	return resp
}

func SaveComment(ctx context.Context, req *core.SaveCommentRequest) *core.SaveCommentResponse {
	resp := new(core.SaveCommentResponse)
	util.FillBizResp(resp, constant.SUCCESS)

	postMgr := po.PostMgr(client.DBClient)
	postPO, err := postMgr.GetFromID(req.GetPostId())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	// check if exists
	if postPO.ID == 0 {
		util.FillBizResp(resp, constant.ERR_RECORD_NOT_FOUND)
		return resp
	}
	// validate
	if !postPO.ValidatePassword(req.GetPassword()) {
		util.FillBizResp(resp, constant.ERR_WRONG_PASSWORD)
		return resp
	}

	// DTO->PO
	commentDTO := req.GetInfo()
	if commentDTO == nil {
		klog.Errorf("[SaveComment] comment info empty")
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}
	commentPO := new(po.Comment)
	err = commentPO.ConvertFromDTO(commentDTO, req.GetPostId(), req.GetPassword())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	commentMgr := po.CommentMgr(client.DBClient)
	res := commentMgr.Save(commentPO)
	if res.Error != nil {
		klog.Errorf("[SaveComment] save comment failed. err:%v", res.Error)
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	return resp
}
