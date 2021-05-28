package comtreeData

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/comtreeData"
)

type Consumer struct {
	client.Baseclient
}

func (c *Consumer) Init() *Consumer {
	c.ServiceName = "rpc_comtreedata"
	c.ServicePath = "Consumer"
	return c
}

func (c *Consumer) AddConsumerRpc(ctx context.Context, args *comtreeData.ArgsAddConsumer, reply *bool) error {
	return c.Call(ctx, "AddConsumerRpc", args, reply)
}
