package insurance

import "git.900sui.cn/kc/rpcinterface/interface/common"

// @author liyang<654516092@qq.com>
// @date  2020/7/24 17:12


const (
	YEAR  = 12 //定义每年划分几个月份
	MRATE = 0.002 //月基准费率-仅限宁波人表
    AHEAD_RENEW_TIME = 172800  //续保的保单设置提前两天进行续保 86400x2
    RENEW_SEVICE_PERIOD = 3 //正常续保保险周期月份
    RENEW_NUM = 3  //正常续保次数
)

//保险保单起止日期、时间
type InsuranceTimeScope struct {
	StartDate string   //起保开始日期
	EndDate   string   //起保结束日期
	StartTime int64    //起保开始时间
	EndTime   int64    //起保结束时间
}

//保险出单扩展信息
type InsuranceExtend struct {
	ServicePeriod int //正常出单、预付卡总保险周期
	InsurancePeriod int //当前只能出单的预付卡周期
	RenewType int     //续保出单的续保类型 0=正常续保 1=分割续保
	StartDate string  //起保开始日期
	EndDate string    //起保结束日期
	StartTime int64   //起保开始时间戳
	EndTime int64     //起保结束时间戳
	RealName string   //被保人真实姓名
	CardNo  string    //被保人身份证号
	Mobile  string    //被保人手机号
	CompanyName string //投保企业/商户名称
	MerchantId  string //商户编号
}

//出单返回信息
type ReplyInsurance struct {
	TransNo string //流水号
	Msg string //信息
}

//用户保险信息入参
type ArgsInsuranceUser struct {
	common.Utoken //用户信息
}

//用户保险信息返回参数
type ReplyInsuranceUser struct {
	TotoalInsurancAmount float64
}









