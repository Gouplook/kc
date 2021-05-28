/**
 * @Author: YangYun
 * @Date: 2020/4/13 19:15
 */
package bus
import (
    "context"
    "git.900sui.cn/kc/rpcinterface/client"
)

type Member struct {
    client.Baseclient
}

func (m *Member)Init() *Member {
    m.ServiceName = "rpc_task"
    m.ServicePath = "Bus/Member"
    return m
}

func (m *Member)MemberAdd(ctx context.Context, memberId *int, reply *bool) error {
    return m.Call(ctx, "MemberAdd", memberId, reply)
}