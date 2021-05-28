//冲值卡服务接口定义
//@author yangzhiwu<578154898@qq.com>
//@date 2020/10/21 14:47
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//充值卡基本信息数据结构
type RcardBase struct {
	Name                string  //名称
	BusID               int     `mapstructure:"bus_id"` //商户ID
	SortDesc            string  //短描述
	RealPrice           float64 //现价（实际支付）
	Price               float64 //标价（面值）
	DiscountType        int     //折扣类型 //参考DISCOUNT_TYPE_price/DISCOUNT_TYPE_item 常量
	Discount            float64 //折扣
	IsDel               int     //是否删除 0=否，1=是
	DelTime             int     // 删除时间
	Sales               int     //销售量
	ServicePeriod       int     //保险时间 月
	SaleShopNum         int     //在售门店数量
	Ctime               int     //发布时间
	IsPermanentValidity int     `mapstructure:is_permanent_validity` // 是否永久有效：1-是；2-否
}

//发布充值卡入参结构体
type ArgsAddRcard struct {
	common.BsToken
	RcardBase
	Notes           []CardNote           //温馨提示
	RcardRule       []ListsRechargeRules //充值卡规则
	IsAllSingle     bool                 //适用于全部单项目
	IncludeSingles  []IncSingle          //包含的单项目
	IsHaveDiscount  int                  //是否享受折扣：1-无折扣；2-有折扣 (适用项目）
	SingleDiscount  float64              //项目折扣
	IsAllProduct    bool                 //适用于全部商品
	IncludeProducts []IncProduct         //包含的商品
	GiveSingles     []IncSingle          //赠送的单项目
	GiveSingleDesc  []GiveSingleDesc     //赠品描述
	ImgHash         string               //封面图片hansh串

}

//编辑充值卡入参
type ArgsEditRcard struct {
	common.BsToken
	RcardId int //冲值卡id
	RcardBase
	Notes           []CardNote           //温馨提示
	RcardRule       []ListsRechargeRules //充值卡规则
	IsAllSingle     bool                 //适用于全部单项目
	IncludeSingles  []IncSingle          //包含的单项目
	IsHaveDiscount  int                  //是否享受折扣：1-无折扣；2-有折扣 (适用项目）
	SingleDiscount  float64              //项目折扣
	IsAllProduct    bool                 //适用于全部商品
	IncludeProducts []IncProduct         //包含的商品
	GiveSingles     []IncSingle          //赠送的单项目
	GiveSingleDesc  []GiveSingleDesc     //赠品描述
	ImgHash         string               //封面图片hansh串
}

//充值卡详情入参
type ArgsRcardInfo struct {
	RcardId int //套餐id
	ShopId  int //门店id非必选，需要获取充值卡在门店的详情时传递
}

//充值卡详情返回数据
type ReplyRcardInfo struct {
	RcardId int //冲值卡id
	RcardBase
	ShareLink      string               //分享链接
	Notes          []CardNote           //温馨提示
	SsId           int                  //充值卡在门店的id
	ImgHash        string               //封面图片hansh串
	ImgUrl         string               //封面图片url
	IsAllSingle    bool                 //适用于全部单项目
	IsAllProduct   bool                 //适用于全部商品
	IncludeSingles []IncSingleDetail    //包含的单项目
	IsHaveDiscount int                  //是否享受折扣：1-无折扣；2-有折扣 (适用项目）
	SingleDiscount float64              //项目折扣
	GiveSingleDesc []GiveSingleDesc     //赠品描述
	IncProducts    []IncProductDetail   //包含的商品
	GiveSingles    []IncSingleDetail    //赠送的单项目
	RcardRules     []ListsRechargeRules //充值卡规则信息
	IsGround       int                  //总店铺是否上架 0=否 1=是
	ShopStatus     int                  //套餐在子店的销售状态 1=下架 2=上架 3=被总店禁用
	BusInfo
	ShopLists []ReplyShopName // 总店充值卡门店添加信息
}

//商家充值卡列表
type ArgsBusRcardPage struct {
	common.Paging
	BusId            int    //商家id
	ShopId           int    //门店id
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
	IsDel            int    // 状态过滤: 是否删除 0=否，1=是
}

//充值卡在列表的数据结构
type ListRcard struct {
	RcardId        int    //充值卡id
	RcardBase             //充值卡基本信息
	CtimeStr       string //发布时间
	Clicks         int
	Sales          int
	IsGround       int //总店铺是否上架 0=否 1=是
	ShopStatus     int //套餐在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopHasAdd     int //子店是否添加 0=否 1=是 只有在门店才有效
	ShopItemId     int //项目在门店的id
	ShopDelStatus  int
	ImgId          int  //图片id
	IsAllSingle    bool //是否适用于全部单项目
	ApplySingleNum int  //适用单项目的个数
	GiveSingleNum  int  //赠送单项目的个数
}

//总店充值卡列表返回数据
type ReplyRcardPage struct {
	TotalNum  int            //总数量
	Lists     []ListRcard    //列表
	IndexImgs map[int]string //封面图
}

//设置适用门店入参
type ArgsSetRcardShop struct {
	common.BsToken
	RcardIds  []int //充值卡ids
	ShopIds   []int //适用的门店ids
	IsAllShop bool  //是否适用所有门店 为true的情况下，ShopIds不用传也不生效
}

//总店上下架入参
type ArgsDownUpRcard struct {
	common.BsToken
	RcardIds []int //充值卡ids
	OptType  uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

// 总店删除充值卡入参
type ArgsDelRcard struct {
	common.BsToken
	ShopId   int   //内部使用
	RcardIds []int // 充值卡ids
	OptType  uint8 // 删除操作

}

//子店获取适用本店的充值卡列表入参
type ArgsShopGetBusRcardPage struct {
	common.Paging
	BusId  int
	ShopId int
}

//子店添加充值卡到自己店铺入参
type ArgsShopAddRcard struct {
	common.BsToken
	RcardIds      []int //充值卡ids
	BusId, ShopId int   //内部使用
}

//获取子店的充值卡列表入参
type ArgsShopRcardPage struct {
	common.Paging
	ShopId int //门店id
	Status int //充值卡状态
}

//子店的套餐列表返回数据
type ReplyShopRcardPage struct {
	TotalNum  int            //总数量
	Lists     []ListRcard    //列表
	IndexImgs map[int]string //封面图
}

//子店上下架套餐
type ArgsShopDownUpRcard struct {
	common.BsToken
	ShopId       int   //内部使用
	ShopRcardIds []int //充值卡在门店的id @kc_shop_rcard.id
	OptType      uint8 //操作类型
}

//获取充值卡基础数据入参
type ArgsGetRcardBaseInfo struct {
	RcardIds []int //充值卡id
	RuleId   int   //充值卡规则id
}

type RcardRulesBase struct {
	Id        int     //充值卡规则id
	RcardId   int     //充值卡id
	Price     float64 //实际充值金额
	GivePrice float64 //赠送金额
}
type GetRcardBaseInfoBase struct {
	RcardId       int
	Name          string  //名称
	BusID         int     `mapstructure:"bus_id"` //商户ID
	SortDesc      string  //短描述
	RealPrice     float64 //现价
	Price         float64 //标价
	DiscountType  int     //折扣类型 //参考DISCOUNT_TYPE_price/DISCOUNT_TYPE_item 常量
	Discount      float64 //折扣
	Sales         int     //销售量
	ServicePeriod int     //保险时间 月
	SaleShopNum   int     //在售门店数量
	Rules         []RcardRulesBase
}

//获取充值卡基础数据出参
type ReplyGetRcardBaseInfo struct {
	Lists []GetRcardBaseInfoBase
}

type ListsRechargeRules struct {
	Id             int
	RechargeAmount float64 //充值金额
	DonationAmount float64 //赠送金额
	Name           string  // 拼接字符串 如充500赠送200

}

// 新增管理充值规则 Recharge rules
type ArgsAddRechargeRules struct {
	common.BsToken
	RcardId int // 充值卡Id @充值卡表中的rcard_id
	ListsRechargeRules
}

// 编辑充值卡规则
type ArgsEditRechargeRules struct {
	common.BsToken
	Id      int // 充值卡规则Id
	RcardId int //充值卡Id @充值卡表中的rcard_id
	ListsRechargeRules
}

// 获取充值卡规则详情
type ArgsRechargeRulesInfo struct {
	Id int // 充值卡规则Id

}
type ReplyRechargeRulesInfo struct {
	RcardId int //充值卡Id @充值卡表中的rcard_id
	ListsRechargeRules
	RechargeAmount float64 //充值金额
	DonationAmount float64 //赠送金额
	IsDel          int
	DiscountType   int     //折扣类型 //参考DISCOUNT_TYPE_price/DISCOUNT_TYPE_item 常量
	Discount       float64 //折扣
}

// 删除充值卡规则
type ArgsDeleRechargeRules struct {
	common.BsToken
	Id int // 充值卡规则Id
}

// 获取充值规则列表
type ArgsRechargeRulesList struct {
	common.Paging
	RcardId int //充值卡id
}

// 返回充值规则数量
type ReplyRechargerRulesList struct {
	TotalNum     int
	Lists        []ListsRechargeRules
	DiscountType int     //折扣类型
	Discount     float64 //折扣
}

type Rcard interface {
	//添加充值卡
	AddRcard(ctx context.Context, args *ArgsAddRcard, rcardId *int) error
	//编辑充值卡
	EditRcard(ctx context.Context, args *ArgsEditRcard, reply *bool) error
	//获取冲值卡详情
	RcardInfo(ctx context.Context, args *ArgsRcardInfo, reply *ReplyRcardInfo) error
	//获取总店的充值卡列表
	BusRcardPage(ctx context.Context, args *ArgsBusRcardPage, reply *ReplyRcardPage) error
	//设置适用门店
	SetRcardShop(ctx context.Context, args *ArgsSetRcardShop, reply *bool) error
	//总店上下架充值卡
	DownUpRcard(ctx context.Context, args *ArgsDownUpRcard, reply *bool) error
	// 总店删除充值卡
	DeleteRcard(ctx context.Context, args *ArgsDelRcard, reply *bool) error
	//子店获取适用本店的充值卡列表
	ShopGetBusRcardPage(ctx context.Context, args *ArgsShopGetBusRcardPage, reply *ReplyRcardPage) error
	//子店添加充值卡到自己的店铺
	ShopAddRcard(ctx context.Context, args *ArgsShopAddRcard, reply *bool) error
	//获取子店的充值卡列表
	ShopRcardPage(ctx context.Context, args *ArgsShopRcardPage, reply *ReplyShopRcardPage) error
	//子店上下架套餐
	ShopDownUpRcard(ctx context.Context, args *ArgsShopDownUpRcard, reply *bool) error
	// 子店删除充值卡
	ShopDeleteRcard(ctx context.Context, args *ArgsDelRcard, reply *bool) error

	//获取充值卡基础数据
	GetRcardBaseInfo(ctx context.Context, args *ArgsGetRcardBaseInfo, reply *ReplyGetRcardBaseInfo) error

	// 总店新增充值规则
	AddRechargeRules(ctx context.Context, args *ArgsAddRechargeRules, id *int) error
	// 充值卡规则编辑
	EditRechargeRules(ctx context.Context, args *ArgsEditRechargeRules, reply *bool) error
	// 获取充值卡规则详情
	RechargeRulesInfo(ctx context.Context, args *ArgsRechargeRulesInfo, reply *ReplyRechargeRulesInfo) error
	// 删除充值卡规则
	DeleRechargeRules(ctx context.Context, args *ArgsDeleRechargeRules, reply *bool) error
	// 获取充值规则列表
	BusRechargeRulesList(ctx context.Context, args *ArgsRechargeRulesList, reply *ReplyRechargerRulesList) error
}
