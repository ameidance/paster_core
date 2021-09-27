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

func GetPost(ctx context.Context, req *core.GetPostRequest) *core.GetPostResponse {
	resp := new(core.GetPostResponse)
	util.FillBizResp(resp, constant.SUCCESS)

	mgr := po.PostMgr(client.DBClient)
	postPO, err := mgr.GetFromID(req.GetId())
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
	if postPO.ValidatePassword(req.GetPassword()) {
		util.FillBizResp(resp, constant.ERR_WRONG_PASSWORD)
		return resp
	}

	// PO->DTO
	postDTO, err := postPO.ConvertToDTO(req.GetPassword())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	resp.Info = postDTO
	return resp
}

func SavePost(ctx context.Context, req *core.SavePostRequest) *core.SavePostResponse {
	resp := new(core.SavePostResponse)
	util.FillBizResp(resp, constant.SUCCESS)

	// DTO->PO
	postDTO := req.GetInfo()
	if postDTO == nil {
		klog.Errorf("[SavePost] post info empty")
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}
	postPO := new(po.Post)
	err := postPO.ConvertFromDTO(postDTO, req.GetPassword())
	if err != nil {
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	mgr := po.PostMgr(client.DBClient)
	res := mgr.Save(postPO)
	if res.Error != nil {
		klog.Errorf("[SavePost] save post failed. err:%v", res.Error)
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	resp.Id = postPO.ID
	return resp
}

func DeletePost(ctx context.Context, req *core.DeletePostRequest) *core.DeletePostResponse {
	resp := new(core.DeletePostResponse)
	util.FillBizResp(resp, constant.SUCCESS)

	mgr := po.PostMgr(client.DBClient)
	res := mgr.Delete(&po.Post{
		ID: req.GetId(),
	})
	if res.Error != nil {
		klog.Errorf("[DeletePost] delete post failed. err:%v", res.Error)
		util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
		return resp
	}

	return resp
}
