package financial

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/financial"
)

type Insurance struct {
	client.Baseclient
}

func (i *Insurance) Init() *Insurance {
	i.ServiceName = "rpc_financial"
	i.ServicePath = "Insurance"
	return i
}

//保险公司上月保费金额
func (i *Insurance) AddInsAssetsMonth(ctx context.Context, args *financial.ArgsInsAssetsMonth, reply *bool) error {
	return i.Call(ctx, "AddInsAssetsMonth", args, reply)

}

//保险公司累计保费金额
func (i *Insurance) AddInsAssetsTotal(ctx context.Context, args *financial.ArgsInsAssetsTotal, reply *bool) error {
	return i.Call(ctx, "AddInsAssetsTotal", args, reply)
}

//================================API接口================================

//GetTopInsBusTotal 获取top保险公司商户统计
func (i *Insurance) GetTopInsBusTotal(ctx context.Context, args *financial.ArgsGetTopInsurance, reply *financial.ReplyTopInsuranceList) error {
	return i.Call(ctx, "GetTopInsBusTotal", args, reply)
}

//GetTopInsAssetsAmount 获取top保险公司月保费金额
func (i *Insurance) GetTopInsAssetsAmount(ctx context.Context, args *financial.ArgsGetTopInsurance, reply *financial.ReplyTopInsuranceList) error {
	return i.Call(ctx, "GetTopInsAssetsAmount", args, reply)
}

//GetMonthInsBusNum 获取每月保险公司商户统计
func (i *Insurance) GetMonthInsBusNum(ctx context.Context, args *financial.ArgsGetAssetsMonth, reply *financial.ReplyGetMonthInsBusNumList) error {
	return i.Call(ctx, "GetMonthInsBusNum", args, reply)
}

//GetMonthInsAssetsMonth 获取每月保险公司保费金额
func (i *Insurance) GetMonthInsAssetsMonth(ctx context.Context, args *financial.ArgsGetAssetsMonth, reply *financial.ReplyGetMonthInsAssetsMonthList) error {
	return i.Call(ctx, "GetMonthInsAssetsMonth", args, reply)
}

// 获取保险公司月保费金额
func (i *Insurance) GetInsAssetsMonth(ctx context.Context, args *financial.ArgsGetAssetsMonth, reply *financial.ReplySalesCardMonth) error {
	return i.Call(ctx, "GetInsAssetsMonth", args, reply)
}

// 获取保险公司累计发卡金额
func (i *Insurance) GetInsAssetsAmount(ctx context.Context, args *financial.ArgsGetAssetsAmount, reply *financial.ReplySalesCardAmount) error {
	return i.Call(ctx, "GetInsAssetsAmount", args, reply)
}

// 获取保险公司商户统计(当前银行， 所有银行商户总数）
func (i *Insurance) GetInsBusTotal(ctx context.Context, args *financial.ArgsInsBusType, reply *financial.ReplyBusNum) error {
	return i.Call(ctx, "GetInsBusTotal", args, reply)
}

// 获取当月保险发卡金额
func (i *Insurance) GetInsIssueCardMonth(ctx context.Context, args *financial.ArgsIssueCardMonth, reply *financial.ReplyIssueCardMonth) (err error) {
	return i.Call(ctx, "GetInsIssueCardMonth", args, reply)
}

// 获取累计保险发卡金额统计
func (i *Insurance) GetInsIssueCardAmount(ctx context.Context, args *financial.ArgsIssueCardAmount, reply *financial.ReplyIssueCardAmount) (err error) {
	return i.Call(ctx, "GetInsIssueCardAmount", args, reply)
}

//获取所有保险公司类型
func (i *Insurance) GetAllInsurance(ctx context.Context, args *int, reply *[]financial.ReplyGetAllInsurance) (err error) {
	return i.Call(ctx, "GetAllInsurance", args, reply)
}
