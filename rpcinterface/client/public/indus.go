package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

type Indus struct {
	client.Baseclient
}

//实例化
func (i *Indus) Init() *Indus {
	i.ServiceName = "rpc_public"
	i.ServicePath = "Indus"
	return i
}

//获取全部领域、行业
func (i *Indus) GetAll(ctx context.Context, reply *[]public.ReplyIndusAll) error {
	return i.Call(ctx, "GetAll", "", reply)
}

//获取全部领域
func (i *Indus) GetIndustry(ctx context.Context, reply *[]public.ReplyIndustryInfo) error {
	return i.Call(ctx, "GetIndustry", "", reply)
}

//获取领域下的行业
func (i *Indus) GetIndusByIndustryId(ctx context.Context, args *public.ArgsIndusInfo, reply *[]public.ReplyIndusInfo) error {
	return i.Call(ctx, "GetIndusByIndustryId", args, reply)
}

//批量获取行业信息
func (i *Indus) GetIndusByBindIds(ctx context.Context, bindIds *[]int, reply *[]public.ReplyIndusInfo) error {
	return i.Call(ctx, "GetIndusByBindIds", bindIds, reply)
}

//添加领域
func (i *Indus) AddIndustry(ctx context.Context, args *public.ArgsIndustryAdd, reply *int) error {
	return i.Call(ctx, "AddIndustry", args, reply)
}

//领域改名
func (i *Indus) UpdateIndustry(ctx context.Context, args *public.ArgsIndustryUpdate, reply *bool) error {
	return i.Call(ctx, "UpdateIndustry", args, reply)
}

//删除领域
func (i *Indus) DelIndustry(ctx context.Context, args *public.ArgsIndustryDel, reply *bool) error {
	return i.Call(ctx, "DelIndustry", args, reply)
}

//查询所有领域
func (i *Indus) GetIndustrys(ctx context.Context, args *public.ArgsIndustryGet, reply *[]public.Industry) error {
	return i.Call(ctx, "GetIndustrys", args, reply)
}

//根据领域IDs获取领域
func (i *Indus) GetIndustrysByIds(ctx context.Context, args *public.ArgsGetIndustrysByIds, reply *[]public.Industry) error {
	return i.Call(ctx, "GetIndustrysByIds", args, reply)
}

//添加行业
func (i *Indus) AddIndus(ctx context.Context, args *public.ArgsIndusAdd, reply *int) error {
	return i.Call(ctx, "AddIndus", args, reply)
}

//批量添加行业
func (i *Indus) AddMoreIndus(ctx context.Context, args *public.ArgsIndusAdds, reply *int) error {
	return i.Call(ctx, "AddMoreIndus", args, reply)
}

//修改行业信息
func (i *Indus) UpdateIndus(ctx context.Context, args *public.ArgsIndusAdd, reply *bool) error {
	return i.Call(ctx, "UpdateIndus", args, reply)
}

//删除一条行业
func (i *Indus) DelIndus(ctx context.Context, args *public.ArgsIndusDel, reply *bool) error {
	return i.Call(ctx, "DelIndus", args, reply)
}

//批量删除行业
func (i *Indus) DelMoreIndus(ctx context.Context, args *public.ArgsIndusDels, reply *bool) error {
	return i.Call(ctx, "DelMoreIndus", args, reply)
}

//根据领域Id查询行业
func (i *Indus) GetInduss(ctx context.Context, args *public.ArgsIndusGet, reply *[]public.IndusRes) error {
	return i.Call(ctx, "GetInduss", args, reply)
}

//获取全部领域和行业
func (i *Indus) GetAllInfos(ctx context.Context, args *public.Args, reply *public.Reply) error {
	return i.Call(ctx, "GetAllInfos", args, reply)
}

//根据行业id查询一条行业详情
func (i *Indus) GetIndusDetail(ctx context.Context, args *int, reply *public.IndusDetail) error {
	return i.Call(ctx, "GetIndusDetail", args, reply)
}
