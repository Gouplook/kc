/******************************************
@Description:监管可视化-充值卡充值订单
@Time : 2020/11/30 13:56
@Author :lixiaojun

*******************************************/
package dataVisualization

import "context"

type ArgsCardPackageRcardRecharge struct {
	CardPackageCardSn string //充值卡充值编号@kc_card_package_rcard_recharge_log.recharge_sn
}

type CardPackageRcardRecharge interface {
	//记录充值订单数据到监管可视化表中
	SetCardPackageRcardRecharge(ctx context.Context,args *ArgsCardPackageRcardRecharge,reply *bool)error
}