package caster

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/caster"
)

type Caster struct {
	client.Baseclient
}
func (c *Caster) Init() *Caster {
	c.ServiceName = "rpc_admin"
	c.ServicePath = "Caster"
	return c
}


func (c *Caster)Verify(ctx context.Context, args *caster.VerifyReq, reply *caster.VerifyResp)(err error){
	err = c.Call(ctx, "Verify", args, reply)
	return
}
func (c *Caster)AddNode(ctx context.Context, args *caster.AddNodeReq, reply *caster.AddNodeResp)(err error){
	err = c.Call(ctx, "AddNode", args, reply)
	return
}
func (c *Caster)DelNode(ctx context.Context, args *caster.NodeID, reply *caster.Empty)(err error){
	err = c.Call(ctx, "DelNode", args, reply)
	return
}
func (c *Caster)EnableNode(ctx context.Context,  args *caster.NodeID, resp *caster.Empty) (err  error){
	err = c.Call(ctx, "EnableNode", args, resp)
	return
}
func (c *Caster)DisableNode(ctx context.Context,  args *caster.NodeID, resp *caster.Empty) (err  error){
	err = c.Call(ctx, "DisableNode", args, resp)
	return
}
func (c *Caster)UpdateNode(ctx context.Context, args *caster.UpdateReq, resp *caster.Empty)(err error){
	err = c.Call(ctx, "UpdateNode", args, resp)
	return
}
func (c *Caster)SubList(ctx context.Context, args *caster.SubListReq, reply *caster.SubListResp )(err error){
	err = c.Call(ctx, "SubList", args, reply)
	return
}
func (c *Caster)SubEnableList(ctx context.Context, args *caster.SubListReq, reply *caster.SubListResp )(err error){
	err = c.Call(ctx, "SubEnableList", args, reply)
	return
}
