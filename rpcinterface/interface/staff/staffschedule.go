package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// @author liyang<654516092@qq.com>
// @date  2020/4/30 9:16

//添加班次入参
type ArgsAddScheduleSetting struct {
	 common.Utoken   //用户信息
	 common.BsToken  //企业/商户/分店信息
	 Name string //班次名称
	 StartTimePoint string //班次开始时间点，格式如:10:00
	 EndTimePoint string   //班次结束时间点，格式如:22:00
	 IsCrossDay int //是否跨日 0=否 1=是
}
//编辑班次入参
type ArgsSetScheduleSetting struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/分店ID
	Id    int    //班次ID
	Name string //班次名称
	StartTimePoint string //班次开始时间点，格式如:10:00
	EndTimePoint string   //班次结束时间点，格式如:22:00
	IsCrossDay int //是否跨日 0=否 1=是
}

//班次设置返回信息
type ReplyScheduleSetting struct {
	Id int //班次ID
}

//获取班次列表
type ArgsGetScheduleSetting struct {
	common.Utoken //用户信息
	common.BsToken//企业/商户/分店ID
}
//班次列表返回信息
type ReplyScheduleSettingInfo struct {
	Id int //班次列表
	Name string //班次名称
	StartTimePoint string //班次开始时间点
	EndTimePoint   string //班次结束时间点
	IsCrossDay     int  //班次是否跨日 0=未跨日 1=已跨日
}
//设置员工排班入参
type ArgsSetStaffSchedule struct {
	 common.Utoken  //用户信息
	 common.BsToken //企业/商户/分店ID
	 StaffId  int //员工ID
	 ScheduleData []Schedule
}
//设置员工排版具体参数信息
type Schedule struct {
	WeekDay int //周几 1=周一 2=周二 3=周三 4=周四 5=周五 6=周六 7=周日
	SchduleSettingId int //班次 当传0时，为休息
}
//设置员工排班返回信息
type ReplySetStaffSchedule struct {
	StaffId int //员工ID
}

//获取员工排班列表入参
type ArgsGetStaffSchedule struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/分店ID
	common.Paging  //分页信息
	Keywords string //搜索关键字，员工手机号、姓名
}

//获取员工排班列表返回信息
type ReplyGetStaffSchedule struct {
	Data []ReplyStaffSchedule  //员工
	TotalNum int 
}
//员工排班列表返回信息
type ReplyStaffSchedule struct {
	StaffId int //员工ID
	StaffName string //员工姓名
	ScheduleData map[int]StaffSchedule //员工排班数据 key=周几的id
}
//员工排班班次信息
type StaffSchedule struct {
	WeekDay int //周几 1=周一 2=周二 3=周三 4=周四 5=周五 6=周六 7=周日
	ScheduleSettingId int //班次ID 当传0时，为休息
	ScheduleSettingStart string //班次起始时间
	ScheduleSettingEnd string //班次结束时间
	ScheduleSettingName string //班次名称
}

//单员工排班详情入参
type ArgsStaffScheduleInfo struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/分店信息
	StaffId int    //员工ID
}

//多个员工当日的排班时间入参
type ArgsStaffTime struct {
	TimeStr string
	StaffIds []int
}
//多个员工当日的排班时间返回 并集
type ReplyStaffTime struct {
	StartTime int
	StartTimeStr string
	EndTime int
	EndTimeStr string
}

type StaffSchdule interface {
	//添加班次
	AddScheduleSetting(ctx context.Context,args *ArgsAddScheduleSetting,reply *ReplyScheduleSetting) error
	//编辑班次
	SetScheduleSetting(ctx context.Context,args *ArgsSetScheduleSetting,reply *ReplyScheduleSetting) error
	//获取班次列表
	GetScheduleSetting(ctx context.Context,args *ArgsGetScheduleSetting,reply *[]ReplyScheduleSettingInfo) error
	//设置员工排版
	SetStaffSchedule(ctx context.Context,args *ArgsSetStaffSchedule,reply *ReplySetStaffSchedule) error
	//获取员工排班列表
	GetStaffSchedule(ctx context.Context,args *ArgsGetStaffSchedule,reply *[]ReplyGetStaffSchedule) error
}

