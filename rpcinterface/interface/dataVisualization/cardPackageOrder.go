/******************************************
@Description:
@Time : 2020/11/30 13:57
@Author :lixiaojun

*******************************************/
package dataVisualization
import (
	"context"
)
type CardPackageOrder interface {
	//根据relationId创建监管可视化-卡包订单信息
	CreateOrderInfoByRelationId(ctx context.Context, relationId *int, reply *bool) error
}
