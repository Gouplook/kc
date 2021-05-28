package order

//定义卡包相关常量
// @author liyang<654516092@qq.com>
// @date  2020/7/22 18:35

const (
	//#支付渠道
	//原生支付宝
	CARDPACKAGE_pay_channel_alipay = 1
	//原生微信
	CARDPACKAGE_pay_channel_wx = 2
	//上海杉德支付
	CARDPACKAGE_pay_channel_sandpay = 3
	//杭州建行直连支付
	CARDPACKAGE_pay_channel_ccbpay = 4
	//#资金存管方式
	//无资金管理方式
	CARDPACKAGE_fund_mode_none = 0
	//存管
	CARDPACKAGE_fund_mode_custod = 1
	//保险
	CARDPACKAGE_fund_mode_insurance = 2
	//#保险渠道
	//无保险
	CARDPACKAGE_insurance_channel_none = 0
	//长安保险
	CARDPACKAGE_insurance_channel_ca = 1
	//人保保险
	CARDPACKAGE_insurance_channel_picc = 2
	//安信保险
	CARDPACKAGE_insurance_channel_aicc = 3
	//#卡包状态
	//待消费
	CARDPACKAGE_status_pay = 1
	//消费中
	CARDPACKAGE_status_ing = 2
	//已完成
	CARDPACKAGE_status_com = 3
	//已关闭
	CARDPACKAGE_status_close = 4
	//已退款
	CARDPACKAGE_status_refund = 5

	//#卡包类型
	//单项目
	CARDPACKAGE_order_type_single = 1
	//套餐
	CARDPACKAGE_order_type_sm = 2
	//综合卡
	CARDPACKAGE_order_type_card = 3
	//限时卡
	CARDPACKAGE_order_type_hcard = 4
	//限次卡
	CARDPACKAGE_order_type_ncard = 5
	//限时限次卡
	CARDPACKAGE_order_type_hncard = 6
	//充值卡
	CARDPACKAGE_order_type_rcard = 7
	//身份卡
	CARDPACKAGE_order_type_icard = 8

	//#删除状态
	//正常
	CARDPACKAGE_deleted_no = 0
	//已删除
	CARDPACKAGE_deleted_yes = 1

	//#适用门店
	//适用全部
	CARDPACKAGE_cable_shop_all = 1
	//适用局部
	CARDPACKAGE_cable_shop_part = 2

	//#确认类型，针对限时卡和限时限次卡
	//线上确认
	CARDPACKAGE_consume_type_normal = 0
	//系统确认
	CARDPACKAGE_consume_type_system = 1

	//#消费服务or消费产品
	//消费服务
	CARDPACKAGE_consume_service = 0
	//消费产品
	CARDPACKAGE_consume_product = 1

	//#消费记录中记录卡包消费状态
	//未消费完
	CARDPACKAGE_consume_uncomp = 0
	//已经消费完
	CARDPACKAGE_consume_comp = 1

	//#充值卡充值单类型
	//面值折扣
	CARDPACKAGE_DISACCOUNT_face = 1
	//包含项目折扣
	CARDPACKAGE_DISACCOUNT_single = 2
	//#充值卡充值单枚举
	//可消费
	CARDPACKAGE_RECHARGE_canuse = 0
	//已消费完毕
	CARDPACKAGE_RECHARGE_comp = 1

	//#限时卡和限时限次卡 系统自动划转状态
	//未处理
	CARDPACKAGE_CONFIRM_STATUS_untreated = 1
	//处理成功
	CARDPACKAGE_CONFIRM_STATUS_process_success = 2
	//处理失败
	CARDPACKAGE_CONFIRM_STATUS_process_fail = 3
	//无需处理
	CARDPACKAGE_CONFIRM_STATUS_no_process = 4
)
