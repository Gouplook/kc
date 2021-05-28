package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

type Shop struct {
	client.Baseclient
}

//初始化
func (s *Shop) Init() *Shop {
	s.ServiceName = "rpc_bus"
	s.ServicePath = "Shop"
	return s
}

//分店信息申请
func (s *Shop) ShopSettled(ctx context.Context, args *bus.ArgsBusShopReg, reply *bus.ReplyBusShopRegUp) error {
	return s.Call(ctx, "ShopSettled", args, reply)
}

//分店信息申请-开放平台
func (s *Shop) OpenPlatFormV1ShopSettled(ctx context.Context, args *bus.ArgsBusShopReg, reply *bus.ReplyBusShopRegUp) error {
	return s.Call(ctx, "OpenPlatFormV1ShopSettled", args, reply)
}

//审核失败重新提交
func (s *Shop) ShopUpdate(ctx context.Context, args *bus.ArgsBusShopRepeat, reply *bus.ReplyBusShopRegUp) error {
	return s.Call(ctx, "ShopUpdate", args, reply)
}

//获取分店-根据分店ID
func (s *Shop) GetShopByShopid(ctx context.Context, args *bus.ArgsGetShop, reply *bus.ReplyShopInfo) error {
	return s.Call(ctx, "GetShopByShopid", args, reply)
}

//获取分店-获取审核通过的分店
func (s *Shop) GetAvailableShopByBusId(ctx context.Context, args *bus.ArgsAvaBusId, reply *[]bus.ReplyShopInfos) error {
	return s.Call(ctx, "GetAvailableShopByBusId", args, reply)
}

//批量获取分店-根据分店IDS
func (s *Shop) GetShopByShopids(ctx context.Context, args *bus.ArgsGetShops, reply *[]bus.ReplyShopInfos) error {
	return s.Call(ctx, "GetShopByShopids", args, reply)
}

//批量获取审核后的分店-根据分店IDS
func (s *Shop) GetAvailableShopByShopids(ctx context.Context, args *bus.ArgsAvaGetShops, reply *[]bus.ReplyShopInfos) error {
	return s.Call(ctx, "GetAvailableShopByShopids", args, reply)
}

//批量获取分店-企业/商户ID
func (s *Shop) GetShopByBusId(ctx context.Context, args *bus.ArgsBusId, reply *[]bus.ReplyShopInfos) error {
	return s.Call(ctx, "GetShopByBusId", args, reply)
}

//检测门店信息的合法性
func (s *Shop) CheckBusShop(ctx context.Context, args *bus.ArgsCheckShop, reply *bus.ReplyCheckShop) error {
	return s.Call(ctx, "CheckBusShop", args, reply)
}

//更新分店设置入参
func (s *Shop) UpdateBusShopSetting(ctx context.Context, args *bus.ArsgUpdateBusShopSetting, reply *bool) error {
	return s.Call(ctx, "UpdateBusShopSetting", args, reply)
}

//实现批量获取分店详细地址和封面图片-根据批量分店id-rpc内部调用
//map[1:map[Address:永乐路1908号 BranchName:宝山店 Image:https://img.900sui.cn/f59b9ad87f8342148729a8fe2389fd7d Phone:13916379354 ShopName:美丽家]
//2:map[Address:永乐路1908号 BranchName:徐汇店 Image:https://img.900sui.cn/1971e438658b4ca6af88dc6c09afba3c Phone:13916379354 ShopName:足浴养生]]
func (s *Shop) GetShopAddressAndImgByIdS(ctx context.Context, args *[]int, reply *map[int]map[string]interface{}) error {
	return s.Call(ctx, "GetShopAddressAndImgByIdS", args, reply)
}

//	附近门店详情
func (s *Shop) GetNearbyShopInfo(ctx context.Context, args *bus.ArgsGetNearbyShopInfo, reply *bus.ReplyGetNearbyShopInfo) error {
	return s.Call(ctx, "GetNearbyShopInfo", args, reply)
}

//	门店详情-其他分店
func (s *Shop) GetOthersShopList(ctx context.Context, args *bus.ArgsOthersShopList, reply *bus.ReplyOthersShopList) error {
	return s.Call(ctx, "GetOthersShopList", args, reply)
}

//增加分店评价统计Rpc
func (s *Shop) AddBusShopReportRpc(ctx context.Context, args *bus.ArgsAddBusShopReportRpc, reply *bool) error {
	return s.Call(ctx, "AddBusShopReportRpc", args, reply)
}

//根据分店ID获取分店评价数据
func (s *Shop) GetShopReportRankByShopID(ctx context.Context, shopId *int, reply *bus.ReplyGetShopReportRankByShopID) error {
	return s.Call(ctx, "GetShopReportRankByShopID", shopId, reply)
}

//获取分店人均消费数据
func (s *Shop) GetShopReportCapitaByShopId(ctx context.Context, args *bus.ArgsGetShopReportCapita, reply *bus.ReplyGetShopReportCapita) error {
	return s.Call(ctx, "GetShopReportCapitaByShopId", args, reply)
}

//商户服务-分店人均消费统计
func (s *Shop) AddShopReportCapitaRpc(ctx context.Context, args *bus.ArgsAddShopReportCapita, reply *bool) error {
	return s.Call(ctx, "GetShopReportCapita", args, reply)
}

//添加门店包含的优势标签
func (s *Shop) AddShopAdvantage(ctx context.Context, args *bus.ArgsAddShopAdvantage, reply *bool) error {
	return s.Call(ctx, "AddShopAdvantage", args, reply)
}

//获取门店包含的标签
func (s *Shop) GetAdvantageIdByShopId(ctx context.Context, args *bus.ArgsGetAdvantageIdByShopId, reply *bus.ReplyGetAdvantageIdByShopId) error {
	return s.Call(ctx, "GetAdvantageIdByShopId", args, reply)
}

//用户收藏
func (s *Shop) UserShopCollect(ctx context.Context, args *bus.ArgsUserShopCollect, reply *bool) error {
	return s.Call(ctx, "UserShopCollect", args, reply)
}

//获取用户收藏列表
func (s *Shop) ShopCollectList(ctx context.Context, args *bus.ArgsShopCollectList, reply *bus.ReplyShopCollectList) error {
	return s.Call(ctx, "ShopCollectList", args, reply)
}

//店铺收藏状态
func (s *Shop) ShopCollectStatus(ctx context.Context, args *bus.ArgsShopCollectStatus, reply *bus.ReplyShopCollectStatus) error {
	return s.Call(ctx, "ShopCollectStatus", args, reply)
}

//根据多个分店id获取分店名称-rpc
func (s *Shop) GetShopNameByShopIds(ctx context.Context, args *[]int, reply *[]bus.ReplyShopName) error {
	return s.Call(ctx, "GetShopNameByShopIds", args, reply)
}

// 获取风控系统商铺ID 门店总数
func (s *Shop) GetRiskShopNum(ctx context.Context, args *bus.ArgsRiskShopNum, reply *bus.ReplyRiskShopNum) error {
	return s.Call(ctx, "GetRiskShopNum", args, reply)
}

//监管平台直连接口-商户下的分店
func (s *Shop) GetGovShopLists(ctx context.Context, args *bus.ArgsGetGovShopLists, reply *bus.ReplyGetGovShopLists) error {
	return s.Call(ctx, "GetGovShopLists", args, reply)
}
