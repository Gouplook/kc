package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Note struct {
	client.Baseclient
}

func (n *Note)Init() *Note {
	n.ServiceName = "rpc_cards"
	n.ServicePath = "Note"
	return n
}

//获取注意事项列表
func (n* Note)GetNotes(ctx context.Context, params *cards.EmptyParams, replies *cards.GetNotesReplies) (err error) {
	err = n.Call(ctx, "GetNotes", params, replies)
	return
}
//获取要求列表
func (n *Note)GetRequirements(ctx context.Context, params *cards.EmptyParams, replies *cards.GetRequirementsReplies) (err error){
	err = n.Call(ctx, "GetRequirements", params, replies)
	return
}

