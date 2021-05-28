/**
 * @Author: YangYun
 * @Date: 2020/4/16 9:50
 */
package bus

import (
    "context"
    "git.900sui.cn/kc/rpcinterface/client"
    "git.900sui.cn/kc/rpcinterface/interface/bus"
    "git.900sui.cn/kc/rpcinterface/interface/common"
)

type MemberLevel struct {
    client.Baseclient
}
//初始化
func (m *MemberLevel)Init() *MemberLevel {
    m.ServiceName = "rpc_bus"
    m.ServicePath = "MemberLevel"
    return m
}

func (m *MemberLevel) GetLevel(ctx context.Context, args *common.BsToken, reply *[]bus.Level) error {
    return m.Call(ctx, "GetLevel", args, reply)
}

func (m *MemberLevel) GetBusLevel(ctx context.Context, args *common.BsToken, reply *[]bus.BusLevel) error {
    return m.Call(ctx, "GetBusLevel", args, reply)
}

func (m *MemberLevel) GetBusLevel2(ctx context.Context, args *int, reply *[]bus.BusLevel) error {
    return m.Call(ctx, "GetBusLevel2", args, reply)
}

func (m *MemberLevel) GetBusLevelDetail(ctx context.Context, args *bus.ArgsBusLevel, reply *bus.BusLevelDetail) error {
    return m.Call(ctx, "GetBusLevelDetail", args, reply)
}

// 获取特定会员等级详细信息
func (m *MemberLevel)GetBusLevelDetail2(ctx context.Context, args *bus.ArgsGetBusLevelDetail2, reply *bus.BusLevelDetail) error{
    return m.Call(ctx, "GetBusLevelDetail2", args, reply)
}

func (m *MemberLevel) UpdateBusLevel(ctx context.Context, args *bus.ArgsBusLevelDetail, reply *bool) error {
    return m.Call(ctx, "UpdateBusLevel", args, reply)
}

func (m *MemberLevel) VerfiyBusLevel(ctx context.Context, level *int, reply *bool) error {
    return m.Call(ctx, "VerfiyBusLevel", level, reply)
}

//批量获取用户会员等级
func (m *MemberLevel)GetUserLevelByUidsRpc(ctx context.Context,args *bus.ArgsGetUserLevelByUids,reply *[]bus.ReplyGetUserLevelByUids)error{
    return m.Call(ctx, "GetUserLevelByUidsRpc", args, reply)
}