package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/product"
)

type Product struct {
	client.Baseclient
}

func (p *Product) Init() *Product {
	p.ServiceName = "rpc_product"
	p.ServicePath = "Product"
	return p
}

//添加商品分类
func (p *Product) AddCategory(ctx context.Context, args *product.ArgsCategoryAdd, categoryId *int) error {
	return p.Call(ctx, "AddCategory", args, categoryId)
}

//更改商品类名
func (p *Product) UpdateCategory(ctx context.Context, args *product.ArgsCategoryUpdate, reply *bool) error {
	return p.Call(ctx, "UpdateCategory", args, reply)
}

//删除 一条 商品分类
func (p *Product) DelCategory(ctx context.Context, args *product.ArgsCategoryDel, reply *bool) error {
	return p.Call(ctx, "DelCategory", args, reply)
}

//查询 对应商户 的商品分类
func (p *Product) GetCategories(ctx context.Context, args *product.ArgsCategoryGet, reply *product.ReplyCategoryPage) error {
	return p.Call(ctx, "GetCategories", args, reply)
}

//添加商品标签
func (p *Product) AddTag(ctx context.Context, args *product.ArgsTagAdd, tagId *int) error {
	return p.Call(ctx, "AddTag", args, tagId)
}

//更改商品标签名
func (p *Product) UpdateTag(ctx context.Context, args *product.ArgsTagUpdate, reply *bool) error {
	return p.Call(ctx, "UpdateTag", args, reply)
}

//删除 一条 商品标签
func (p *Product) DelTag(ctx context.Context, args *product.ArgsTagDel, reply *bool) error {
	return p.Call(ctx, "DelTag", args, reply)
}

//查询 对应商户 的商品标签
func (p *Product) GetTags(ctx context.Context, args *product.ArgsTagGet, reply *product.ReplyTagPage) error {
	return p.Call(ctx, "GetTags", args, reply)
}


//添加商品
func (p *Product) AddProduct(ctx context.Context, args *product.ArgsProductAdd, detailId *int) error {
	return p.Call(ctx, "AddProduct", args, detailId)
}

//修改商品
func (p *Product) UpdateProduct(ctx context.Context, args *product.ArgsProductAdd, reply *bool) error {
	return p.Call(ctx, "UpdateProduct", args, reply)
}

//异步修改预警
func (p *Product) MqUpdateWarn(ctx context.Context, args int,reply *bool) error {
	return p.Call(ctx, "MqUpdateWarn", args, reply)
}

//根据ids 查询商品信息
func (p *Product) GetProductByIds(ctx context.Context, args *product.ArgsProductGetByIds, reply *[]product.ReplyProductGetByIds) error {
	return p.Call(ctx, "GetProductByIds", args, reply)
}

//删除一条 商品
func (p *Product) DelProduct(ctx context.Context, args *product.ArgsProductDel, reply *bool) error {
	return p.Call(ctx, "DelProduct", args, reply)
}

//异步删除商品
func (p *Product) MqDelProduct(ctx context.Context,args []int,reply *bool) error {
	return p.Call(ctx, "MqDelProduct", args, reply)
}

//批量删除  商品
func (p *Product) DelMoreProduct(ctx context.Context, args *product.ArgsProductDelMore, reply *bool) error {
	return p.Call(ctx, "DelMoreProduct", args, reply)
}

//根据 商品分类 和 商品标签 查询商品
func (p *Product) GetProducts(ctx context.Context, args *product.ArgsProductGet, reply *product.ReplyProductPage) error {
	return p.Call(ctx, "GetProducts", args, reply)
}

//查询一条
func (p *Product) GetProductOne(ctx context.Context, args *product.ArgsProductOneGet, reply *product.ReplyProductOne) error {
	return p.Call(ctx, "GetProductOne", args, reply)
}

//根据 商品分类 和 商品名称 查询商品详情
func (p *Product) GetDetail(ctx context.Context, args *product.ArgsStockGet, reply *product.ReplyDetailPage) error {
	return p.Call(ctx, "GetDetail", args, reply)
}

//添加商品规格
func (p *Product) AddSpec(ctx context.Context, args *product.ArgsSpecAdd, reply *int) error {
	return p.Call(ctx, "AddSpec", args, reply)
}

//根据busId查询所属
func (p *Product) GetSpecs(ctx context.Context, args *int, reply *[]product.ReplySpec) error {
	return p.Call(ctx, "GetSpecs", args, reply)
}

//根据一级规格id查询所有二级规格 //传0 获取所有一级规格
func (p *Product) GetSpecsById(ctx context.Context, args *product.ArgsSpecGet, reply *[]product.ReplySpec) error {
	return p.Call(ctx, "GetSpecsById", args, reply)
}

//根据busId查询所有标签id和标签名字
func (p *Product) GetTagsByBusId(ctx context.Context, args *int, reply *[]product.ReplyTag) error {
	return p.Call(ctx, "GetTagsByBusId", args, reply)
}

//根据detailIds查询售价-RPC调用
func (p *Product) GetByDetailIds(ctx context.Context, detailIds *[]int, reply *[]product.ReplyShopSell) error {
	return p.Call(ctx, "GetByDetailIds", detailIds, reply)
}

//根据DetailId获取商品规格-rpc
func(p *Product)GetSpecsByDeatailIds(ctx context.Context,detailIds *[]int,reply *product.ReplyGetSpecs)error{
	return p.Call(ctx, "GetSpecsByDeatailIds", detailIds, reply)
}

//商品明细-rpc
func(p *Product)GetDetailByIds(ctx context.Context,detailIds *[]int, reply *[]product.ReplyGetDetailById)error{
	return p.Call(ctx, "GetDetailByIds", detailIds, reply)
}
////根据多个商品明细id 查询信息-RPC
//func(p *Product)GetDetailInfosByDetailIds(ctx context.Context,detailIds *[]int,reply *map[int]map[string]interface{})error {
//	return p.Call(ctx, "GetDetailInfosByDetailIds", detailIds, reply)
//}

//根据多个商品id查询照片
func(p *Product) GetImageByProductIds(ctx context.Context,productIds *[]int, reply *[]product.ReplyImage)error {
	return p.Call(ctx, "GetImageByProductIds", productIds, reply)
}

//根据商品id获取子规格商品
func(p *Product) GetDetailsByProductId(ctx context.Context, args *cards.ArgsSubServer, reply *product.ReplySubServer) error {
	return p.Call(ctx, "GetDetailsByProductId", args, reply)
}

//购买成功，设置产品的销量和库存
func(p *Product)  ChangeProductSalesAndStack(ctx context.Context, orderSn *string, reply *bool) error{
	return p.Call(ctx, "ChangeProductSalesAndStack", orderSn, reply)
}

//异步修改总库存和预警
func (p *Product) MqUpdateStock (ctx context.Context,args *product.ArgsMqUpdateStock,reply *bool) error {
	return p.Call(ctx, "MqUpdateStock", args, reply)
}
// 获取添加商品总数
func (p *Product)GetProductNum(ctx context.Context, args *product.ArgsProductNum, reply *product.ReplyProductNum) error {
	return p.Call(ctx, "GetProductNum", args, reply)
}
//根据busid获取商品列表 --rpc用
func (p *Product)GetProductByBusId(ctx context.Context, args *product.ArgsGetProductByBusId, reply *product.ReplyGetProductByBusId )error{
	return p.Call(ctx, "GetProductByBusId", args, reply)
}
//根据busId或者shopId获取商品id --rpc
func (p *Product)GetProductIds(ctx context.Context, args *product.ArgsGetProductIds, reply *product.ReplyGetProductIds)error{
	return p.Call(ctx, "GetProductIds", args, reply)
}

//判断分店是否添加总店所有商品
func (p *Product)IsShopProductEqBus(ctx context.Context,args *product.ArgsIsShopProductEqBus,reply *bool)error{
	return p.Call(ctx, "IsShopProductEqBus", args, reply)
}
//判断分店下是否含有指定商品
func (p *Product)IsShopIncProducts(ctx context.Context,args *product.ArgsIsShopIncProducts,reply *bool)error{
	return p.Call(ctx, "IsShopIncProducts", args, reply)
}