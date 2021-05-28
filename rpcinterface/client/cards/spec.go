//单项目规格
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/7 18:39
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Spec struct {
	client.Baseclient
}

func (s *Spec) Init() *Spec {
	s.ServiceName = "rpc_cards"
	s.ServicePath = "Spec"
	return s
}

//添加规格
func (s *Spec) AddSpec( ctx context.Context, args *cards.SpecInfo, reply *int ) error{
	return s.Call(ctx, "AddSpec", args, reply )
}

//获取子规格
func (s *Spec) GetByParentSpecId( ctx context.Context, args *cards.GetSubSpecParams, reply *[]cards.SubSpec ) error {
	return s.Call(ctx, "GetByParentSpecId", args, reply )
}

//根据sspid 查询子规格名字和id
func (s *Spec)  GetBySspIds(ctx context.Context, args *[]int, reply *map[int][]cards.SubSpec ) (err error) {
	return s.Call(ctx, "GetBySspIds", args, reply )
}