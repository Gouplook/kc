package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comtreeData"
)

type MerchantSecurityCode struct {
	client.Baseclient
}

func (c *MerchantSecurityCode) Init() *MerchantSecurityCode {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "MerchantSecurityCode"
	return c
}

func (c *MerchantSecurityCode) AddMerchantSecurityCodeRpc(ctx context.Context, args *comtreeData.ArgsMerchantSecurityCode, reply *bool) error {
	return c.Call(ctx, "AddMerchantSecurityCodeRpc", args, reply)
}
