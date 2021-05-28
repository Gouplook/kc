package news

import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	// 状态待审核
	STATUS_WAITING = 0
	// 状态通过
	STATUS_PASS = 1
	// 状态不通过
	STATUS_NOT_PASS = 2

	// 来源后台
	SOURE_ADMIN = 0
	// 来源用户
	SOURE_USER = 1

	// 是否推荐:不推荐
	RECOMMEND_NO = 0
	// 是否推荐:推荐
	RECOMMEND_YES = 1

	//	模块类型
	// 攻略
	MODULE_TYPE_RAIDERS = 1
	// 探店
	MODULE_TYPE_EXPLORATION = 2
	// 问答
	MODULE_TYPE_QUIZ = 3
	// 头条
	MODULE_TYPE_HEADLINE=4

	// 是否关联商户
	// 否
	RELATION_BUS_NO = 0
	// 是
	RELATION_BUS_YES = 1
)

func getStatusList() []int {
	return []int{
		STATUS_WAITING,
		STATUS_PASS,
		STATUS_NOT_PASS,
	}
}
func getModuleTypeList() []int {
	return []int{
		MODULE_TYPE_RAIDERS,
		MODULE_TYPE_EXPLORATION,
		MODULE_TYPE_QUIZ,
	}
}
func getRecommendList() []int {
	return []int{
		RECOMMEND_YES,
		RECOMMEND_NO,
	}
}
func getSoureList() []int {
	return []int{
		SOURE_USER,
		SOURE_ADMIN,
	}
}

// VerifyStatus 验证状态的有效性
func VerifyStatus(status int) bool {
	return functions.InArray(status, getStatusList())
}

// VerifyModuleType 验证模块类型的有效性
func VerifyModuleType(moduleType int) bool {
	return functions.InArray(moduleType, getModuleTypeList())
}

// VerifyRecommend 验证推荐的有效性
func VerifyRecommend(recommend int) bool {
	return functions.InArray(recommend, getRecommendList())
}

// VerifySoure 验证来源的有效性
func VerifySoure(soure int) bool {
	return functions.InArray(soure, getSoureList())
}

// RaidersBase 攻略base
type RaidersBase struct {
	Id         int    `mapstructure:id`
	Title      string `mapstructure:"title"`      // 攻略文章标题
	AuthorId   int    `mapstructure:"author_id"`  // 文章作者
	Recommend  int    `mapstructure:"raider_id"`  // 是否推荐 0-不推荐 1-推荐
	ShareDesc  string `mapstructure:"share_desc"` // 分享简介
	Body       string `mapstructure:"body"`       // 正文
	ImgId      int    `mapstructure:"img_id"`     // 封面图片ID,不能为空
	ImgHash    string `mapstructure:"img_hash"`   // 封面图片hash
	ImgUrl     string `mapstructure:"img_url"`
	StartTime  int64  `mapstructure:"start_time"`  // 起始时间
	EndTime    int64  `mapstructure:"end_time"`    // 结束时间
	IsRelation int    `mapstructure:"is_relation"` // 是否已关联商户
	BusBody    string `mapstructure:"bus_body"`    // 关联的商户正文
	Soure      int    `mapstructure:"soure"`       // 来源 0-后台 1-用户
	Ctime      int64  `mapstructure:"ctime"`       // 创建时间
	CtimeStr   string `mapstructure:"ctime_str"`   // 创建时间字符串
	Attention  bool   `mapstructure:"attention"`   // 是否收藏
}

// RaidersBusRelationBase 攻略关联的商户
type RaidersBusRelationBase struct {
	Id       int
	RaiderId int    `mapstructure:"raider_id"` // 攻略ID
	BusId    int    `mapstructure:"bus_id"`    // 商户ID
	ImgId    int    `mapstructure:"img_id"`    // 图片ID
	ImgUrl   string `mapstructure:"img_url"`
	ImgHash  string `mapstructure:"img_hash"` // 图片Hash
	ImgDesc  string `mapstructure:"img_desc"` // 图片内容
}

// RaidersArticleBase 攻略文章base
type RaidersArticleBase struct {
	Id          int    //
	RaiderId    int    `mapstructure:"raider_id"` // 攻略ID
	ImgId       int    `mapstructure:"img_id"`    // 文章图片id
	ImgUrl      string `mapstructure:"img_url"`
	ImgHash     string `mapstructure:"img_hash"`     // 文章图片hash
	ArticleDesc string `mapstructure:"article_desc"` // 文章内容
}

// ArgsAdminAddRaiders 添加攻略入参
type ArgsAdminAddRaiders struct {
	common.Autoken
	RaidersBase
	RaidersArticle     []RaidersArticleBase     // 攻略文章base
	RaidersBusRelation []RaidersBusRelationBase // 攻略关联的商户
}

// ReplyAdminAddRaiders 添加攻略返回参数
type ReplyAdminAddRaiders struct {
	RaiderId int
}

// ArgsAdminUpdateRaiders 更新攻略入参
type ArgsAdminUpdateRaiders struct {
	common.Autoken
	RaidersBase
	RaiderId           int                      // 攻略ID
	RaidersArticle     []RaidersArticleBase     // 攻略文章base
	RaidersBusRelation []RaidersBusRelationBase // 攻略关联的商户
}

// ArgsAdminRaidersList 后台攻略列表入参
type ArgsAdminRaidersList struct {
	common.Autoken
	common.Paging
	Soure     string // 来源 0-后台 1-用户
	Recommend string // 是否推荐 0-不推荐 1-推荐
}

// ReplyAdminRaidersList 后台攻略列表返回参数
type ReplyAdminRaidersList struct {
	TotalNum int //
	Lists     []RaidersListBase
}

// RaidersListBase RaidersListBase
type RaidersListBase struct {
	Id           int    `mapstructure:"id"`     // id
	Title        string `mapstructure:"title"`  // 文章标题
	Soure        int    `mapstructure:"soure"`  // 来源 0-后台 1-用户
	Author       string `mapstructure:"author"` // 作者
	AuditStatus  int    `mapstructure:"audit_status"`
	AuthorId     int    `mapstructure:"author_id"`
	AttentionNum int    `mapstructure:"attention_num"` // 搜藏量
	ViewsNum     int    `mapstructure:"views_num"`     // 浏览量
	CommentsNum  int    `mapstructure:"comments_num"`  // 评论量
	Recommend    int    `mapstructure:"recommend"`     // 是否推荐 0-不推荐 1-推荐
	IsRelation   int    `mapstructure:"is_relation"`   // 是否已关联商户
	Ctime        int64  `mapstructure:"ctime"`         // 创建时间
	CtimeStr     string `mapstructure:"ctime_str"`     // 创建时间字符串
}

// ArgsAdminRaiderInfo ArgsAdminRaiderInfo
type ArgsAdminRaiderInfo struct {
	RaiderId int
	Soure    int // 来源
}

// ReplyAdminRaiderInfo ReplyAdminRaiderInfo
type ReplyAdminRaiderInfo struct {
	RaidersBase
	Author string // 作者
	//RaiderId           int                      // 攻略ID
	RaidersArticle     []RaidersArticleBase     // 攻略文章base
	RaidersBusRelation []RaidersBusRelationBase // 攻略关联的商户
}

// ArgsAuditRaider 审核攻略
type ArgsAuditRaider struct {
	common.Autoken
	RaiderId int
	Status   int
}

// -----------------------------前台攻略rpc-------------------------------
// ArgsUserRaiderList 前台攻略列表入参
type ArgsUserRaiderList struct {
	common.Utoken
	common.Paging
	Uid int //内部传输用
}
type UserRaiderListBase struct {
	Id           int    `mapstructure:"id"`            // id
	Title        string `mapstructure:"title"`         // 文章标题
	AttentionNum int    `mapstructure:"attention_num"` // 搜藏量
	ViewsNum     int    `mapstructure:"views_num"`     // 浏览量
	CommentsNum  int    `mapstructure:"comments_num"`  // 评论量
	Attention    bool   `mapstructure:"attention"`     // 是否收藏
	ImgUrl       string       // 封面url
	ImgHash      string     // 封面hash
	ImgId        int    `mapstructure:"img_id"`
	Ctime        int64    `mapstructure:"ctime"`         // 创建时间
	CtimeStr     string      // 创建时间字符串
}

// ReplyUserRaiderList 前台攻略列表出参
type ReplyUserRaiderList struct {
	TotalNum int `mapstructure:"total_num"`
	Lists     []UserRaiderListBase
}

// ArgsUserRaiderInfo 前台攻略详情入参
type ArgsUserRaiderInfo struct {
	common.Utoken
	RaiderId int
}

// ReplyUserRaiderInfo 前台攻略详情出参
type ReplyUserRaiderInfo struct {
	RaidersBase
	AttentionNum   int                  `mapstructure:"attention_num"` // 搜藏量
	ViewsNum       int                  `mapstructure:"views_num"`     // 浏览量
	CommentsNum    int                  `mapstructure:"comments_num"`  // 评论量
	Author         string               // 作者
	AvatarUrl string //用户头像
	ShareLink string //分享链接
	RaidersArticle []RaidersArticleBase // 攻略文章base
	//Recommend   []UserRaiderListBase // 建议阅读
}

// ArgsRaiderRecommend 推荐阅读入参
type ArgsRaiderRecommend struct {
	common.Paging
	common.Utoken
	Count int // 展示量
	RaiderId   int
	ContainOwn bool // 如果不包含当前文章时,须将当前文章的id传递过来
}
// ReplyRaiderRecommend 推荐阅读出参
type ReplyRaiderRecommend struct {
	TotalNum int
	Lists []UserRaiderListBase
}

// ArgsUserAddRaider 前台用户添加攻略入参
type ArgsUserAddRaider struct {
	common.Utoken
	Title          string               // 攻略文章标题
	Body           string               // 正文
	ImgHash        string               // 封面图片ID,不能为空
	RaidersArticle []RaidersArticleBase // 攻略文章base
}

// ReplyUserAddRaider 前台用户添加攻略出参
type ReplyUserAddRaider struct {
	Id int
}

// ArgsGetShop 获取攻略关联的店铺入参
type ArgsGetShop struct {
	common.Paging
	RaiderId int
}

// ReplyGetShop 获取攻略关联的店铺出参
type ReplyGetShop struct {
	BusBody string `mapstructure:"bus_body"`
	TotalNum int
	Lists    []RaidersBusRelationBase
}
type Raiders interface {
	//	AdminAddRaiders 后台添加攻略
	AdminAddRaiders(ctx context.Context, args *ArgsAdminAddRaiders, reply *ReplyAdminAddRaiders) error
	// AdminUpdateRaiders 后台更新攻略
	AdminUpdateRaiders(ctx context.Context, args *ArgsAdminUpdateRaiders, reply *bool) error
	// AdminRaidersList 后台攻略列表
	AdminRaidersList(ctx context.Context, args *ArgsAdminRaidersList, reply *ReplyAdminRaidersList) error
	// AdminRaiderInfo　后台攻略详情
	AdminRaiderInfo(ctx context.Context, args *ArgsAdminRaiderInfo, reply *ReplyAdminRaiderInfo) error
	// AuditRaider 审核攻略
	AuditRaider(ctx context.Context, args *ArgsAuditRaider, reply *bool) error

	// UserAddRaider 前台用户添加攻略
	UserAddRaider(ctx context.Context, args *ArgsUserAddRaider, reply *ReplyUserAddRaider) error
	// UserRaidersList 前台台攻略列表
	UserRaidersList(ctx context.Context, args *ArgsUserRaiderList, reply *ReplyUserRaiderList) error
	//GetUserPublishRaiderList 用户发布的攻略
	GetUserPublishRaiderList(ctx context.Context,args *ArgsUserRaiderList,reply *ReplyUserRaiderList)error
	// UserRaiderInfo 前台攻略详情
	UserRaiderInfo(ctx context.Context, args *ArgsUserRaiderInfo, reply *ReplyUserRaiderInfo) error
	// RecommendList 推荐阅读
	RecommendList(ctx context.Context, args *ArgsRaiderRecommend, reply *ReplyRaiderRecommend) error
	// GetShop 获取攻略关联的店铺
	GetShop(ctx context.Context, args *ArgsGetShop, reply *ReplyGetShop) error
}
