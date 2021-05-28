package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
)

type PingAn struct {
	client.Baseclient
}

//初始化
func (p *PingAn) Init() *PingAn {
	p.ServiceName = "rpc_order"
	p.ServicePath = "Pingan"
	return p
}

//确认消费 - 平安银行解冻存管资金
func (p *PingAn) ThawDeposAmount(ctx context.Context, args *int, reply *bool) (err error) {
	return p.Call(ctx, "ThawDeposAmount", args, reply)
}

//第三方在途充值-支付完成后，将订单信息归集在待充值表中
func (p *PingAn) PaySuc(ctx context.Context,orderSn *string,reply *bool)(err error){
	return p.Call(ctx, "PaySuc", orderSn, reply)
}

//第三方在途充值-执行充值
func (p *PingAn) PayThirdPayment(ctx context.Context,paymentId *int,reply *bool)(err error){
	return p.Call(ctx, "PayThirdPayment", paymentId, reply)
}
