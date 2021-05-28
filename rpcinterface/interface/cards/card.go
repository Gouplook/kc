package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//综合卡基本信息数据结构
type CardBase struct {
	Name                string  `mapstructure:"name"`           //名称
	BusID               int     `mapstructure:"bus_id"`         //商户ID
	ShortDesc           string  `mapstructure:"sort_desc"`      //短描述short
	RealPrice           float64 `mapstructure:"real_price"`     //现价
	Price               float64 `mapstructure:"price"`          //标价
	ServicePeriod       int     `mapstructure:"service_period"` //保险时间 月
	SaleShopNum         int     `mapstructure:"sale_shop_num"`  //在售门店数量
	Sales               int     `mapstructure:"sales"`          //销量
	ImgID               int     `mapstructure:"img_id"`         //图片ID
	Ctime               int     `mapstructure:"ctime"`          //发布时间
	CtimeStr            string  //create time 字符串格式
	IsPermanentValidity int     // 是否永久有效：1-是；2-否
}

//添加综合卡入参
type ArgsAddCard struct {
	common.BsToken
	CardBase
	Notes           []CardNote       //温馨提示
	IsAllSingle     bool             //包含全部单项目
	IncludeSingles  []IncInfSingle   //包含的单项目
	IsAllProduct    bool             //包含全部商品
	IncludeProducts []IncProduct     //包含的商品
	GiveSingles     []IncSingle      //赠送的单项目
	GiveSingleDesc  []GiveSingleDesc //赠品描述
	ImgHash         string           `mapstructure:"img_hash"` //封面图片hash串
}

//添加综合卡出参
type RepliesAddCard struct {
	CardID int
}

//修改综合卡入参
type ArgsEditCard struct {
	common.BsToken
	CardID int `mapstructure:"card_id"` //综合卡ID
	CardBase
	Notes           []CardNote       //温馨提示
	IsAllSingle     bool             //包含全部单项目
	IncludeSingles  []IncInfSingle   //包含的单项目
	IsAllProduct    bool             //包含全部商品
	IncludeProducts []IncProduct     //包含的商品
	GiveSingles     []IncSingle      //赠送的单项目
	GiveSingleDesc  []GiveSingleDesc //赠品描述
	ImgHash         string           `mapstructure:"img_hash"` //封面图片hash串
}

//综合卡详情入参
type ArgsCardInfo struct {
	CardID int `mapstructure:"card_id"` //综合卡ID
	ShopID int `mapstructure:"shop_id"` //门店ID非必选，需要获取综合卡在门店的详情时传递
}

//综合卡详情返回数据
type ReplyCardInfo struct {
	CardBase
	ShareLink string     //分享链接
	Notes     []CardNote //温馨提示
	CardID    int        `mapstructure:"card_id"` //综合卡ID
	BindID    int        `mapstructure:"bind_id"` //商家主营行业ID
	SsId      int        //在门店的id

	ImgHash        string               //封面图片hash串
	ImgUrl         string               //封面图片url
	IncludeSingles []IncInfSingleDetail //包含的单项目
	IncProducts    []IncProductDetail   //包含的商品
	IsAllSingle    bool                 //适用于全部单项目
	IsAllProduct   bool                 //适用于全部商品
	GiveSingles    []IncSingleDetail    //赠送的单项目
	GiveSingleDesc []GiveSingleDesc     //赠品描述
	IsGround       int                  `mapstructure:"is_ground"` //总店铺是否上架 0=否 1=是
	ShopStatus     int                  //综合卡在子店的销售状态 1=下架 2=上架 3=被总店禁用
	BusInfo
	SingleTotalNum int             //卡项包含项目的总次数
	ShopLists      []ReplyShopName // 总店综合卡门店添加信息
	//Discount       float64         //折扣率 7.5 表示75折
}

//总店综合卡列表入参
type ArgsBusCardPage struct {
	common.Paging
	common.BsToken
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
}

//综合卡在列表的数据结构
type CardDesc struct {
	CardBase     //综合卡基本信息
	CardID   int `mapstructure:"card_id"` //综合卡ID
	BindID   int `mapstructure:"bind_id"` //商家主营行业ID

	Clicks         int
	Sales          int  `mapstructure:"sales"`
	IsGround       int  `mapstructure:"is_ground"` //总店铺是否上架 0=否 1=是
	ShopStatus     int  `mapstructure:"status"`    //综合卡在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopDelStatus  int  //在店铺的删除状态
	ShopHasAdd     int  //子店是否添加 0=否 1=是 只有在门店才有效
	ShopItemId     int  //项目在门店的id
	IsAllSingle    bool //是否适用于全部单项目
	ApplySingleNum int  //适用单项目的个数
	GiveSingleNum  int  //赠送单项目的个数
}

//综合卡列表返回数据
type ReplyCardPage struct {
	TotalNum int            //综合卡总数量
	List     []CardDesc     //综合卡列表
	IndexImg map[int]string //综合卡封面图
}

//设置适用门店
type ArgsSetCardShop struct {
	common.BsToken
	CardIDs   []int `mapstructure:"card_ids"` //综合卡IDs
	ShopIDs   []int `mapstructure:"shop_ids"` //适用的门店IDs
	IsAllShop bool  `mapstructure:"all_shop"` //是否适用所有门店 为true的情况下，ShopIDs不用传也不生效
}

//总店上下架综合卡入参
type ArgsDownUpCard struct {
	common.BsToken
	CardIDs []int `mapstructure:"card_ids"` //综合卡IDs
	OptType uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

//子店获取适用本店的综合卡列表入参
type ArgsShopGetBusCardPage struct {
	common.Paging
	common.BsToken
}

//子店添加综合卡到自己店铺入参
type ArgsShopAddCard struct {
	common.BsToken
	CardIDs []int `mapstructure:"card_ids"` //综合卡IDs
}

//获取子店的综合卡列表入参
type ArgsShopCardPage struct {
	common.Paging
	ShopID   int    `mapstructure:"shop_id"`   //门店ID
	ShopCall bool   `mapstructure:"shop_call"` // 门店调用
	Status   string `mapstructure:"status"`    // 门店status
}

//子店上下架综合卡
type ArgsShopDownUpCard struct {
	common.BsToken
	CardIDs []int
	OptType uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

// ArgsDeleteCard 总店删除综合卡
type ArgsDeleteCard struct {
	common.BsToken
	CardIds []int `mapstructure:"card_ids"` // 综合卡id
}

// ArgsDeleteShopCard 门店删除综合卡
type ArgsDeleteShopCard struct {
	common.BsToken
	CardIds []int `mapstructure:"card_ids"` // 综合卡id
}

type ArgsCardsInfo struct {
	common.BsToken
	CardIds []int // 卡ids
}

type ReplyCardsInfo struct {
	CardID         int                `mapstructure:"card_id"`    //综合卡ID
	Name           string             `mapstructure:"name"`       //名称
	RealPrice      float64            `mapstructure:"real_price"` //现价
	Price          float64            `mapstructure:"price"`      //标价
	IncludeSingles []IncSingleDetail  //包含的单项目
	IncProducts    []IncProductDetail //包含的商品
	GiveSingles    []IncSingleDetail  //赠送的单项目
}

type ArgsShopCardListRpc struct {
	ShopId  int
	CardIds []int
}
type ReplyShopCardListRpc struct {
	List []CardDesc //综合卡列表
}
type ArgsAllCardsNum struct {
	BusId int
}
type ReplyAllCardsNum struct {
	AllCardsNum int // 所发布的卡的总数量
}

type Card interface {
	//添加综合卡
	AddCard(ctx context.Context, args *ArgsAddCard, replies *RepliesAddCard) error
	//编辑综合卡
	EditCard(ctx context.Context, args *ArgsEditCard, replies *EmptyReplies) error
	// DeleteCard 总店删除综合卡
	DeleteCard(ctx context.Context, args *ArgsDeleteCard, reply *bool) error
	//获取综合卡的详情
	CardInfo(ctx context.Context, args *ArgsCardInfo, reply *ReplyCardInfo) error
	//获取总店的综合卡列表
	BusCardPage(ctx context.Context, args *ArgsBusCardPage, reply *ReplyCardPage) error
	//设置适用门店
	SetCardShop(ctx context.Context, args *ArgsSetCardShop, reply *EmptyReplies) error
	//总店上下架综合卡
	DownUpCard(ctx context.Context, args *ArgsDownUpCard, reply *EmptyReplies) error
	//子店获取适用本店的综合卡列表
	ShopGetBusCardPage(ctx context.Context, args *ArgsShopGetBusCardPage, reply *ReplyCardPage) error
	//子店添加综合卡到自己的店铺
	ShopAddCard(ctx context.Context, args *ArgsShopAddCard, reply *EmptyReplies) error
	//获取子店的综合卡列表
	ShopCardPage(ctx context.Context, args *ArgsShopCardPage, reply *ReplyCardPage) error
	//子店上下架综合卡
	ShopDownUpCard(ctx context.Context, args *ArgsShopDownUpCard, reply *EmptyReplies) error
	// 获取门限可售综合卡详情
	CardsInfo(ctx context.Context, args *ArgsCardsInfo, reply *map[int]*ReplyCardsInfo) error
	// 门店综合卡数据rpc内部调用
	ShopCardListRpc(ctx context.Context, args *ArgsShopCardListRpc, reply *ReplyShopCardListRpc) error

	// 获取所有卡的发布数量 rpc内部调用
	GetAllCardsNum(ctx context.Context, args *ArgsAllCardsNum, reply *ReplyAllCardsNum) error
	//delete shop card 分店删除卡
	DeleteShopCard(ctx context.Context, args *ArgsDeleteShopCard, reply *bool) error
}
