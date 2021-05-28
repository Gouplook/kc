package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

//前台
type Front struct {
	client.Baseclient
}

func (s *Front) Init() *Front {
	s.ServiceName = "rpc_bus"
	s.ServicePath = "Front"
	return s
}

//获取 支出类目
func (s *Front) GetDisburseCategory(ctx context.Context, args *bus.ArgsDisburseCategoryInfo, reply *[]bus.ReplyCategoryType) error {
	return s.Call(ctx, "GetDisburseCategory", args, reply)
}

//添加 支出类目
func (s *Front) AddDisburseCategory(ctx context.Context, args *bus.ArgsDisburseCategoryInfo, reply *bus.ReplyDisburseCategoryRes) error {
	return s.Call(ctx, "AddDisburseCategory", args, reply)
}

//修改 支出类目
func (s *Front) UpdateDisburseCategory(ctx context.Context, args *bus.ArgsDisburseCategoryInfo, reply *bus.ReplyDisburseCategoryRes) error {
	return s.Call(ctx, "UpdateDisburseCategory", args, reply)
}

//删除 支出类目
func (s *Front) DeleteDisburseCategory(ctx context.Context, args *bus.ArgsDisburseCategoryInfo, reply *bus.ReplyDisburseCategoryRes) error {
	return s.Call(ctx, "DeleteDisburseCategory", args, reply)
}

//实现 支出明细的 添加
func (s *Front) AddDisburseDetail(ctx context.Context, args *bus.ArgsDisburseDetailInfo, reply *bus.ReplyDisburseDetailRes) error {
	return s.Call(ctx, "AddDisburseDetail", args, reply)
}

//实现 支出明细的 获取
func (s *Front) GetDisburseDetail(ctx context.Context, args *bus.ArgsDisburseDetailReq, reply *bus.ReplyDisburseDetailInfo) error {
	return s.Call(ctx, "GetDisburseDetail", args, reply)
}

//实现 获取u一条支出明细
func (s *Front) GetOneDisburseDetail(ctx context.Context, args int, reply *bus.DisburseDetailInfo) error {
	return s.Call(ctx, "GetOneDisburseDetail", args, reply)
}

//实现 支出明细的 修改
func (s *Front) UpdateDisburseDetail(ctx context.Context, args *bus.ArgsDisburseDetailInfo, reply *bus.ReplyDisburseDetailRes) error {
	return s.Call(ctx, "UpdateDisburseDetail", args, reply)
}

//实现 支出明细的 删除
func (s *Front) DeleteDisburseDetail(ctx context.Context, args *bus.ArgsDisburseDetailInfo, reply *bus.ReplyDisburseDetailRes) error {
	return s.Call(ctx, "DeleteDisburseDetail", args, reply)
}

//添加/更新消毒地点
func (s *Front) AddOrUpdateDisinfectAddress(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, replyId *int) error {
	return s.Call(ctx, "AddOrUpdateDisinfectAddress", args, replyId)
}

//删除消毒地点-软删除
func (s *Front) DelDisinfectAddress(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, reply *bool) error {
	return s.Call(ctx, "DelDisinfectAddress", args, reply)
}

//消毒地点数据获取
func (s *Front) GetDisinfectAddress(ctx context.Context, args *bus.GetDisinfectFrontBase, reply *bus.ReplyGetDisinfectFrontCommon) error {
	return s.Call(ctx, "GetDisinfectAddress", args, reply)
}

//添加/更新消毒类型
func (s *Front) AddOrUpdateDisinfectType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, replyId *int) error {
	return s.Call(ctx, "AddOrUpdateDisinfectType", args, replyId)
}

//删除消毒类型-软删除
func (s *Front) DelDisinfectType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, reply *bool) error {
	return s.Call(ctx, "DelDisinfectType", args, reply)
}

//消毒类型数据获取
func (s *Front) GetDisinfectType(ctx context.Context, args *bus.GetDisinfectFrontBase, reply *bus.ReplyGetDisinfectFrontCommon) error {
	return s.Call(ctx, "GetDisinfectType", args, reply)
}

//添加/更新消毒物品
func (s *Front) AddOrUpdateDisinfectGood(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, replyId *int) error {
	return s.Call(ctx, "AddOrUpdateDisinfectGood", args, replyId)
}

//删除消毒物品-软删除
func (s *Front) DelDisinfectGood(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, reply *bool) error {
	return s.Call(ctx, "DelDisinfectGood", args, reply)
}

//消毒物品数据获取
func (s *Front) GetDisinfectGoods(ctx context.Context, args *bus.GetDisinfectFrontBase, reply *bus.ReplyGetDisinfectFrontCommon) error {
	return s.Call(ctx, "GetDisinfectGoods", args, reply)
}

//添加/更新垃圾处理方式
func (s *Front) AddOrUpdateGarbageDisposetType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, replyId *int) error {
	return s.Call(ctx, "AddOrUpdateGarbageDisposetType", args, replyId)
}

//删除垃圾处理方式-软删除
func (s *Front) DelGarbageDisposetType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, reply *bool) error {
	return s.Call(ctx, "DelGarbageDisposetType", args, reply)
}

//垃圾处理方式数据获取
func (s *Front) GetGarbageDisposetType(ctx context.Context, args *bus.GetDisinfectFrontBase, reply *bus.ReplyGetDisinfectFrontCommon) error {
	return s.Call(ctx, "GetGarbageDisposetType", args, reply)
}

//添加/更新垃圾类型
func (s *Front) AddOrUpdateGarbageType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, replyId *int) error {
	return s.Call(ctx, "AddOrUpdateGarbageType", args, replyId)
}

//删除垃圾类型-软删除
func (s *Front) DelGarbageType(ctx context.Context, args *bus.ArgsAddOrUpdateDelDisinfectFrontBase, reply *bool) error {
	return s.Call(ctx, "DelGarbageType", args, reply)
}

//垃圾类型数据获取
func (s *Front) GetGarbageType(ctx context.Context, args *bus.GetDisinfectFrontBase, reply *bus.ReplyGetDisinfectFrontCommon) error {
	return s.Call(ctx, "GetGarbageType", args, reply)
}

//添加 环境消毒明细
func (s *Front) AddSetting(ctx context.Context, args *bus.ArgsEpidemicSettingInfo, reply *bus.ReplyEpidemicSetting) error {
	return s.Call(ctx, "AddSetting", args, reply)
}

//删除 一条环境消毒明细
func (s *Front) DeleteSetting(ctx context.Context, args *bus.ArgsEpidemicSettingInfo, reply *bus.ReplyEpidemicSetting) error {
	return s.Call(ctx, "DeleteSetting", args, reply)
}

//查询 环境消毒明细
func (s *Front) GetSetting(ctx context.Context, args *bus.ArgsEpidemicSettingInfo, reply *bus.ReplyEpidemicSettingGet) error {
	return s.Call(ctx, "GetSetting", args, reply)
}

//添加 用品消毒明细
func (s *Front) AddTackle(ctx context.Context, args *bus.ArgsEpidemicTackleInfo, reply *bus.ReplyEpidemicTackle) error {
	return s.Call(ctx, "AddTackle", args, reply)
}

//删除 一条用品消毒记录
func (s *Front) DeleteTackle(ctx context.Context, args *bus.ArgsEpidemicTackleInfo, reply *bus.ReplyEpidemicTackle) error {
	return s.Call(ctx, "DeleteTackle", args, reply)
}

//查询 用品消毒 记录
func (s *Front) GetTackle(ctx context.Context, args *bus.ArgsEpidemicTackleInfo, reply *bus.ReplyEpidemicTackleGet) error {
	return s.Call(ctx, "GetTackle", args, reply)
}

//添加 垃圾处理 记录
func (s *Front) AddGarbage(ctx context.Context, args *bus.ArgsEpidemicGarbageInfo, reply *bus.ReplyEpidemicGarbage) error {
	return s.Call(ctx, "AddGarbage", args, reply)
}

//删除 一条垃圾处理 记录
func (s *Front) DeleteGarbage(ctx context.Context, args *bus.ArgsEpidemicGarbageInfo, reply *bus.ReplyEpidemicGarbage) error {
	return s.Call(ctx, "DeleteGarbage", args, reply)
}

//查询 垃圾处理 记录
func (s *Front) GetGarbage(ctx context.Context, args *bus.ArgsEpidemicGarbageInfo, reply *bus.ReplyEpidemicGarbageGet) error {
	return s.Call(ctx, "GetGarbage", args, reply)
}

//添加 技师健康 记录
func (s *Front) AddTechnician(ctx context.Context, args *bus.ArgsEpidemicTechnicianInfo, reply *bus.ReplyEpidemicTechnician) error {
	return s.Call(ctx, "AddTechnician", args, reply)
}

//删除 一条技师健康 记录
func (s *Front) DeleteTechnician(ctx context.Context, args *bus.ArgsEpidemicTechnicianInfo, reply *bus.ReplyEpidemicTechnician) error {
	return s.Call(ctx, "DeleteTechnician", args, reply)
}

//查询 技师健康 记录
func (s *Front) GetTechnician(ctx context.Context, args *bus.ArgsEpidemicTechnicianInfo, reply *bus.ReplyEpidemicTechnicianGet) error {
	return s.Call(ctx, "GetTechnician", args, reply)
}

//查看门店当天的防疫情况
func (s *Front) GetShopEpidemic(ctx context.Context, shopId *int, reply *bus.ReplyGetShopEpidemic) error {
	return s.Call(ctx, "GetTechnician", shopId, reply)
}

//前台总部支出汇总查询
func (s *Front) GetBusDisburseSum(ctx context.Context, args *bus.ArgsGetBusDisburse, reply *bus.ReplyGetBusDisburse) error {
	return s.Call(ctx, "GetBusDisburseSum", args, reply)
}

//添加 复工申请
func (s *Front) AddWork(ctx context.Context, args *bus.ArgsReturnWork, reply *bus.ReplyReturnWork) error {
	return s.Call(ctx, "AddWork", args, reply)
}

//查询 审核状态
func (s *Front) GetVerify(ctx context.Context, args *bus.ArgsReturnWork, reply *bus.ReplyReturnWorkGet) error {
	return s.Call(ctx, "GetVerify", args, reply)
}

//门店添加员工登记
func (s *Front) AddStaffRegister(ctx context.Context, args *bus.ArgsAddStaffRegister, reply *int) error {
	return s.Call(ctx, "AddStaffRegister", args, reply)
}

//门店查询员工登记
func (s *Front) GetStaffRegister(ctx context.Context, args *bus.ArgsGetStaffRegister, reply *bus.ReplyGetStaffRegister) error {
	return s.Call(ctx, "GetStaffRegister", args, reply)
}

//门店添加防控日报
func (s *Front) AddDefendDaily(ctx context.Context, args *bus.ArgsAddDefendDaily, reply *int) error {
	return s.Call(ctx, "AddDefendDaily", args, reply)
}

//门店查询防控日报
func (s *Front) GetDefendDaily(ctx context.Context, args *bus.ArgsGetDefendDaily, reply *bus.ReplyGetDefendDaily) error {
	return s.Call(ctx, "GetDefendDaily", args, reply)
}
