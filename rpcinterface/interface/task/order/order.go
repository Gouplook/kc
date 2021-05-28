//订单相关任务入队列
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/5 14:31

package order

import "context"

type ArgsTimeOutOrderid struct {
	PayOrderId int //支付订单id
	ExpireTimeOut int //超时时间, 单位：秒， 如：30秒就传30000
}

type ArgsTimeOutTemp struct {
	TempId int //挂单id
	ExpireTimeOut int //超时时间, 单位：秒， 如：30秒就传30000
}

type Order interface {
	//将卡包主表的主键id加入交换机
	SetCardPackageRelationId(ctx context.Context, cardPackageRelationId *int, reply *bool ) error
	//超时支付订单id加入交换机
	TimeOutOrderid(ctx context.Context, args *ArgsTimeOutOrderid, reply *bool) error
	//挂单超时取消挂单
	TimeOutTemp(ctx context.Context, args *ArgsTimeOutTemp, reply *bool) error
}

