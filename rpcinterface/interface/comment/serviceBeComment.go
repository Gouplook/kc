package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//待评价
	CommentStatusWaiting = 1
	//已评价
	CommentStatusComplete = 2

	//评论回复状态
	// 未回复
	CommentReplyStatusNo = 1
	// 已回复
	CommentReplyStatusYes = 2
)

func GetCommentStatusList() []int {
	return []int{CommentStatusWaiting, CommentStatusComplete}
}

type ServiceBeCommentBase struct {
	Id              int
	SingleId        int    //单项目ID
	SingleName      string //单项目名
	SspId           int    //规格ID
	SspName         string //规格名称
	StaffId         int    //技师ID
	BusId           int    //企业/商户ID
	ShopId          int    //分店ID
	Uid             int    //用户ID
	RelationId      int    //卡包关联ID
	CardPackageId   int    //卡包ID
	CardPackageType int    //卡包类型 1=单项目 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡
	CommentId       int    //评论ID
	Status          int    //状态 1=待评价 2=已评价
	Ctime           int64  //创建时间
}

//确认消费的单项目
type ConsumeSingle struct {
	SingleId   int     //单项目id
	SingleName string  //单项目名
	SspId      int     //单项目规格组合id
	SspName    string  //单项目规格名称
	Num        int     //消费次数
	StaffId    int     //服务技师
	Price      float64 //消费的价格
	ImgId      int     //无需传，方法内部自己根据SingleId去查找
}

//添加待评价服务入参
type ArgsAddServiceBeComment struct {
	BusId           int //企业/商户ID
	ShopId          int //分店ID
	Uid             int //用户ID
	RelationId      int //卡包关联ID
	CardPackageId   int //卡包ID
	CardPackageType int //卡包类型 1=单项目 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡
	CommentId       int //评论ID(默认是0)
	Status          int //状态 1=待评价 2=已评价(默认是1)
	Singles         []ConsumeSingle
}

//添加待评价服务出参
type ReplyAddServiceBeComment struct {
	Result bool
}

//根据状态获取待评价/已评价的服务列表入参
type ArgsGetServiceBeComments struct {
	common.Utoken
	common.Paging
	Status          int
	Uid             int //内部使用
	CardPackageType int // 卡包类型
}
type ReplyGetServiceBeCommentsBase struct {
	ServiceBeCommentBase
	SingleImgId int
	ImgPath     string
	Num         int     //消费次数
	StaffId     int     //服务技师
	StaffName string //服务技师名
	SinglePrice float64 // 消费的价格
	ShopName    string
	BranchName  string
	CtimeStr    string

	PriceScore      int
	ServiceScore    int
	EnvirScore      int
	ComplexScore      float64
	ScoreType      int
	Comment         string
	Assist          int
	ReplyStatus        int
	CommentCtime           string
	CommentCtimeStr           string
	//UserPraise bool // 已登录用户是否点赞
	CommentTags []CommentTagBase
	CommentImgs []string
	ShopReplies        Reply                        //商家回复表*/
}

//根据状态获取待评价/已评价的服务列表出参
type ReplyGetServiceBeComments struct {
	TotalNum int
	Lists    []ReplyGetServiceBeCommentsBase
}

type ServiceBeComment interface {
	//添加待评价服务-rpc
	AddServiceBeCommentRpc(ctx context.Context, args *ArgsAddServiceBeComment, reply *ReplyAddServiceBeComment) error
	//根据状态获取待评价/已评价的服务列表
	GetServiceBeComments(ctx context.Context, args *ArgsGetServiceBeComments, reply *ReplyGetServiceBeComments) error
}
