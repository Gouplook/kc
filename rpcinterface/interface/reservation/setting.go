package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"time"
)

//总店开启/关闭预约功能-入参
type ArgsBusOpenCloseReservSwitch struct {
	common.BsToken
	Status int //否开启预约功能 0 关闭, 1 开启
}

//获取预约开启状态-入参
type ArgsGetReservSwitchStatus struct {
	BusId int
}

//获取预约开启状态-出参
type ReplyGetReservSwitchStatus struct {
	Status int
}

//添加/编辑预约设置-入参
type CreateAndEditSettingParams struct {
	common.BsToken                 //店面id
	SettingId       int            // 预约设置ID
	ShopId          int            //此参数只有当总店设置指定分店时传入
	WeekDate        WeekDateParams //日期开启状态
	BusinessBegin   string         //营业开始时间 '09:35'
	BusinessEnd     string         //营业结束时间 '24:45'
	IsCrossDay      int            // 是否跨日 0:否,1:是
	TimeGranularity int            //预约时间粒度 eg:15,30,45m
	MaxCount        int            //每天预约的最大人数
	NoticeTime      float64        // 提前多长时间通知预约用户，单位小时
	//NumberPerTime        int    // 店铺每段时间可预约人数
	//NumberPerTimePerTech int    // 每个手艺人在每段时间最多可预约人数
	//MaxTime              int    //可提前预约的最大时间:m
	MinTime              int    //可提前预约的最小时间:m
}

//日期
type WeekDateParams struct {
	WeekMonStatus  int //周一状态，0：关闭，1-开启
	WeekTuesStatus int //周二状态，0：关闭，1-开启
	WeekWedStatus  int //周三状态，0：关闭，1-开启
	WeekThurStatus int //周四状态，0：关闭，1-开启
	WeekFriStatus  int //周五状态，0：关闭，1-开启
	WeekSatStatus  int //周六状态，0：关闭，1-开启
	WeekSunStatus  int //周日状态，0：关闭，1-开启
}

type GetSettingParams struct {
	common.BsToken
	ShopID int //店面ID
}

type ChangeSettingStatus struct {
	common.BsToken //店面ID
	ShopId         int
	Status         int //0-关闭;1-开启
}

type GetSettingReplies struct {
	SettingID         int            `mapstructure:"setting_id"`     //预约设置id
	ShopID            int            `mapstructure:"shop_id"`        //店面ID
	SettingEnable     int            `mapstructure:"setting_enable"` //setting_enable:0-关闭;1-开启
	WeekDate          WeekDateParams //日期开启状态
	BusinessBeginTime string         `mapstructure:"business_begin_time"` //预约开始时间
	BusinessEndTime   string         `mapstructure:"business_end_time"`   //预约结束时间
	TimeGranularity   int            `mapstructure:"time_granularity"`    //时间粒度
	IsCrossDay        int            `mapstructure:"is_cross_day"`
	MaxCount          int            `mapstructure:"max_count"` //店铺每天预约的最大人数
	NoticeTime        float64            // 提前多长时间通知预约用户，单位小时
	//MaxTime              int    `mapstructure:"max_time"`                 //可提前预约的最大时间
	MinTime              int    `mapstructure:"min_time"`                 //可提前预约的最小时间
	//RemainTime           int    `mapstructure:"remain_time"`              //提前remain_time提醒预约用户
	//NumberPerTime        int    `mapstructure:"number_per_time"`          // 店铺每段时间可预约人数,
	//NumberPerTimePerTech int    `mapstructure:"number_per_time_per_tech"` // 每个手艺人在每段时间最多可预约人数
	//TimeDuration         int    `mapstructure:"time_duration"`            // 客户预约服务时固定时长,默认15m
}

type SetSettingParams struct {
	common.BsToken                    //店面id
	BeginTime               time.Time //营业开始时间
	BusinessDuration        int       //营业时长
	DistributionGranularity int       //时间分配粒度
	MaxTime                 int       //最大提前预约时间
	MinTime                 int       //最下提前预约时间
	MaxCount                int       //每天可预约的最大数量
	RemainderTime           int       //预约提前多长时间提醒
}

//批量获取分店预约设置返回信息
type ReplySettingShops struct {
	ShopId        int //分店ID
	SettingEnable int //开启状态 0=未开启/关闭 1=开启
}

type Setting interface {
	//总店开启/关闭预约功能
	BusOpenCloseReservSwitch(ctx context.Context, args *ArgsBusOpenCloseReservSwitch, reply *bool) error
	//获取总店预约功能开启状态
	GetReservSwitchStatus(ctx context.Context, args *ArgsGetReservSwitchStatus, reply *ReplyGetReservSwitchStatus) error
	// 获取预约配置
	GetSetting(ctx context.Context, args *GetSettingParams, reply *GetSettingReplies) error
	// 店面创建预约设置
	CreateSetting(ctx context.Context, args *CreateAndEditSettingParams, reply *bool) error
	// 店面编辑预约设置
	EditSetting(ctx context.Context, args *CreateAndEditSettingParams, reply *bool) error
	// 店面开通预约设置
	EnableSetting(ctx context.Context, args *CreateAndEditSettingParams, reply *bool) error
	// 更改预约设置状态
	ChangeSettingStatus(ctx context.Context, args *ChangeSettingStatus, reply *bool) error
}
