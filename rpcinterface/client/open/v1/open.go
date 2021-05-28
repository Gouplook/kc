package open

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	v1 "git.900sui.cn/kc/rpcinterface/interface/open/v1"
)

type Open struct {
	client.Baseclient
}

func (o *Open) Init() *Open {
	o.ServiceName = "rpc_open"
	o.ServicePath = "V1/Open"
	return o
}

//获取指定的商户签约信息
func (o *Open) GetOpenMerchantContract(ctx context.Context, args *string, reply *v1.ReplyGetOneOpenMerchantContract) error {
	return o.Call(ctx, "GetOpenMerchantContract", args, reply)
}

//上传影像
func (o *Open) UploadFile(ctx context.Context, args *v1.ArgsUploadFile, reply *v1.ReplyUploadFile) error {
	return o.Call(ctx, "UploadFile", args, reply)
}

//商家发布预付充值卡
func (o *Open) MerchantCreateRcard(ctx context.Context, args *string, reply *v1.ReplyMerchantCreateRcard) error {
	return o.Call(ctx, "MerchantCreateRcard", args, reply)
}

//商家获取预付充值卡列表
func (o *Open) GetMerchantRcardLists(ctx context.Context, args *string, reply *v1.ReplyMerchantRcardLists) error {
	return o.Call(ctx, "GetMerchantRcardLists", args, reply)
}

//商家删除预付充值卡
func (o *Open) MerchantDelRcard(ctx context.Context, args *string, reply *v1.ReplyMerchantDelRcard) error {
	return o.Call(ctx, "MerchantDelRcard", args, reply)
}

//商家查询账户信息及资金
func (o *Open) GetMerchantDeposInfo(ctx context.Context, args *string, reply *v1.ReplyMerchantDeposInfo) error {
	return o.Call(ctx, "GetMerchantDeposInfo", args, reply)
}

//商家查询留存资金明细
func (o *Open) GetMerchantDeposLogs(ctx context.Context, args *string, reply *v1.ReplyMerchantDeposLogs) error {
	return o.Call(ctx, "GetMerchantDeposLogs", args, reply)
}

//商家绑定结算账户-个体
func (o *Open) BindEntitySettleAcct(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "BindEntitySettleAcct", args, reply)
}

//商家绑定结算账户-企业
func (o *Open) BindCompanySettleAcct(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "BindCompanySettleAcct", args, reply)
}

//商家绑定结算账户-小额鉴权回填
func (o *Open) SmalAmountAuthBackfill(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "SmalAmountAuthBackfill", args, reply)
}

//商家解绑结算账户
func (o *Open) MerchantUnbindRelateAcct(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "MerchantUnbindRelateAcct", args, reply)
}

//商家申请提现
func (o *Open) MerchantApplyWithdraw(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "MerchantApplyWithdraw", args, reply)
}

//商家查询提现记录
func (o *Open) GetMerchantWithdrawLogs(ctx context.Context, args *string, reply *v1.ReplyGetMerchantWithdrawLogs) error {
	return o.Call(ctx, "GetMerchantWithdrawLogs", args, reply)
}

//门店进件
func (o *Open) ApplyShop(ctx context.Context, args *string, reply *v1.ReplyApplyShop) error {
	return o.Call(ctx, "ApplyShop", args, reply)
}

//门店获取信息对接唯一 ID
func (o *Open) GetGovBusId(ctx context.Context, args *string, reply *v1.ReplyGetGovBusId) error {
	return o.Call(ctx, "GetGovBusId", args, reply)
}

//门店添加预付充值卡
func (o *Open) ShopAddRcard(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "ShopAddRcard", args, reply)
}

//门店获取预付充值卡列表
func (o *Open) GetShopRcardLists(ctx context.Context, args *string, reply *v1.ReplyGetShopRcardLists) error {
	return o.Call(ctx, "GetShopRcardLists", args, reply)
}

//门店上、下架预付充值卡
func (o *Open) ShopDownUpRcard(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "ShopDownUpRcard", args, reply)
}

//门店删除预付卡充值卡
func (o *Open) ShopDelRcard(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "ShopDelRcard", args, reply)
}

//查询会员卡包信息
func (o *Open) GetUserCardPackageLists(ctx context.Context, args *string, reply *v1.ReplyGetUserCardPackageLists) error {
	return o.Call(ctx, "GetUserCardPackageLists", args, reply)
}

//门店兑付预付充值卡
func (o *Open) ShopConsumeRcardPackage(ctx context.Context, args *string, reply *bool) error {
	return o.Call(ctx, "ShopConsumeRcardPackage", args, reply)
}

//查询预付充值卡兑付记录
func (o *Open) GetRardPackageConsumeLog(ctx context.Context, args *string, reply *v1.ReplyGetRardPackageConsumeLog) error {
	return o.Call(ctx, "GetRardPackageConsumeLog", args, reply)
}
