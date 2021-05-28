//单项目订单接口定义
//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/22 17:14

package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//下单单项目数据结构
type SingleData struct {
	SingleId int //单项目id
	SsId int 	//单项目在门店的id
	SspId int 	//子规格组合主键id
	StaffIds []int //销售ids
	CraftStaffIds []int //手艺人ids
	Num int //购买数量
	ChangePrice float64 //手动改价后的单件价格
	//一期优化
	IcardPackageId int //身份卡卡包id
	CardPackageType int //卡包类型
	RelationId int //卡包关联id
}

//生成单项目订单参数数据接口体
type SingleOrderData struct {
	Singles []SingleData //购买的单项目
	OrderSource int //订单渠道 使用channel

}

//saas后台下单的入参数据
type ArgsSaasCreateSingleOrder struct {
	common.BsToken
	common.Utoken //当前操作人员登录信息
	SingleOrderData
	Uid int  //购买人uid
	ReservationId int //预约id
	IcardPackageId int //身份卡卡包id
	TempId int //挂单id
}

//saas后台下单返回数据结构
type ReplySaasCreateSingleOrder struct {
	OrderSn string //订单号
	Ctime int64 //下单时间
	//BuyerName string //客户姓名
	//Cashier string //收银员
	PayTypes []PayType //saas的付款方式
	TotalAmount float64 //应支付总金额
}

//前端用户下单的入参数据
type ArgsUserCreateSingleOrder struct {
	common.Utoken
	ShopId int //门店id
	SingleOrderData
	IcardPackageId int //身份卡卡包id
}

//前端用户下单返回的数据格式
type ReplyUserCreateSingleOrder struct {
	OrderSn string //订单号
	TotalAmount float64 //应支付总金额
	PayTypes []PayType //前端的付款方式
}


type SingleOrder interface {
	//SaaS创建单项目订单方法
	SaasCreateSingleOrder(ctx context.Context, args *ArgsSaasCreateSingleOrder, reply *ReplySaasCreateSingleOrder ) error
	//用户创建订单
	UserCreateSingleOrder(ctx context.Context, args *ArgsUserCreateSingleOrder, reply *ReplyUserCreateSingleOrder ) error
}


