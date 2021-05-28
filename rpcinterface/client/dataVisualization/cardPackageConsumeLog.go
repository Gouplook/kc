/******************************************
@Description:监管可视化-卡包消费记录
@Time : 2020/11/30 13:58
@Author :lixiaojun

*******************************************/
package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"
)

type CardPackageConsumeLog struct {
	client.Baseclient
}

func (c *CardPackageConsumeLog) Init() *CardPackageConsumeLog {
	c.ServiceName = "rpc_visualization"
	c.ServicePath = "CardPackageConsumeLog"
	return c
}

//记录卡包消费记录数据到监管可视化表中
func (c *CardPackageConsumeLog) SetCardPackageConsumeLog(ctx context.Context, args *dataVisualization.ArgsSetCardPackageConsumeLog, reply *bool) error {
	return c.Call(ctx, "SetCardPackageConsumeLog", args, reply)
}
