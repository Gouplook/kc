package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// 定义提成设置相关rpc客户端调用方法
// @author yinjinlin<yinjinlin_uplook@163.com>
// @date  2020/10/12 16:23

type StaffRecruit struct {
	client.Baseclient
}

func (s *StaffRecruit) Init() *StaffRecruit {
	s.ServiceName = "rpc_staff"
	s.ServicePath = "StaffRecruit"
	return s
}

//添加招聘信息
func (s *StaffRecruit) AddStaffRecruit(ctx context.Context, args *staff.ArgsAddRecruitInfo, reply *staff.ReplyStaffRecruit) error {
	return s.Call(ctx, "AddStaffRecruit", args, reply)
}

//获取招聘信息列表
func (s *StaffRecruit) GetStaffRecruitList(ctx context.Context, args *staff.ArgsGetStaffRecruitList, reply *staff.ReplyGetStaffRecruitList) error {
	return s.Call(ctx, "GetStaffRecruitList", args, reply)
}

//编辑招聘信息
func (s *StaffRecruit) EditStaffRecruit(ctx context.Context, args *staff.ArgsEditStaffRecruit, reply *staff.ReplyStaffRecruit) error {
	return s.Call(ctx, "EditStaffRecruit", args, reply)
}

//获得招聘信息详情
func (s *StaffRecruit) GetStaffRecruitInfo(ctx context.Context, args *staff.ArgsStaffRecruitInfo, reply *staff.ReplyStaffRecruitInfo) error {
	return s.Call(ctx, "GetStaffRecruitInfo", args, reply)
}

//删除招聘信息
func (s *StaffRecruit) DelStaffRecruit(ctx context.Context, args *staff.ArgsDelStaffRecruit, reply *staff.ReplyStaffRecruit) error {
	return s.Call(ctx, "DelStaffRecruit", args, reply)
}

//批量删除招聘信息
func (s *StaffRecruit) BatchDelStaffRecruit(ctx context.Context, args *staff.ArgsBatchDelStaffRecruit, reply *bool) error {
	return s.Call(ctx, "BatchDelStaffRecruit", args, reply)
}

//暂停招聘信息
func (s *StaffRecruit) SuspendStaffRecruit(ctx context.Context, args *staff.ArgsSuspendReleaseStaffRecruit, reply *staff.ReplySuspendReleaseStaffRecuit) error {
	return s.Call(ctx, "SuspendStaffRecruit", args, reply)
}

//批量暂停招聘信息
func (s *StaffRecruit) BatchSuspendStaffRecruit(ctx context.Context, args *staff.ArgsBatchSusReltaffRecruit, reply *bool) error {
	return s.Call(ctx, "BatchSuspendStaffRecruit", args, reply)
}

//发布招聘信息
func (s *StaffRecruit) ReleaseStaffRecruit(ctx context.Context, args *staff.ArgsSuspendReleaseStaffRecruit, reply *staff.ReplySuspendReleaseStaffRecuit) error {
	return s.Call(ctx, "ReleaseStaffRecruit", args, reply)
}

//批量发布招聘信息
func (s *StaffRecruit) BatchReleaseStaffRecruit(ctx context.Context, args *staff.ArgsBatchSusReltaffRecruit, reply *bool) error {
	return s.Call(ctx, "BatchReleaseStaffRecruit", args, reply)
}

//获取招聘职位
func (s *StaffRecruit) GetRecruitPositions(ctx context.Context, args *staff.ArgsEmpty, reply *staff.ReplyGetRecruitPositions) error {
	return s.Call(ctx, "GetRecruitPositions", args, reply)
}
