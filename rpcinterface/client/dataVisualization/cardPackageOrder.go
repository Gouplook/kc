/******************************************
@Description:
@Time : 2020/11/30 13:57
@Author :lixiaojun

*******************************************/
package dataVisualization
import (
	"git.900sui.cn/kc/rpcinterface/client"
	"context"
	//"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"

)
type CardPackageOrder struct {
	client.Baseclient
}
func (c *CardPackageOrder) Init() *CardPackageOrder {
	c.ServiceName = "rpc_visualization"
	c.ServicePath = "CardPackageOrder"
	return c
}
//根据relationId创建监管可视化-卡包订单信息
func (c *CardPackageOrder)CreateOrderInfoByRelationId(ctx context.Context, relationId *int, reply *bool) error{
	return c.Call(ctx,"CreateOrderInfoByRelationId",relationId,reply)
}
