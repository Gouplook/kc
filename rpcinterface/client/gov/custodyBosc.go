package gov

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/gov"
)

//CustodyBoscClient 管存开户[上海银行]
type CustodyBoscClient struct {
	client.Baseclient
}

//Init 初始化
func (c *CustodyBoscClient) Init() *CustodyBoscClient {
	c.ServiceName = "rpc_bus"
	c.ServicePath = "CustodyBosc"
	return c
}

//SendSms 开设账户,发送短信验证码
func (c *CustodyBoscClient) SendSms(ctx context.Context, args *gov.InputParamsSendSms, reply *gov.OutputParamsSendSms) error {
	return c.Call(ctx, "SendSms", args, reply)
}

//OpenAccount 开设账户
func (c *CustodyBoscClient) OpenAccount(ctx context.Context, args *gov.InputParamsOpenAccount, reply *gov.OutputParamsOpenAccount) error {
	return c.Call(ctx, "OpenAccount", args, reply)
}

//ActiveAccount 激活帐号
func (c *CustodyBoscClient) ActiveAccount(ctx context.Context, args *gov.InputParamsActiveAccount, reply *gov.OutputParamsActiveAccount) error {
	return c.Call(ctx, "ActiveAccount", args, reply)
}

//ApplyContractSign 申请签约，发送授权码
func (c *CustodyBoscClient) ApplyContractSign(ctx context.Context, args *gov.InputParamsApplyContractSign, reply *gov.OutputParamsApplyContractSign) error {
	return c.Call(ctx, "ApplyContractSign", args, reply)
}

//ContractSign 获取授权码，开始签约
func (c *CustodyBoscClient) ContractSign(ctx context.Context, args *gov.InputParamsContractSign, reply *gov.OutputParamsContractSign) error {
	return c.Call(ctx, "ContractSign", args, reply)
}

//OpenAcctNotify 开户异步通知
func (c *CustodyBoscClient) OpenAcctNotify(ctx context.Context, args *gov.InputParamsOpenAcctNotify, reply *gov.OutputParamsOpenAcctNotify) error {
	return c.Call(ctx, "OpenAcctNotify", args, reply)
}
