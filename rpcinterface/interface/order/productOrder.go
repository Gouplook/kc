//产品订单接口定义
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/25 16:09

package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//提货状态
	PICK_UP_GOODS_STATUS_WAIT = 0	//未提货状态
	PICK_UP_GOODS_STATUS_DONE = 1	//已经提货状态
)

//下单产品数据结构
type ProductData struct {
	DetailId int 	//产品的明细id
	StaffIds []int //销售ids
	Num int //购买数量
	ChangePrice float64 //手动改价后的单件价格
}

//saas后台下单的入参数据
type ArgsSaasCreateProductOrder struct {
	common.BsToken
	common.Utoken //当前操作人员登录信息
	ProductOrderData []ProductData //购买的产品数据
	Uid int  //购买人uid
	OrderSource int //订单渠道 使用channel
	ReservationId int //预约记录id
	IcardPackageId int //身份卡卡包id
	TempId int //挂单id
}

//saas后台下单返回数据结构
type ReplySaasCreateProductOrder struct {
	OrderSn string //订单号
	Ctime int64 //下单时间
	TotalAmount float64 //支付金额
	PayTypes []PayType //saas的付款方式
}

//前端用户下单的入参数据
type ArgsUserCreateProductOrder struct {
	common.Utoken
	ShopId int //门店id
	ProductOrderData []ProductData //购买的产品
	OrderSource int //订单渠道 使用channel
	IcardPackageId int //身份卡卡包id
}

//前端用户下单返回的数据格式
type ReplyUserCreateProductOrder struct {
	OrderSn string //订单号
	TotalAmount float64 //支付金额
	PayTypes []PayType //前端的付款方式
}

//根据用户提货单号修改用户提货状态码 入参
type ArgsSetUserPickUpGoodsStatus struct {
	common.BsToken
	//PickUpGoodsCode string
	OrderId int
	ConsumeGood string
	Uid int
	PickUpGoodsStatus int
}

type ReplyBool struct {
	Status bool
}

type ProductOrderInterface interface {
	//SaaS创建单项目订单方法
	SaasCreateProductOrder(ctx context.Context, args *ArgsSaasCreateProductOrder, reply *ReplySaasCreateProductOrder ) error
	//用户创建订单
	UserCreateProductOrder(ctx context.Context, args *ArgsUserCreateProductOrder, reply *ReplyUserCreateProductOrder ) error
	// 根据用户商品提货单号修改用户商品提货状态
	SetUserPickUpGoodsStatus(ctx context.Context, args *ArgsSetUserPickUpGoodsStatus, reply *ReplyBool) error
}