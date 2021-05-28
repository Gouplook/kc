package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// -----------------------------评论-------------------------------

// ArgsAddCommnets 添加评论入参
type ArgsAddCommnets struct {
	common.Utoken     // 必须登陆才可以评论
	CommentId     int //
	Uid           int //
	Type          int //
	Content       string
}

// ReplyAddComments 添加评论出参
type ReplyAddComments struct {
	CommentId int
}

// ArgsAdminCommentsList 后台攻略/探店评论列表入参
type ArgsAdminCommentsList struct {
	common.Paging
	common.Autoken
	Type   int
	Status string //
}
type AdminCommentsListBase struct {
	Id        int    `mapstructure:"id"`         // 评论id
	CommentId int    `mapstructure:"comment_id"` // 模块id
	Uid       int    `mapstructure:"uid"`        // uid
	Nick      string `mapstructure:"nick"`       // 用户昵称
	Content   string `mapstructure:"content"`    // 评论内容
	Title     string `mapstructure:"title"`      // 主题:攻略,探店
	Type      int    `mapstructure:"type"`       // 板块类型 1-攻略 2-探店
	Ctime     int64    `mapstructure:"ctime"`      // 评论时间
	CtimeStr  string `mapstructure:"ctime_str"`  //
	Status    int    `mapstructure:"status"`     // 状态
}

// ReplyAdminCommentsList 后台攻略/探店评论列表出参
type ReplyAdminCommentsList struct {
	TotalNum int //
	Lists     []AdminCommentsListBase
}

// ArgsAuditComment 审核用户评论
type ArgsAuditComment struct {
	common.Autoken
	Id     int
	Status int
}

// ArgsUserComentList 前台获取评论列表入参
type ArgsUserComentList struct {
	common.Utoken
	common.Paging
	RaiderID int // 攻略id
}

// ReplyUserComentList 前台获取评论出参
type ReplyUserComentList struct {
}

// ArgsAddReplyComment 添加评论回复入参
type ArgsAddReplyComment struct {
	common.Utoken
	CommentId int    // 评论id
	Content   string // 回复内容
	Type      int    // 板块类型 1-攻略 2-探店
	ReplyUid  int    // 回复人id
}

// ReplyAddReplyComment 添加评论回复出参
type ReplyAddReplyComment struct {
	ReplyId int // 评论回复id
}

// ArgsAdminDelReplyComment 管理员删除回复评论
type ArgsAdminDelReplyComment struct {
	common.Autoken
	Id int //
}

// ArgsAdminReplyCommentList 后台评论回复列表入参
type ArgsAdminReplyCommentList struct {
	common.Autoken
	common.Paging
	CommentId int // 评论id
}

// AdminReplyCommentLsitBase 后台评论回复列表base
type AdminReplyCommentLsitBase struct {
	Id        int
	CommentId int `mapstructure:"comment_id"` // 评论的模块id
	Content   string
	ReplyUid  int `mapstructure:"reply_uid"`
	Author    string
	Ctime     int64
	CtimeStr  string
}

// ReplyAdminReplyCommentLsit　后台评论回复列表出参
type ReplyAdminReplyCommentLsit struct {
	TotalNum int
	Lists     []AdminReplyCommentLsitBase
}

// ArgsUserCommentList 前台用户评论列表入参
type ArgsUserCommentList struct {
	common.Paging
	Id   int // 探店/攻略id
	Type int // 评论模块类型必传
}

// UserCommentListBase UserCommentListBase
type UserCommentListBase struct {
	Id       int
	Uid      int
	Nick     string             // 用户昵称
	ImgUrl   string             // 头像url
	ImgHash  string             // 头像hash
	ImgId    int                // 头像id
	Content  string             // 内容
	Ctime    int64                // 评论时间
	CtimeStr string             //
	NewReply    *ReplyCommentListBase // 最新的评论回复
	ReplyTotal int // 回复评论量
}

// ReplyUserCommentList 前台用户评论列表出参
type ReplyUserCommentList struct {
	TotalNum int
	Lists     []UserCommentListBase
}

// ArgsReplyCommentList 前台评论回复列表
type ArgsReplyCommentList struct {
	common.Paging
	Type      int
	CommentId int // 评论id
}
type ReplyCommentList struct {
	TotalNum int
	Lists []ReplyCommentListBase
}
// ReplyCommentList ReplyCommentList
type ReplyCommentListBase struct {
	Id        int
	CommentId int    `mapstructure:"comment_id"` // 主评论id
	Nick      string // 用户昵称
	ImgUrl    string // 头像url
	ImgHash   string // 头像hash
	ImgId     int    // 头像id
	Content   string // 内容
	Ctime     int64    // 评论时间
	CtimeStr  string //
	ReplyUid  int    `mapstructure:"reply_uid"` // 回复人
}

type Comments interface {
	// AdminCommentsList 后台攻略/探店评论列表
	AdminCommentsList(ctx context.Context, args *ArgsAdminCommentsList, reply *ReplyAdminCommentsList) error
	// AuditComment 后台审核用户评论
	AuditComment(ctx context.Context, args *ArgsAuditComment, reply *bool) error

	// AddCommnets 添加评论
	AddCommnets(ctx context.Context, args *ArgsAddCommnets, reply *ReplyAddComments) error
	// UserCommentList 前台评论列表
	UserCommentList(ctx context.Context, args *ArgsUserCommentList, reply *ReplyUserCommentList) error
	// ReplyCommentList 前台评论回复列表
	ReplyCommentList(ctx context.Context, args *ArgsReplyCommentList, reply *ReplyCommentList) error
	// AddReplyComments 添加评论回复
	AddReplyComments(ctx context.Context, args *ArgsAddReplyComment, reply *ReplyAddReplyComment) error
	//　AdminDelReplyComment 管理员删除评论回复
	AdminDelReplyComment(ctx context.Context, args *ArgsAdminDelReplyComment, reply *bool) error
	//　AdminReplyCommentList　后台评论回复列表
	AdminReplyCommentList(ctx context.Context, args *ArgsAdminReplyCommentList, reply *ReplyAdminReplyCommentLsit) error
}
