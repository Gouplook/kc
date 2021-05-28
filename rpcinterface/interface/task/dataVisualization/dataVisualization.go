/******************************************
@Description:监管可视化-充值卡充值订单表
@Time : 2020/11/30 17:01
@Author :lixiaojun

*******************************************/
package dataVisualization

import "context"

type CardPackageRcardRecharge struct {
	CardPackageCardSn string //充值卡充值编号@kc_card_package_rcard_recharge_log.recharge_sn
}

type DataVisualization interface {
	//设置充值卡充值订单到交换机
	SetCardPackageRcardRecharge(ctx context.Context,args *CardPackageRcardRecharge,reply *bool)error
}
