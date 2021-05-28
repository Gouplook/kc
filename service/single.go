//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/9 13:12
package service

import (
	"context"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Single struct {
}

//查询单项目-rpc
func (s *Single) GetSimpleSingleInfo(ctx context.Context, singleId *int, reply *cards.SimpleSingleInfo) (err error) {
	mSingle := new(logics.SingleLogic)
	res := mSingle.GetSimpleSingleInfo(ctx, *singleId)
	if err = mapstructure.WeakDecode(res, reply); err != nil {
		return
	}
	return nil
}

//批量查询单项目-rpc
func (s *Single) GetSimpleSingleInfos(ctx context.Context, singleIds *[]int, reply *[]cards.SimpleSingleInfo) (err error) {
	mSingle := new(logics.SingleLogic)
	res := mSingle.GetSimpleSingleInfos(ctx, *singleIds)
	if err = mapstructure.WeakDecode(res, reply); err != nil {
		return
	}
	return nil
}

//添加单项目
func (s *Single) AddSingle(ctx context.Context, single *cards.ArgsAddSingle, singleId *int) (err error) {
	mSingle := new(logics.SingleLogic)
	*singleId, err = mSingle.AddSingle(ctx, single)
	if err != nil {
		return
	}
	return nil
}

//获取单项目信息
func (s *Single) GetSingleInfo(ctx context.Context, args *cards.ArgsGetSingleInfo, reply *cards.SingleDetail) (err error) {
	mSingle := new(logics.SingleLogic)
	uid := 0
	uid, _ = args.GetUid()
	if uid == 0 && args.Uid > 0 {
		uid = args.Uid
	}

	*reply, err = mSingle.SingleInfo(ctx, args.SingleId, uid, args.ShopId)
	if err != nil {
		return
	}
	return nil
}

//修改单项目信息
func (s *Single) EditSingle(ctx context.Context, single *cards.ArgsEditSingle, reply *bool) error {
	mSingle := new(logics.SingleLogic)
	*reply = true
	err := mSingle.EditSingle(ctx, single)
	if err != nil {
		*reply = false
		return err
	}
	return nil
}

//子店添加单项目
func (s *Single) ShopAddSingle(ctx context.Context, single *cards.ArgsShopAddSingle, reply *bool) error {
	mSingle := new(logics.SingleLogic)
	*reply = true
	err := mSingle.ShopAddSingle(single)
	if err != nil {
		*reply = false
		return err
	}
	return nil
}

//获取商家的单项目列表
func (s *Single) BusSinglePage(ctx context.Context, args *cards.ArgsBusSinglePage, reply *cards.ReplyBusSinglePage) (err error) {
	mSingle := new(logics.SingleLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mSingle.GetBusSingles(ctx, args.BusId, args.ShopId, start, limit, args.IsGround, args.IsDel,args.FilterShopHasAdd)
	if err != nil {
		return
	}
	return nil
}

//子店铺设置单项目价格
func (s *Single) ShopChangePrice(ctx context.Context, args *cards.ArgsShopChangePrice, reply *bool) (err error) {
	mSingle := new(logics.SingleLogic)
	*reply = true
	err = mSingle.ShopChangePrice(args)
	if err != nil {
		*reply = false
		return
	}
	return
}

/*//总店上下架单项目操作
func (s *Single) DownUpSingle(ctx context.Context, args *cards.ArgsDownUpSingle, reply *bool) (err error) {
	mSingle := new(logics.SingleLogic)
	*reply = true
	err = mSingle.BusUpDownSingle(ctx, args)
	if err != nil {
		*reply = false
		return
	}
	return
}*/

//总店 删除单项目
func (s *Single) DelSingle(ctx context.Context, args *cards.ArgsDelSingle, reply *bool) (err error) {
	delSingle := new(logics.SingleLogic)
	*reply = true
	err = delSingle.DelSingle(ctx, args)
	if err != nil {
		*reply = false
	}
	return err
}

//分店 删除单项目
func (s *Single) DelShopSingle(ctx context.Context, args *cards.ArgsDelSingle, reply *bool) (err error) {
	shopId, err := args.GetShopId()
	if err != nil {
		return err
	}
	dss := new(logics.SingleLogic)
	*reply = true
	if err = dss.DelShopSingle(ctx, shopId, args); err != nil {
		*reply = false
	}
	return
}

//获取子店的单项目
func (s *Single) ShopSinglePage(ctx context.Context, args *cards.ArgsShopSinglePage, reply *cards.ReplyShopSinglePage) (err error) {
	mSingle := new(logics.SingleLogic)
	start := args.GetStart()
	limit := args.GetPageSize()
	*reply, err = mSingle.GetShopSingles(ctx, args.ShopId, start, limit, args.SingleIds, args.Status, args.IsDel)
	if err != nil {
		return err
	}
	return nil
}

//子店上下架单项目
func (s *Single) ShopDownUpSingle(ctx context.Context, args *cards.ArgsShopDownUpSingle, reply *bool) (err error) {
	mSingle := new(logics.SingleLogic)
	*reply = true
	err = mSingle.ShopDownUpSingle(ctx, args)
	if err != nil {
		*reply = false
		return
	}
	return
}

//获取项目在子店的价格
func (s *Single) GetShopSinglePrice(ctx context.Context, args *cards.ArgsGetShopSinglePrice, price *float64) (err error) {
	mSingle := new(logics.SingleLogic)
	*price, err = mSingle.GetShopSinglePrice(args.SsId, args.SspId)
	if err != nil {
		return
	}
	return
}

//获取所有的属性标签
func (s *Single) GetAttrs(ctx context.Context, emptyStr string, reply *cards.ReplyGetAttrs) (err error) {
	mSingle := new(logics.SingleLogic)
	*reply = mSingle.GetAttrs()

	return nil
}

// 根据单项目id批量获取基础价格信息
func (s *Single) GetSinglePriceListsBySingleIds(ctx context.Context, singleIds []int, reply *map[int]cards.SinglePriceInfo) (err error) {
	*reply = new(logics.SingleLogic).GetSinglePriceListsBySingleIds(singleIds)
	return nil
}

// 根据手艺人ID获取关联的单项目
func (s *Single) GetSignlesByStaffId(ctx context.Context, args *cards.ArgsGetSignlesByStaffID, reply *cards.ReplyGetSignlesByStaffID) (err error) {
	return new(logics.SingleLogic).GetSignlesByStaffId(ctx, args, reply)
}

//获取门店的单项目-rpc内部调用
func (s *Single) GetShopSingleBySingleIdsRpc(ctx context.Context, args *cards.ArgsGetShopSingleBySingleIdsRpc, reply *cards.ReplyGetShopSingleBySingleIdsRpc) (err error) {
	return new(logics.SingleLogic).GetShopSingleBySingleIdsRpc(ctx, args, reply)
}

//根据门店和标签id查询服务
func (s *Single) GetSingleByShopIdAndTagId(ctx context.Context, args *cards.ArgsShopSingleByPage, reply *cards.ReplyShopSingle) error {
	return new(logics.SingleLogic).GetSingleByShopIdAndTagId(ctx, args, reply)
}

//根据单项目id获取子服务
func (s *Single) GetSubServerBySingleId(ctx context.Context, args *cards.ArgsSubServer, reply *cards.ReplySubServer) error {
	return new(logics.SingleLogic).GetSubServerBySingleId(args, reply)
}

//根据门店id和单项目Id获取单项目数据
func (s *Single) GetSingleByShopIdAndSingleIds(ctx context.Context, args *cards.ArgsGetSingleByShopIdAndSingleIds, reply *cards.ReplyShopSingle) error {
	return new(logics.SingleLogic).GetSingleByShopIdAndSingleIds(ctx, args, reply)
}

func (s *Single) GetBySsidsRpc(ctx context.Context, ssIds *[]int, reply *[]cards.ReplyGetBySsidsRpc) error {
	return new(logics.SingleLogic).GetBySsidsRpc(ssIds, reply)
}

//根据指定的规格组合id和门店id 获取数据
func (s *Single) GetShopSpecs(ctx context.Context, args *cards.ArgsGetShopSpecs, reply *[]cards.ReplyGetShopSpecs) error {
	return new(logics.SingleLogic).GetShopSpecs(args, reply)
}

//根据规格ID获取规格数据
func (s *Single) GetSingleSpecBySspId(ctx context.Context, args *cards.ArgsSubSpecID, reply *cards.ReplySubServer2) (err error) {
	return new(logics.SingleLogic).GetSingleSpecBySspId(args, reply)
}

//根据sspIds获取对应的singleId -rpc内部使用 消费确认时检查sspid和singleid的匹配
func (s *Single) GetBySspids(ctx context.Context, args *[]int, reply *[]cards.ReplyGetBySspids) error {
	return new(logics.SingleLogic).GetBySspids(args, reply)
}

//根据shopId和批量组合规格查询-rpc确认消费
func (s *Single) GetByShopSspIds(ctx context.Context, args *cards.ArgsGetShopSpecs, reply *[]cards.ReplyCommonSingleSpec) (err error) {
	singleLogic := new(logics.SingleLogic)
	res := singleLogic.GetByShopSspIds(args.ShopId, args.SspIds)
	if err = mapstructure.WeakDecode(&res, reply); err != nil {
		return
	}
	return
}

//根据singleids批量获取门店单项目-rpc确认消费
func (s *Single) GetByShopSingle(ctx context.Context, args *cards.ArgsCommonShopSingle, reply *[]cards.ReplyCommonShopSingle) (err error) {
	singleLogic := new(logics.SingleLogic)
	res := singleLogic.GetByShopSingle(args.ShopId, args.SingleIds, args.Status)
	if err = mapstructure.WeakDecode(&res, reply); err != nil {
		return
	}
	return
}

//获取门店指定单项目规格的价格
func (s *Single) GetPriceByShopIdAndSsspId(ctx context.Context, args *cards.ArgsCommonShopSingle, reply *[]cards.ReplyGetPriceByShopIdAndSsspId) (err error) {
	singleLogic := new(logics.SingleLogic)
	res := singleLogic.GetPriceByShopIdAndSsspId(args.ShopId, args.SspIds)
	if err = mapstructure.WeakDecode(&res, reply); err != nil {
		return
	}
	return
}

//根据singleids批量获取单项目-rpc确认消费
func (s *Single) GetBySingle(ctx context.Context, singleIds *[]int, reply *[]cards.ReplyCommonSingle) (err error) {
	singleLogic := new(logics.SingleLogic)
	res := singleLogic.GetBySingle(*singleIds)
	if err = mapstructure.WeakDecode(&res, reply); err != nil {
		return
	}
	return
}

//九百岁首页精选服务
func (s *Single) GetSelectServices(ctx context.Context, args *cards.ArgsGetSelectServices, reply *[]cards.ReplyGetSelectServices) error {
	return new(logics.SingleLogic).GetSelectServices(ctx, args, reply)
}
