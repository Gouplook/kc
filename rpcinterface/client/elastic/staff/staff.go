package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// @author liyang<654516092@qq.com>
// @date  2020/5/9 10:38

type Staff struct {
	client.Baseclient
}

func (s *Staff) Init() *Staff {
	s.ServiceName = "rpc_elastic"
	s.ServicePath = "Staff/Staff"
	return s
}

//新增、更新员工
func (s *Staff) SetStaff(ctx context.Context, staffId *int, reply *bool) error {
	return s.Call(ctx, "SetStaff", staffId, reply)
}
//搜索员工
func (s *Staff) SearchStaff(ctx context.Context,args *staff.ArgsSearchWhere,reply *staff.ReplySearch) error{
	return s.Call(ctx, "SearchStaff", args, reply)
}

