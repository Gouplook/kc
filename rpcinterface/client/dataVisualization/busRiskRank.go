/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 16:29
@Description:

*********************************************/
package dataVisualization

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/dataVisualization"
)

type BusRiskRank struct {
	client.Baseclient
}

func (b *BusRiskRank)Init()*BusRiskRank {
	b.ServiceName = "rpc_visualization"
	b.ServicePath = "BusRiskRank"
	return b
}

func (b *BusRiskRank)UpdateBusRisRandSafeCode(ctx context.Context, args *dataVisualization.ArgsBusRisRandSafeCode, reply *bool) error{
	return b.Call(ctx, "UpdateBusRisRandSafeCode", args, reply)
}

//修改商家的区域信息
func (b *BusRiskRank) ChangeBusArea(ctx context.Context, args *dataVisualization.ArgsChangeBusArea, reply *bool) error  {
	return b.Call(ctx, "ChangeBusArea", args, reply)
}