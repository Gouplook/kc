package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//保单接口定义
// @author liyang<654516092@qq.com>
// @date  2020/7/28 9:21

//卡包保单入参
type ArgsCardpackagePolicy struct {
	common.Utoken     //用户信息
	RelationId    int //卡包关联ID
}

//卡包保单返回信息
type ReplyCardPackageList struct {
	Id               int             //保单ID
	PolicySn         string          //保单号
	Uid              int             //保单所属用户ID
	RelationId       int             //保单所属卡包关联ID
	InsuranceChannel int             //保单承保渠道
	ServicePeriod    int             //保险周期、单位：月
	Price            float64         //保额
	CalPrice         float64         //保费
	CalPriceCapital  string          //保费大写
	IsRenew          int             //是否为续保保单 0=否 1=是
	StartDate        string          //起保开始日期
	EndDate          string          //起保截止日期
	PolicyDetailUrl  string          //保单详情H5url
	Applicant        CommonApplicant //公共信息
}

//公共投保人与被保人信息
type CommonApplicant struct {
	RealName    string //被保人真实姓名
	RealNo      string //被保人身份证号
	Mobile      string //被保人手机号
	CompanyName string //投保企业名称
}

//保单详情入参
type ArgsSinglePolicy struct {
	common.Utoken     //用户信息
	Id            int //保单ID
}

//卡包详情返回信息
type ReplySinglePolicy struct {
	Id               int              //保单ID
	PolicySn         string           //保单号
	Uid              int              //保单所属用户ID
	RelationId       int              //保单所属卡包关联ID
	InsuranceChannel int              //保单承保渠道
	Price            float64          //保额
	ServicePeriod    int              //保险周期、单位：月
	CalPrice         float64          //保费
	CalPriceCapital  string           //保费大写
	IsRenew          int              //是否为续保保单 0=否 1=是
	StartDate        string           //起保开始日期
	EndDate          string           //起保截止日期
	Applicant        CommonApplicant  //公共信息
	Service          InsuranceService //服务信息
}

//服务信息
type InsuranceService struct {
	ServiceUrl     string //承保渠道网站地址
	ServicePhone   string //承保渠道电话
	ServiceAddress string //承保渠道地址
	Remark         string //平台描述
	ServiceExpress string //明示告知
}

// 获取保单任务信息/续保 入参数
type ArgsPolicyTask struct {
	TransNo string // 流水号
}

// 获取保单任务信息 返回信息
type ReplyPolicyTask struct {
	BusId            int     // 店铺/企业Id
	RiskBusId        int     `json:"risk_bus_id"`       //企业/商户ID
	Uid              int     `json:"uid"`               //用户ID
	RelationId       int     `json:"relation_id"`       //卡包关联ID
	CardPackageId    int     `json:"card_package_id"`   //卡包ID
	CardPackageSn    string  `json:"card_package_sn"`   //预付卡编号
	CardPackageType  int     `json:"card_package_type"` //卡包类型 0=套餐 1=综合卡 2=限时卡 3=限次卡 4=限时限次卡
	PayChannel       int     `json:"pay_channel"`       //支付渠道
	InsuranceChannel int     `json:"insurance_channel"` //承保渠道 1=长安保险 2=人保保险 3=安信保险
	Pid              int     `json:"pid"`               //省/直辖市ID
	Cid              int     `json:"cid"`               //城市ID
	Did              int     `json:"did"`               //区
	Tid              int     `json:"tid"`               //区下属镇/街道ID
	Price            float64 `json:"price"`             //保单保额
	IsRanew          int     `json:"is_ranew"`          //是否为续保保单 0=否 1=是
	CalPrice         float64 `json:"cal_price"`         //保单保费
	ServicePeriod    int     `json:"service_period"`    //保险周期，单位：月
	StartDate        string  `json:"start_date"`        //保单有效起始日期，格式如2020-07-13
	EndDate          string  `json:"end_date"`          //保单有效截止日期，格式如2020-07-13
	StartTime        int64   `json:"start_time"`        //保单有效起始时间戳
	EndTime          int64   `json:"end_time"`          //保单有效截止时间戳
	Ctime            int64   `json:"ctime"`             //创建时间
}

// 返回续保任务信息
type ReplyRenewPolicyTask struct {
	RelationId      int    `json:"relation_id"`       //卡包关联ID
	CardPackageId   int    `json:"card_package_id"`   //卡包ID
	CardPackageType int    `json:"card_package_type"` //卡包类型 0=套餐 1=综合卡 2=限时卡 3=限次卡 4=限时限次卡
	ServicePeriod   int    `json:"service_period"`    //保险周期，单位：月
	StartDate       string `json:"start_date"`        //承保保单有效起始日期，格式如2020-07-13
	EndDate         string `json:"end_date"`          //承保保单有效截止日期，格式如2020-07-13
	StartTime       int64  `json:"start_time"`        //承保保单有效起始时间戳
	EndTime         int64  `json:"end_time"`          //承保保单有效截止时间戳
	Ctime           int64  `json:"ctime"`             //创建时间
}

// 保单续保获取支付渠道 入参数
type ArgsPayChannel struct {
	CardPackageId int //卡包ID
}

// 续保保单支付渠道返回
type ReplyPayChannel struct {
	PayChannel int `json:"pay_channel"` //支付渠道
}

type Policy interface {
	//获取卡包保单信息
	GetPolicyByRelationId(ctx context.Context, args *ArgsCardpackagePolicy, reply *[]ReplyCardPackageList) error
	//获取卡包保单信息
	GetPolicyById(ctx context.Context, args *ArgsSinglePolicy, reply *ReplySinglePolicy) error
	// 获取保单任务信息
	GetPolicyTaskInfo(ctx context.Context, args *ArgsPolicyTask, reply *ReplyPolicyTask) error
	// 获取续保任务信息
	GetRenewPolicyTaskInfo(ctx context.Context, args *ArgsPolicyTask, reply *ReplyRenewPolicyTask) error
	// 获取保单续保获取支付渠道
	//GetRenewPayChannel(ctx context.Context, args *ArgsPayChannel, reply *ReplyPayChannel) error
}
