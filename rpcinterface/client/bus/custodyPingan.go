package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/gov"
	v1 "git.900sui.cn/kc/rpcinterface/interface/open/v1"

	"git.900sui.cn/kc/rpcinterface/client"
)

//CustodyPinganClient 平安存管
type CustodyPinganClient struct {
	client.Baseclient
}

//Init 初始化
func (c *CustodyPinganClient) Init() *CustodyPinganClient {
	c.ServiceName = "rpc_bus"
	c.ServicePath = "CustodyPingan"
	return c
}

//第三方渠道在途充值
func (c *CustodyPinganClient) ThirdRechargeWay(ctx context.Context, args *gov.RechargeWay, reply *gov.ReplyRechargeWay) (err error) {
	return c.Call(ctx, "ThirdRechargeWay", args, reply)
}

//会员绑卡前需要的数据信息
func (c *CustodyPinganClient) GetPinganPreBindCardInfo(ctx context.Context, args *gov.ArgsPinganBindInfo, reply *gov.ReplyPinganBindInfo) (err error) {
	return c.Call(ctx, "GetPinganPreBindCardInfo", args, reply)
}

//绑定结算账户
func (c *CustodyPinganClient) BindAmount(ctx context.Context, args gov.ArgsPinganBindAmount, reply *bool) error {
	return c.Call(ctx, "BindAmount", args, reply)
}

//平安绑卡验证
func (c *CustodyPinganClient) CheckAmount(ctx context.Context, args gov.ArgsCheckAmountData, reply *bool) error {
	return c.Call(ctx, "CheckAmount", args, reply)
}

//解绑提现账户
func (c *CustodyPinganClient) UnbindRelateAcct(ctx context.Context, args gov.ArgsUnbindRelateAcct, reply *bool) error {
	return c.Call(ctx, "UnbindRelateAcct", args, reply)
}

//生成请求流水号
func (c *CustodyPinganClient) CreateTransNo(ctx context.Context, cid *int, transNo *string) error {
	return c.Call(ctx, "CreateTransNo", cid, transNo)
}

//会员资金冻结
func (c *CustodyPinganClient) MembershipTrancheFreeze(ctx context.Context, args gov.MembershipTrancheFreeze, reply *gov.ReplyMembershipTrancheFreeze) (err error) {
	return c.Call(ctx, "MembershipTrancheFreeze", args, reply)
}

//核对提现结果
func (c *CustodyPinganClient) GetCashOutResult(ctx context.Context, args string, reply *bool) error {
	return c.Call(ctx, "GetCashOutResult", args, reply)
}

//开放平台-绑定结算账户-个体工商户
func (c *CustodyPinganClient) OpenPlatFormV1BindAmountGt(ctx context.Context,args v1.ArgsBindEntitySettleAcct,reply *bool)error{
	return c.Call(ctx, "OpenPlatFormV1BindAmountGt", args, reply)
}

//开放平台-绑定结算账户-企业
func (c *CustodyPinganClient) OpenPlatFormV1BindAmountQy(ctx context.Context,args v1.ArgsBindCompanySettleAcct,reply *bool)error{
	return c.Call(ctx, "OpenPlatFormV1BindAmountQy", args, reply)
}

//开放平台-绑定结算账户-回填金额
func (c *CustodyPinganClient)OpenPlatFormV1CheckAmount(ctx context.Context, args v1.ArgsSmalAmountAuthBackfill, reply *bool) (err error) {
	return c.Call(ctx, "OpenPlatFormV1CheckAmount", args, reply)
}

//开放平台-会员解绑提现账户
func (c *CustodyPinganClient)OpenPlatFormV1UnbindRelateAcct(ctx context.Context, args v1.ArgsMerchantUnbindRelateAcct, reply *bool) (err error) {
	return c.Call(ctx, "OpenPlatFormV1UnbindRelateAcct", args, reply)
}