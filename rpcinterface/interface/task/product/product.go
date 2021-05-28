package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/product"
)

type Produuct interface {
	AddProduct(ctx context.Context, args *product.ArgsAddProduct, reply *bool) error

	//异步删除门店商品
	DelProduct(ctx context.Context, args []int, reply *bool) error

	//异步修改商品预警
	UpdateWarnDetails(ctx context.Context, args int, reply *bool) error

	//异步修改总库存和预警
	UpdateProductStock(ctx context.Context, args *product.ArgsMqUpdateStock, reply *bool) error

	//商品上下架统计
	ProductUpDownStatistics(ctx context.Context, shopId *int, reply *bool) error

	//商品新增删除统计
	ProductAddOrDelStatistics(ctx context.Context,args []int,reply *bool) error
}
