package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

type Bank struct {
	client.Baseclient
}

//实例化
func (b *Bank) Init() *Bank {
	b.ServiceName = "rpc_public"
	b.ServicePath = "Bank"
	return b
}

//获取网银信息
func (b *Bank) GetInternetBankList(ctx context.Context,args *string ,reply *[]public.ReplyInternetBankList) error {
	return b.Call(ctx, "GetInternetBankList", args, reply)
}

//根据超级网银号获取网银信息返回参数
func (b *Bank) InternetBankInfo(ctx context.Context,EiconBankBranchId string,reply *public.ReplyInternetBankInfo)error {
	return b.Call(ctx, "InternetBankInfo", EiconBankBranchId, reply)
}

