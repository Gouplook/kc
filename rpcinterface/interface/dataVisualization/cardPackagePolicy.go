/******************************************
@Description:
@Time : 2020/11/30 13:56
@Author :lixiaojun

*******************************************/
package dataVisualization

import "context"

// 入参数
type ArgsCardPolicy struct {
	TransNo  string // 流水号
	IsRenew int    // 出单类型 0=正常出单 1= 续保出单
}

// 返回添加信息
type ReplyCardPolicy struct {
	PolicySn         string  `json:"policy_sn"`         //保单单号
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

type Policy interface {
	// 添加监管可视化 保险保单信息
	AddVisualizationPolicyRpc(ctx context.Context, args *ArgsCardPolicy, reply *bool) error
}
