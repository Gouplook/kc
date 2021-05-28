package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type RelateServiceBase struct {
	SspIds   string // 单项目规格IDs
	SingleId int    // 单项目ID
}

// 关联员工(手艺人)服务入参
type ArgsRelateStaffService struct {
	common.BsToken
	PostId        int  // 岗位ID
	StaffId       int  // 员工ID(前提是该员工的岗位为手艺人)
	AllService    bool // 是否为选择所有服务
	RelateService []RelateServiceBase
}

type ArgsGetStaffByService struct {
	//common.BsToken
	ShopId int
	SingleId int // 默认为0代表全部服务
	SspId    int // 单项目规格
}
type StaffServiceBase struct {
	ID       int    `mapstructure:"id"`
	StaffId  int    `mapstructure:"staff_id"`  // 员工ID(前提是该员工的岗位为手艺人)
	SingleId int    `mapstructure:"single_id"` // 单项目ID,默认为0代表全部服务
	SspIds   string `mapstructure:"ssp_ids"`   // 单项目规格IDs,如果有多个规格中间用","隔开
}

// 根据服务和服务规格获取手艺人
type ReplyGetStaffByService struct {
	StaffId   int    `mapstructure:"staff_id"` //员工ID
	Name      string //员工姓名
	ImgId     int    `mapstructure:"img_id"`
	AvatarUrl string //头像地址
}

type ArgsGetServiceByStaffId struct {
	StaffId int
}
type ReplyGetServiceByStaffId struct {
	GetAll    bool  // 是否关联所有服务
	SspIds    []int // 单项目规格IDs
	SingleIds []int // 单项目IDs
}

type StaffService interface {
	// 关联员工(手艺人)服务
	RelateStaffService(ctx context.Context, args *ArgsRelateStaffService, reply *bool) error
	// 根据商品服务和服务规格获取手艺人
	GetStaffByService(ctx context.Context, args *ArgsGetStaffByService, reply *[]ReplyGetStaffByService) error
	// 员工关联商品服务情况,内部使用
	GetServiceByStaffId(ctx context.Context, args *ArgsGetServiceByStaffId, reply *ReplyGetServiceByStaffId) error
}
