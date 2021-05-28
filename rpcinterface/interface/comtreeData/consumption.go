/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/3/25 14:28
@Description:

*********************************************/
package comtreeData

import "context"

// 预付卡消费入场
type ArgsConsumption struct {
	ConsumeLogId int // 消费记录号
}

type  Consumption interface {
	// 添加 预付卡消费 信息
	AddConsumptionRpc(ctx context.Context, consumeLogId *int, reply *bool) error
	//添加  预付卡保险出单 信息
	AddInsuranceBillRpc(ctx context.Context, transNo *string, reply *bool)error
	// 添加 预付卡交易 信息
	AddTransactionsRpc(ctx context.Context, orderSn *string, reply *bool) error
	// 添加 全国市场规模 信息
	AddScaleNationalMarketRpc(ctx context.Context, busId *int , reply *bool) error
	//添加 业务规模 信息
	AddScaleBusinessRpc(ctx context.Context, busId *int, reply *bool) error
}