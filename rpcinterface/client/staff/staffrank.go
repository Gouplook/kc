package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

type StaffRank struct {
	client.Baseclient
}

func (s *StaffRank) Init() *StaffRank {
	s.ServiceName = "rpc_staff"
	s.ServicePath = "StaffRank"
	return s
}

//增加员工评价统计Rpc
func (s *StaffRank) AddStaffRankRpc(ctx context.Context, args *staff.StaffRankBase, reply *bool) error {
	return s.Call(ctx, "AddStaffRankRpc", args, reply)
}
