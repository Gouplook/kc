package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comtreeData"
)

type MerchantSettled struct {
	client.Baseclient
}

func (c *MerchantSettled) Init() *MerchantSettled {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "MerchantSettled"
	return c
}

func (c *MerchantSettled) AddMerchantSettledRpc(ctx context.Context, args *comtreeData.ArgsMerchantSettled, reply *bool) error {
	return c.Call(ctx, "AddMerchantSettledRpc", args, reply)
}
