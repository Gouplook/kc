package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comtreeData"
)

type MerchantEmployee struct {
	client.Baseclient
}

func (c *MerchantEmployee) Init() *MerchantEmployee {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "MerchantEmployee"
	return c
}

func (c *MerchantEmployee) AddMerchantEmployeeRpc(ctx context.Context, args *comtreeData.ArgsMerchantEmployee, reply *bool) error {
	return c.Call(ctx, "AddMerchantEmployeeRpc", args, reply)
}
