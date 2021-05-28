package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//退款状态
	RefundStatusPending = 1 //退款中
	RefundStatusSuc     = 2 //退款成功
	RefundStatusFail    = 3 //退款失败
)

//根据支付订单id和订单类型 计算当前订单可退款总金额入参
type ArgsCalculateSingleOrCardRefundAmount struct {
	common.BsToken
	PayOrderId   int //支付订单id
	PayOrderType int //支付订单类型
	SubOrderSubIds []int//子订单id
	ReleationIds []int
}

//根据支付订单id和订单类型 计算当前订单可退款总金额出参
type ReplyCalculateSingleOrCardRefundAmount struct {
	RefundTotalAmount float64
}

//退款申请
type ArgsRefundApply struct {
	common.BsToken
	PayOrderId         int     //支付订单id
	SubOrderSubIds []int//子订单id
	ReleationIds []int
	PayOrderType       int     //支付订单类型（1=单项目订单 2=卡项订单  3=商品订单）
	RefundWay          int     //退款方式：1=支付宝 2=微信 3=现金
	ActualRefundAmount float64 //实际退款金额
	Explain            string  //退款说明
}

//退款订单列表入参
type ArgsGetRefundOrderList struct {
	common.BsToken
	common.Paging
	Status          int
	ShopId          int
	RefundStartTime string //退款开始时间，2006-01-02
	RefundEndTime   string //退款结束时间
	DateType        int    //时间类型 1:今天  2:近3天  3:近7天
}
type GetRefundOrderListBase struct {
	Id                 int
	PayOrderId         int //支付订单id
	Name               []string
	RefundSn           string //退款订单编号
	BusId              int
	ShopId             int
	ShopName           string  //分店门店名称
	BranchName         string  //分店名称
	RefundWay          int     //退款方式：1=支付宝 2=微信 3=现金
	OrderType          int     //支付订单类型（1=单项目订单 2=卡项订单  3=商品订单）
	OrderAmount        float64 //订单金额
	ActualRefundAmount float64 //实际退款金额
	Status             int     //状态：1-退款中 2-退款成功 3-退款失败
	Explain            string  //退款说明
	CtimeStr           string  //退款时间
}

//退款订单列表出参
type ReplyGetRefundOrderList struct {
	TotalNum int
	Lists    []GetRefundOrderListBase
}

type ReplyGetRefundOrderInfoById struct {
	Id                 int
	PayOrderId         int    //支付订单id
	RefundSn           string //退款订单编号
	BusId              int
	ShopId             int
	RefundWay          int     //退款方式：1=支付宝 2=微信 3=现金
	OrderType          int     //支付订单类型（1=单项目订单 2=卡项订单  3=商品订单）
	OrderAmount        float64 //订单金额
	ActualRefundAmount float64 //实际退款金额
	Status             int     //状态：1-退款中 2-退款成功 3-退款失败
	Explain            string  //退款说明
	Ctime int64
	CtimeStr           string  //退款时间
}

type RefundOrder interface {
	//根据支付订单id和订单类型 计算当前订单可退款总金额
	CalculateSingleOrCardRefundAmount(ctx context.Context, args *ArgsCalculateSingleOrCardRefundAmount, reply *ReplyCalculateSingleOrCardRefundAmount) error
	//退款申请
	RefundApply(ctx context.Context, args *ArgsRefundApply, reply *bool) error
	//退款订单列表
	GetRefundOrderList(ctx context.Context, args *ArgsGetRefundOrderList, reply *ReplyGetRefundOrderList) error
	//退款详情
	GetRefundOrderInfoById(ctx context.Context, refundOrderId *int, reply *ReplyGetRefundOrderInfoById) error
}
