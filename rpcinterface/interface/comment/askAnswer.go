package comment

/*
	1.问答模块：提问的答案
	2.问答模块：答案的评论
*/
import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//最新的回答结构
type NewAnswerBase struct {
	Id         int
	Uid        int
	NickName   string // 用户昵称
	AvatarUrl  string // 头像地址
	Content    string // 回答内容
	AssistNum  int    // 点赞量
	CommentNum int    // 评论量
	UserAssist bool   // 当前用户是否点赞过
	Ctime      int64
	CtimeStr   string // 时间
}

// 答案和评论组合结构
type AnswerListsBase struct {
	Id           int
	Uid          int
	NickName     string // 用户昵称
	AvatarUrl    string // 头像地址
	Content      string // 回答内容
	AssistNum    int    // 点赞量
	CommentNum   int    // 评论量
	UserAssist   bool   // 当前用户是否点赞过
	Ctime        int64
	CtimeStr     string             // 时间
	CommentLists []AskAnswerComment // 最新的三条答案评论
}

//提问-发表答案入参
type ArgsAddAskAnswer struct {
	common.Utoken
	QuizId  int    // 提问id
	Content string // 回答内容
}

//查看全部回答入参
type ArgsGetAllAskAnswer struct {
	common.Utoken // 可选参数
	common.Paging
	QuizId int // 提问id
}

//提问-全部回答
type ReplyGetAllAskAnswer struct {
	TotalNum    int
	AnswerLists []AnswerListsBase
}

//答案评论结构体
type AskAnswerComment struct {
	AnswerId  int
	Uid       int
	NickName  string
	AvatarUrl string // 头像地址
	Content   string //回到内容
	Ctime     int64
	CtimeStr  string
}

//查看指定答案下的全部评论入参
type ArgsGetAnswerComments struct {
	common.Utoken // 可选参数
	common.Paging
	AnswerId int // 答案id

}

//查看指定答案下的全部评论出参
type ReplyGetAnswerComments struct {
	NewAnswerBase                    // 答案数据结构
	TotalNum int
	CommentLists  []AskAnswerComment // 答案评论数据结构
}

//点赞答案
type ArgsPraiseAnswer struct {
	common.Utoken
	AnswerId int
}

//发表答案评论
type ArgsAddAnswerComment struct {
	common.Utoken
	AnswerId int
	Content string
}

//==============================================后台管理接口============================================

//后台答案列表接口出参
type ReplyAdminGetAnswerList struct {
	TotalNum int
	List     []AdminGetAnswerListBase
}
type AdminGetAnswerListBase struct {
	Id         int    //答案id
	Content    string //回答内容
	AssistNum  int    // 点赞量
	CommentNum int    // 评论量
	Ctime      int64
	CtimeStr   string //回答时间
	Status     int    //审核状态
}

//审核答案入参
type ArgsAdminAuditAnswerOrComment struct {
	common.Autoken
	AnswerId int // 答案id
	CommentId int // 评论id
	Status   int
}

//后台评论列表出参
type ReplyAdminGetCommentList struct {
	TotalNum int
	Lists []AdminGetCommentList
}
type AdminGetCommentList struct{
	Id int
	Content string
	Status int
	Ctime int64
	CtimeStr string
}

type AskAnswer interface {
	//用户回答提问
	AddAskAnswer(ctx context.Context, args *ArgsAddAskAnswer, reply *ReplyAddAskId) error
	//查看指定提问下的全部回答(包含三条最新的答案评论)
	GetAllAskAnswer(ctx context.Context, args *ArgsGetAllAskAnswer, reply *ReplyGetAllAskAnswer) error
	//查看答案详情
	GetAnswerComments(ctx context.Context, args *ArgsGetAnswerComments, reply *ReplyGetAnswerComments) error
	//点赞答案
	PraiseAnswer(ctx context.Context, args *ArgsPraiseAnswer, reply *bool) error
	//发表答案评论
	AddAnswerComment(ctx context.Context, args *ArgsAddAnswerComment, reply *bool) error

	//==============================================后台管理接口============================================

	//后台答案列表接口出参
	AdminGetAnswerList(ctx context.Context, args *AdminAuditListBase, reply *ReplyAdminGetAnswerList) error
	//审核答案/评论入参
	AdminAuditAnswer(ctx context.Context, args *ArgsAdminAuditAnswerOrComment, reply *bool) error
	//后台评论列表出参
	AdminGetCommentList(ctx context.Context,args *AdminAuditListBase,reply *ReplyAdminGetCommentList)error
}
