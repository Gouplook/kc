/**
 * @Author: YangYun
 * @Date: 2020/8/6 9:18
 */
package bus

import "context"

const (
	TYPE_charge_in  = 1 // 入账
	TYPE_charge_out = 2 // 出账

	ORDER_TYPE_buy      = 1 // 购买
	ORDER_TYPE_agent    = 2 // 资金清算
	ORDER_TYPE_withdraw = 3 // 商户提现

	FUND_TYPE_pay_in          = 1 // 用户支付
	FUND_TYPE_bankfee_out     = 2 // 银行手续费
	FUND_TYPE_insurance_out   = 3 // 保险费用
	FUND_TYPE_reinsurance_out = 4 // 续保费用
	FUND_TYPE_plat_out        = 5 // 平台手续费
	FUND_TYPE_agent_fee_out   = 6 // 清算手续费
	FUND_TYPE_withdraw_out    = 7 // 商户提现

	FUND_TYPE_depository = 8 // 清分到存管
	FUND_TYPE_usable     = 9 // 清分到可提现

	FUND_TYPE_consume_out = 10 // 确认消费

	// 存管订单类型
	DEPOSITORY_ORDER_TYPE_agent   = 1 // 清算入账
	DEPOSITORY_ORDER_TYPE_consume = 2 // 耗卡出账
	DEPOSITORY_ORDER_TYPE_refund  = 3 //退款出账

	// 可提现订单类型
	USABLE_ORDER_TYPE_agent    = 1 // 清算入账
	USABLE_ORDER_TYPE_consume  = 2 // 耗卡入账
	USABLE_ORDER_TYPE_withdraw = 3 // 提现出账
	USABLE_ORDER_TYPE_refund   = 4 // 退款入账
)

type FundRecord struct {
	BusId          int
	FundRecordList []FundRecordItem
}
type FundRecordItem struct {
	OrderSn  string // 订单编号
	Amount   string // 变动金额
	FundType int    // 费用类型
}

//充值卡的消费信息结构体
type RcardRechargeLog struct {
	RechargeLogId   int    //充值记录id
	RealPrice       string //真实金额
	SubOrderId      int    //订单ID@kc_card_order_card表中的id字段
	PayOrderSn      string //支付订单号
	PayOrderId      int    //支付订单id
	PayChannel      int    //支付渠道
	DepositRatio    string //存管留存比例
	DepositAmount   string //总留存金额
	HasRelaseAmount string //已经释放的留存资金
	Amount          string //确认消费使用的实际金额
	LastConsumeFlag int    //是否最后一次消费 1=是 0=否
}

//确认消费商家资金变动信息维护入参
type ArgsConsumeCharge struct {
	BusId                 int
	OrderSn               string             //订单编号
	PayChannel            int                //支付渠道
	PayOrderId            int                //支付订单的id
	Amount                string             //确认消费的实际金额
	CardPackageRelationid int                //卡包关系id
	RealPrice             string             //卡包的真实金额
	DepositRatio          string             //存管留存比例
	DepositAmount         string             //卡包总留存金额
	HasRelaseAmount       string             //已经释放的留存资金
	LastConsumeFlag       int                //是否为最后一次消费 1=是 0=否
	DepositoryOrderType   int                //存管订单类型:1-清算入账;2-耗卡出账;3-退款出账
	RcardRechargeLogInfo  []RcardRechargeLog //充值卡消费信息

}

type Fund interface {
	// 商户资金变动
	PayCharge(ctx context.Context, fundRecord *FundRecord, reply *bool) error
	//确认消费，商家存管资金变动
	ConsumeCharge(ctx context.Context, args *ArgsConsumeCharge, reply *bool) error
	//处理出账中的存管消费金额
	DeposNeedOut(ctx context.Context, clearId *int, reply *bool) error
}
