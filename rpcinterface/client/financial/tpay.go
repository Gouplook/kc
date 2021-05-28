package financial

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/financial"
)

type Tpay struct {
	client.Baseclient
}

func (t *Tpay) Init() *Tpay {
	t.ServiceName = "rpc_financial"
	t.ServicePath = "Tpay"
	return t
}

//GetTopTpayBusNum 获取top第三方支付的商户总数
func (t *Tpay) GetTopTpayBusNum(ctx context.Context, tpayType *financial.ArgsGetTopTpay, reply *financial.ReplyTopTpayList) (err error) {
	return t.Call(ctx, "GetTopTpayBusNum", tpayType, reply)
}

//GetTopTpayAssetsTotal 获取top第三方支付的总收单金额
func (t *Tpay) GetTopTpayAssetsTotal(ctx context.Context, tpayType *financial.ArgsGetTopTpay, reply *financial.ReplyTopTpayList) (err error) {
	return t.Call(ctx, "GetTopTpayAssetsTotal", tpayType, reply)
}

//GetTpayBusNumMonth 获取每月第三方支付的商户总数
func (t *Tpay) GetTpayBusNumMonth(ctx context.Context, tpayType *financial.ArgsGetTpay, reply *financial.ReplyGetTpayBusNumMonthList) (err error) {
	return t.Call(ctx, "GetTpayBusNumMonth", tpayType, reply)
}

//GetTpayAssetsMonth 获取每月第三方支付的总收单金额
func (t *Tpay) GetTpayAssetsMonth(ctx context.Context, tpayType *financial.ArgsGetTpay, reply *financial.ReplyGetTpayAssetsMonthList) (err error) {
	return t.Call(ctx, "GetTpayAssetsMonth", tpayType, reply)
}

//获取单个第三方支付的数据
func (t *Tpay) GetOneTpayData(ctx context.Context, tpayType *int, reply *financial.TpayData) (err error) {
	return t.Call(ctx, "GetOneTpayData", tpayType, reply)
}

//获取全部第三方支付的数据
func (t *Tpay) GetAllData(ctx context.Context, tpayType *int, reply *[]financial.TpayData) (err error) {
	return t.Call(ctx, "GetAllData", tpayType, reply)
}

//获取所有第三方支付类型
func (t *Tpay) GetAllTpay(ctx context.Context, args *int, reply *[]financial.ReplyGetAllTpay) (err error) {
	return t.Call(ctx, "GetAllTpay", args, reply)
}
