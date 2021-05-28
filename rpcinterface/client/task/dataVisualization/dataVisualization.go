/******************************************
@Description:监管可视化-充值卡充值订单 task
@Time : 2020/11/30 17:00
@Author :lixiaojun

*******************************************/
package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/dataVisualization"
)

type DataVisualization struct {
	client.Baseclient
}

func (c *DataVisualization) Init() *DataVisualization {
	c.ServiceName = "rpc_task"
	c.ServicePath = "DataVisualization"
	return c
}

//设置充值卡充值订单到交换机
func (c *DataVisualization) SetCardPackageRcardRecharge(ctx context.Context, args *dataVisualization.CardPackageRcardRecharge, reply *bool) error {
	return c.Call(ctx, "SetCardPackageRcardRecharge", args, reply)
}
