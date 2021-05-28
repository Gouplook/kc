package financial

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/financial"
)

type Bank struct {
	client.Baseclient
}

func (b *Bank) Init() *Bank {
	b.ServiceName = "rpc_financial"
	b.ServicePath = "Bank"
	return b
}

//统计银行/保险-商户数量
func (b *Bank) StatisticsBankOrInsuranceBusNum(ctx context.Context, busId *int, reply *bool) error {
	return b.Call(ctx, "StatisticsBankOrInsuranceBusNum", busId, reply)
}

//统计银行/保险-发卡金额(购卡成功后消费调用)
func (b *Bank) StatisticsBankOrInsuranceSalesCardAssets(ctx context.Context, orderSn *string, reply *bool) error {
	return b.Call(ctx, "StatisticsBankOrInsuranceSalesCardAssets", orderSn, reply)
}

//统计银行-存管金额（购卡/复充成功后消费调用）
func (b *Bank) StatisticsBankDepositAssets(ctx context.Context, orderSn *string, reply *bool) error {
	return b.Call(ctx, "StatisticsBankDepositAssets", orderSn, reply)
}

//确认消费时，根据付款时间释放银行的存管金额
func (b *Bank) ConsumeRelaseDepositAssetsRpc(ctx context.Context, args *financial.ArgsConsumeRelaseDepositAssets, reply *bool) error {
	return b.Call(ctx, "ConsumeRelaseDepositAssetsRpc", args, reply)
}

//================================API接口================================

//GetTopBankBusNum 获取top银行商户总数
func (b *Bank) GetTopBankBusNum(ctx context.Context, args *financial.ArgsGetTopBank, reply *financial.ReplyTopBankList) error {
	return b.Call(ctx, "GetTopBankBusNum", args, reply)
}

//GetTopBankDepositAssets 获取当前银行存管金额
func (b *Bank) GetTopBankDepositAssets(ctx context.Context, args *financial.ArgsGetTopBank, reply *financial.ReplyTopBankList) error {
	return b.Call(ctx, "GetTopBankDepositAssets", args, reply)
}

//GetMonthBankBusNum 获取月银行商户总数
func (b *Bank) GetMonthBankBusNum(ctx context.Context, args *financial.ArgsGetBankBusMonthNum, reply *financial.ReplyMonthTotalNumList) error {
	return b.Call(ctx, "GetMonthBankBusNum", args, reply)
}

//获取银行商户总数
func (b *Bank) GetBankBusNum(ctx context.Context, args *financial.ArgsGetBankBusNum, reply *financial.ReplyTotalNum) error {
	return b.Call(ctx, "GetBankBusNum", args, reply)
}

//GetBankDepositAssetsMonth 获取每月银行存管金额
func (b *Bank) GetBankDepositAssetsMonth(ctx context.Context, args *financial.ArgsGetBankCurrentAssets, reply *financial.ReplyTotalAssetsMonthList) error {
	return b.Call(ctx, "GetBankDepositAssetsMonth", args, reply)
}

//获取当前银行存管金额
func (b *Bank) GetBankCurrentDepositAssets(ctx context.Context, args *financial.ArgsGetBankCurrentAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankCurrentDepositAssets", args, reply)
}

//获取银行累计存管金额
func (b *Bank) GetBankTotalDepositAssets(ctx context.Context, args *financial.ArgsGetBankTotalAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankTotalDepositAssets", args, reply)
}

//获取银行当前发卡金额
func (b *Bank) GetBankCurrentSalesCardAssets(ctx context.Context, args *financial.ArgsGetBankCurrentAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankCurrentSalesCardAssets", args, reply)
}

//获取银行累计发卡金额
func (b *Bank) GetBankTotalSalesCardAssets(ctx context.Context, args *financial.ArgsGetBankTotalAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankTotalSalesCardAssets", args, reply)
}

//获取银行当前普惠金融业绩
func (b *Bank) GetBankCurrentInclusiveFinanceAssets(ctx context.Context, args *financial.ArgsGetBankCurrentAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankCurrentInclusiveFinanceAssets", args, reply)
}

//获取银行累计普惠金融业绩
func (b *Bank) GetBankTotalInclusiveFinanceAssets(ctx context.Context, args *financial.ArgsGetBankTotalAssets, reply *financial.ReplyTotalAssets) error {
	return b.Call(ctx, "GetBankTotalInclusiveFinanceAssets", args, reply)
}


//获取所有支持的银行
func (b *Bank)  GetAllBank(ctx context.Context, args *int, reply *[]financial.ReplyGetAllBank ) (err error){
	return b.Call(ctx, "GetAllBank", args, reply)
}