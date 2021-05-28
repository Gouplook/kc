/******************************************
@Description:监管可视化-卡包消费记录
@Time : 2020/11/30 13:58
@Author :lixiaojun

*******************************************/
package dataVisualization

import "context"

type ArgsSetCardPackageConsumeLog struct {
	RelationLogId int //订单服务-卡项-消费记录索引id @kc_card_package_relation_log.id
}

type CardPackageConsumeLog interface {
	//记录卡包消费记录数据到监管可视化表中
	SetCardPackageConsumeLog(ctx context.Context,args *ArgsSetCardPackageConsumeLog,reply *bool)error
}
