package product

import (
	"context"
)

type ArgsProductEs struct {
	Id int
	CateId int
	Name string
	TagIds string
	BusId int
	ShopId int
	Details []Detail
}
type Detail struct {
	BarCode string
	SpecIds string
	DetailId int
	Cost float64
}

type ArgsShopAddProduct struct {
	ProductIds []int
	ShopId int
}

type Product interface {

	//总店添加 商品
	AddProduct(ctx context.Context, args *ArgsProductEs, reply *bool) error

	//门店批量添加 商品
	AddShopProducts(ctx context.Context, args *ArgsShopAddProduct, reply *bool) error

	//删除商品
	DelProduct(ctx context.Context, args *int, reply *bool) error
	/*
	//总店修改 商品
	UpdateProduct(ctx context.Context, args *product.ArgsProductEs, reply *bool) error

	//检索商品
	SearchProduct(ctx context.Context, args *product.ArgsProductEsGet, reply *product.ReplyProductEs) error

	//添加 商品详情
	AddDetail(ctx context.Context, args *product.ArgsDetailEs, reply *bool) error

	//总店添加商品
	AddProductBatch(ctx context.Context, args *product.ArgsAddProductEs, reply *bool) error

	//门店添加商品 或者批量添加
	AddShopProductBatchs(ctx context.Context, args *product.ArgsShopAddProductEs, reply *bool) error

	//总店修改 商品
	UpdateProductBatch(ctx context.Context, args *product.ArgsAddProductEs, reply *bool) error

	//删除商品 或者 批量删除商品
	DelProductBatch(ctx context.Context, args *product.ArgsDelProductEs, reply *bool) error

	//检索规格明细商品
	SearchProductDetail(ctx context.Context, args *product.ArgsProductDetailEsGet, reply *product.ReplyProductEs) error*/
}
