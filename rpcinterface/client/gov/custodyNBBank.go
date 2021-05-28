package gov

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/gov"
)

//CustodyNBBankClient [宁波银行]
type CustodyNBBankClient struct {
	client.Baseclient
}

//Init 初始化
func (c *CustodyNBBankClient) Init() *CustodyNBBankClient {
	c.ServiceName = "rpc_bus"
	c.ServicePath = "CustodyNBBank"
	return c
}

//OpenAccount 开户申请
func (c *CustodyNBBankClient) OpenAccount(ctx context.Context, args *gov.InputParamsCustodyNBBankOpenAccount, reply *gov.OutputParamsCustodyNBBankOpenAccount) error {
	return c.Call(ctx, "OpenAccount", args, reply)
}

//BindingAccount 账户绑定
func (c *CustodyNBBankClient) BindingAccount(ctx context.Context, args *gov.InputParamsCustodyNBBankBindingAccount, reply *gov.OutputParamsCustodyNBBankBindingAccount) error {
	return c.Call(ctx, "BindingAccount", args, reply)
}

//CashApply 资金提现
func (c *CustodyNBBankClient) CashApply(ctx context.Context, args *gov.InputParamsCustodyNBBankCashApply, reply *gov.OutputParamsCustodyNBBankCashApply) error {
	return c.Call(ctx, "CashApply", args, reply)
}

//OpenAccountNotify 开户状态通知
func (c *CustodyNBBankClient) OpenAccountNotify(ctx context.Context, args string, reply *gov.OutputParamsCustodyNBBankOpenAccountNotify) error {
	return c.Call(ctx, "OpenAccountNotify", args, reply)
}

//CashApplyNotify 资金提现通知
func (c *CustodyNBBankClient) CashApplyNotify(ctx context.Context, args string, reply *gov.OutputParamsCustodyNBBankCashApplyNotify) error {
	return c.Call(ctx, "CashApplyNotify", args, reply)
}
