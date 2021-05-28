package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

//IcardOrderClient IcardOrderClient
type IcardOrderClient struct {
	client.Baseclient
}

//Init Init
func (i *IcardOrderClient) Init() *IcardOrderClient {
	i.ServiceName = "rpc_order"
	i.ServicePath = "IcardOrder"
	return i
}

//GetIcardListByUserID GetIcardListByUserID
func (i *IcardOrderClient) GetIcardListByUserID(ctx context.Context, args *order.InputParamsICardCanUse, reply *order.OutputParamsICardCanUse) error {
	return i.Call(ctx, "GetIcardListByUserID", args, reply)
}
