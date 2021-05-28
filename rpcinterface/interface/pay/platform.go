/**
 * @Author: YangYun
 * @Date: 2020/8/25 16:56
 */
package pay

const (
	TYPE_charge_in  = 1 // 入账
	TYPE_charge_out = 2 // 出账

	ORDER_TYPE_pay                   = 1 // 收取支付手续费
	ORDER_TYPE_insure_fund_fee       = 2 // 保险清分手续费
	ORDER_TYPE_renew_insure_fund_fee = 3 // 续保清分手续费

	FUND_SOURCE_TYPE_bus          = 1 // 商家来源
	FUND_SOURCE_TYPE_insure       = 2 // 保险公司来源
	FUND_SOURCE_TYPE_renew_insure = 3 // 续保公司来源

	AGENT_TYPE_sand = 2 //上海杉德清算
	AGENT_TYPE_ccb  = 3 //杭州建行清算
)
