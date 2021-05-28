package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/risk"
)

type busBasicIndusAvgMonth struct {
	client.Baseclient
}

func (b *busBasicIndusAvgMonth)Init()*busBasicIndusAvgMonth  {
	b.ServiceName = "rpc_risk"
	b.ServicePath = "busBasicIndusAvgMonth"
	return b
}
//行业发卡数量月度统计
func (b *busBasicIndusAvgMonth)RiskForIncService(ctx context.Context,busId *int, reply *bool)error  {
	return b.Call(ctx, "RiskForIncService", busId, reply)
}
//行业发卡数量月度统计
func (b *busBasicIndusAvgMonth)RiskForSaleCard(ctx context.Context,args *risk.ArgsSalesCardNum,reply *bool)error  {
	return b.Call(ctx, "RiskForSaleCard", args, reply)
}
