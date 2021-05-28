package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//sass创建充值订单入参
type ArgsSaasCreateRechargeOrder struct {
	common.BsToken
	common.Utoken     //当前操作人员登录信息
	Uid           int //购买人uid
	RechargeOrderData
}
type RechargeOrderData struct {
	CardPackageId int    //充值卡包id
	RcardRuleId   int    //充值规则id
	RechargeAmount float64//充值金额（商家自定义充值金额）
	GiveAmount float64 //赠送的金额
	StaffIds      string //销售人员
	OrderSource   int    //订单渠道 使用channel
	Gives []GiveSingleBase //赠送的服务
}

type GiveSingleBase struct {
	SingleId int
	Num int//项目次数
	ExpireDateStr string //过期时间,权益过期时间
}

//用户创建充值订单入参
type ArgsUserCreateRechargeOrder struct {
	common.Utoken
	ShopId int //门店id
	RechargeOrderData
}

//创建充值订单出参
type ReplyCreateRechargeOrder struct {
	OrderSn     string    //订单号
	Ctime       int64     //下单时间
	CtimeStr string
	TotalAmount float64   //支付金额
	PayTypes    []PayType //saas的付款方式
}

type RcardOrder interface {
	//sass创建充值订单
	SaasCreateRechargeOrder(ctx context.Context, args *ArgsSaasCreateRechargeOrder, reply *ReplyCreateRechargeOrder) error
	//user创建充值订单
	UserCreateRechargeOrder(ctx context.Context, args *ArgsUserCreateRechargeOrder, reply *ReplyCreateRechargeOrder) error
}
