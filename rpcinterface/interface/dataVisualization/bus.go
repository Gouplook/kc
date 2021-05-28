/******************************************
@Description:
@Time : 2020/11/30 13:57
@Author :lixiaojun

*******************************************/
package dataVisualization

import (
	"context"
)

type Bus interface {
	//根据busId创建监管可视化-企业/商户基础信息表
	CreateBusInfoByBusId(ctx context.Context, busId *int, reply *bool) error
	//创建/更新 企业/商户分店基础信息
	CreateOrUpdateBusShopInfo(ctx context.Context, shopId *int, reply *bool) error
}
