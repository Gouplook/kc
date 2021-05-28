package risk

import "context"

/*
	预付卡风险管理系统-不同行业统计的经营流水、会员数量、服务人次、企业总资产表
*/



type BusBasicIndusAvgMonth interface {
	//行业接待人次月度统计
	RiskForIncService(ctx context.Context,busId *int,reply *bool)error
	//行业发卡数量月度统计
	RiskForSaleCard(ctx context.Context,args ArgsSalesCardNum,reply *bool)
}


