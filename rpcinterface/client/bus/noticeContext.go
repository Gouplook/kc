package bus

//定义rpc调用
// @author yinjinlin<yinjinlin_uplook@163.com>
// @date  2020/10/9/ 10:30

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type NoticeContext struct {
	client.Baseclient
}

//初始化
func (n *NoticeContext) Init() *NoticeContext {
	n.ServiceName = "rpc_bus"
	n.ServicePath = "NoticeContext"
	return n
}

//添加商户Erp通知公告
func (n *NoticeContext)AddNoticeContext(ctx context.Context, args *bus.ArgsAddNoticeContext, reply *bool) error{
	return n.Call(ctx, "AddNoticeContext", args, reply)
}
//修改通知公告
func (n *NoticeContext)UpdateNoticeContext(ctx context.Context, args *bus.ArgsAddNoticeContext,reply *bool) error{
	return n.Call(ctx, "UpdateNoticeContext", args, reply)
}
//删除通知公告
func (n *NoticeContext)DelNoticeContext(ctx context.Context, args *bus.ArgsDelNoticeContext,reply *bool) error{
	return n.Call(ctx, "DelNoticeContext", args, reply)
}
//查看商户Erp通知公告
func (n *NoticeContext) GetNoticeContext(ctx context.Context, args *bus.ArgsNoticeContext, reply *bus.ReplyNoticeContext) error {
	return n.Call(ctx, "GetNoticeContext", args, reply)
}

//后台 通知公告列表
func (n *NoticeContext)NoticeContextList(ctx context.Context, args *bus.ArgsNoticeList, reply *bus.ReplyNoticeList) error{
	return n.Call(ctx, "NoticeContextList", args, reply)
}
//后台 通知公告详情
func(n *NoticeContext)NoticeContextInfo(ctx context.Context, args *bus.ArgsNoticeInfo, reply *bus.ReplyNoticeInfo) error{
	return n.Call(ctx, "NoticeContextInfo", args, reply)
}
//获取通知列表信息
func(n *NoticeContext)GetNoticeListInfo(ctx context.Context,args *bus.ArgsNoticeListInfo,reply *bus.ReplyNoticeListInfo) (err error){
	return n.Call(ctx, "GetNoticeListInfo",args,reply)
}
