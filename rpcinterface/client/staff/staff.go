package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

// 定义rpc客户端调用方法
// @author liyang<654516092@qq.com>
// @date  2020/4/1 18:33

type Staff struct {
	client.Baseclient
}

//初始化
func (s *Staff)Init() *Staff {
	s.ServiceName = "rpc_staff"
	s.ServicePath = "Staff"
	return s
}

//添加员工
func (s *Staff) AddStaff(ctx context.Context, args *staff.ArgsAddStaff, reply *staff.ReplyStaff) error {
	return s.Call(ctx, "AddStaff", args, reply)
}
//修改员工
func (s *Staff) EditStaff(ctx context.Context, args *staff.ArgsEditStaff, reply *staff.ReplyStaff) error {
	return s.Call(ctx, "EditStaff", args, reply)
}
//删除员工
func (s *Staff) DeleteStaff(ctx context.Context, args *staff.ArgsDelStaff, reply *staff.ReplyStaff) error {
	return s.Call(ctx, "DeleteStaff", args, reply)
}

//员工详情-企业/商户
func (s *Staff) GetStaffInfoByBus(ctx context.Context,args *staff.ArgsStaffInfo,reply *staff.ReplyStaffInfo) error{
	return s.Call(ctx,"GetStaffInfoByBus",args,reply)
}

//员工详情-分店
func (s *Staff) GetStaffInfoByShop(ctx context.Context,args *staff.ArgsStaffInfo,reply *staff.ReplyStaffInfo) error{
	return s.Call(ctx,"GetStaffInfoByShop",args,reply)
}

//员工详情-rpc
func (s *Staff) GetStaffDetail(ctx context.Context,args *staff.ArgsStaffInfo,reply *staff.ReplyStaffInfo) error{
	return s.Call(ctx,"GetStaffDetail",args,reply)
}

//员工列表-企业/商户
func (s *Staff) GetStaffByBusId(ctx context.Context,args *staff.ArgsGetStaffList,reply *staff.ReplyStaffList) error{
	return s.Call(ctx,"GetStaffByBusId",args,reply)
}


//员工列表-分店
func (s *Staff) GetStaffByShopId(ctx context.Context,args *staff.ArgsGetStaffList,reply *staff.ReplyStaffList) error{
	return s.Call(ctx,"GetStaffByShopId",args,reply)
}

//员工列表-前端
func (s *Staff) GetStaffForShopId(ctx context.Context,args *staff.ArgsStaffList,reply *staff.ReplyStaffList) error{
	return s.Call(ctx,"GetStaffForShopId",args,reply)
}

//批量获取员工评分
func (s *Staff) GetRankByStaffIds(ctx context.Context,staffIds *[]int,reply *map[string]staff.RankInfo) error{
	return s.Call(ctx,"GetRankByStaffIds",staffIds,reply)
}


//添加员工岗位
func (s *Staff) AddPost(ctx context.Context,args *staff.ArgsAddPost,reply *staff.ReplyPost) error{
	return s.Call(ctx,"AddPost",args,reply)
}
//修改员工岗位
func (s *Staff) EditPost(ctx context.Context,args *staff.ArgsEditPost,reply *staff.ReplyPost) error{
	return s.Call(ctx,"EditPost",args,reply)
}
//设置员工岗位职称
func (s *Staff) SetJobTitle(ctx context.Context,args *staff.ArgsSetJobTitle,reply *staff.ReplyPost) error{
	return s.Call(ctx,"SetJobTitle",args,reply)
}
//删除员工岗位
func (s *Staff) DeletePost(ctx context.Context,args *staff.ArgsDelPost,reply *staff.ReplyPost) error{
	return s.Call(ctx,"DeletePost",args,reply)
}
//获取职称
func (s *Staff) GetPostList(ctx context.Context,args *staff.ArgsGetPostList,reply *[]staff.ReplyInfo) error{
	return s.Call(ctx,"GetPostList",args,reply)
}
//获取员工列表-rpc内部调用
func (s *Staff)GetListByStaffIds(ctx context.Context,args *staff.ArgsGetListByStaffIds,reply *[]staff.ReplyGetListByStaffIds)error{
	return s.Call(ctx,"GetListByStaffIds",args,reply)
}

//获取员工列表-rpc内部调用
func (s *Staff) GetListByStaffIds2(ctx context.Context, staffIds *[]int, reply *[]staff.ReplyGetListByStaffIds2) error {
	return s.Call(ctx,"GetListByStaffIds2",staffIds,reply)
}

//员工列表-无需验证-分店
func (s *Staff)GetStaffListByShopId(ctx context.Context,args *staff.ArgsGetStaffListByShopId,reply *staff.ReplyGetStaffListByShopId)error{
	return s.Call(ctx,"GetStaffListByShopId",args,reply)
}
//员工详情-无需验证-分店
func (s *Staff)GetStaffInfoById(ctx context.Context,args *int,reply *staff.ReplyGetStaffInfoById)error{
	return s.Call(ctx,"GetStaffInfoById",args,reply)
}
//根据多个员工id 查询员工名称和对应岗位信息 rpc 调用
func (s *Staff) GetStaffNameAndPostByIds(ctx context.Context,args *[]int,reply *[]staff.ReplyStaffNameAndPost)error {
	return s.Call(ctx,"GetStaffNameAndPostByIds",args,reply)
}

//获取当月员工新增或离职率-rpc
func (s *Staff) GetStaffAddMinusByBusIdRpc(ctx context.Context,args *staff.ArgsGetStaffAddMinus,reply *staff.ReplyGetStaffAddMinus)error{
	return s.Call(ctx,"GetStaffAddMinusByBusIdRpc",args,reply)
}

//根据id获取员工状态（新增，离职，删除）
func (s *Staff) GetStaffStatusById(ctx context.Context,args *staff.ArgsGetStaffStatusById,reply *staff.ReplyGetStaffStatusById)error{
	return s.Call(ctx,"GetStaffStatusById",args,reply)
}
// 根据busId 获取在在职员工总数量
func (s *Staff) GetStaffByBusIdNum(ctx context.Context, args *staff.ArgsGetStaffByBusIdNum, reply *staff.ReplyGetStaffByBusIdNum) error {
	return s.Call(ctx, "GetStaffByBusIdNum",args, reply)
}