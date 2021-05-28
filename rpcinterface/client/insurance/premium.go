package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/insurance"
)

// 保费试算
// @author liyang<654516092@qq.com>
// @date  2020/7/27 16:38

type Premium struct {
	client.Baseclient
}

func (p *Premium) Init() *Premium {
	p.ServiceName = "rpc_insurance"
	p.ServicePath = "Premium"
	return p
}

//保费试算
func (p *Premium) GetCalcuPremium(ctx context.Context, args *insurance.ArgsCalcuPremium, reply *insurance.ReplyCalcuPremium) error {
	return p.Call(ctx, "GetCalcuPremium", args, reply)
}
