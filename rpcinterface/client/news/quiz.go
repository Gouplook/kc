package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Quiz struct {
	client.Baseclient
}

func (q *Quiz) Init() *Quiz {
	q.ServiceName = "rpc_news"
	q.ServicePath = "Quiz"
	return q
}

//后台

//添加一级分类
func (q *Quiz) AddCate(ctx context.Context, args *news.ArgsCateAdd, reply *int) error {
	return q.Call(ctx, "AddCate", args, reply)
}

//添加二级分类
func (q *Quiz) AddSubCate(ctx context.Context, args *news.ArgsSubCateAdd, reply *int) error {
	return q.Call(ctx, "AddSubCate", args, reply)
}

//修改二级分类
func (q *Quiz) UpdateSubCate(ctx context.Context, args *news.ArgsSubCateAdd, reply *bool) error {
	return q.Call(ctx, "UpdateSubCate", args, reply)
}

//查询所有一级分类一级下属二级分类
func (q *Quiz) GetCateAll(ctx context.Context, args *news.Args, reply *news.ReplyCateAll) error {
	return q.Call(ctx, "GetCateAll", args, reply)
}

//按分页 推荐 审核 查询提问 后台查询
func (q *Quiz) GetQuizByPage(ctx context.Context, args *news.ArgsQuizGetP, reply *news.ReplyQuizPage) error {
	return q.Call(ctx, "GetQuizByPage", args, reply)
}

//添加提问
func (q *Quiz) AddQuiz(ctx context.Context, args *news.ArgsQuizAdd, reply *int) error {
	return q.Call(ctx, "AddQuiz", args, reply)
}

//修改提问
func (q *Quiz) UpdateQuiz(ctx context.Context, args *news.ArgsQuizAdd, reply *bool) error {
	return q.Call(ctx, "UpdateQuiz", args, reply)
}

//通过审核 或者 拒绝审核
func (q *Quiz) CheckQuiz(ctx context.Context, args *news.ArgsQuizCheck, reply *bool) error {
	return q.Call(ctx, "CheckQuiz", args, reply)
}

//按 分页 审核 查询答案 后台用
func (q *Quiz) GetAnswerByPage(ctx context.Context, args *news.ArgsAnswerGetP, reply *news.ReplyAnswerPage) error {
	return q.Call(ctx, "GetAnswerByPage", args, reply)
}

//通过审核 或者 拒绝审核 答案 or 评论
func (q *Quiz) CheckAnswer(ctx context.Context, args *news.ArgsAnswerCheck, reply *bool) error {
	return q.Call(ctx, "CheckAnswer", args, reply)
}

//按 分页 审核 查询评论
func (q *Quiz) GetCommentByPage(ctx context.Context, args *news.ArgsAnswerGetP, reply *news.ReplyCommentPage) error {
	return q.Call(ctx, "GetCommentByPage", args, reply)
}

//添加 关键字
func (q *Quiz) AddKeyWord(ctx context.Context, args *news.ArgsKeyWordAdd, reply *int) error {
	return q.Call(ctx, "AddKeyWord", args, reply)
}

//修改 关键字
func (q *Quiz) UpdateKeyWord(ctx context.Context, args *news.ArgsKeyWordAdd, reply *bool) error {
	return q.Call(ctx, "UpdateKeyWord", args, reply)
}

//删除 关键字
func (q *Quiz) DelKeyWord(ctx context.Context, args *news.ArgsKeyWordDel, reply *bool) error {
	return q.Call(ctx, "DelKeyWord", args, reply)
}

//根据二级分类id查询详情
func (q *Quiz) GetCateDetailById(ctx context.Context, args *int, reply *news.ReplyCateDetail) error {
	return q.Call(ctx, "GetCateDetailById", args, reply)
}

//前后台 共用

//按模块查询关键字
func (q *Quiz) GetKeyWordByModId(ctx context.Context, args *news.ArgsKeyWordGet, reply *news.ReplyKeyWord) error {
	return q.Call(ctx, "GetKeyWordByModId", args, reply)
}

//前台

//添加提问
func (q *Quiz) AddQuiz2(ctx context.Context, args *news.ArgsQuizAdd2, reply *int) error {
	return q.Call(ctx, "AddQuiz2", args, reply)
}

//根据用户Id查询 我的问答 数据
func (q *Quiz) GetQuizByUid(ctx context.Context, args *news.ArgsQuizGet, reply *news.ReplyQuiz) error {
	return q.Call(ctx, "GetQuizByUid", args, reply)
}

//App端查询提问
func (q *Quiz) GetQuizByApp(ctx context.Context, args *news.ArgsQuizGetA, reply *news.ReplyQuizGetA) error {
	return q.Call(ctx, "GetQuizByApp", args, reply)
}

//app端根据问题id查询答案和评论  //从问答列表 进入
func (q *Quiz) GetQuizDetail(ctx context.Context, args *news.ArgsQuizD, reply *[]news.ReplyInfo) error {
	return q.Call(ctx, "GetQuizDetail", args, reply)
}

//app端根据问题id查询答案和评论  //从我的问答 问题id 查询详情
func (q *Quiz) GetQuizDetailB(ctx context.Context, args *news.ArgsQuizD, reply *news.ReplyQuizB) error {
	return q.Call(ctx, "GetQuizDetailB", args, reply)
}

//用户 添加 回复
func (q *Quiz) AddAnswer(ctx context.Context, args *news.ArgsAnswerAdd, reply *int) error {
	return q.Call(ctx, "AddAnswer", args, reply)
}

//用户 收藏 提问
func (q *Quiz) CollectAnswer(ctx context.Context, args *news.ArgsAnswerCollect, reply *int) error {
	return q.Call(ctx, "CollectAnswer", args, reply)
}

//用户 取消 收藏提问
func (q *Quiz) CancelCollect(ctx context.Context, args *news.ArgsCollectCancel, reply *bool) error {
	return q.Call(ctx, "CancelCollect", args, reply)
}

//用户点赞    包括 提问点赞 答案点赞
func (q *Quiz) AddGiveALike(ctx context.Context, args *news.ArgsGiveALike, reply *bool) error {
	return q.Call(ctx, "AddGiveALike", args, reply)
}

//根据提问id查询详情
func (q *Quiz) GetQuizOne(ctx context.Context, args *int, reply *news.ReplyQuizOne) error {
	return q.Call(ctx, "GetQuizOne", args, reply)
}

//根据2级分类查询1级分类id
func (q *Quiz) GetCateBySubId(ctx context.Context, args *int, reply *int) error {
	return q.Call(ctx, "GetCateBySubId", args, reply)
}