package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// 定义提成设置相关rpc客户端调用方法
// @author liyang<654516092@qq.com>
// @date  2020/4/14 16:23

type Commission struct {
	client.Baseclient
}

//初始化
func (c *Commission)Init() *Commission {
	c.ServiceName = "rpc_staff"
	c.ServicePath = "Commission"
	return c
}

//新增项目提成
func (c *Commission) AddProject(ctx context.Context, args *staff.ArgsAddProject, reply *staff.ReplyProject) error {
	return c.Call(ctx, "AddProject", args, reply)
}

//更新项目提成
func (c *Commission) SetProject(ctx context.Context, args *staff.ArgsSetProject, reply *staff.ReplyProject) error {
	return c.Call(ctx, "SetProject", args, reply)
}

//删除项目提成
func (c *Commission) DeleteProject(ctx context.Context, args *staff.ArgsDelProject, reply *staff.ReplyProject) error {
	return c.Call(ctx, "DeleteProject", args, reply)
}

//获取项目提成列表
func (c *Commission) GetProjectList(ctx context.Context,args *staff.ArgsCommList,reply *staff.ReplyCommList) error{
	return c.Call(ctx, "GetProjectList", args, reply)
}

//新增销售提成
func (c *Commission) AddSales(ctx context.Context, args *staff.ArgsAddSales, reply *staff.ReplySales) error {
	return c.Call(ctx, "AddSales", args, reply)
}

//更新销售提成
func (c *Commission) SetSales(ctx context.Context, args *staff.ArgsSetSales, reply *staff.ReplySales) error {
	return c.Call(ctx, "SetSales", args, reply)
}

//删除销售提成
func (c *Commission) DeleteSales(ctx context.Context, args *staff.ArgsDelSales, reply *staff.ReplySales) error {
	return c.Call(ctx, "DeleteSales", args, reply)
}

//获取销售提成列表
func (c *Commission) GetSalesList(ctx context.Context,args *staff.ArgsCommList,reply *staff.ReplyCommList) error{
	return c.Call(ctx, "GetSalesList", args, reply)
}

//项目提成-根据岗位ID和项目ID获取对应的提成
func (c *Commission) GetProjectByBusId(ctx context.Context,args *staff.ArgsCommSingle,reply *staff.ReplyCommSingle) error{
	return c.Call(ctx, "GetProjectByBusId", args, reply)
}
//销售提成-根据岗位ID和项目ID获取对应的提成
func (c *Commission) GetSalesByBusId(ctx context.Context,args *staff.ArgsCommSingle,reply *staff.ReplyCommSingle) error{
	return c.Call(ctx, "GetSalesByBusId", args, reply)
}