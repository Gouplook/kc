/******************************************
@Description:
@Time : 2020/11/30 13:57
@Author :lixiaojun

*******************************************/
package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	//"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"
)

type Bus struct {
	client.Baseclient
}

func (b *Bus) Init() *Bus {
	b.ServiceName = "rpc_visualization"
	b.ServicePath = "Bus"
	return b
}

//根据busId创建监管可视化-企业/商户基础信息表
func (b *Bus) CreateBusInfoByBusId(ctx context.Context, busId *int, reply *bool) error {
	return b.Call(ctx, "CreateBusInfoByBusId", busId, reply)
}

//创建/更新 企业/商户分店基础信息
func (b *Bus) CreateOrUpdateBusShopInfo(ctx context.Context, shopId *int, reply *bool) error {
	return b.Call(ctx, "CreateOrUpdateBusShopInfo", shopId, reply)
}
