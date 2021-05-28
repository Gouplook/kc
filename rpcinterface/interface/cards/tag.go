//管理标签
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/10 15:52
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//标签是否被删除 0=否 1=是
	TAG_IS_DEL_no  = 0
	TAG_IS_DEL_yes = 1
)

//添加标签入参
type ArgAddTag struct {
	common.BsToken
	Name string //标签名称
}

//修改标签入参
type ArgEditTag struct {
	common.BsToken
	Name  string //标签名称
	TagId int    //标签id
}

//删除标签入参
type ArgDelTag struct {
	common.BsToken
	TagId int //标签id
}

//标签信息
type TagInfo struct {
	TagId int
	Name  string
	Ctime int
}

type Tags interface {
	//添加标签
	AddTag(ctx context.Context, args *ArgAddTag, tagid *int) error
	//修改标签
	EditTag(ctx context.Context, args *ArgEditTag, reply *bool) error
	//删除标签
	DelTag(ctx context.Context, args *ArgDelTag, reply *bool) error
	//获取商家的所有标签
	BusTags(ctx context.Context, args *common.BsToken, reply *[]TagInfo) error
}
