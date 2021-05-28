/******************************************
@Description:
@Time : 2020/11/30 13:56
@Author :lixiaojun

*******************************************/

package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"
)

type CardPackagePolicy struct {
	client.Baseclient
}

func (c *CardPackagePolicy)Init() *CardPackagePolicy{
	c.ServiceName = "rpc_visualization"
	c.ServicePath = "CardPackagePolicy"
	return c
}

// 添加监管可视化 保险保单信息
func (c *CardPackagePolicy)AddVisualizationPolicyRpc(ctx context.Context, args *dataVisualization.ArgsCardPolicy, reply *bool) error {
	return c.Call(ctx, "AddVisualizationPolicyRpc",args, reply)
}

