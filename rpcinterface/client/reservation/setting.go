package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type Setting struct {
	client.Baseclient
}

func (s *Setting) Init() *Setting {
	s.ServiceName = "rpc_reservation"
	s.ServicePath = "Setting"
	return s
}

//总店开启/关闭预约功能
func (s *Setting) BusOpenCloseReservSwitch(ctx context.Context, args *reservation.ArgsBusOpenCloseReservSwitch, reply *bool) (err error) {
	err = s.Call(ctx, "BusOpenCloseReservSwitch", args, reply)
	return
}

//获取预约功能开启状态
func (s *Setting) GetReservSwitchStatus(ctx context.Context, args *reservation.ArgsGetReservSwitchStatus, reply *reservation.ReplyGetReservSwitchStatus) (err error) {
	err = s.Call(ctx, "GetReservSwitchStatus", args, reply)
	return
}

// 获取配置
func (s *Setting) GetSetting(ctx context.Context, args *reservation.GetSettingParams, reply *reservation.GetSettingReplies) (err error) {
	err = s.Call(ctx, "GetSetting", args, reply)
	return
}

// 店面创建预约设置
func (s *Setting) CreateSetting(ctx context.Context, args *reservation.CreateAndEditSettingParams, reply *bool) (err error) {
	err = s.Call(ctx, "CreateSetting", args, reply)
	return
}

// 店面编辑预约设置
func (s *Setting) EditSetting(ctx context.Context, args *reservation.CreateAndEditSettingParams, reply *bool) (err error) {
	err = s.Call(ctx, "EditSetting", args, reply)
	return
}

// 店面开通预约设置
func (s *Setting) EnableSetting(ctx context.Context, args *reservation.CreateAndEditSettingParams, reply *bool) (err error) {
	err = s.Call(ctx, "EnableSetting", args, reply)
	return
}

// disable店面预约设置
func (s *Setting) ChangeSettingStatus(ctx context.Context, args *reservation.ChangeSettingStatus, reply *bool) (err error) {
	err = s.Call(ctx, "ChangeSettingStatus", args, reply)
	return
}

// 批量获取分店预约设置信息
func (s *Setting) GetSetttingShops(ctx context.Context, shopIds *[]int, reply *[]reservation.ReplySettingShops) (err error) {
	err = s.Call(ctx, "GetSetttingShops", shopIds, reply)
	return
}
