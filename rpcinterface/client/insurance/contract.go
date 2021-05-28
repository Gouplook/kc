package insurance

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/insurance"
)

//Contract 保险合同
type Contract struct {
	client.Baseclient
}

//Init Init
func (c *Contract) Init() *Contract {
	c.ServiceName = "rpc_insurance"
	c.ServicePath = "Contract"
	return c
}

//InsuranceSignUp 保险签约
func (c *Contract) InsuranceSignUp(ctx context.Context, args *insurance.ContractParams, reply *insurance.ContractReply) error {
	return c.Call(ctx, "InsuranceSignUp", args, reply)
}
