/******************************************
@Description:监管可视化-充值卡充值订单
@Time : 2020/11/30 13:56
@Author :lixiaojun

*******************************************/
package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"
)

type CardPackageRcardRecharge struct {
	client.Baseclient
}

func (c *CardPackageRcardRecharge) Init() *CardPackageRcardRecharge {
	c.ServiceName = "rpc_visualization"
	c.ServicePath = "CardPackageRcardRecharge"
	return c
}

//记录充值订单数据到监管可视化表中
func (c *CardPackageRcardRecharge) SetCardPackageRcardRecharge(ctx context.Context, args *dataVisualization.ArgsCardPackageRcardRecharge, reply *bool) error {
	return c.Call(ctx, "SetCardPackageRcardRecharge", args, reply)
}
