package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/product"
)

type ShopProduct struct {
	client.Baseclient
}

func (s *ShopProduct) Init() *ShopProduct {
	s.ServiceName = "rpc_product"
	s.ServicePath = "ShopProduct"
	return s
}

//总部商品添加到门店
func (s *ShopProduct) AddProduct(ctx context.Context, args *product.ArgsShopProductAdd, id *int) error {
	return s.Call(ctx, "AddProduct", args, id)
}


//异步回调 添加商品
func (s *ShopProduct) MqAddProduct(ctx context.Context,args *product.ArgsAddProduct,reply *bool) error {
	return s.Call(ctx, "MqAddProduct", args, reply)
}

//批量添加
func (s *ShopProduct) AddMoreProduct(ctx context.Context, args *product.ArgsShopProductMoreAdd, id *int) error {
	return s.Call(ctx, "AddMoreProduct", args, id)
}

//商品上下架
func (s *ShopProduct) UPDownProduct(ctx context.Context, args *product.ArgsShopProductUpDown, reply *bool) error {
	return s.Call(ctx, "UPDownProduct", args, reply)
}

//商品批量上下架
func (s *ShopProduct) MoreUpDownProduct(ctx context.Context, args *product.ArgsShopProductUpDowns, reply *bool) error {
	return s.Call(ctx, "MoreUpDownProduct", args, reply)
}

//查询门店商品
func (s *ShopProduct) GetShopProduct(ctx context.Context, args *product.ArgsShopProductGet, reply *product.ReplyShopProductPage) error {
	return s.Call(ctx, "GetShopProduct", args, reply)
}

// 出入库添加商品 查询 根据分类和名称
func (s *ShopProduct) GetShopDetails(ctx context.Context, args *product.ArgsStockGet, reply *product.ReplyDetailPage) error {
	return s.Call(ctx, "GetShopDetails", args, reply)
}

//门店查询出入库明细
func (s *ShopProduct) GetShopInOutStock(ctx context.Context, args *product.ArgsInOutStockGet, reply *product.ReplyInOutStockPage) error {
	return s.Call(ctx, "GetShopInOutStock", args, reply)
}

/*//统一设置库存预警数量
func (s *ShopProduct) SetShopWarnNum(ctx context.Context, args *product.ArgsWarnNumSet, reply *bool) error {
	return s.Call(ctx, "SetShopWarnNum", args, reply)
}

//关闭或开启 明细商品 预警
func (s *ShopProduct) SetShopWarnDetail(ctx context.Context, args *product.ArgsDetailSet, reply *bool) error {
	return s.Call(ctx, "SetShopWarnDetail", args, reply)
}

//自定义预警值
func (s *ShopProduct) SetShopCustom(ctx context.Context, args *product.ArgsCustom, reply *bool) error {
	return s.Call(ctx, "SetShopCustom", args, reply)
}

//查询所有预警 商品明细
func (s *ShopProduct) GetShopWarn(ctx context.Context, args *product.ArgsWarnDetail, reply *product.ReplyWarnDetailPage) error {
	return s.Call(ctx, "GetShopWarn", args, reply)
}*/

//门店要货申请
func (s *ShopProduct) AddShopRequire(ctx context.Context, args *product.ArgsShopRequireAdd, reply *int) error {
	return s.Call(ctx, "AddShopRequire", args, reply)
}

/*//门店一条要货申请详情查询
func (s *ShopProduct) GetRequireDetail(ctx context.Context, args *product.ArgsRequireDetailGet, reply *product.ReplyRequireDetailPage) error {
	return s.Call(ctx, "GetRequireDetail", args, reply)
}*/

//门店要货申请查询
func (s *ShopProduct) GetShopRequire(ctx context.Context, args *product.ArgsShopRequireGet, reply *product.ReplyShopRequirePage) error {
	return s.Call(ctx, "GetShopRequire", args, reply)
}

//门店要货申请修改状态 //取消操作  待审核 修改为 已关闭
func (s *ShopProduct) CancelRequire(ctx context.Context, args *product.ArgsRequireCancel, reply *bool) error {
	return s.Call(ctx, "CancelRequire", args, reply)
}

//门店要货申请修改状态 //入库操作  待入库 修改为 已完成
func (s *ShopProduct) InRequire(ctx context.Context, args *product.ArgsRequireUpdate, reply *bool) error {
	return s.Call(ctx, "InRequire", args, reply)
}


// 附近-门店详情-本店商品
func (s *ShopProduct)ShopInfoProductList(ctx context.Context, args *product.ArgsShopInfoProductList, reply *product.ReplyShopInfoProductList) error{
	return s.Call(ctx, "ShopInfoProductList", args, reply)
}

//门店修改商品价格
func (s *ShopProduct) UpdateShopPrice(ctx context.Context, args *product.ShopPriceUpdate, reply *bool) error {
	return s.Call(ctx, "UpdateShopPrice", args, reply)
}

//根据门店id和detailIds 查询售价-RPC调用
func (s *ShopProduct) GetByShopIdAndDetailIds(ctx context.Context, args *product.ArgsShopSellGet, reply *[]product.ReplyShopSell) error {
	return s.Call(ctx, "GetByShopIdAndDetailIds", args, reply)
}

//根据门店id和商品ID查询是否存在-RPC调用
func (s *ShopProduct)GetProductByIds(ctx context.Context,args *product.ArgsGetProductByIds,reply *[]product.ReplyGetProductByIds)error{
	return s.Call(ctx,"GetProductByIds",args,reply)
}

//根据多个商品id查询 价格区间
func (s *ShopProduct)GetShopSpecPrice(ctx context.Context,args *product.ArgsGetImages,reply *[]product.ReplySpecPrice)error {
	return s.Call(ctx,"GetShopSpecPrice",args,reply)
}

//根据门店id和detailIds 获取门店产品信息
func (s *ShopProduct) GetShopProdcuts(ctx context.Context, args *product.ArgsGetShopProdcuts, reply *[]product.ReplyGetShopProdcuts) error  {
	return s.Call(ctx,"GetShopProdcuts",args,reply)
}

//获取企业当月商品下架率
func (s *ShopProduct)GetBusProductUnderRate(ctx context.Context,args *product.ArgsGetBusProductUnderRate,reply *product.ReplyGetBusProductUnderRate)error{
	return s.Call(ctx,"GetBusProductUnderRate",args,reply)
}