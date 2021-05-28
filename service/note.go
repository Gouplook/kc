package service

import (
	"context"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Note struct {
	nl *logics.NodeLogic
}

//初始化服务
func (n *Note) Init()*Note{
	n.nl = new(logics.NodeLogic)
	return n
}

//获取注意事项列表
func (n *Note)GetNotes(ctx context.Context, params *cards.EmptyParams, replies *cards.GetNotesReplies) (err error){
	replies.List,err = n.nl.GetNotes()
	return
}

//获取要求列表
func (n *Note)GetRequirements(ctx context.Context, params *cards.EmptyParams, replies *cards.GetRequirementsReplies)(err error){
	replies.List,err = n.nl.GetRequirements()
	return
}


