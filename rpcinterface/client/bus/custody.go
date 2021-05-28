package bus

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

//CustodyClient 管存开户2
type CustodyClient struct {
	client.Baseclient
}

//Init 初始化
func (c *CustodyClient) Init() *CustodyClient {
	c.ServiceName = "rpc_bus"
	c.ServicePath = "Custody"
	return c
}

//OpenAccount 开设账户
func (c *CustodyClient) OpenAccount(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "OpenAccount", args, reply)
}

//SendSms 发送短信验证码
func (c *CustodyClient) SendSms(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "SendSms", args, reply)
}

//ActiveAccount 激活帐号
func (c *CustodyClient) ActiveAccount(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "ActiveAccount", args, reply)
}

//ApplyContractSign 申请签约，发送授权码
func (c *CustodyClient) ApplyContractSign(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "ApplyContractSign", args, reply)
}

//ContractSign 获取授权码，开始签约
func (c *CustodyClient) ContractSign(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "ContractSign", args, reply)
}

//OpenAcctNotify 开户异步通知
func (c *CustodyClient) OpenAcctNotify(ctx context.Context, args *bus.AccountParams, reply *bus.AccountReply) error {
	return c.Call(ctx, "OpenAcctNotify", args, reply)
}
