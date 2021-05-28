/**
 * @Author: YangYun
 * @Date: 2020/4/15 17:35
 */
package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	RebateTypeNone = 0
	RebateTypeAll  = 1
	RebateTypePart = 2

	InitLevel = 1
)

type Level struct {
	Level int    // 会员等级
	Name  string // 会员名称
}
type BusLevel struct {
	Level         int    // 会员等级
	Name          string // 会员名称
	Growth        int    // 成长值
	ServiceRebate byte   // 服务折扣
	ProductRebate byte   // 商品折扣
	ServiceType   byte   // 折扣类型  0=无折扣 1=全部折扣 2=部分折扣
	ProductType   byte   // 折扣类型  0=无折扣 1=全部折扣 2=部分折扣
}
type RebateService struct {
	ServiceId int
	Rebate    byte
}
type RebateProduct struct {
	ProductId int
	Rebate    byte
}
type BusLevelDetail struct {
	Mid int //会员id
	BusLevel
	RebateServices []RebateService
	RebateProducts []RebateProduct
}
type ArgsBusLevel struct {
	common.BsToken

	Level int
}
type ArgsBusLevelDetail struct {
	common.BsToken

	BusLevelDetail
}

type ArgsGetUserLevelByUids struct {
	Uids []int
	ShopId int
	BusId int
}
type ReplyGetUserLevelByUids struct {
	Uid int
	Level int
	Name string
	MemberId int
	MemberName string
	Phone string
}

type ArgsGetBusLevelDetail2 struct {
	Uid int
	ShopId int
}

type MemberLevel interface {
	// 获取会员等级列表
	GetLevel(ctx context.Context, args *common.BsToken, reply *[]Level) error
	// 获取会员等级详情列表
	GetBusLevel(ctx context.Context, args *common.BsToken, reply *[]BusLevel) error
	// 获取会员等级详情列表2
	GetBusLevel2(ctx context.Context, args *int, reply *[]BusLevel) error
	// 获取特定会员等级详细信息
	GetBusLevelDetail(ctx context.Context, args *ArgsBusLevel, reply *BusLevelDetail) error
	// 获取特定会员等级详细信息
	GetBusLevelDetail2(ctx context.Context, args *ArgsGetBusLevelDetail2, reply *BusLevelDetail) error
	// 修改会员等级信息
	UpdateBusLevel(ctx context.Context, args *ArgsBusLevelDetail, reply *bool) error
	// 验证会员等级
	VerfiyBusLevel(ctx context.Context, level *int, reply *bool) error
	//批量获取用户会员等级
	GetUserLevelByUidsRpc(ctx context.Context,args *ArgsGetUserLevelByUids,reply *[]ReplyGetUserLevelByUids)error
}
