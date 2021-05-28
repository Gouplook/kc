package order

import "git.900sui.cn/kc/rpcinterface/interface/task/cards"

//定义常量
const (
	//一次购买的最大数量
	BUY_MAX_NUM = 10

	//总订单类型 1=单项目订单 2=卡项订单  3=商品订单
	PAY_ORDER_TYPE_SINGLE = 1
	PAY_ORDER_TYPE_CARD   = 2
	PAY_ORDER_TYPE_GOODS  = 3

	//卡项订单类型 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡 7=充值卡 8=身份卡
	ORDER_TYPE_SM     = 2
	ORDER_TYPE_CARD   = 3
	ORDER_TYPE_HCARD  = 4
	ORDER_TYPE_NCARD  = 5
	ORDER_TYPE_HNCARD = 6
	ORDER_TYPE_RCARD  = 7
	ORDER_TYPE_ICARD  = 8
	ORDER_TYPE_EQUITY = 9

	//单项目订单类型 1=会员订单 2=游客订单
	SINGLE_ORDER_TYPE_member = 1
	SINGLE_ORDER_TYPE_GUEST  = 2

	//支付渠道 1=原生支付（支付宝，微信） 2=杉德支付 3=建行直连 4=平安银行 5=工商银行 6=宁波银行
	PAY_CHANNEL_ALIWX  = 1
	PAY_CHANNEL_sand   = 2
	PAY_CHANNEL_ccb    = 3
	PAY_CHANNEL_pingan = 4
	PAY_CHANNEL_icbc   = 5
	PAY_CHANNEL_nbcb   = 6

	//付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
	PAY_TYPE_ALI  = 1
	PAY_TYPE_WX   = 2
	PAY_TYPE_CASH = 3
	PAY_TYPE_YUAN = 4

	//项目订单来源 0=普通订单 1=一元体验订单 2=拼团订单
	SINGLE_SOURCE_NAM   = 0
	SINGLE_SOURCE_TIYAN = 1
	SINGLE_SOURCE_TUAN  = 2

	//支付状态 0=待支付 1=已支付 2=支付失败 3=订单关闭
	PAY_STATUS_NO     = 0
	PAY_STATUS_SUC    = 1
	PAY_STATUS_FAIL   = 2
	PAY_STATUS_CLOSE  = 3
	PAY_STATUS_REFUND = 4 //已退款

	//订单来源 0=普通订单 1=一元体验订单 2=拼团订单 3=充值卡复充
	SOURCE_NOMAL = 0
	SOURCE_ONE   = 1
	SOURCE_TUAN  = 2
	SOURCE_RCARD = 3

	//saas/康享保的支付方式 1=支付宝扫描  2=微信扫描 3=现金 4=杉德支付宝 5=杉德微信 6=建行直连 7=杉德直连 8=工行支付宝App支付 9=工行微信App支付 10=工行二维码支付 11=工行H5支付
	//12=平安银行渠道-微信app支付直连 13=平安银行渠道-支付宝app支付直连
	CHOSE_PAY_TYPE_ALI         = 1
	CHOSE_PAY_TYPE_WX          = 2
	CHOSE_PAY_TYPE_CASH        = 3
	CHOSE_PAY_TYPE_SAND_ALI    = 4
	CHOSE_PAY_TYPE_SAND_WX     = 5
	CHOSE_PAY_TYPE_SAND_CCB    = 6
	CHOSE_PAY_TYPE_SAND        = 7
	CHOSE_PAY_TYPE_ICBC_ALI    = 8
	CHOSE_PAY_TYPE_ICBC_WX     = 9
	CHOSE_PAY_TYPE_ICBC_QRCODE = 10
	CHOSE_PAY_TYPE_ICBC_H5     = 11
	CHOSE_PAY_TYPE_PINGAN_WX   = 12
	CHOSE_PAY_TYPE_PINGAN_ALI  = 13
	CHOSE_PAY_TYPE_PINGAN      = 14

	//订单前缀
	PREFIX_SINGLE = "JS"
	PREFIX_SM     = "JM"
	PREFIX_CARD   = "JC"
	PREFIX_HCARD  = "JH"
	PREFIX_NCARD  = "JN"
	PREFIX_HNCARD = "JP"
	PREFIX_GOODS  = "JG"
	PREFIX_RCARD  = "JR"
	PREFIX_ICARD  = "JI"

	//平安解冻任务处理状态  0=待处理  1=处理中 2=成功 3=失败
	PINGAN_STATUS_wait = 0
	PINGAN_STATUS_ing  = 1
	PINGAN_STATUS_suc  = 2
	PINGAN_STATUS_fail = 3
	//平安解冻任务的来源类型 1=消费 2=退款
	PINGAN_THAW_TYPE_consume = 1
	PINGAN_THAW_TYPE_refund  = 2
	//平安支付子订单的支付模式 0-冻结支付 1-普通支付
	PINGAN_PAY_MODE_dj    = 0
	PINGAN_PAY_MODE_nomal = 1

	//平安第三方在途充值状态 0=待处理  1=处理中 2=成功 3=失败
	PINGAN_THIRD_wait = 0
	PINGAN_THIRD_ing  = 1
	PINGAN_THIRD_suc  = 2
	PINGAN_THIRD_fail = 3
)

var PrefixArr = map[int]string{
	cards.ITEMTYPE_single: PREFIX_SINGLE,
	cards.ITEMTYPE_sm:     PREFIX_SM,
	cards.ITEMTYPE_card:   PREFIX_CARD,
	cards.ITEMTYPE_hcard:  PREFIX_HCARD,
	cards.ITEMTYPE_ncard:  PREFIX_NCARD,
	cards.ITEMTYPE_hncard: PREFIX_HNCARD,
	cards.ITEMTYPE_rcard:  PREFIX_RCARD,
	cards.ITEMTYPE_icard:  PREFIX_ICARD,
}

//不同支付渠道的支付手续费
var PayChannelFree = map[int]string{
	PAY_CHANNEL_ALIWX: "0.006",
	PAY_CHANNEL_sand:  "0.005",
	PAY_CHANNEL_ccb:   "0.006",
	PAY_CHANNEL_icbc:  "0.006",
}

type PayOrder struct {
	OrderId           int     //订单id
	OrderSn           string  //订单号
	BusId             int     //企业id
	ShopId            int     //门店id
	OrderType         int     //订单类型
	OrderSource       int     //订单来源
	Uid               int     //下单人
	PaySn             string  //支付流水号
	PayTime           int     //支付时间
	PayStatus         int     //支付状态
	PayChannel        int     //支付渠道
	PayFee            float64 //支付手续费
	InsuranceChannel  int     //保险渠道
	TotalAmount       float64 //应付款金额
	RealPrice         float64 //实际付款金额
	FundMode          int     //资金管理
	DeposBankChannel  int     //资金存管银行渠道 1=上海银行 2=交通银行
	DepositRatio      float64 //留存比例
	DepositAmount     float64 //留存金额
	InsureAmount      float64 //保险费用
	RenewInsureAmount float64 //续保费用
	PlatformAmount    float64 //平台手续费
	BusAmount         float64 //商户应收起金额
	CreateTime        int64   //创建时间
}

//支付方式结构体数据
type PayType struct {
	Type int    //支付方式
	Name string //支付方式名称
	Logo string //支付方式的图片
}

//充值卡充值记录的支付订单信息
type RechargeLogPayOrderInfo struct {
	OrderId    int    //订单id
	OrderSn    string //订单号
	PayChannel int    //支付渠道
}
