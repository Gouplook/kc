package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

/*
最问答-分类
*/
type ArgsEmptyParams struct {
}
type AskCateBase struct {
	Id    int
	ImgId int
	Name  string
	Path  string // 分类图片
	ImgHash string
}

type SubAskCate struct {
	Id    int
	ImgId int
	Name  string
	Path  string // 分类图片
	Sub   []AskCateBase
}

//前台-获取分类出参
type ReplyGetAskCate struct {
	TotalNum int // 一级分类
	Lists    []SubAskCate
}

//==============================================后台管理接口============================================
//添加一级分类入参
type ArgsAdminAddParentCate struct {
	common.Autoken
	Name string
}

//添加二级分类入参
type ArgsAdminAddSubCate struct {
	common.Autoken
	Id int // 二级分类id
	ParentCateId int    //一级分类id
	Name         string //二级分类名字
	ImgHash      string //二级分类图片
}

//添加分类出参
type ReplyAdminAddCate struct {
	CateId int
}

type AskCate interface {
	//前台-获取问答分类数据
	GetAskCate(ctx context.Context, args *ArgsEmptyParams, reply *ReplyGetAskCate) error

	//==============================================后台管理接口============================================
	//添加一级分类
	AdminAddParentCate(ctx context.Context, args *ArgsAdminAddParentCate, reply *ReplyAdminAddCate) error
	//添加二级分类
	AdminAddSubCate(ctx context.Context, args *ArgsAdminAddSubCate, reply *ReplyAdminAddCate) error
	//修改二级分类
	AdminUpdateSubCate(ctx context.Context, args *ArgsAdminAddSubCate, reply *bool) error
	//查询所有一级分类一级下属二级分类
	AdminGetCate(ctx context.Context, args *ArgsEmptyParams, reply *ReplyGetAskCate) error
	//根据二级分类查询1级分类id
	GetCateBySubId(ctx context.Context, args *int, reply *int) error
	//根据二级分类id查询详情
	GetCateDetailById(ctx context.Context, args *int, reply *AskCateBase) error
}
