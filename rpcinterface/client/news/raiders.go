package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Raiders struct {
	client.Baseclient
}

func (a *Raiders) Init() *Raiders {
	a.ServicePath = "Raiders"
	a.ServiceName = "rpc_news"
	return a
}

// AdminAddRaiders 后台添加攻略
func (a *Raiders) AdminAddRaiders(ctx context.Context, args *news.ArgsAdminAddRaiders, reply *news.ReplyAdminAddRaiders) error {
	return a.Call(ctx, "AdminAddRaiders", args, reply)
}

// AdminUpdateRaiders 后台更新攻略
func (a *Raiders) AdminUpdateRaiders(ctx context.Context, args *news.ArgsAdminUpdateRaiders, reply *bool) error {
	return a.Call(ctx, "AdminUpdateRaiders", args, reply)
}

// AdminRaidersList 后台攻略列表
func (a *Raiders) AdminRaidersList(ctx context.Context, args *news.ArgsAdminRaidersList, reply *news.ReplyAdminRaidersList) error {
	return a.Call(ctx, "AdminRaidersList", args, reply)
}

// AdminRaiderInfo 后台攻略详情
func (a *Raiders) AdminRaiderInfo(ctx context.Context, args *news.ArgsAdminRaiderInfo, reply *news.ReplyAdminRaiderInfo) error {
	return a.Call(ctx, "AdminRaiderInfo", args, reply)
}

// AuditRaider 审核攻略
func (a *Raiders) AuditRaider(ctx context.Context, args *news.ArgsAuditRaider, reply *bool) error {
	return a.Call(ctx, "AuditRaider", args, reply)
}

// UserAddRaider 前台用户添加攻略
func (a *Raiders) UserAddRaider(ctx context.Context, args *news.ArgsUserAddRaider, reply *news.ReplyUserAddRaider) error {
	return a.Call(ctx, "UserAddRaider", args, reply)
}

// UserRaidersList 前台台攻略列表
func (a *Raiders) UserRaidersList(ctx context.Context, args *news.ArgsUserRaiderList, reply *news.ReplyUserRaiderList) error {
	return a.Call(ctx, "UserRaidersList", args, reply)
}

//GetUserPublishRaiderList 用户发布的攻略
func (a *Raiders)GetUserPublishRaiderList(ctx context.Context,args *news.ArgsUserRaiderList,reply *news.ReplyUserRaiderList)error{
	return a.Call(ctx, "GetUserPublishRaiderList", args, reply)
}

// UserRaiderInfo 前台攻略详情
func (a *Raiders) UserRaiderInfo(ctx context.Context, args *news.ArgsUserRaiderInfo, reply *news.ReplyUserRaiderInfo) error {
	return a.Call(ctx, "UserRaiderInfo", args, reply)
}

// RecommendList 推荐阅读
func (a *Raiders) RecommendList(ctx context.Context, args *news.ArgsRaiderRecommend, reply *news.ReplyRaiderRecommend) error {
	return a.Call(ctx, "RecommendList", args, reply)
}

// GetShop 获取攻略关联的店铺
func (a *Raiders) GetShop(ctx context.Context, args *news.ArgsGetShop, reply *news.ReplyGetShop) error {
	return a.Call(ctx, "GetShop", args, reply)
}
