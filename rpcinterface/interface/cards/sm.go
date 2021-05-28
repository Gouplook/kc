//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 11:40
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//套餐基本信息数据结构
type SmBase struct {
	Name                string  //名称
	BusID               int     `mapstructure:"bus_id"` //商户ID
	SortDesc            string  //短描述
	RealPrice           float64 //现价
	Price               float64 //标价
	Sales               int     //销售量
	ServicePeriod       int     //保险时间 月
	SaleShopNum         int     //在售门店数量
	ValidCount          int     `mapstructure:"validcount"` //总次数
	Ctime               int     //发布时间
	IsPermanentValidity int     // 是否永久有效：1-是；2-否

}

//套餐在列表的数据结构
type ListSm struct {
	SmId           int    //套餐id
	SmBase                //套餐基本信息
	CtimeStr       string //发布时间
	Clicks         int
	Sales          int
	IsGround       int //总店铺是否上架 0=否 1=是
	ShopStatus     int //套餐在子店的销售状态 1=下架 2=上架 3=被总店禁用 只有在门店才有效
	ShopHasAdd     int //子店是否添加 0=否 1=是 只有在门店才有效
	ShopItemId     int //项目在门店的id
	ShopDelStatus  int
	ImgId          int //图片id
	ApplySingleNum int //适用单项目的个数
	GiveSingleNum  int //赠送单项目的个数
}

//添加套餐入参
type ArgsAddSm struct {
	common.BsToken
	SmBase
	Notes          []CardNote       //温馨提示
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           //封面图片hansh串
}

//修改套餐入参
type ArgsEditSm struct {
	common.BsToken
	SmId int //套餐id
	SmBase
	Notes          []CardNote       //温馨提示
	IncludeSingles []IncSingle      //包含的单项目
	GiveSingles    []IncSingle      //赠送的单项目
	GiveSingleDesc []GiveSingleDesc //赠品描述
	ImgHash        string           //封面图片hansh串
}

//套餐详情入参
type ArgsSmInfo struct {
	SmId   int //套餐id
	ShopId int //门店id非必选，需要获取套餐在门店的详情时传递
}

//套餐详情返回数据
type ReplySmInfo struct {
	SmBase
	ShareLink      string             //分享链接
	Notes          []CardNote         //温馨提示
	SmId           int                //套餐id
	SsId           int                //套餐在门店的id
	ImgHash        string             //封面图片hansh串
	ImgUrl         string             //封面图片url
	IncludeSingles []IncSingleDetail2 //包含的单项目
	GiveSingles    []IncSingleDetail2 //赠送的单项目
	IsAllSingle    bool               //适用于全部单项目
	IsAllProduct   bool               //适用于全部商品
	GiveSingleDesc []GiveSingleDesc   //赠品描述
	IsGround       int                //总店铺是否上架 0=否 1=是
	ShopStatus     int                //套餐在子店的销售状态 1=下架 2=上架 3=被总店禁用
	ShopLists      []ReplyShopName    // 总店套餐门店添加信息
	BusInfo
	SingleTotalNum int //卡项包含项目的总次数
}

type ReplyShopName struct {
	ShopId     int
	ShopName   string
	BranchName string
}

//总店套餐列表入参
type ArgsBusSmPage struct {
	common.Paging
	BusId            int    //商家id
	ShopId           int    //门店id
	FilterShopHasAdd bool   //false-获取全部，true-过滤添加过的数据
	IsGround         string //状态过滤：默认全部，1=下架 2=上架
}

//套餐列表返回数据
type ReplySmPage struct {
	TotalNum  int            //套餐总数量
	List      []ListSm       //套餐列表
	IndexImgs map[int]string //套餐封面图
}

//设置适用门店
type ArgsSetSmShop struct {
	common.BsToken
	SmIds     []int //套餐ids
	ShopIds   []int //适用的门店ids
	IsAllShop bool  //是否适用所有门店 为true的情况下，ShopIds不用传也不生效
}

//总店-删除套餐卡
type ArgsDeleteSm struct {
	common.BsToken
	SmIds []int
}

//分店-删除套餐卡
type ArgsDeleteShopSm struct {
	common.BsToken
	SmIds []int
}

//总店上下架套餐入参
type ArgsDownUpSm struct {
	common.BsToken
	SmIds   []int //套餐ids
	OptType uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

//子店获取适用本店的套餐列表入参
type ArgsShopGetBusSmPage struct {
	common.Paging
	ShopId int //门店id
	BusId  int //企业id
}

//子店获取适用本店的套餐列表返回数据
type ReplyShopGetBusSmPage struct {
	ReplySmPage
}

//子店添加套餐到自己店铺入参
type ArgsShopAddSm struct {
	common.BsToken
	SmIds []int //套餐ids
}

//获取子店的套餐列表入参
type ArgsShopSmPage struct {
	common.Paging
	ShopId int //门店id
	Status int //套餐状态
}

//套餐在门店列表的数据结构
type ListShopSm struct {
	ShopSmId   int //套餐在门店的id
	ShopItemId int //项目在门店的id
	ListSm
}

//子店的套餐列表返回数据
type ReplyShopSmPage struct {
	TotalNum  int            //套餐总数量
	Lists     []ListShopSm   //套餐列表
	IndexImgs map[int]string //套餐封面图
}

//子店上下架套餐
type ArgsShopDownUpSm struct {
	common.BsToken
	ShopSmIds []int //套餐在门店的id @kc_shop_sm.id
	OptType   uint8 //操作类型 参考常量OPT_UP/OPT_DOWN
}

type Sm interface {
	//添加套餐
	AddSm(ctx context.Context, sm *ArgsAddSm, smId *int) error
	//编辑套餐
	EditSm(ctx context.Context, sm *ArgsEditSm, reply *bool) error
	//获取套餐的详情
	SmInfo(ctx context.Context, args *ArgsSmInfo, reply *ReplySmInfo) error
	//获取总店的套餐列表
	BusSmPage(ctx context.Context, args *ArgsBusSmPage, reply *ReplySmPage) error
	//设置适用门店
	SetSmShop(ctx context.Context, args *ArgsSetSmShop, reply *bool) error
	//总店上下架套餐
	DownUpSm(ctx context.Context, args *ArgsDownUpSm, reply *bool) error
	//子店获取适用本店的套餐列表
	ShopGetBusSmPage(ctx context.Context, args *ArgsShopGetBusSmPage, reply *ReplyShopGetBusSmPage) error
	//子店添加套餐到自己的店铺
	ShopAddSm(ctx context.Context, args *ArgsShopAddSm, reply *bool) error
	//获取子店的套餐列表
	ShopSmPage(ctx context.Context, args *ArgsShopSmPage, reply *ReplyShopSmPage) error
	//子店上下架套餐
	ShopDownUpSm(ctx context.Context, args *ArgsShopDownUpSm, reply *bool) error
	//删除总店套餐卡
	DeleteSm(ctx context.Context, args *ArgsDeleteSm, reply *bool) error
	//删除分店套餐卡
	DeleteShopSm(ctx context.Context, args *ArgsDeleteShopSm, reply *bool) error
	// 子店添加套餐（一期优化，去掉总店推送，改为门店自动拉取
	ShopAddToSm(ctx context.Context, args *ArgsShopAddSm, reply *bool) error
}
