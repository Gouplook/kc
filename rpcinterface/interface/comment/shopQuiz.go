package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//添加提问入参
type ArgsAddQuiz struct {
	common.Utoken
	Content string
	ShopId int
}

//添加回答入参
type ArgsAddReply struct {
	common.Utoken
	Content string
	QuizId int
	ReplyId int
}

//返回提问
type ResQuiz struct {
	Id int
	Uid int
	Nick string
	ImgUrl string
	Content string
	ShopId int
	ReplyCount int
	CreateTime int
	CreateTimeStr string
}
//返回答案
type ResReply struct {
	Id int
	QuizId int
	ReplyId int
	ReplyNick string
	Uid int
	Nick string
	ImgUrl string
	Content string
	GiveALike int
	CreateTime int
	CreateTimeStr string
}

//获取门店问答列表入参
type ArgsGetQuizList struct {
	common.Paging
	ShopId int
}

//获取门店问答列表返回
type ReplyGetQuizList struct {
	TotalNum int
	Lists []ResQuizStruct
}
type ResQuizStruct struct {
	ResQuiz ResQuiz
	ResReply ResReply
}

//获取详细问答入参
type ArgsGetQuizDetails struct {
	common.Paging
	QuizId int
}

//获取详细问答返回
type ReplyGetQuizDetails struct {
	ResQuiz ResQuiz
	ResReplyTotalNum int
	ResReplyLists []ResReply
}

type ArgsGiveALick struct {
	common.Utoken
	ReplyId int
}


type ShopQuiz interface {
	//用户添加提问
	AddQuiz(ctx context.Context, args *ArgsAddQuiz,reply *int) error
	//用户回答提问
	AddReply(ctx context.Context, args *ArgsAddReply, reply *int) error

	//获取门店问答列表
	GetQuizList(ctx context.Context, args *ArgsGetQuizList, reply *ReplyGetQuizList) error
	//获取详细问答
	GetQuizDetails(ctx context.Context, args *ArgsGetQuizDetails, reply *ReplyGetQuizDetails) error

	//门店回答点赞
	GiveALike(ctx context.Context, args *ArgsGiveALick, reply *bool) error
}
