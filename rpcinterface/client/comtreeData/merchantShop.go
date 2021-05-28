package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comtreeData"
)

type MerchantShop struct {
	client.Baseclient
}

func (c *MerchantShop) Init() *MerchantShop {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "MerchantShop"
	return c
}

func (c *MerchantShop) AddMerchantShopRpc(ctx context.Context, args *comtreeData.ArgsMerchantShop, reply *bool) error {
	return c.Call(ctx, "AddMerchantShopRpc", args, reply)
}
