package service

import (
    "context"

    "github.com/ameidance/paster_core/client"
    "github.com/ameidance/paster_core/constant"
    "github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core"
    "github.com/ameidance/paster_core/model/po"
    "github.com/ameidance/paster_core/util"
    "github.com/apache/thrift/lib/go/thrift"
    "github.com/bytedance/gopkg/util/logger"
)

func GetPost(ctx context.Context, req *core.GetPostRequest) *core.GetPostResponse {
    resp := core.NewGetPostResponse()
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
    if len(postPO.Passwd) > 0 && util.GetMd5String([]byte(req.GetPassword())) != postPO.Passwd {
        util.FillBizResp(resp, constant.ERR_WRONG_PASSWORD)
        return resp
    }

    // PO->DTO
    postDTO, err := postPO.ConvertToDTO(ctx, req.GetPassword())
    if err != nil {
        util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
        return resp
    }

    resp.SetInfo(postDTO)
    return resp
}

func SavePost(ctx context.Context, req *core.SavePostRequest) *core.SavePostResponse {
    resp := core.NewSavePostResponse()
    util.FillBizResp(resp, constant.SUCCESS)

    // DTO->PO
    postDTO := req.GetInfo()
    if postDTO == nil {
        logger.CtxErrorf(ctx, "[SavePost] post info empty")
        util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
        return resp
    }
    postPO := new(po.Post)
    err := postPO.ConvertFromDTO(ctx, postDTO, req.GetPassword())
    if err != nil {
        util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
        return resp
    }

    mgr := po.PostMgr(client.DBClient)
    res := mgr.Save(postPO)
    if res.Error != nil {
        logger.CtxErrorf(ctx, "[SavePost] save post failed. err:%v", res.Error)
        util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
        return resp
    }

    resp.SetId(thrift.Int64Ptr(postPO.ID))
    return resp
}

func DeletePost(ctx context.Context, req *core.DeletePostRequest) *core.DeletePostResponse {
    resp := core.NewDeletePostResponse()
    util.FillBizResp(resp, constant.SUCCESS)

    mgr := po.PostMgr(client.DBClient)
    res := mgr.Delete(&po.Post{
        ID: req.GetId(),
    })
    if res.Error != nil {
        logger.CtxErrorf(ctx, "[DeletePost] delete post failed. err:%v", res.Error)
        util.FillBizResp(resp, constant.ERR_SERVICE_INTERNAL)
        return resp
    }

    return resp
}
