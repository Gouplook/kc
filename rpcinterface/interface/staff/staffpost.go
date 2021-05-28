package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)
// 定义员工岗位接口
// @author liyang<654516092@qq.com>
// @date  2020/4/1 11:11

//添加员工岗位入参
type ArgsAddPost struct{
	common.Utoken     //用户信息
	common.BsToken    //企业/商户/分店ID
	PostName string   //岗位名称
	IsPerformance int //岗位是否产生业绩 0=否 1=是
}
//编辑员工岗位入参
type ArgsEditPost struct{
	common.Utoken     //用户信息
	common.BsToken    //企业/商户/分店ID
	PostId int        //岗位ID
	PostName string   //岗位名称
	IsPerformance int //岗位是否产生业绩 0=否 1=是
	JobTitle  string //岗位职称，多个以","号隔开
	OpType   int //操作类型 0=编辑 1=设置单个业绩
}
//设置员工岗位职称入参
type ArgsSetJobTitle struct {
	common.Utoken     //用户信息
	common.BsToken    //企业/商户/分店ID
	PostId  int    //岗位ID
	JobTitle string  //岗位职称名称
}
//删除员工岗位入参
type ArgsDelPost struct {
	common.Utoken     //用户信息
	common.BsToken    //企业/商户/分店ID
	PostId  int    //岗位ID
}
//操作返回值
type ReplyPost struct {
	PostId int //岗位ID
}

//获取岗位列表入参
type ArgsGetPostList struct {
	common.Utoken
	common.BsToken
}
//返回岗位信息
type ReplyInfo struct {
	PostId int `mapstructure:"id"`//岗位ID
	PostName string //岗位名称
	IsPerformance int //是否产生业绩 0=否 1=是
	JobTitle  string  //岗位职称，多个以","号隔开
}

//定义岗位接口
type StaffPost interface {
	//添加岗位
	AddPost(ctx context.Context,args *ArgsAddStaff,reply *ReplyPost) error
	//编辑岗位
	EditPost(ctx context.Context,args *ArgsEditStaff,reply *ReplyPost) error
	//设置岗位职称
	SetJobTitle(ctx context.Context,args *ArgsSetJobTitle,reply *ReplyPost) error
	//删除岗位
	DeletePost(ctx context.Context,args *ArgsDelPost,reply *ReplyPost) error
	//获取岗位列表
	GetPostList(ctx context.Context,args *ArgsGetPostList,reply *[]ReplyInfo) error
}