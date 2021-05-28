package risk

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/task/bus"
	"git.900sui.cn/kc/rpcinterface/interface/task"
)

/*
	预付卡风险管理系统-不同行业统计的门店数量、会员数量、员工数量
*/



type BusBasicIndusAvg interface {
	//审核商户
	RiskForAuditBus(ctx context.Context,busId *int,reply *bool)error
	//审核店铺
	RiskForAuditShop(ctx context.Context,shopId *int,reply *bool)error
	//员工新增、离职、删除
	RiskForSetStaff(ctx context.Context,staffId *int,reply *bool)error
	//会员新增
	RiskForAddMember(ctx context.Context,memberId *int,reply *bool)error
	//订单支付成功
	RiskForPaySuccess(ctx context.Context,orderSn *string, reply *bool)error
	//店铺面积更新
	RiskForShopAreaUpdate(ctx context.Context,args *bus.ArgsShopAreaUpdate, reply *bool)error
	//新增删除卡项/商品
	RiskForAddDelGoods(ctx context.Context,args *task.ArgsAddDelGoods ,reply *bool)error
}


