package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

/*
	问答模块：提问
*/
const (
	//	审核状态： 1待审核 2已通过 3未通过
	AuditStatusPending = 1
	AuditStatusPass    = 2
	AuditStatusNotPass = 3

	// 删除的状态
	DeleteStatusNo  = 1 // 未删除
	DeleteStatusYes = 2 // 已删除

	//关键字所属类型
	KeywordsTypeAsk = 1 // 问答

	//提问来源
	SourceUser  = 0 // 用户
	SourceAdmin = 1 // 后台

	//是否推荐
	RecNo = 1 //不推荐
	RecYes  = 2 //推荐
)

//最问答-提问基础结构
type AskQuizBase struct {
	Id         int
	Uid        int
	NickName   string // 用户昵称
	AvatarUrl  string // 头像地址
	Content    string // 提问的内容
	CateId     int    // 问题分类
	CateName   string
	IsRec      int   // 是否推荐
	AnswerNum  int   // 回答量
	AssistNum  int   // 点赞量
	CollectNum int   // 收藏量
	Status int
	Ctime      int64 // 提问时间
	CtimeStr   string
	//RecStartTime    int64 // 推荐开始日期
	//RecStartTimeStr string
	//RecEndTime      int64 // 推荐结束日期
	//RecEndTimeStr   string
}

//
type ReplyAddAskId struct {
	Id int
}
type ImgIdExplain struct {
	Id      int
	Explain string // 图片说明
}
type ImgJosn struct {
	ImgHash string
	ImgPath string // 图片路径
	Explain string // 图片说明
}

//图片Hash和说明
type ImgHashExplain struct {
	ImgHash string // 图片hash
	Explain string // 图片说明
}

//提问入参
type ArgsAddAskQuiz struct {
	common.Utoken
	CateId  int
	Content string
	Imgs    []ImgHashExplain // 图片hash最多三张
}

//提问列表入参
type ArgsGetAskQuiz struct {
	common.Utoken // 可选参数，如果传入需要判断当前用户是否收藏和点赞过
	common.Paging
	KeyWords string // 关键字
}

//提问列表出参Base
type GetAskQuizBase struct {
	Id          int
	Uid         int
	NickName    string // 用户昵称
	AvatarUrl string
	Content     string // 提问的内容
	CateId      int    // 问题分类
	CateName    string
	IsRec       int   // 是否推荐
	AnswerNum   int   // 回答量
	AssistNum   int   // 点赞量
	CollectNum  int   // 收藏量
	Ctime       int64 // 提问时间
	CtimeStr    string
	UserAssist  bool          // 用户是否点赞过
	UserCollect bool          // 用户是否收藏过
	Imgs        []ImgJosn     // 图片数据
	LastAnswer  NewAnswerBase // 最新的一条答案
}

//提问列表出参
type ReplyGetAskQuiz struct {
	TotalNum int
	Lists    []GetAskQuizBase
}

//提问详情入参
type ArgsGetAskQuizInfo struct {
	common.Utoken     // 可选参数
	QuizId        int // 提问Id
}

//提问详情出参
type ReplyGetAskQuizInfo struct {
	AskQuizBase
	UserAssist  bool              // 用户是否点赞过
	UserCollect bool              // 用户是否收藏过
	Imgs        []ImgJosn         // 图片数据
	AnswerLists []AnswerListsBase // 默认展示最新的三条提问答案
}

type ArgsDisposeAskQuizRpc struct {
	Uid     int
	AqmMaps []map[string]interface{}
}
type ReplyDisposeAskQuizRpc struct {
	AqmMaps []map[string]interface{}
}

//点赞提问
type ArgsPraiseQuiz struct {
	common.Utoken
	QuizId int
}

//收藏提问
type ArgsCollectQuiz struct {
	common.Utoken
	QuizId int
}

//我的收藏的提问入参
type ArgsGetMyCollectQuiz struct {
	common.Utoken
	common.Paging
}
type GetMyCollectQuizBase struct {
	Id        int    // 提问id
	Content   string // 提问内容
	AnswerNum int    // 回答量
}

//我的收藏的提问出参
type ReplyGetMyCollectQuiz struct {
	TotalNum int
	Lists    []GetMyCollectQuizBase
}

//关键字结构
type QuizKeywordsBase struct {
	Id      int
	Keyword string
	ModId   int
	IsDel   int
}

//获取问答-关键字入参
type ArgsGetQuizKeywords struct {
	common.Paging
}

//获取问答-关键字出参
type ReplyGetQuizKeywords struct {
	TotalNum int
	Lists    []string
}

//我的回复(相当于提问的答案)入参
type ArgsGetMyQuizAnswer struct {
	common.Utoken
	common.Paging
}
type GetMyQuizAnswerBase struct {
	Id        int
	Content   string // 答案内容
	AnswerNum int    // 回答量
	Ctime int64
	CtimeStr string
}

//我的回复(相当于提问的答案)出参
type ReplyGetMyQuizAnswer struct {
	TotalNum int
	Lists    []GetMyQuizAnswerBase
}

//我的提问入参
type ArgsGetMyQuiz struct {
	common.Utoken
	common.Paging
}

//我的提问出参
type ReplyGetMyQuiz struct {
	TotalNum int
	Lists    []GetMyQuizAnswerBase
}

//==============================================后台管理接口============================================

//后台添加提问
type ArgsAdminAddOrUpdateQuiz struct {
	common.Autoken
	Id           int
	CateId       int    //二级分类id
	Content      string //内容
	AuthorId     int
	Source       int              //来源
	IsRec        int              //是否推荐 默认 1不推荐 2推荐
	Imgs         []ImgHashExplain // 图片hash最多三张
	RecStartTime string           //推荐日期 开始时间,格式：2006-01-06
	RecEndTime   string           //推荐日期 结束时间
}

//后台列表参数入参
type AdminAuditListBase struct {
	common.Autoken
	common.Paging
	Status int //审核状态： 1待审核 2已通过 3未通过
}

//后台提问列表接口入参
type ArgsAdminGetQuizList struct {
	AdminAuditListBase
	IsRec int //1不推荐  2推荐
}
type AdminGetQuizListBase struct {
	Id         int
	Uid        int
	NickName   string // 用户昵称
	Content    string // 提问的内容
	CateId     int    // 问题分类
	CateName   string
	IsRec      int   // 是否推荐:1不推荐 2推荐
	Status     int   //审核状态: 1待审核 2已通过 3未通过
	AnswerNum  int   // 回答量
	AssistNum  int   // 点赞量
	CollectNum int   // 收藏量
	Ctime      int64 // 提问时间
	CtimeStr   string
	Source     string //来源默认0  0是用户 1是后台
}

//后台提问列表接口出参
type ReplyAdminGetQuizList struct {
	TotalNum int
	Lists    []AdminGetQuizListBase
}

//审核提问入参
type ArgsAdminAuditQuiz struct {
	common.Autoken
	QuizId int
	Status int
}

//后台提问详情入参
type ArgsAdminGetQuizInfo struct {
	common.Autoken
	QuizId int
}

//后台提问详情出参
type ReplyAdminGetQuizInfo struct {
	Id              int
	Content         string
	Uid             int
	NickName        string
	CateId          int
	CateName        string
	Status          int
	IsRec           int
	Imgs            []ImgJosn // 图片数据
	RecStartTimeStr string
	RecEndTimeStr   string
}

//后台获取关键字列表出参
type ReplyAdminGetKeywordLists struct {
	TotalNum int
	Lists    []AdminGetKeyword
}

type AdminGetKeyword struct {
	Id      int
	Keyword string
	ModId   int
}

//后台添加/修改关键字入参
type ArgsAdminAddOrUpdateKeyword struct {
	common.Autoken
	Id      int
	ModId   int    // 所属模块，默认是问答
	Keyword string //关键字
}

//后台删除关键字入参
type ArgsAdminDelKeyword struct {
	common.Autoken
	Id int
}

type ArgsSetQuiz struct {
	Id int
}

type AskQuiz interface {
	//用户发表提问
	AddAskQuiz(ctx context.Context, args *ArgsAddAskQuiz, reply *ReplyAddAskId) error
	//查看提问列表
	GetAskQuiz(ctx context.Context, args *ArgsGetAskQuiz, reply *ReplyGetAskQuiz) error
	//内部处理提问列表数据rpc
	DisposeAskQuizRpc(ctx context.Context, args *ArgsDisposeAskQuizRpc, reply *ReplyDisposeAskQuizRpc) error
	//查看提问详情(包含三条默认的提问答案)
	GetAskQuizInfo(ctx context.Context, args *ArgsGetAskQuizInfo, reply *ReplyGetAskQuizInfo) error
	//点赞提问
	PraiseQuiz(ctx context.Context, args *ArgsPraiseQuiz, reply *bool) error
	//收藏提问
	CollectQuiz(ctx context.Context, args *ArgsCollectQuiz, reply *bool) error
	//我收藏的提问
	GetMyCollectQuiz(ctx context.Context, args *ArgsGetMyCollectQuiz, reply *ReplyGetMyCollectQuiz) error
	//获取问答关键字
	GetQuizKeywords(ctx context.Context, args *ArgsGetQuizKeywords, reply *ReplyGetQuizKeywords) error
	//我的回复(相当于提问的答案)
	GetMyQuizAnswer(ctx context.Context, args *ArgsGetMyQuizAnswer, reply *ReplyGetMyQuizAnswer) error
	//我的提问
	GetMyQuiz(ctx context.Context, args *ArgsGetMyQuiz, reply *ReplyGetMyQuiz) error
	//提问数据-rpc
	GetQuizListByIdsRpc(ctx context.Context, quizIds *[]int, reply *[]map[string]interface{}) error
	//根据id获取简单的提问数据
	GetSimpleQuizById(ctx context.Context,quizId *int,reply *AskQuizBase)error

	//==============================================后台管理接口============================================
	//获取提问列表
	AdminGetQuizList(ctx context.Context, args *ArgsAdminGetQuizList, reply *ReplyAdminGetQuizList) error
	//添加/修改提问
	AdminAddOrUpdateQuiz(ctx context.Context, args *ArgsAdminAddOrUpdateQuiz, reply *ReplyAddAskId) error
	//审核提问
	AdminAuditQuiz(ctx context.Context, args *ArgsAdminAuditQuiz, reply *bool) error
	//提问详情
	AdminGetQuizInfo(ctx context.Context, args *ArgsAdminGetQuizInfo, reply *ReplyAdminGetQuizInfo) error
	//添加提问关键字
	AdminAddKeyword(ctx context.Context, args *ArgsAdminAddOrUpdateKeyword, reply *bool) error
	//修改提问关键字
	AdminUpdateKeyword(ctx context.Context, args *ArgsAdminAddOrUpdateKeyword, reply *bool) error
	//关键字列表
	AdminGetKeywordLists(ctx context.Context, args *ArgsGetQuizKeywords, reply *ReplyAdminGetKeywordLists) error
	//删除提问关键字
	AdminDelKeyword(ctx context.Context, args *ArgsAdminDelKeyword, reply *bool) error
}
