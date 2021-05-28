/**
 * 门店数据统计
 * @Author: yangzhiwu
 * @Date: 2020/9/2 16:32
 */

package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type ShopStatistics struct {
	client.Baseclient
}

func (s *ShopStatistics) Init() *ShopStatistics {
	s.ServiceName = "rpc_bus"
	s.ServicePath = "ShopStatistics"
	return s
}

//支付成功，统计门店的经营数据
func (s *ShopStatistics) StatisBuy(ctx context.Context, orderSn *string, reply *bool) error{
	return s.Call(ctx, "StatisBuy", orderSn, reply)
}

//确认消费完成，统计门店的消费数据
func (s *ShopStatistics) StatisConsume(ctx context.Context, consumeLogId *int, reply *bool) error{
	return s.Call(ctx, "StatisConsume", consumeLogId, reply)
}

//获取门店的营业数据
func (s *ShopStatistics)  ShopStatisData(ctx context.Context, args *bus.ArgsShopStatisData, reply *bus.ReplyShopStatisData) error{
	return s.Call(ctx, "ShopStatisData", args, reply)
}

//获取总店/分店新增会员数
func (s *ShopStatistics)MemberStatisData(ctx context.Context, args *bus.ArgsMemberStatisData, reply *bus.ReplyMemberStatisData) error{
	return s.Call(ctx, "MemberStatisData", args, reply)
}