package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/elastic/product"
)

type Product struct {
	client.Baseclient
}

func (p *Product) Init() *Product {
	p.ServiceName = "rpc_elastic"
	p.ServicePath = "Product/Product"
	return p
}

//总店添加 商品
func (p *Product) AddProduct(ctx context.Context, args *product.ArgsProductEs, reply *bool) error {
	return p.Call(ctx, "AddProduct", args, reply)
}

//门店批量添加 商品
func (p *Product) AddShopProducts(ctx context.Context, args *product.ArgsShopAddProduct, reply *bool) error {
	return p.Call(ctx, "AddShopProducts", args, reply)
}

//删除商品
func (p *Product) DelProduct(ctx context.Context, args *int, reply *bool) error {
	return p.Call(ctx, "DelProduct", args, reply)
}
/*
//总店修改 商品
func (p *Product) UpdateProduct(ctx context.Context, args *product.ArgsProductEs, reply *bool) error {
	return p.Call(ctx, "UpdateProduct", args, reply)
}

//检索商品
func (p *Product) SearchProduct(ctx context.Context, args *product.ArgsProductEsGet, reply *product.ReplyProductEs) error {
	return p.Call(ctx, "SearchProduct", args, reply)
}

//添加 商品详情
func (p *Product) AddDetail(ctx context.Context, args *product.ArgsDetailEs, reply *bool) error {
	return p.Call(ctx, "AddDetail", args, reply)
}

//总店添加商品
func (p *Product) AddProductBatch(ctx context.Context, args *product.ArgsAddProductEs, reply *bool) error {
	return p.Call(ctx, "AddProductBatch", args, reply)
}

//门店添加商品 或者批量添加
func (p *Product) AddShopProductBatchs(ctx context.Context, args *product.ArgsShopAddProductEs, reply *bool) error {
	return p.Call(ctx, "AddShopProductBatchs", args, reply)
}

//总店修改 商品
func (p *Product) UpdateProductBatch(ctx context.Context, args *product.ArgsAddProductEs, reply *bool) error {
	return p.Call(ctx, "UpdateProductBatch", args, reply)
}

//删除商品 或者 批量删除商品
func (p *Product) DelProductBatch(ctx context.Context, args *product.ArgsDelProductEs, reply *bool) error {
	return p.Call(ctx, "DelProductBatch", args, reply)
}

//检索规格明细商品
// flag==true 如果有值表示 是预警在调这个接口//预警调接口 取消分页 查询所有
func (p *Product) SearchProductDetail(ctx context.Context, args *product.ArgsProductDetailEsGet, reply *product.ReplyProductEs) error {
	return p.Call(ctx, "SearchProductDetail", args, reply)
}
*/