package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/product"
)

type Stock struct {
	client.Baseclient
}

func (s *Stock) Init() *Stock {
	s.ServiceName = "rpc_product"
	s.ServicePath = "Stock"
	return s
}

//库存查询
func (s *Stock) GetStock(ctx context.Context, args *product.ArgsStockGet, reply *product.ReplyDetailPage) error {
	return s.Call(ctx, "GetStock", args, reply)
}

//新增入库
func (s *Stock) AddInStock(ctx context.Context, args *product.ArgsStockInAdd, reply *int) error {
	return s.Call(ctx, "AddInStock", args, reply)
}

//查询所有入库
func (s *Stock) GetInStock(ctx context.Context, args *product.ArgsInStockGet, reply *product.ReplyInStockPage) error {
	return s.Call(ctx, "GetInStock", args, reply)
}

//查询一条出入库详情
func (s *Stock) GetOneInOutStockDetail(ctx context.Context, args *product.ArgsGetOneStock, reply *product.ReplyOneStockPage) error {
	return s.Call(ctx, "GetOneInOutStockDetail", args, reply)
}

//新增出库
func (s *Stock) AddOutStock(ctx context.Context, args *product.ArgsOutStockAdd, reply *int) error {
	return s.Call(ctx, "AddOutStock", args, reply)
}

//查询所有出库
func (s *Stock) GetOutStock(ctx context.Context, args *product.ArgsOutStockGet, reply *product.ReplyOutStockPage) error {
	return s.Call(ctx, "GetOutStock", args, reply)
}

//新建 盘点 单
func (s *Stock) AddStockCheck(ctx context.Context, args *product.ArgsCheckAdd, reply *int) error {
	return s.Call(ctx, "AddStockCheck", args, reply)
}

//查询盘点单
func (s *Stock) GetStockCheck(ctx context.Context, args *product.ArgsCheckGet, reply *product.ReplyCheckPage) error {
	return s.Call(ctx, "GetStockCheck", args, reply)
}

//作废盘点单
func (s *Stock) SetStockCheck(ctx context.Context, args *product.ArgsCheckSet, reply *bool) error {
	return s.Call(ctx, "SetStockCheck", args, reply)
}

//盘点单详情 - 恢复盘点单  共用接口
func (s *Stock) GetStockCheckDetail(ctx context.Context, args *product.ArgsCheckDetail, reply *product.ReplyCheckDetailPage) error {
	return s.Call(ctx, "GetStockCheckDetail", args, reply)
}

//按条件查询出入库明细
func (s *Stock) GetInOutStock(ctx context.Context, args *product.ArgsInOutStockGet, reply *product.ReplyInOutStockPage) error {
	return s.Call(ctx, "GetInOutStock", args, reply)
}

//统一设置库存预警数量
func (s *Stock) SetWarnNum(ctx context.Context, args *product.ArgsWarnNumSet, reply *bool) error {
	return s.Call(ctx, "SetWarnNum", args, reply)
}

//查询所有预警 商品明细
func (s *Stock) GetWarn(ctx context.Context, args *product.ArgsWarnDetail, reply *product.ReplyWarnDetailPage) error {
	return s.Call(ctx, "GetWarn", args, reply)
}

//关闭或开启 明细商品 预警
func (s *Stock) SetWarnDetail(ctx context.Context, args *product.ArgsDetailSet, reply *bool) error {
	return s.Call(ctx, "SetWarnDetail", args, reply)
}

//自定义预警值
func (s *Stock) SetCustom(ctx context.Context, args *product.ArgsCustom, reply *bool) error {
	return s.Call(ctx, "SetCustom", args, reply)
}

//恢复默认预警值
func (s *Stock) RegainDefault(ctx context.Context, args *product.ArgsRegainDefault, reply *bool) error {
	return s.Call(ctx, "RegainDefault", args, reply)
}

//添加供应商
func (s *Stock) AddSub(ctx context.Context, args *product.ArgsSubAdd, reply *int) error {
	return s.Call(ctx, "AddSub", args, reply)
}

//修改供应商
func (s *Stock) UpdateSub(ctx context.Context, args *product.ArgsSubUpdate, reply *bool) error {
	return s.Call(ctx, "UpdateSub", args, reply)
}

//查询一条供应商信息
func (s *Stock) GetSubOne(ctx context.Context, args *product.ArgsSubGetOne, reply *product.SupInfo) error {
	return s.Call(ctx, "GetSubOne", args, reply)
}

//查询供应商
func (s *Stock) GetSub(ctx context.Context, args *product.ArgsSubGet, reply *product.ReplySup) error {
	return s.Call(ctx, "GetSub", args, reply)
}

//采购入库
func (s *Stock) AddPur(ctx context.Context, args *product.ArgsPurAdd, reply *int) error {
	return s.Call(ctx, "AddPur", args, reply)
}

//查询采购信息
func (s *Stock) GetPur(ctx context.Context, args *product.ArgsPurGet, reply *product.ReplyPur) error {
	return s.Call(ctx, "GetPur", args, reply)
}

//查询一条采购详情信息
func (s *Stock) GetPurDetail(ctx context.Context, args *product.ArgsPurDetail, reply *product.ReplyPurDetail) error {
	return s.Call(ctx, "GetPurDetail", args, reply)
}

//总部要货申请修改状态 //审核操作  待审核 修改为 待入库
func (s *Stock) CheckRequire(ctx context.Context, args *product.ArgsCheckRequire, reply *bool) error {
	return s.Call(ctx, "CheckRequire", args, reply)
}

//总部要货申请 审核驳回   待审核 修改为  已驳回
func (s *Stock) SetRequire(ctx context.Context, args *product.ArgsSetRequire, reply *bool) error {
	return s.Call(ctx, "SetRequire", args, reply)
}

//查询供应商名称和id接口
func (s *Stock) GetSubNameAndId(ctx context.Context, args *int, reply *[]product.ReplySubInfo) error {
	return s.Call(ctx, "GetSubNameAndId", args, reply)
}

/*//查询门店要货入库
func (s *Stock) GetRequireInStock(ctx context.Context, args *product.ArgsRequireInGet, reply *product.ReplyRequirePage) error {
	return s.Call(ctx, "GetRequireInStock", args, reply)
}*/

//查询一条要货入库详情
func (s *Stock) GetOneRequireInStock(ctx context.Context, args *product.ArgsRequireOneGet, reply *product.ReplyRequireDetail) error {
	return s.Call(ctx, "GetOneRequireInStock", args, reply)
}

//开启预警或者关闭预警
func (s *Stock) SetWarn(ctx context.Context, args *product.ArgsWarnSet, reply *bool) error {
	return s.Call(ctx, "SetWarn", args, reply)
}

//返回预警状态
func (s *Stock) WarnGet(ctx context.Context, args *product.ArgsWarnGet, reply *product.ReplyWarnGet) error {
	return s.Call(ctx, "WarnGet", args, reply)
}

//根据供应商Id 查询供应商名
func (s *Stock) GetSupName(ctx context.Context, args *[]int, reply *map[int]string) error {
	return s.Call(ctx, "GetSupName", args, reply)
}

//获取要货提货待处理数量和预警待处理数量
func (s *Stock) GetPending(ctx context.Context,args *product.ArgsGetPending,reply *product.ReplyGetPending) error {
	return s.Call(ctx, "GetPending", args, reply)
}