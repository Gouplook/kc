package comment

import "context"

type CommentTagBase struct {
	TagId int
	Name string
}

type ArgsGetCommentTag struct {

}

type ReplyGetCommentTag struct {
	Lists []CommentTagBase
}

type CommentTag interface {
	//评价标签
	GetCommentTag(ctx context.Context,args *ArgsGetCommentTag,reply *ReplyGetCommentTag)error
}