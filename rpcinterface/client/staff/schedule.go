package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// @author liyang<654516092@qq.com>
// @date  2020/4/30 15:49

type Schedule struct {
	client.Baseclient
}

//初始化
func (s *Schedule)Init() *Schedule {
	s.ServiceName = "rpc_staff"
	s.ServicePath = "Schedule"
	return s
}

//新增排班班次-分店
func (s *Schedule) AddScheduleSetting(ctx context.Context, args *staff.ArgsAddScheduleSetting, reply *staff.ReplyScheduleSetting) error {
	return s.Call(ctx, "AddScheduleSetting", args, reply)
}

//更新排班班次-分店
func (s *Schedule) SetScheduleSetting(ctx context.Context, args *staff.ArgsSetScheduleSetting, reply *staff.ReplyScheduleSetting) error {
	return s.Call(ctx, "SetScheduleSetting", args, reply)
}

//获取排班班次-分店
func (s *Schedule) GetScheduleSetting(ctx context.Context, args *staff.ArgsGetScheduleSetting, reply *[]staff.ReplyScheduleSettingInfo) error {
	return s.Call(ctx, "GetScheduleSetting", args, reply)
}

//更新/设置员工排班
func (s *Schedule) SetStaffSchedule(ctx context.Context,args *staff.ArgsSetStaffSchedule,reply *staff.ReplySetStaffSchedule) error{
	return s.Call(ctx, "SetStaffSchedule", args, reply)
}

//获取单个员工排班
func (s *Schedule) GetSingleStaffSchedule(ctx context.Context,args *staff.ArgsStaffScheduleInfo,reply *staff.ReplyStaffSchedule) error{
	return s.Call(ctx, "GetSingleStaffSchedule", args, reply)
}

//获取单个员工排班-rpc内部
func (s *Schedule) GetSingleStaffScheduleRpc(ctx context.Context,staffId *int,reply *staff.ReplyStaffSchedule) error{
	return s.Call(ctx, "GetSingleStaffScheduleRpc", staffId, reply)
}

//获取员工排班列表
func (s *Schedule) GetStaffSchedule(ctx context.Context,args *staff.ArgsGetStaffSchedule,reply *staff.ReplyGetStaffSchedule) error{
	return s.Call(ctx, "GetStaffSchedule", args, reply)
}

//获取多个员工排班-RPC
func (s *Schedule) GetStaffSchedules(ctx context.Context,args *[]int,reply *[]staff.ReplyStaffSchedule) error{
	return s.Call(ctx, "GetStaffSchedules", args, reply)
}

//获取多个员工当日上班时间-RPC
func (s *Schedule) GetStaffTime(ctx context.Context,args *staff.ArgsStaffTime,reply *staff.ReplyStaffTime) error{
	return s.Call(ctx, "GetStaffTime", args, reply)
}