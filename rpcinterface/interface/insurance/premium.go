package insurance

import "context"

// 保费试算定义
// @author liyang<654516092@qq.com>
// @date  2020/7/27 16:23

type ArgsCalcuPremium struct {
	BusId int     //企业/商户ID
	Price float64 //保额
	InsuranceChannel int //保险渠道 1=长安 2=人保 3=安信
	ServicePeriod  int   //保险周期,单位：月
	CutServicePeriod int //续保每张保单周期,不计算续保保费，忽略
}

type ReplyCalcuPremium struct {
	Permium float64
}

type Premium interface {
	//保费试算
    GetCalcuPremium(ctx context.Context,args *ArgsCalcuPremium,reply *ReplyCalcuPremium) error
}