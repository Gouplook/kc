package bus

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/bus"
)

//Bus Bus
type Bus struct {
	client.Baseclient
}

//Init Init
func (b *Bus) Init() *Bus {
	b.ServiceName = "rpc_bus"
	b.ServicePath = "Bus"
	return b
}

//Reg 企业/商户入驻
func (b *Bus) Reg(ctx context.Context, args *bus.InputParamsReg, reply *bus.OutputReplyReg) error {
	return b.Call(ctx, "Reg", args, reply)
}

//企业/商户申请/更新主体
func (b *Bus) BusSettled(ctx context.Context, args *bus.ArgsBusReg, reply *bus.ReplyBusReg) error {
	return b.Call(ctx, "BusSettled", args, reply)
}

//企业/商户详情
func (b *Bus) GetByBusid(ctx context.Context, args *bus.ArgsSingleBus, reply *bus.ReplySingleBus) error {
	return b.Call(ctx, "GetByBusid", args, reply)
}

//企业/商户详情-登录验证
func (b *Bus) GetByBusidForUser(ctx context.Context, args *bus.ArgsSingleBusUser, reply *bus.ReplySingleBus) error {
	return b.Call(ctx, "GetByBusidForUser", args, reply)
}

//批量企业/商户信息
func (b *Bus) GetByBusids(ctx context.Context, args *bus.ArgsBatchBus, reply *[]bus.ReplyBatchBus) error {
	return b.Call(ctx, "GetByBusids", args, reply)
}

//检测用户绑定企业/商户信息
func (b *Bus) GetBusUserType(ctx context.Context, args *bus.ArgsBusUserType, reply *bus.ReplyBusUserType) error {
	return b.Call(ctx, "GetBusUserType", args, reply)
}

//获取企业/商户性质信息
func (b *Bus) GetBusinessType(ctx context.Context, reply *[]bus.ReplyBusinessType) error {
	return b.Call(ctx, "GetBusinessType", "", reply)
}

//变更企业/商户总账号
func (b *Bus) ReplaceAccount(ctx context.Context, args *bus.ArgsReplaceAccount, reply *bus.ReplyReplaceAccount) error {
	return b.Call(ctx, "ReplaceAccount", args, reply)
}

//检查行业是否属于商家
func (b *Bus) CheckBindid(ctx context.Context, args *bus.ArgsCheckBindid, reply *bool) error {
	return b.Call(ctx, "CheckBindid", args, reply)
}

//卡项服务获取商家显示信息
func (b *Bus) BusInfo(ctx context.Context, args *bus.ArgsBusInfo, reply *bus.ReplyBusInfo) error {
	return b.Call(ctx, "BusInfo", args, reply)
}

// AdminBusAudit 后台商户审核
func (b *Bus) AdminBusAudit(ctx context.Context, args *bus.ArgsAdminBusAudit, reply *bus.ReplyAdminBusAudit) error {
	return b.Call(ctx, "AdminBusAudit", args, reply)
}

// AdminBusAuditPage 后台审核管理-商户列表
func (b *Bus) AdminBusAuditPage(ctx context.Context, args *bus.ArgsAdminBusAuditPage, reply *bus.ReplyAdminBusAuditPage) error {
	return b.Call(ctx, "AdminBusAuditPage", args, reply)
}

// EsAdminBusInfo 后台获取商户详情(更新Es的时候调用)
func (b *Bus) EsAdminBusInfo(ctx context.Context, args *bus.ArgsEsAdminBusInfo, reply *bus.ReplyEsAdminBusInfo) error {
	return b.Call(ctx, "EsAdminBusInfo", args, reply)
}

//获取商家-服务设置选项
func (b *Bus) GetBusServiceSwitch(ctx context.Context, args *bus.ArgsGetBusServiceSwitch, reply *[]bus.ReplyGetBusServiceSwitch) error {
	return b.Call(ctx, "GetBusServiceSwitch", args, reply)
}

//更新商家-服务设置选项入参
func (b *Bus) UpdateBusServiceSwitch(ctx context.Context, args *bus.ArgsUpdateBusServiceSwitch, reply *bool) error {
	return b.Call(ctx, "UpdateBusServiceSwitch", args, reply)
}

//获取分店管理页面
func (b *Bus) GetBranchPageManagement(ctx context.Context, args *bus.ArgsBranchPageMgt, reply *bus.ReplyBranchPageMgt) error {
	return b.Call(ctx, "GetBranchPageManagement", args, reply)
}

//分店/店铺详情审核
func (b *Bus) GetBranchExamine(ctx context.Context, args *bus.ArgsBranchExamine, reply *bus.ReplyBranchExamine) (err error) {
	return b.Call(ctx, "GetBranchExamine", args, reply)
}

//更新分店/店铺详情
func (b *Bus) UpadatebranchPage(ctx context.Context, args *bus.ArgsUpdatPage, reply *bus.ReplyUpdatPage) (err error) {
	return b.Call(ctx, "UpadatebranchPage", args, reply)
}

//根据企业编号获取企业/商户主体详情
func (b *Bus) GetByMerchantId(ctx context.Context, merchantId *string, reply *bus.ReplySingleBus) error {
	return b.Call(ctx, "GetByMerchantId", merchantId, reply)
}

//更改商家安全码颜色
func (b *Bus) UpdateRiskBusSafeCode(ctx context.Context, args *bus.ArgsUpdateRiskBusSafeCode, reply *bus.ReplyUpdateRiskBusSafeCode) error {
	return b.Call(ctx, "UpdateRiskBusSafeCode", args, reply)
}

//监管平台直连接口-商户主体用户评论
func (b *Bus) GetGovBusComment(ctx context.Context, args *bus.ArgsGetGovBusComment, reply *bus.ReplyGetGovBusComment) error {
	return b.Call(ctx, "GetGovBusComment", args, reply)
}
