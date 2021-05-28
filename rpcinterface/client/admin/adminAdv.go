package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/admin"
)

type AdminAdv struct {
	client.Baseclient
}

func (adminAdv *AdminAdv) Init() *AdminAdv {
	adminAdv.ServiceName = "rpc_admin"
	adminAdv.ServicePath = "AdminAdv"
	return adminAdv
}

/*//添加广告位
func (a *AdminAdv) AddAdvSpa(ctx context.Context, args *admin.ArgsAdvSpaAdd, reply *int) error {
	return a.Call(ctx, "AddAdvSpa", args, reply)
}*/

//修改广告位
func (a *AdminAdv) UpdateAdvSpa(ctx context.Context, args *admin.ArgsAdvSpaAdd, reply *bool) error {
	return a.Call(ctx, "UpdateAdvSpa", args, reply)
}

//删除广告位
func (a *AdminAdv) DelAdvSpa(ctx context.Context, args *admin.ArgsAdvSpaDel, reply *bool) error {
	return a.Call(ctx, "DelAdvSpa", args, reply)
}

//按分页查询广告位
func (a *AdminAdv) GetAdvSpaByPage(ctx context.Context, args *admin.ArgsAdvSpaGet, reply *admin.ReplyAdvSpaPage) error {
	return a.Call(ctx, "GetAdvSpaByPage", args, reply)
}

//无条件查询广告位
func (a *AdminAdv) GetAdvSpas(ctx context.Context, args *admin.Args, reply *map[string]string) error {
	return a.Call(ctx, "GetAdvSpas", args, reply)
}

//按id查询一条广告位
func (a *AdminAdv) GetAdvSpaOne(ctx context.Context, args *admin.ArgsAdvSpaGetOne, reply *admin.AdvSpaInfo) error {
	return a.Call(ctx, "GetAdvSpaOne", args, reply)
}

//添加 广告
func (a *AdminAdv) AddAdv(ctx context.Context, args *admin.ArgsAdvAdd, reply *int) error {
	return a.Call(ctx, "AddAdv", args, reply)
}

//修改 广告
func (a *AdminAdv) UpdateAdv(ctx context.Context, args *admin.ArgsAdvAdd, reply *bool) error {
	return a.Call(ctx, "UpdateAdv", args, reply)
}

//删除 广告
func (a *AdminAdv) DelAdv(ctx context.Context, args *admin.ArgsAdvDel, reply *bool) error {
	return a.Call(ctx, "DelAdv", args, reply)
}

//查询广告
func (a *AdminAdv) GetAdvs(ctx context.Context, args *admin.ArgsAdvGet, reply *admin.ReplyAdvPage) error {
	return a.Call(ctx, "GetAdvs", args, reply)
}

//查询一条广告
func (a *AdminAdv) GetAdv(ctx context.Context, args *admin.ArgsAdvGetOne, reply *admin.ReplyAdvInfo) error {
	return a.Call(ctx, "GetAdv", args, reply)
}

//查询所有广告位和广告
func (a *AdminAdv) GetAdvBySpaId(ctx context.Context, args *admin.ArgsBySpaId, reply *admin.ReplyBySpaId) error {
	return a.Call(ctx, "GetAdvBySpaId", args, reply)
}
