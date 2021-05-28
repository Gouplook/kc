package order

/**
 * @className PinganThirdPayOrder
 * 平安第三方在途充值
 * @author liyang<654516092@qq.com>
 * @date 2021/4/26 17:15
 */

//子订单数组信息
type PinganThirdSubOrders struct {
	SubOrders []PinganThirdSubOrder
}
//子订单信息
type PinganThirdSubOrder struct {
	RechargeSubAcctNo    string `mapstructure:"sub_acct_no"` //充值子账号
	SubOrderFillMemberCd string `mapstructure:"merchant_id"` //充值会员代码
	SubOrderTranAmt      string `mapstructure:"amount"`  //交易金额
	SubOrderTranFee      string `mapstructure:"tran_fee"` //交易手续费-平安
	SubOrderNo           string `mapstructure:"order_sn"` //子订单号
	SubOrderPayMode      string `mapstructure:"pay_mode"` //交易模式  0-冻结支付 1-普通支付
	SubOrderContent      string `mapstructure:"order_desc"` //子订单描述
}
