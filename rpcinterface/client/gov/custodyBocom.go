package gov

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/gov"
)

//CustodyBocomClient 管存开户[交通银行]
type CustodyBocomClient struct {
	client.Baseclient
}

//Init 初始化
func (c *CustodyBocomClient) Init() *CustodyBocomClient {
	c.ServiceName = "rpc_bus"
	c.ServicePath = "CustodyBocom"
	return c
}

// 上传图片到交通银行
func (c *CustodyBocomClient) UploadImgs(ctx context.Context, args *gov.InputParamsUploadImgs, reply *gov.OutputParamsUploadImgs) error {
	return c.Call(ctx, "UploadImgs", args, reply)
}

// 开户申请
func (c *CustodyBocomClient) OpenAccount(ctx context.Context, args *gov.InputParamsOpenAcc, reply *gov.OutputParamsOpenAcc) error {
	return c.Call(ctx, "OpenAccount", args, reply)
}

// 交行绑卡
func (c *CustodyBocomClient) BindBankCard(ctx context.Context, args *gov.InputParamsBindBankCard, reply *bool) error {
	return c.Call(ctx, "BindBankCard", args, reply)
}

// 开户异步通知
func (c *CustodyBocomClient) OpenNotify(ctx context.Context, args *gov.CallbackParams, reply *string) error {
	return c.Call(ctx, "OpenNotify", args, reply)
}

// 提现异步通知
func (c *CustodyBocomClient) CashNotify(ctx context.Context, args *gov.CallbackParams, reply *string) error {
	return c.Call(ctx, "CashNotify", args, reply)
}
