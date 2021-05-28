package caster

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/caster"
)

type Role struct {
	client.Baseclient
}

func (r *Role) Init() *Role {
	r.ServiceName = "rpc_admin"
	r.ServicePath = "Role"
	return r
}



func (r *Role)AddRole(ctx context.Context,  args *caster.AddRoleReq, reply *caster.RoleID)(err error){
	err = r.Call(ctx,"AddRole", args, reply)
	return
}
func (r *Role)DelRole(ctx context.Context, args *caster.RoleID, reply *caster.Empty)(err error){
	err = r.Call(ctx,"DelRole", args, reply)
	return
}
func (r *Role)UpdateRole(ctx context.Context, args *caster.RoleStruct, reply *caster.Empty)(err error){
	err = r.Call(ctx,"UpdateRole", args, reply)
	return
}
func (r *Role)RoleInfo(ctx context.Context, args *caster.RoleID, reply *caster.RoleStruct)(err error){
	err = r.Call(ctx,"RoleInfo", args, reply)
	return
}
func (r *Role)RoleList(ctx context.Context, args *caster.Page, reply *caster.RoleList)(err error){
	err = r.Call(ctx,"RoleList", args, reply)
	return
}

