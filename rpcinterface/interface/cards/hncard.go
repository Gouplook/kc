package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//限时限次卡基本信息数据结构
type HNCardBase struct {
	Name                string  `mapstructure:"name"`           //名称
	BusID               int     `mapstructure:"bus_id"`         //商户ID
	ShortDesc           string  `mapstructure:"sort_desc"`      //短描述short
	RealPrice           float64 `mapstructure:"real_price"`     //现价
	Price               float64 `mapstructure:"price"`          //标价
	ServicePeriod       int     `mapstructure:"service_period"` //保险时间 月
	SaleShopNum         int     `mapstructure:"sale_shop_num"`  //在售门店数量
	IsDel               int     `mapstructure:"is_del"`         //删除状态：0-否，1-是
	DelTime             int     `mapstructure:"del_time"`       //删除时间
	ImgID               int     `mapstructure:"img_id"`         //图片ID
	Sales               int     `mapstructure:"sales"`          //销量
	Ctime               int     `mapstructure:"ctime"`          //发布时间
	ValidCount          int     `mapstructure:"validcount"`     //包含单项目总次数
	CtimeStr            string  //create time  字符串格式
	IsPermanentValidity int     `mapstructure:is_permanent_validity` // 是否永久有效：1-是；2-否
}

//添加限时限次卡入参
type ArgsAddHNCard struct {
	common.BsToken
	HNCardBase
	Notes          []CardNote       //温馨提示
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           `mapstructure:"img_hash"` //封面图片hash串
}

//添加限时限次卡出参
type RepliesAddHNCard struct {
	HNCardID int
}

//修改限时限次卡入参
type ArgsEditHNCard struct {
	common.BsToken
	HNCardBase
	Notes          []CardNote       //温馨提示
	CardID         int              `mapstructure:"card_id"` //限时限次卡ID
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           `mapstructure:"img_hash"` //封面图片hansh串
}

//限时限次卡详情入参
type ArgsHNCardInfo struct {
	HNCardID int `mapstructure:"card_id"` //限时限次卡ID
	ShopID   int `mapstructure:"shop_id"` //门店ID非必选，需要获取限时限次卡在门店的详情时传递
}

//限时限次卡详情返回数据
type ReplyHNCardInfo struct {
	HNCardBase
	ShareLink string     //分享链接
	Notes     []CardNote //温馨提示
	HNCardID  int        `mapstructure:"hncard_id"` //限时限次卡ID
	SsId      int        //在门店的id
	BindID    int        `mapstructure:"bind_id"` //商家主营行业ID

	ImgHash        string             //封面图片hash串
	ImgUrl         string             //封面图片url
	IncludeSingles []IncSingleDetail2 //包含的单项目
	GiveSingles    []IncSingleDetail2 //赠送的单项目
	IsAllSingle    bool               //适用于全部单项目
	IsAllProduct   bool               //适用于全部商品
	GiveSingleDesc []GiveSingleDesc   //赠品描述
	IsGround       int                `mapstructure:"is_ground"` //总店铺是否上架 0=否 1=是
	ShopStatus     int                //限时限次卡在子店的销售状态 1=下架 2=上架 3=被总店禁用
	BusInfo
	SingleTotalNum int             //卡项包含项目的总次数
	ShopLists      []ReplyShopName // 总店限时限次卡门店添加信息
}

//总店限时限次卡列表入参
type ArgsBusHNCardPage struct {
	common.Paging
	common.BsToken
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
}

//限时限次卡在列表的数据结构
type HNCardDesc struct {
	HNCardBase     //限时限次卡基本信息
	HNCardID   int `mapstructure:"hncard_id"` //限时限次卡ID
	BindID     int `mapstructure:"bind_id"`   //商家主营行业ID

	Clicks         int
	Sales          int `mapstructure:"sales"`
	IsGround       int `mapstructure:"is_ground"` //总店铺是否上架 0=否 1=是
	ShopStatus     int //限时限次卡在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopHasAdd     int //子店是否添加 0=否 1=是 只有在门店才有效
	ShopDelStatus  int //在店铺的删除状态
	ShopItemId     int //项目在门店的id
	ApplySingleNum int //适用单项目的个数
	GiveSingleNum  int //赠送单项目的个数
}

//限时限次卡列表返回数据
type ReplyHNCardPage struct {
	TotalNum int            //限时限次卡总数量
	List     []HNCardDesc   //限时限次卡列表
	IndexImg map[int]string //限时限次卡封面图
}

//设置适用门店
type ArgsSetHNCardShop struct {
	common.BsToken
	HNCardIDs []int `mapstructure:"card_ids"` //限时限次卡IDs
	ShopIDs   []int `mapstructure:"shop_ids"` //适用的门店IDs
	IsAllShop bool  `mapstructure:"all_shop"` //是否适用所有门店 为true的情况下，ShopIDs不用传也不生效
}

//总店上下架限时限次卡入参
type ArgsDownUpHNCard struct {
	common.BsToken
	HNCardIDs []int `mapstructure:"card_ids"` //限时限次卡IDs
	OptType   uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

//子店获取适用本店的限时限次卡列表入参
type ArgsShopGetBusHNCardPage struct {
	common.Paging
	common.BsToken
}

//子店添加限时限次卡到自己店铺入参
type ArgsShopAddHNCard struct {
	common.BsToken
	HNCardIDs []int `mapstructure:"card_ids"` //限时限次卡IDs
}

//获取子店的限时限次卡列表入参
type ArgsShopHNCardPage struct {
	common.Paging
	ShopID int `mapstructure:"shop_id"` //门店ID
	Status int //卡项上下架状态
}

//子店上下架限时限次卡
type ArgsShopDownUpHNCard struct {
	common.BsToken
	CardIDs []int
	OptType uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

type ArgsShopHncardRpc struct {
	ShopId    int
	HNCardIds []int
}

type ReplyShopHncardRpc struct {
	List []HNCardDesc //限时限次卡列表
}

// 总店删除限时限次卡入参
type ArgsDelHNCard struct {
	common.BsToken
	HNCardIds []int // 限时限次IDs
	OptType   uint8 // 删除操作

}

type HNCard interface {
	//添加限时限次卡
	AddHNCard(ctx context.Context, args *ArgsAddHNCard, replies *RepliesAddHNCard) error
	//编辑限时限次卡
	EditHNCard(ctx context.Context, args *ArgsEditHNCard, replies *EmptyReplies) error
	//获取限时限次卡的详情
	HNCardInfo(ctx context.Context, args *ArgsHNCardInfo, reply *ReplyHNCardInfo) error
	//获取总店的限时限次卡列表
	BusHNCardPage(ctx context.Context, args *ArgsBusHNCardPage, reply *ReplyHNCardPage) error
	//设置适用门店
	SetHNCardShop(ctx context.Context, args *ArgsSetHNCardShop, reply *EmptyReplies) error
	//总店上下架限时限次卡
	DownUpHNCard(ctx context.Context, args *ArgsDownUpHNCard, reply *EmptyReplies) error
	//子店获取适用本店的限时限次卡列表
	ShopGetBusHNCardPage(ctx context.Context, args *ArgsShopGetBusHNCardPage, reply *ReplyHNCardPage) error
	//子店添加限时限次卡到自己的店铺
	ShopAddHNCard(ctx context.Context, args *ArgsShopAddHNCard, reply *EmptyReplies) error
	//获取子店的限时限次卡列表
	ShopHNCardPage(ctx context.Context, args *ArgsShopHNCardPage, reply *ReplyHNCardPage) error
	//子店上下架限时限次卡
	ShopDownUpHNCard(ctx context.Context, args *ArgsShopDownUpHNCard, reply *EmptyReplies) error
	//ShopHncardRpc
	ShopHncardRpc(ctx context.Context, args *ArgsShopHncardRpc, reply *ReplyShopHncardRpc) error

	// 总店删除限时限次卡
	DeleteHNCard(ctx context.Context, args *ArgsDelHNCard, reply *bool) error
	// 子店删除限时限次卡
	ShopDeleteHNCard(ctx context.Context, args *ArgsDelHNCard, reply *bool) error
}
