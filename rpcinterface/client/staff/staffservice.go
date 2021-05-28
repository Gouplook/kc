package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

type StaffService struct {
	client.Baseclient
}

//初始化
func (s *StaffService) Init() *StaffService {
	s.ServiceName = "rpc_staff"
	s.ServicePath = "StaffService"
	return s
}

// 关联员工(手艺人)服务
func (s *StaffService) RelateStaffService(ctx context.Context, args *staff.ArgsRelateStaffService, reply *bool) error {
	return s.Call(ctx, "RelateStaffService", args, reply)
}

// 根据服务和服务规格获取手艺人
func (s *StaffService) GetStaffByService(ctx context.Context, args *staff.ArgsGetStaffByService, reply *[]staff.ReplyGetStaffByService) error {
	return s.Call(ctx, "GetStaffByService", args, reply)
}

// 员工关联商品服务情况,内部使用
func (s *StaffService) GetServiceByStaffId(ctx context.Context, args *staff.ArgsGetServiceByStaffId, reply *staff.ReplyGetServiceByStaffId) error {
	return s.Call(ctx, "GetServiceByStaffId", args, reply)
}
