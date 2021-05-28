//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/9 13:42
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Single struct {
	client.Baseclient
}

func (s *Single) Init() *Single {
	s.ServiceName = "rpc_cards"
	s.ServicePath = "Single"
	return s
}

//单项目详情-rpc
func (s *Single) GetSimpleSingleInfo(ctx context.Context, singleId *int, reply *cards.SimpleSingleInfo) error {
	return s.Call(ctx, "GetSimpleSingleInfo", singleId, reply)
}

//批量单项目详情-rpc
func (s *Single) GetSimpleSingleInfos(ctx context.Context, singleIds *[]int, reply *[]cards.SimpleSingleInfo) error {
	return s.Call(ctx, "GetSimpleSingleInfos", singleIds, reply)
}

//添加单项目数据
func (s *Single) AddSingle(ctx context.Context, single *cards.ArgsAddSingle, singleId *int) error {
	return s.Call(ctx, "AddSingle", single, singleId)
}

//获取单项目数据
func (s *Single) GetSingleInfo(ctx context.Context, args *cards.ArgsGetSingleInfo, reply *cards.SingleDetail) error {
	return s.Call(ctx, "GetSingleInfo", args, reply)
}

//修改单项目数据
func (s *Single) EditSingle(ctx context.Context, single *cards.ArgsEditSingle, reply *bool) error {
	return s.Call(ctx, "EditSingle", single, reply)
}

//子店添加单项目
func (s *Single) ShopAddSingle(ctx context.Context, single *cards.ArgsShopAddSingle, reply *bool) error {
	return s.Call(ctx, "ShopAddSingle", single, reply)
}

//获取商家的单项目列表
func (s *Single) BusSinglePage(ctx context.Context, args *cards.ArgsBusSinglePage, reply *cards.ReplyBusSinglePage) error {
	return s.Call(ctx, "BusSinglePage", args, reply)
}

//子店铺设置单项目价格
func (s *Single) ShopChangePrice(ctx context.Context, args *cards.ArgsShopChangePrice, reply *bool) error {
	return s.Call(ctx, "ShopChangePrice", args, reply)
}

//总店上下架单项目操作
func (s *Single) DownUpSingle(ctx context.Context, args *cards.ArgsDownUpSingle, reply *bool) error {
	return s.Call(ctx, "DownUpSingle", args, reply)
}

//获取子店的单项目列表
func (s *Single) ShopSinglePage(ctx context.Context, args *cards.ArgsShopSinglePage, reply *cards.ReplyShopSinglePage) error {
	return s.Call(ctx, "ShopSinglePage", args, reply)
}

//子店上下架单项目
func (s *Single) ShopDownUpSingle(ctx context.Context, args *cards.ArgsShopDownUpSingle, reply *bool) error {
	return s.Call(ctx, "ShopDownUpSingle", args, reply)
}

//获取单项目在子店的价格
func (s *Single) GetShopSinglePrice(ctx context.Context, args *cards.ArgsGetShopSinglePrice, price *float64) error {
	return s.Call(ctx, "GetShopSinglePrice", args, price)
}

//获取所有的属性标签数据
func (s *Single) GetAttrs(ctx context.Context, reply *cards.ReplyGetAttrs) error {
	return s.Call(ctx, "GetAttrs", "", reply)
}

//根据单项目id批量获取基础价格信息
func (s *Single) GetSinglePriceListsBySingleIds(ctx context.Context, singleIds []int, reply *map[int]cards.SinglePriceInfo) error {
	return s.Call(ctx, "GetSinglePriceListsBySingleIds", singleIds, reply)
}

//总店 删除单项目
func (s *Single) DelSingle(ctx context.Context, args *cards.ArgsDelSingle, reply *bool)  error {
	return s.Call(ctx, "DelSingle", args, reply)
}

//分店 删除单项目
func (s *Single) DelShopSingle(ctx context.Context, args *cards.ArgsDelSingle, reply *bool)  error {
	return s.Call(ctx, "DelShopSingle", args, reply)
}


// 根据手艺人ID获取关联的单项目
func (s *Single) GetSignlesByStaffId(ctx context.Context, args *cards.ArgsGetSignlesByStaffID, reply *cards.ReplyGetSignlesByStaffID) error {
	return s.Call(ctx, "GetSignlesByStaffId", args, reply)
}

//获取门店的单项目-rpc内部调用
func (s *Single) GetShopSingleBySingleIdsRpc(ctx context.Context, args *cards.ArgsGetShopSingleBySingleIdsRpc, reply *cards.ReplyGetShopSingleBySingleIdsRpc) error {
	return s.Call(ctx, "GetShopSingleBySingleIdsRpc", args, reply)
}

//根据门店id查询单项目服务
func (s *Single) GetSingleByShopIdAndTagId(ctx context.Context, args *cards.ArgsShopSingleByPage, reply *cards.ReplyShopSingle) error {
	return s.Call(ctx, "GetSingleByShopIdAndTagId", args, reply)
}

//根据单项目id获取子规格服务
func (s *Single) GetSubServerBySingleId(ctx context.Context, args *cards.ArgsSubServer, reply *cards.ReplySubServer) error {
	return s.Call(ctx, "GetSubServerBySingleId", args, reply)
}

//根据门店id和单项目Id获取单项目数据
func (s *Single)GetSingleByShopIdAndSingleIds(ctx context.Context,args *cards.ArgsGetSingleByShopIdAndSingleIds,reply *cards.ReplyShopSingle)error{
	return s.Call(ctx, "GetSingleByShopIdAndSingleIds", args, reply)
}

//根据单项目门店ids 获取信息-rpc内部调用
func (s *Single) GetBySsidsRpc(ctx context.Context, ssIds *[]int, reply *[]cards.ReplyGetBySsidsRpc) error {
	return s.Call(ctx, "GetBySsidsRpc", ssIds, reply)
}

///根据规格ID获取门店规格数据
func (s *Single) GetShopSpecs(ctx context.Context, args *cards.ArgsGetShopSpecs, reply *[]cards.ReplyGetShopSpecs) error {
	return s.Call(ctx, "GetShopSpecs", args, reply)
}

//根据规格ID获取门店规格数据
func (s *Single)GetSingleSpecBySspId(ctx context.Context,args *cards.ArgsSubSpecID,reply *cards.ReplySubServer2)error{
	return  s.Call(ctx, "GetSingleSpecBySspId", args, reply)
}


//根据sspIds获取对应的singleId -rpc内部使用 消费确认时检查sspid和singleid的匹配
func (s *Single) GetBySspids(ctx context.Context, args *[]int, reply *[]cards.ReplyGetBySspids) error {
	return s.Call(ctx, "GetBySspids", args, reply)
}

//根据shopId和批量组合规格查询-rpc确认消费
func (s *Single) GetByShopSspIds(ctx context.Context,args *cards.ArgsGetShopSpecs,reply *[]cards.ReplyCommonSingleSpec)error{
	return  s.Call(ctx, "GetByShopSspIds", args, reply)
}

//根据singleids批量获取门店单项目-rpc确认消费
func (s *Single) GetByShopSingle(ctx context.Context,args *cards.ArgsCommonShopSingle,reply *[]cards.ReplyCommonShopSingle)error{
	return  s.Call(ctx, "GetByShopSingle", args, reply)
}

//获取门店指定单项目规格的价格
func (s *Single)GetPriceByShopIdAndSsspId(ctx context.Context,args *cards.ArgsCommonShopSingle,reply *[]cards.ReplyGetPriceByShopIdAndSsspId)error{
	return s.Call(ctx, "GetPriceByShopIdAndSsspId", args, reply)
}

//根据singleids批量获取单项目-rpc确认消费
func (s *Single) GetBySingle(ctx context.Context,singleIds *[]int,reply *[]cards.ReplyCommonSingle)error{
	return  s.Call(ctx, "GetBySingle", singleIds, reply)
}

//九百岁首页精选服务
func (s *Single) GetSelectServices(ctx context.Context, args *cards.ArgsGetSelectServices, reply *[]cards.ReplyGetSelectServices) error {
	return  s.Call(ctx, "GetSelectServices", args, reply)
}
