//标签
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/10 16:32
package service

import (
	"context"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type Tag struct {
}

//添加标签
func (t *Tag) AddTag( ctx context.Context, args *cards.ArgAddTag, tagid *int ) ( err error ) {
	mTag := new(logics.TagsLogic)
	*tagid, err = mTag.AddTag(args)
	if err != nil{
		return
	}
	return
}

//修改标签
func (t *Tag) EditTag( ctx context.Context, args *cards.ArgEditTag, reply *bool ) ( err error ) {
	mTag := new(logics.TagsLogic)
	err = mTag.EditTag(args)
	if err != nil{
		return
	}
	return
}

//删除标签
func (t *Tag) DelTag( ctx context.Context, args *cards.ArgDelTag, reply *bool ) (err error) {
	mTag := new(logics.TagsLogic)
	err = mTag.DelTag(args)
	if err != nil{
		return
	}
	return
}

//获取商家的所有标签
func (t *Tag) BusTags( ctx context.Context, args *common.BsToken, reply *[]cards.TagInfo ) (err error) {
	mTag := new(logics.TagsLogic)
	*reply, err = mTag.GetBusTags(*args)
	if err != nil{
		return
	}
	return
}

