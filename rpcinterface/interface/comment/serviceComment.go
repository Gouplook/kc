package comment

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//根据门店id和卡包id获取评论入参
type ArgsComment struct {
	common.Paging
	common.BsToken
	common.Utoken
	Status int // 1-未回复；2-已回复
	ScoreType int // 1好评 2中评 3差评
	ShopId int  //门店id
	BusId int
	SingleId int //单项目id
	StartTime string
	EndTime string
}

//用户评论返回
type ReplyComment struct {
	TotalNum int
	ComprehensiveScore float64//门店评价综合得分
	ApplauseRate float64//好评率
	UserComments      []UserComment                  //用户评价
	//ShopReplies        []Reply                        //商家回复表
	//UserInfos    map[int]map[string]interface{} //用户名称+头像
	//ContentImage map[int]file2.ReplyFileInfo    //评论内容图片
	//CommentTags  []CommentTagBase
}

//用户评价
type UserComment struct {
	Id              int    `mapstructure:"id"`
	SingleId        int    `mapstructure:"single_id"`
	SingleName        string    `mapstructure:"single_name"`
	SspId           int    `mapstructure:"ssp_id"`
	SspName         string `mapstructure:"ssp_name"`

	StaffId         int    `mapstructure:"staff_id"`
	StaffName  string
	StaffNick string
	SalesStaffId         int    `mapstructure:"sales_staff_id"`
	SalesStaffName string
	SalesStaffNick string

	BusId           int    `mapstructure:"bus_id"`
	ShopId          int    `mapstructure:"shop_id"`
	ShopName          string
	BranchName        string

	Uid             int    `mapstructure:"uid"`
	Unick  	string
	Usex  int
	UImgPath string
	UPhone string

	RelationId      int    `mapstructure:"relation_id"`
	CardPackageId   int    `mapstructure:"card_package_id"`
	CardPackageType int    `mapstructure:"card_package_type"`
	PriceScore      int    `mapstructure:"price_score"`
	ServiceScore    int    `mapstructure:"service_score"`
	EnvirScore      int    `mapstructure:"envir_score"`
	ComplexScore      float64    `mapstructure:"complex_score"`
	ScoreType      int    `mapstructure:"score_type"`
	Comment         string `mapstructure:"comment"`
	Assist          int    `mapstructure:"assist"`

	Status        string `mapstructure:"status"`
	Ctime           string `mapstructure:"ctime"`
	CtimeStr        string `mapstructure:"CtimeStr"`
	UserPraise bool // 已登录用户是否点赞

	//TagIds          string `mapstructure:"tag_ids"`
	CommentTags []CommentTagBase

	//ImgIds          string  `mapstructure:"img_ids"`
	CommentImgs []string
	ShopReplies        Reply                        //商家回复表

}

//商家回复表
type Reply struct {
	CommentId    int    `mapstructure:"comment_id"`
	ReplyContent string `mapstructure:"reply_content"`
	Ctime        string `mapstructure:"ctime"`
	CtimeStr     string `mapstructure:"CtimeStr"`
}
type ServiceCommentBase struct {
	Id               int
	WaitingCommentID int    // 待评价ID
	SingleId         int    //单项目ID
	SingleName string //单项目名
	SspId            int    //规格ID
	SspName          string //规格名称
	StaffId          int    //技师ID
	BusId            int    //企业/商户ID
	ShopId           int    //分店ID
	Uid              int    //评论者ID
	RelationId       int    //卡包关联ID
	CardPackageId    int    //卡包ID
	CardPackageType  int    //卡包类型 1=单项目 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡
	PriceScore       int    //价格评分(满分5分)
	ServiceScore     int    //服务评分(满分5分)
	EnvirScore       int    //环境评分(满分5分)
	Comment          string //评论信息
	Assist           int    //点赞数
	ImgIds           string //晒图 多个以''逗号''隔开
	TagIds           string // 评价标签ID 多个以''逗号''隔开
	Ctime            int64  //评论时间
}

type ArgsAddServiceComment struct {
	common.Utoken
	ImgHashs string // 多个图片用‘逗号’隔开
	ServiceCommentBase
}

type ReplyAddServiceComment struct {
	Id int
}

//评价点赞（如果点赞过的话就取消点赞）
type ArgsCommentPraise struct {
	common.Utoken
	CommentId int // 评论的Id
}
//获取评价得分
type ReplyGetServiceCommentScore struct {
	Id int
	ShopId int
	BusId int
	StaffId int
	PriceScore      int
	ServiceScore    int
	EnvirScore      int
}

// 获取消费者评分
type ArgsCousumerEvalutaion struct {
	BusId int //

}
type ReplyCousumerEvalutaion struct {
	CousumerEvalutaion  float64 // 消费者评分
}

type ServiceComment interface {

	//根据Saas门店id查询用户评论
	GetCommentBySaas(ctx context.Context, args *ArgsComment, reply *ReplyComment) error
	//根据门店id查询用户评论
	GetCommentByShopId(ctx context.Context, args *ArgsComment, reply *ReplyComment) error
	//根据单项目Id查询用户评论
	GetCommentBySingleId(ctx context.Context, args *ArgsComment, reply *ReplyComment) error

	//用户评论
	AddServiceComment(ctx context.Context, args *ArgsAddServiceComment, reply *ReplyAddServiceComment) error
	//根据评价ID获取评价得分
	GetServiceCommentScoreById(ctx context.Context,serviceCommentId *int ,reply *ReplyGetServiceCommentScore)error
	//根据评价ID获取评价内容
	GetServiceCommentByID(ctx context.Context, serviceCommentId *int, reply *UserComment) error
	//评价点赞
	CommentPraise(ctx context.Context,args *ArgsCommentPraise,reply *bool)error

	// 获取消费者评分
	GetCousumerEvalutaion(ctx context.Context, args *ArgsCousumerEvalutaion, reply *ReplyCousumerEvalutaion) error

}
