package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//空的入参
type EmptyParams struct {
}

//空的出参
type EmptyReplies struct {
}

//限次卡基本信息数据结构
type NCardBase struct {
	Name                string  `mapstructure:"name"`           //名称
	BusID               int     `mapstructure:"bus_id"`         //商户ID
	ShortDesc           string  `mapstructure:"sort_desc"`      //短描述short
	RealPrice           float64 `mapstructure:"real_price"`     //现价
	Price               float64 `mapstructure:"price"`          //标价
	ServicePeriod       int     `mapstructure:"service_period"` //保险时间 月
	SaleShopNum         int     `mapstructure:"sale_shop_num"`  //在售门店数量
	ImgID               int     `mapstructure:"img_id"`         //图片ID
	Sales               int     `mapstructure:"sales"`          //销量
	Ctime               int     `mapstructure:"ctime"`          //发布时间
	ValidCount          int     `mapstructure:"validcount"`     //包含单项目总次数
	CtimeStr            string  // create time 字符串格式
	IsPermanentValidity int     `mapstructure:is_permanent_validity` // 是否永久有效：1-是；2-否
}

//添加限次卡入参
type ArgsAddNCard struct {
	common.BsToken
	NCardBase
	Notes          []CardNote       //温馨提示
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           //封面图片hash串
}

//添加限次卡出参
type RepliesAddNCard struct {
	NCardID int
}

//修改限次卡入参
type ArgsEditNCard struct {
	common.BsToken
	NCardBase
	Notes          []CardNote       //温馨提示
	CardID         int              `mapstructure:"card_id"` //限次卡ID
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           `mapstructure:"img_hash"` //封面图片hansh串
}

//限次卡详情入参
type ArgsNCardInfo struct {
	NCardID int `mapstructure:"card_id"` //限次卡ID
	ShopID  int `mapstructure:"shop_id"` //门店ID非必选，需要获取限次卡在门店的详情时传递
}

type BusInfo struct {
	BusIcon        string //bus Icon url
	BusCompanyName string //bus name
	BusBrandName   string //bus name
}

//限次卡详情返回数据
type ReplyNCardInfo struct {
	NCardBase
	ShareLink string     //分享链接
	Notes     []CardNote //温馨提示
	NCardID   int        `mapstructure:"ncard_id"` //限次卡ID
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
	ShopStatus     int                //限次卡在子店的销售状态 1=下架 2=上架 3=被总店禁用
	BusInfo
	SingleTotalNum int             //卡项包含项目的总次数
	ShopLists      []ReplyShopName // 总店限次卡门店添加信息
}

//总店限次卡列表入参
type ArgsBusNCardPage struct {
	common.Paging
	common.BsToken
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
}

//限次卡在列表的数据结构
type NCardDesc struct {
	NCardBase     //限次卡基本信息
	NCardID   int `mapstructure:"ncard_id"` //限次卡ID
	BindID    int `mapstructure:"bind_id"`  //商家主营行业ID

	Clicks         int
	Sales          int `mapstructure:"sales"`
	IsGround       int `mapstructure:"is_ground"` //总店铺是否上架 0=否 1=是
	ShopStatus     int //限次卡在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopHasAdd     int //子店是否添加 0=否 1=是 只有在门店才有效
	ShopDelStatus  int //在店铺的删除状态
	ShopItemId     int //项目在门店的id
	ApplySingleNum int //适用单项目的个数
	GiveSingleNum  int //赠送单项目的个数
}

//限次卡列表返回数据
type ReplyNCardPage struct {
	TotalNum int            //限次卡总数量
	List     []NCardDesc    //限次卡列表
	IndexImg map[int]string //限次卡封面图
}

//设置适用门店
type ArgsSetNCardShop struct {
	common.BsToken
	NCardIDs  []int `mapstructure:"card_ids"` //限次卡IDs
	ShopIDs   []int `mapstructure:"shop_ids"` //适用的门店IDs
	IsAllShop bool  `mapstructure:"all_shop"` //是否适用所有门店 为true的情况下，ShopIDs不用传也不生效
}

//总店上下架限次卡入参
type ArgsDownUpNCard struct {
	common.BsToken
	NCardIDs []int `mapstructure:"card_ids"` //限次卡IDs
	OptType  uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

//子店获取适用本店的限次卡列表入参
type ArgsShopGetBusNCardPage struct {
	common.Paging
	common.BsToken
}

//子店添加限次卡到自己店铺入参
type ArgsShopAddNCard struct {
	common.BsToken
	NCardIDs []int `mapstructure:"card_ids"` //限次卡IDs
}

//获取子店的限次卡列表入参
type ArgsShopNCardPage struct {
	common.Paging
	ShopID int `mapstructure:"shop_id"` //门店ID
	Status int //限次卡上架状态
}

//子店上下架限次卡
type ArgsShopDownUpNCard struct {
	common.BsToken
	CardIDs []int
	OptType uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

type ArgsShopNcardRpc struct {
	ShopId   int
	NcardIds []int
}
type ReplyShopNcardRpc struct {
	List []NCardDesc //限次卡列表
}

//总店-删除
type ArgsDeleteNCard struct {
	common.BsToken
	NcardIds []int
}

//分店-删除
type ArgsDeleteShopNCard struct {
	common.BsToken
	NcardIds []int
}

type NCard interface {
	//添加限次卡
	AddNCard(ctx context.Context, args *ArgsAddNCard, replies *RepliesAddNCard) error
	//编辑限次卡
	EditNCard(ctx context.Context, args *ArgsEditNCard, replies *EmptyReplies) error
	//获取限次卡的详情
	NCardInfo(ctx context.Context, args *ArgsNCardInfo, reply *ReplyNCardInfo) error
	//获取总店的限次卡列表
	BusNCardPage(ctx context.Context, args *ArgsBusNCardPage, reply *ReplyNCardPage) error
	//设置适用门店
	SetNCardShop(ctx context.Context, args *ArgsSetNCardShop, reply *EmptyReplies) error
	//总店上下架限次卡
	DownUpNCard(ctx context.Context, args *ArgsDownUpNCard, reply *EmptyReplies) error
	//子店获取适用本店的限次卡列表
	ShopGetBusNCardPage(ctx context.Context, args *ArgsShopGetBusNCardPage, reply *ReplyNCardPage) error
	//子店添加限次卡到自己的店铺
	ShopAddNCard(ctx context.Context, args *ArgsShopAddNCard, reply *EmptyReplies) error
	//获取子店的限次卡列表
	ShopNCardPage(ctx context.Context, args *ArgsShopNCardPage, reply *ReplyNCardPage) error
	//子店上下架限次卡
	ShopDownUpNCard(ctx context.Context, args *ArgsShopDownUpNCard, reply *EmptyReplies) error
	//ShopNcardRpc
	ShopNcardRpc(ctx context.Context, args *ArgsShopNcardRpc, reply *ReplyShopNcardRpc) error
	//总店-软删除
	DeleteNCard(ctx context.Context, args *ArgsDeleteNCard, reply *bool) error
	//分店-软删除
	DeleteShopNCard(ctx context.Context, args *ArgsDeleteShopNCard, reply *bool) error
}
