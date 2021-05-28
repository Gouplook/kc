/******************************************
@Description:
@Time : 2021/1/7 16:58
@Author :lixiaojun

*******************************************/
package reservation

const (
	//类型
	CallTypeReservation = 0 // 1-预约
	CallTypeOrder       = 1 // 开单

	//预约状态
	ReservationStatusTimeOut     = 0 //超时
	ReservationStatusWaitConfirm = 1 //待确认
	ReservationStatusCanceled    = 2 //取消（用户/商家取消）
	ReservationStatusConfirmed   = 3 //待服务
	ReservationStatusOrdered     = 4 //待结算
	ReservationStatusCompleted   = 5 //已完成

	//预约类型
	ReservationTypeProduct = 0 //商品
	ReservationTypeSingle  = 1 //单项目
	ReservationTypeSm      = 2 //套餐
	ReservationTypeCard    = 3 //综合卡
	ReservationTypeHcard   = 4 //限时卡
	ReservationTypeNcard   = 5 //限次卡
	ReservationTypeHncard  = 6 //限时限次卡
	ReservationTypeRcard   = 7 //充值卡

	//预约人类型
	UserTypeNormal = 1 // 散客
	UserTypeMember = 2 // 会员

	//预约条目的有效性
	ReservationItemsVaild   = 0 // 有效
	ReservationItemsInvalid = 1 // 无效

	//取消类型
	CancelTypeUser = 1 //用户取消
	CancelTypeShop = 2 //商家取消

	//默认预约人数
	DefaultReservationPeopleNum = 1

	//结算类型
	SettleTypeZero           = 0
	SettleTypeWaitConsumCard = 1 //待耗卡
	SettleTypeWaitPay        = 2 //待支付
	SettleTypeFinish         = 3 //结算完成
	//订单类型
	OrderTypeProduct = 1 //商品
	OrderTypeService = 2 //服务
	//预约功能状态开关
	ReservSwitchDisableStatus = 0 //关闭
	ReservSwitchEnableStatus  = 1 //开启

	//#周一到周日
	Week_mon  = 1
	Week_tues = 2
	Week_wed  = 3
	Week_thur = 4
	Week_fri  = 5
	Week_sat  = 6
	Week_sun  = 7
)

func GetWeekTypeLists() []int {
	return []int{Week_mon, Week_tues, Week_wed, Week_thur, Week_fri, Week_sat, Week_sun}
}
