package staff
// @author liyang<654516092@qq.com>
// @date  2020/5/8 11:09
import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/task/staff"
)
type Staff struct {
	client.Baseclient
}

func (s *Staff)Init() *Staff {
	s.ServiceName = "rpc_task"
	s.ServicePath = "Staff/Staff"
	return s
}

//消息队列（员工新增、更新）
func (s *Staff)SetStaff(ctx context.Context, staffId *int, reply *bool) error {
	return s.Call(ctx, "SetStaff", staffId, reply)
}

//员工 新增/离职/删除
func (s *Staff)SetAddLeaveLeDeleteStaff(ctx context.Context,args *staff.ArgsSetAddLeaveLeDeleteStaff,reply *bool)error{
	return s.Call(ctx, "SetAddLeaveLeDeleteStaff", args, reply)
}




