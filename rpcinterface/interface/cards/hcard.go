package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// HcardBase 限时卡基本信息数据结构
type HcardBase struct {
	BindID        int     `mapstructure:"bind_id"`        //商家主营行业ID
	BusID         int     `mapstructure:"bus_id"`         //商户ID
	Name          string  `mapstructure:"name"`           //名称
	SortDesc      string  `mapstructure:"sort_desc"`      //短描述short
	RealPrice     float64 `mapstructure:"real_price"`     //现价
	Price         float64 `mapstructure:"price"`          //标价
	Sales         int     `mapstructure:"sales"`          //销量
	ServicePeriod int     `mapstructure:"service_period"` //保险时间 月
	SaleShopNum   int     `mapstructure:"sale_shop_num"`  //在售门店数量
	Ctime         int     `mapstructure:"ctime"`          //发布时间
	IsDel         int     `mapstructure:"is_del"`         //删除状态：0-否，1-是
	DelTime       int     `mapstructure:"del_time"`       //删除时间
	CtimeStr      string  `mapstructure:"ctime_str"`      //create time 字符串格式
	ImgID         int     `mapstructure:"img_id"`         //图片ID
}

// HcardSingle 限时卡包含的单项目
type HcardSingle struct {
	SingleID         int `mapstructure:"single_id"`          // 单项目ID
	Num              int `mapstructure:"num"`                // 单项目次数
	PeriodOfValidity int `mapstructure:"period_of_validity"` //有效期，单位天(只有赠送项目有效)
}

// ArgsAddHcard 添加限时卡入参
type ArgsAddHcard struct {
	common.BsToken
	HcardBase
	Notes          []CardNote       //温馨提示
	IncludeSingles []IncInfSingle   // 包含的单项目
	GiveSingles    []HcardSingle    // 赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           `mapstructure:"img_hash"` // 封面图片hansh串
}

// ReplyAddHcard 添加限时卡出参
type ReplyAddHcard struct {
	HcardID int `mapstructure:"hcard_id"`
}

// ArgsEditHcard 修改限时卡入参
type ArgsEditHcard struct {
	common.BsToken
	HcardBase
	Notes          []CardNote       //温馨提示
	HcardID        int              `mapstructure:"hcard_id"` // 限时卡id
	IncludeSingles []IncInfSingle   // 包含的单项目
	GiveSingles    []HcardSingle    // 赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           `mapstructure:"img_hash"` // 封面图片hansh串
}

// ArgsHcardInfo 限时卡详情入参
type ArgsHcardInfo struct {
	HcardID int `mapstructure:"hcard_id"` // 限时卡id
	ShopID  int `mapstructure:"shop_id"`  // 门店id,需要获取限次卡在门店的详情时传递
}

// ReplyHcardInfo 限时卡详情出参
type ReplyHcardInfo struct {
	HcardBase
	ShareLink      string               //分享链接
	Notes          []CardNote           //温馨提示
	HcardID        int                  `mapstructure:"hcard_id"` // 限时卡id
	SsId           int                  //在门店的id
	ImgHash        string               `mapstructure:"img_hash"` // 封面图片hansh串
	ImgUrl         string               `mapstructure:"img_url"`  // 封面图片url
	IsAllSingle    bool                 //适用于全部单项目
	IsAllProduct   bool                 //适用于全部商品
	IncludeSingles []IncInfSingleDetail // 包含的单项目
	GiveSingles    []IncSingleDetail    // 赠送的单项目
	GiveSingleDesc []GiveSingleDesc     //赠品描述
	IsGround       int                  `mapstructure:"is_ground"`   // 总店铺是否上架 0=否 1=是
	ShopStatus     int                  `mapstructure:"shop_status"` // 限时卡在子店的销售状态 1=下架 2=上架 3=被总店禁用
	BusInfo
	ShopLists []ReplyShopName // 总店限时卡门店添加信息

}

// HcardDesc 限次卡在列表的数据结构
type HcardDesc struct {
	HcardBase          // 限次卡基本信息
	HcardID        int `mapstructure:"hcard_id"`  // 限次卡ID
	Clicks         int `mapstructure:"clicks"`    // 点击量
	Sales          int `mapstructure:"sales"`     // 销量
	IsGround       int `mapstructure:"is_ground"` // 总店铺是否上架 0=否 1=是
	ShopStatus     int // 限次卡在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopHasAdd     int // 子店是否添加 0=否 1=是 只有在门店才有效
	ShopDelStatus  int //在店铺的删除状态
	ShopItemId     int //项目在门店的id
	ApplySingleNum int  //适用单项目的个数
	GiveSingleNum  int  //赠送单项目的个数
}

// ArgsBusHcardPage 总店限时卡列表入参
type ArgsBusHcardPage struct {
	common.Paging
	BusID            int    `mapstructure:"bus_id"` // 商家ID
	ShopId           int    //门店id
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
}

// ReplyHcardPage 限时卡列表返回数据
type ReplyHcardPage struct {
	TotalNum int            `mapstructure:"total_num"` // 限时卡总数量
	List     []HcardDesc    `mapstructure:"list"`      // 限时卡列表
	IndexImg map[int]string `mapstructure:"index_img"` // 限时卡封面图
}

// ArgsSetHcardShop 设置限时卡适用门店入参
type ArgsSetHcardShop struct {
	common.BsToken
	HcardIDs  []int `mapstructure:"hcard_ids"`   // 限时卡ids
	ShopIDs   []int `mapstructure:"shop_ids"`    // 适用的门店ids
	IsAllShop bool  `mapstructure:"is_all_shop"` // 是否适用所有门店 为true的情况下，ShopIds不用传也不生效
}

// ArgsDownUpHcard 总店上下架限时卡入参
type ArgsDownUpHcard struct {
	common.BsToken
	HcardIDs []int `mapstructure:"hcard_ids"` // 限时卡ids
	OptType  uint8 `mapstructure:"opt_type"`  // 操作类型 参考常量OPT_UP/OPT_DOWN
}

// ArgsShopGetBusHcardPage 子店获取适用本店的限时卡列表入参
type ArgsShopGetBusHcardPage struct {
	common.Paging
	ShopID int `mapstructure:"shop_id"` // 门店ID
	BusID  int `mapstructure:"bus_id"`  // 企业ID
}

// ReplyShopGetBusHcardPage 子店获取适用本店的限时卡列表出参
type ReplyShopGetBusHcardPage struct {
	ReplyHcardPage
}

// ArgsShopAddHcard 子店添加总部限时卡到自己的店铺入参
type ArgsShopAddHcard struct {
	common.BsToken
	HcardIDs []int `mapstructure:"hcard_ids"` // 限时卡IDs
}

// ArgsShopDownUpHcard 子店上下架自己店铺中的限时卡入参
type ArgsShopDownUpHcard struct {
	common.BsToken
	ShopHcardIDs []int `mapstructure:"shop_hcard_ids"` // 限时卡在门店的ID
	OptType      uint8 `mapstructure:"opt_type"`       // 操作类型 参考常量OPT_UP/OPT_DOWN
}

// ArgsShopHcardPage 获取子店的限时卡列表入参
type ArgsShopHcardPage struct {
	common.Paging
	ShopID   int    `mapstructure:"shop_id"`   //门店ID
	ShopCall bool   `mapstructure:"shop_call"` // 门店调用
	Status   string `mapstructure:"status"`    // 门店状态
}

// ArgsDeleteHcard 总店删除限时卡
type ArgsDeleteHcard struct {
	common.BsToken
	HcardIds []int `mapstructure:"hcard_ids"`
}

type ArgsShopHcardListRpc struct {
	ShopId   int
	HcardIds []int
}
type ReplyShopHcardListRpc struct {
	List []HcardDesc
}

// Hcard 限时卡
type Hcard interface {
	// AddHcard 总店添加限时卡
	AddHcard(ctx context.Context, args *ArgsAddHcard, reply *ReplyAddHcard) error
	// EditHcard 总店修改限时卡
	EditHcard(ctx context.Context, args *ArgsEditHcard, reply *bool) error
	// DeleteHcard 总店删除限时卡
	DeleteHcard(ctx context.Context, args *ArgsDeleteHcard, reply *bool) error
	// HcardInfo 获取限时卡详情(总店和分店共用)
	HcardInfo(ctx context.Context, args *ArgsHcardInfo, reply *ReplyHcardInfo) error
	// BusHcardPage 获取总店限时卡列表
	BusHcardPage(ctx context.Context, args *ArgsBusHcardPage, reply *ReplyHcardPage) error
	// SetHcardShop 设置总店限时卡的适用门店
	SetHcardShop(ctx context.Context, args *ArgsSetHcardShop, reply *bool) error
	// DownUpHcard 总店上下架限时卡
	DownUpHcard(ctx context.Context, args *ArgsDownUpHcard, reply *bool) error
	// ShopGetBusHcardPage 子店获取适用本店的限时卡列表(总店分配给自己店铺的限时卡)
	ShopGetBusHcardPage(ctx context.Context, args *ArgsShopGetBusHcardPage, reply *ReplyShopGetBusHcardPage) error
	// ShopHcardPage 子店限时卡列表
	ShopHcardPage(cxt context.Context, args *ArgsShopHcardPage, reply *ReplyHcardPage) error
	// ShopAddHcard 子店添加总部限时卡到自己的店铺
	ShopAddHcard(ctx context.Context, args *ArgsShopAddHcard, reply *bool) error
	// ShopDownUpHcard 子店上下架自己店铺中的限时卡
	ShopDownUpHcard(ctx context.Context, args *ArgsShopDownUpHcard, reply *bool) error
	//子店删除限时卡
	ShopDeleteHcard(ctx context.Context, args *ArgsDeleteHcard, reply *bool) error
	// ShopHcardListRpc 子店现时卡列表rpc内部调用
	ShopHcardListRpc(ctx context.Context, args *ArgsShopCardListRpc, reply *ReplyShopHcardListRpc) error
}
