package public

import "context"

type EmptyParams struct {
}
type AdvantageBase struct {
	Id   int
	Name string
}

//所有的优势出参
type ReplyGetAdvantageList struct {
	Lists []AdvantageBase
}

type ArgsGetAdvantageByIds struct {
	Ids []int
}

//指定的优势出参
type ReplyGetAdvantageByIds struct {
	Lists []AdvantageBase
}

type Advantage interface {
	//获取所有的优势标签
	GetAdvantageList(ctx context.Context, args *EmptyParams, reply *ReplyGetAdvantageList) error
	//获取指定的优势标签-RPC
	GetAdvantageByIds(ctx context.Context, args *ArgsGetAdvantageByIds, reply *ReplyGetAdvantageByIds) error
}
