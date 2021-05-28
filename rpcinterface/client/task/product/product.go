package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/product"
)

type Product struct {
	client.Baseclient
}

func (p *Product) Init() *Product {
	p.ServiceName = "rpc_task"
	p.ServicePath = "product/Product"
	return p
}

func (p *Product) AddProduct(ctx context.Context,args *product.ArgsAddProduct,reply *bool) error  {
	return p.Call(ctx,"AddProduct",args,reply)
}

//异步删除门店商品
func (p *Product) DelProduct(ctx context.Context,args []int,reply *bool) error {
	return p.Call(ctx,"DelProduct",args,reply)
}

//异步修改商品预警
func (p *Product) UpdateWarnDetails(ctx context.Context,args int, reply *bool) error {
	return p.Call(ctx,"UpdateWarnDetails",args,reply)
}

//异步修改总库存和预警
func (p *Product) UpdateProductStock (ctx context.Context,args *product.ArgsMqUpdateStock,reply *bool) error {
	return p.Call(ctx,"UpdateProductStock",args,reply)
}

//商品上下架统计
func (p *Product) ProductUpDownStatistics(ctx context.Context, shopId *int, reply *bool) error{
	return p.Call(ctx,"ProductUpDownStatistics",shopId,reply)
}

//商品新增删除统计
func (p *Product)ProductAddOrDelStatistics(ctx context.Context,args []int,reply *bool) error{
	return p.Call(ctx,"ProductAddOrDelStatistics",args,reply)
}
