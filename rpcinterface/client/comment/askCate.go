package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comment"
)

type AskCate struct {
	client.Baseclient
}

func (a *AskCate) Init() *AskCate {
	a.ServiceName = "rpc_comment"
	a.ServicePath = "AskCate"
	return a
}

//前台-获取问答分类数据
func (a *AskCate) GetAskCate(ctx context.Context, args *comment.ArgsEmptyParams, reply *comment.ReplyGetAskCate) error {
	return a.Call(ctx, "GetAskCate", args, reply)
}

//==============================================后台管理接口============================================
//添加一级分类
func (a *AskCate) AdminAddParentCate(ctx context.Context, args *comment.ArgsAdminAddParentCate, reply *comment.ReplyAdminAddCate) error {
	return a.Call(ctx, "AdminAddParentCate", args, reply)
}

//添加二级分类
func (a *AskCate) AdminAddSubCate(ctx context.Context, args *comment.ArgsAdminAddSubCate, reply *comment.ReplyAdminAddCate) error {
	return a.Call(ctx, "AdminAddSubCate", args, reply)
}

//修改二级分类
func (a *AskCate) AdminUpdateSubCate(ctx context.Context, args *comment.ArgsAdminAddSubCate, reply *bool) error {
	return a.Call(ctx, "AdminUpdateSubCate", args, reply)
}

//查询所有一级分类一级下属二级分类
func (a *AskCate) AdminGetCate(ctx context.Context, args *comment.ArgsEmptyParams, reply *comment.ReplyGetAskCate) error {
	return a.Call(ctx, "AdminGetCate", args, reply)
}

//根据二级分类查询1级分类id
func (a *AskCate) GetCateBySubId(ctx context.Context, args *int, reply *int) error {
	return a.Call(ctx, "GetCateBySubId", args, reply)
}

//根据二级分类id查询详情
func (a *AskCate) GetCateDetailById(ctx context.Context, args *int, reply *comment.AskCateBase) error {
	return a.Call(ctx, "GetCateDetailById", args, reply)
}
