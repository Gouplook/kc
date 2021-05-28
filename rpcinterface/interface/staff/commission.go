package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// 定义提成相关接口
// @author liyang<654516092@qq.com>
// @date  2020/4/14 16:21

//新增项目提成入参
type ArgsAddProject struct {
	 common.Utoken //用户信息
	 common.BsToken //企业/商户信息
	 PostId int  //岗位ID
	 SingleId int //单项目ID
	 CommissionType  int //提成类型 1=比列 2=固定
	 CommValue float64  //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}
//编辑/更新项目提成入参
type ArgsSetProject struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	Id       int   //项目提成ID
	CommissionType  int //提成类型 1=比列 2=固定
	CommValue float64  //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}

//删除项目提成入参
type ArgsDelProject struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	Id       int   //项目提成ID
}

//新增/编辑/删除项目提成返回信息
type ReplyProject struct {
	Id int //项目提成ID
}

//新增销售提成入参
type ArgsAddSales struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	PostId int  //岗位ID
	SingleId int //单项目ID
	CommissionType  int //提成类型 1=比列 2=固定
	CommValue float64  //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}
//编辑/更新销售提成入参
type ArgsSetSales struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	Id       int   //项目提成ID
	CommissionType  int //提成类型 1=比列 2=固定
	CommValue float64  //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}
//删除销售提成入参
type ArgsDelSales struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	Id       int   //项目提成ID
}

//新增/编辑/删除销售提成返回信息
type ReplySales struct {
	Id int //销售提成ID
}

//项目/销售提成设置列表入参
type ArgsCommList struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/分店ID
	common.Paging  //分页信息
}
//项目/销售提成设置列表返回信息
type ReplyCommList struct {
	 Data []ReplyCommInfo  //列表信息
	 TotalNum int //总数量
}

//返回单个项目/销售提成信息
type ReplyCommInfo struct {
	Id int //提成ID
	PostName string //岗位名称
	SingleName string //单项目名称
	CommissionType int  `mapstructure:"type"`//提成类型 1=比例 2=固定金额
	CommValue float64 //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}
//岗位对应的项目信息入参
type ArgsCommSingle struct {
	BusId int //企业/商户ID
	PostId int //岗位ID
	SingleId int //单项目ID
}
//岗位对应的项目返回信息
type ReplyCommSingle struct {
	CommissionType int //提成类型 1=比例 2=固定金额
	CommValue float64 //提成值，CommissionType=1时，提成比例（格式如：0.15则为15%），CommissionType=2时，固定金额（格式如：5则为5元）
}

//定义接口
type Commission interface{
	 //新增项目提成
	 AddProject(ctx context.Context,args *ArgsAddProject,reply *ReplyProject)    error
	 //更新项目提成
	 SetProject(ctx context.Context,args *ArgsSetProject,reply *ReplyProject)    error
	 //删除项目提成
	 DeleteProject(ctx context.Context,args *ArgsDelProject,reply *ReplyProject) error
	 //新增销售提成
	 AddSales(ctx context.Context,args *ArgsAddSales,reply *ReplySales)      error
	 //更新销售提成
	 SetSales(ctx context.Context,args *ArgsSetSales,reply *ReplySales)      error
	 //删除销售提成
	 DeleteSales(ctx context.Context,args *ArgsDelSales,reply *ReplySales)   error
	 //获取项目提成列表
	 GetProjectList(ctx context.Context,args *ArgsCommList,reply *ReplyCommList) error
	 //获取销售提成列表
	 GetSalesList(ctx context.Context,args *ArgsCommList,reply *ReplyCommList) error
	 //获取岗位对应的项目提成信息
	 GetProjectByBusId(ctx context.Context,args *ArgsCommSingle,reply *ReplyCommSingle) error
	 //获取销售对应的项目提成信息
	 GetSalesByBusId(ctx context.Context,args *ArgsCommSingle,reply *ReplyCommSingle) error
}