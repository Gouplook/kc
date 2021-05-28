package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/insurance"
)

type AuthApply struct {
	client.Baseclient
}

func (a *AuthApply) Init() *AuthApply  {
	a.ServiceName = "rpc_insurance"
	a.ServicePath = "AuthApply"
	return a
}

//获取申请列表
func (a *AuthApply) GetApplyLists(ctx context.Context, args *insurance.ArgsGetApplyLists, reply *insurance.ReplyGetApplyLists) error  {
	return a.Call(ctx, "GetApplyLists", args, reply)
}

//申请详情
func (a *AuthApply) ApplyDetail(ctx context.Context, id *int, reply *insurance.ApplyInfo)  error {
	return a.Call(ctx, "ApplyDetail", id, reply)
}

//提交投保申请
func (a *AuthApply) ApplyDo(ctx context.Context, args *insurance.ArgsApplyDo , reply *bool)  error {
	return a.Call(ctx, "ApplyDo", args, reply)
}

//手动通知监管平台
func (a *AuthApply) RetryNotifyToGov(ctx context.Context, id *int, reply *bool)  error {
	return a.Call(ctx, "RetryNotifyToGov", id, reply)
}

//手动同意签约意向
func (a *AuthApply) RetryAgreeToAnx(ctx context.Context, id *int, reply *bool)  error {

	return a.Call(ctx, "RetryAgreeToAnx", id, reply)
}

//处理安信承保回调业务
func (a *AuthApply) AaicMerchantResultNotify(ctx context.Context, notifyData *string, reply *bool) error{
	return a.Call(ctx, "AaicMerchantResultNotify", notifyData, reply)
}

//监管平台商家同步到平台，绑定商家承保信息
func (a *AuthApply) GovAuthSucBindToBus(ctx context.Context, args *insurance.ArgsGovAuthSucBindToBus, reply *insurance.ReplyGovAuthSucBindToBus ) error {
	return a.Call(ctx, "GovAuthSucBindToBus", args, reply)
}