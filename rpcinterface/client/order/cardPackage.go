package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

// @author liyang<654516092@qq.com>
// @date  2020/7/23 14:03
type CardPackage struct {
	client.Baseclient
}

//初始化
func (c *CardPackage) Init() *CardPackage {
	c.ServiceName = "rpc_order"
	c.ServicePath = "CardPackage"
	return c
}

//单条消费记录-RPC
func (c *CardPackage) GetSingleConsumeLog(ctx context.Context, logId *int, reply *order.ReplyConsumeDataList) (err error) {
	return c.Call(ctx, "GetSingleConsumeLog", logId, reply)
}

//卡包消费记录
func (c *CardPackage) GetCardPackageConsumeLog(ctx context.Context, args *order.ArgsConsumeDataLog, reply *order.ReplyConsumeDataLog) error {
	return c.Call(ctx, "GetCardPackageConsumeLog", args, reply)
}

//卡包适用分店
func (c *CardPackage) GetCardPackageCableShop(ctx context.Context, args *order.ArgsCableShopInfo, reply *[]order.ReplyCableShopInfo) error {
	return c.Call(ctx, "GetCardPackageCableShop", args, reply)
}

//获取用户卡包详情-单项目
func (c *CardPackage) GetCardPackageInfoByUserSingle(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplySingleCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserSingle", args, reply)
}

//获取用户卡包详情-套餐
func (c *CardPackage) GetCardPackageInfoByUserSm(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplySmCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserSm", args, reply)
}

//获取用户卡包详情-综合卡
func (c *CardPackage) GetCardPackageInfoByUserCard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyCardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserCard", args, reply)
}

//获取用户卡包详情-限时卡
func (c *CardPackage) GetCardPackageInfoByUserHcard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyHcardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserHcard", args, reply)
}

//获取用户卡包详情-限次卡
func (c *CardPackage) GetCardPackageInfoByUserNcard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyNcardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserNcard", args, reply)
}

//获取用户卡包详情-限时限次卡
func (c *CardPackage) GetCardPackageInfoByUserHncard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyHncardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserHncard", args, reply)
}

//获取用户卡包详情-充值卡
func (c *CardPackage) GetCardPackageInfoByUserRcard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyRcardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserRcard", args, reply)
}

//获取用户卡包详情-身份卡
func (c *CardPackage) GetCardPackageInfoByUserIcard(ctx context.Context, args *order.ArgsUserCardPackageInfo, reply *order.ReplyIcardCardPackageInfo) error {
	return c.Call(ctx, "GetCardPackageInfoByUserIcard", args, reply)
}

//获取用户卡包二维码信息
func (c *CardPackage) GetCardPackageQrcode(ctx context.Context, args *order.ArgsUserCardPageQrcode, reply *string) error {
	return c.Call(ctx, "GetCardPackageQrcode", args, reply)
}

//卡包二维码扫一扫监测
func (c *CardPackage) ScanQrcodeCheck(ctx context.Context, args *order.ArgsCardPackageQrcodeCheck, reply *order.ReplyCardPackageQrcodeCheck) error {
	return c.Call(ctx, "ScanQrcodeCheck", args, reply)
}

//根据消费码获取预付卡信息
func (c *CardPackage) GetQrcodeByConsumeCode(ctx context.Context, args *order.ArgsCardPackageQrcode, reply *order.ReplyCardPackageQrcode) error {
	return c.Call(ctx, "GetQrcodeByConsumeCode", args, reply)
}

//获取用户卡包列表
func (c *CardPackage) GetCardPackageByUser(ctx context.Context, args *order.ArgsUserList, reply *order.ReplyUserList) error {
	return c.Call(ctx, "GetCardPackageByUser", args, reply)
}

//获取用户所在连锁店卡包列表
func (c *CardPackage) GetCardPackageByUserBus(ctx context.Context, args *order.ArgsUserBusList, reply *order.ReplyUserBusList) error {
	return c.Call(ctx, "GetCardPackageByUserBus", args, reply)
}

//获取用户所在连锁店卡包列表-开放平台
func (c *CardPackage) OpenPlatFormV1GetCardPackageByUserBus(ctx context.Context, args *order.ArgsUserBusList, reply *order.ReplyUserBusList) error {
	return c.Call(ctx, "OpenPlatFormV1GetCardPackageByUserBus", args, reply)
}

//获取卡包关联(索引)详情
func (c *CardPackage) GetCardPackageRelationInfo(ctx context.Context, args *order.ArgsGetSimpleRelation, reply *order.CardRelationInfo) error {
	return c.Call(ctx, "GetCardPackageRelationInfo", args, reply)
}

//获取卡包-单项目详情【RPC】
func (c *CardPackage) GetSingleInfo(ctx context.Context, args *order.ArgsGetRpcSingle, reply *order.CardSingleInfo) error {
	return c.Call(ctx, "GetSingleInfo", args, reply)
}

//获取卡包-套餐详情【RPC】
func (c *CardPackage) GetSmInfo(ctx context.Context, args *order.ArgsGetRpcSm, reply *order.CardSmInfo) error {
	return c.Call(ctx, "GetSmInfo", args, reply)
}

//获取卡包-综合卡详情【RPC】
func (c *CardPackage) GetCardInfo(ctx context.Context, args *order.ArgsGetRpcCard, reply *order.CardCardInfo) error {
	return c.Call(ctx, "GetCardInfo", args, reply)
}

//获取卡包-限时卡详情【RPC】
func (c *CardPackage) GetHcardInfo(ctx context.Context, args *order.ArgsGetRpcHcard, reply *order.CardHcardInfo) error {
	return c.Call(ctx, "GetHcardInfo", args, reply)
}

//获取卡包-限次卡详情【RPC】
func (c *CardPackage) GetNcardInfo(ctx context.Context, args *order.ArgsGetRpcNcard, reply *order.CardNcardInfo) error {
	return c.Call(ctx, "GetNcardInfo", args, reply)
}

//获取卡包-限时限次卡详情【RPC】
func (c *CardPackage) GetHncardInfo(ctx context.Context, args *order.ArgsGetRpcHncard, reply *order.CardHncardInfo) error {
	return c.Call(ctx, "GetHncardInfo", args, reply)
}

//获取卡包-充值卡【RPC】
func (c *CardPackage) GetRcardInfo(ctx context.Context, args *order.ArgsGetRpcRcard, reply *order.CardRcardInfo) error {
	return c.Call(ctx, "GetRcardInfo", args, reply)
}

//获取卡包-身份卡【RPC】
func (c *CardPackage) GetIcardInfo(ctx context.Context, args *order.ArgsGetRpcIcard, reply *order.CardIcardInfo) error {
	return c.Call(ctx, "GetIcardInfo", args, reply)
}

//获取用户关联表-rpc
func (c *CardPackage) GetUserCardPackageByUser(ctx context.Context, args *order.ArgsGetUserCardPackageByUser, reply *order.ReplyGetUserCardPackageByUser) error {
	return c.Call(ctx, "GetUserCardPackageByUser", args, reply)
}

//根据卡包关联IDs获取卡包信息-【RPC】
func (c *CardPackage) GetCardPackageListByIdsRpc(ctx context.Context, ids *[]int, reply *order.ReplyGetCardPackageListByIds) error {
	return c.Call(ctx, "GetCardPackageListByIdsRpc", ids, reply)
}

//卡包充值记录
func (c *CardPackage) GetCardPackageRechangeLog(ctx context.Context, args *order.ArgsRechangeLog, reply *order.ReplyRechangeLog) error {
	return c.Call(ctx, "GetCardPackageRechangeLog", args, reply)
}

//获取用户持卡数量出参-rpc
func (c *CardPackage) GetUserCardPackageCountRpc(ctx context.Context, args *order.ArgsGetUserCardPackageCountRpc, reply *order.ReplyGetUserCardPackageCountRpc) error {
	return c.Call(ctx, "GetUserCardPackageCountRpc", args, reply)
}

//根据卡包Id获取的busId--rpc
func (c *CardPackage) GetCardPackageBusIdRpc(ctx context.Context, args *order.ArgsGetCardPackageIdRpc, reply *order.ReplyGetCardPackageIdRpc) error {
	return c.Call(ctx, "GetCardPackageBusIdRpc", args, reply)
}

//获取卡包详情包含的单项目
func (c *CardPackage) GetCardPackageInfoCardSingle(ctx context.Context, args *order.ArgsGetCardPackageInfoCardSingleGoods, reply *order.ReplyGetCardPackageInfoCardSingle) error {
	return c.Call(ctx, "GetCardPackageInfoCardSingle", args, reply)
}

//获取卡包详情包含的商品
func (c *CardPackage) GetCardPackageInfoCardGoods(ctx context.Context, args *order.ArgsGetCardPackageInfoCardSingleGoods, reply *order.ReplyGetCardPackageInfoCardGood) error {
	return c.Call(ctx, "GetCardPackageInfoCardGoods", args, reply)
}

//根据充值编号获取充值记录
func (c *CardPackage) GetRcardRechargeBySnRpc(ctx context.Context, args *order.ArgsGetRcardRechargeBySn, reply *order.ReplyGetRcardRechargeBySn) error {
	return c.Call(ctx, "GetRcardRechargeBySnRpc", args, reply)
}

//根据relationId 获取卡包基础信息 --rpc
func (c *CardPackage) GetCardPackageByRelation(ctx context.Context, relationId *int, reply *order.ReplyGetCardPackageByRelation) error {
	return c.Call(ctx, "GetCardPackageByRelation", relationId, reply)
}

//用户预付卡/预约中单子数量
func (c *CardPackage) GetUserCardPackageNum(ctx context.Context, args *order.ArgsGetUserCardPackageNum, reply *order.ReplyGetUserCardPackageNum) error {
	return c.Call(ctx, "GetUserCardPackageNum", args, reply)
}

//更新身份卡折扣同步标识到卡包中
func (c *CardPackage) UpdateIcardPackageSyncLogo(ctx context.Context, icardId *int, reply *bool) error {
	return c.Call(ctx, "UpdateIcardPackageSyncLogo", icardId, reply)
}

//获取用户权益及充值卡数据
func (c *CardPackage) GetUserEquityByRpc(ctx context.Context, args *order.ArgsGetChooseMarket, reply *order.ResponseChooseMarketList) error {
	return c.Call(ctx, "GetUserEquityByRpc", args, reply)
}
