package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//是否删除
	IS_DEL_YES = 1 //已删除
	IS_DEL_NO  = 0 //未删除
)

//后台 添加/修改 通知公告入参
type ArgsAddNoticeContext struct {
	common.Autoken // 后台管理员信息
	common.Paging
	Id            int    //自增ID
	NoticeTitle   string //通知公告标题
	NoticeContext string //通知公告内容
}

//删除通知公告
type ArgsDelNoticeContext struct {
	common.Autoken     // 后台管理员信息
	Id             int //自增ID
}

//查看商户Erp通知公告入参数
type ArgsNoticeContext struct {
	common.BsToken
	Id            int //通知公告编号
}

type NoticeContextInfo struct {
	Id            int    //自增ID
	NoticeTitle   string //通知公告标题
	Ctime         string //发布通知公告时间
	CtimeStr      string  //转化好的时间
	NoticeContext string //通知公告内容
}

//查看商户Erp通知公告返回参数
type ReplyNoticeContext struct {
	TotalNum int
	NoticeContextInfo
}

//后台 通知公告列表入参数
type ArgsNoticeList struct {
	common.Paging
}

type ReplyNoticeList struct {
	TotalNum int
	Lists    []NoticeContextInfo
}
type ArgsNoticeInfo struct {
	Id int
}

type ReplyNoticeInfo struct {
	Id            int    //自增ID
	NoticeTitle   string //通知公告标题
	Ctime         string //发布通知公告时间
	CtimeStr      string //转化好的时间
	NoticeContext string //通知公告内容
}

//返回 获取通知列表信息
type ReplyNoticeListInfo struct {
	TotalNum int
	Lists    []NoticeContextInfo
}
type ArgsNoticeListInfo struct {
}

//通知公告服务接口
type NoticeContext interface {
	//后台 添加通知公告
	AddNoticeContext(ctx context.Context, args *ArgsAddNoticeContext, reply *bool) error
	//后台 通知公告列表
	NoticeContextList(ctx context.Context, args *ArgsNoticeList, reply *ReplyNoticeList) error
	//后台 通知公告详情
	NoticeContextInfo(ctx context.Context, args *ArgsNoticeInfo, reply *ReplyNoticeInfo) error
	//后台 修改通知公告
	UpdateNoticeContext(ctx context.Context, args *ArgsAddNoticeContext, reply *bool) error
	//后台 删除通知公告
	DelNoticeContext(ctx context.Context, args *ArgsDelNoticeContext, reply *bool) error
	//查看商户Erp通知公告
	GetNoticeContext(ctx context.Context, args *ArgsNoticeContext, reply *ReplyNoticeContext) error
	//获取通知列表信息
	GetNoticeListInfo(ctx context.Context, args *ArgsNoticeListInfo, reply *ReplyNoticeListInfo) error
}
