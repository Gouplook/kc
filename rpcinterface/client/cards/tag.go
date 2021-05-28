//标签
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/10 16:41
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type Tag struct {
	client.Baseclient
}

func (t *Tag) Init() *Tag {
	t.ServiceName = "rpc_cards"
	t.ServicePath = "Tag"
	return t
}

//添加标签
func (t *Tag) AddTag(ctx context.Context, args *cards.ArgAddTag, tagid *int) (err error) {
	return t.Call(ctx, "AddTag", args, tagid)
}

//修改标签
func (t *Tag) EditTag(ctx context.Context, args *cards.ArgEditTag, reply *bool) (err error) {
	return t.Call(ctx, "EditTag", args, reply)
}

//删除标签
func (t *Tag) DelTag(ctx context.Context, args *cards.ArgDelTag, reply *bool) (err error) {
	return t.Call(ctx, "DelTag", args, reply)
}

//获取商家的所有标签
func (t *Tag) BusTags(ctx context.Context, args *common.BsToken, reply *[]cards.TagInfo) (err error) {
	return t.Call(ctx, "BusTags", args, reply)
}
