//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/09 17:07
package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/file"
	"git.900sui.cn/kc/rpcinterface/interface/order"
)

//九百岁查询附近筛选条件
const (
	//项目类型
	ITEM_TYPE_single = 1 //单项目
	ITEM_TYPE_sm     = 2 //套餐
	ITEM_TYPE_card   = 3 //综合卡
	ITEM_TYPE_hcard  = 4 //限时卡
	ITEM_TYPE_ncard  = 5 //限次卡
	ITEM_TYPE_hncard = 6 //限时限次卡
	ITEM_TYPE_rcard  = 7 //充值卡
	ITEM_TYPE_icard  = 8 //身份卡

	//排序
	SORT_DISTANCE        = 0 //距离优先
	SORT_SALES_HIGH      = 1 //销量最高
	SORT_PRICE_LOW       = 2 //价格优先
	SORT_EVALUATION_BEST = 3 //评价最好

	//活动
	ACTIVITY_HAVA_RED_ENVELOPE  = 1 //有红包
	ACTIVITY_AI_UNITY           = 2 //ai团
	ACTIVITY_STAGING            = 3 //分期
	ACTIVITY_CAN_BE_BOOKED      = 4 //可预约
	ACTIVITY_MLC                = 5 //mlc
	ACTIVITY_ASSOCIATION_MEMBER = 6 //协会会员

	//价格
	PRICE_BELOW_50       = 1 //50以下
	PRICE_BETEEN_51_100  = 2 //51-100
	PRICE_BETEEN_101_200 = 3 //101-200
	PRICE_BETEEN_201_300 = 4 //201-300
	PRICE_OVER_300       = 5 //300以上

	//安全码
	SAFECODE_GRAY   = 1 //灰色
	SAFECODE_RED    = 2 //红色
	SAFECODE_YELLOW = 3 //黄色
	SAFECODE_GREEN  = 4 //绿色
	SAFECODE_BLACK  = 5 //黑色
)

var CardsScreen = map[int]string{
	//卡项
	ITEM_TYPE_single: "单项目",
	ITEM_TYPE_sm:     "套餐",
	ITEM_TYPE_card:   "综合卡",
	ITEM_TYPE_hcard:  "限时卡",
	ITEM_TYPE_ncard:  "限次卡",
	ITEM_TYPE_hncard: "限时限次卡",
	ITEM_TYPE_icard:  "身份卡",
	ITEM_TYPE_rcard:  "充值卡",
}
var SortScreen = map[int]string{
	//排序
	SORT_DISTANCE:        "距离优先",
	SORT_SALES_HIGH:      "销量最高",
	SORT_PRICE_LOW:       "价格优先",
	SORT_EVALUATION_BEST: "评价最好",
}
var ActivityScreen = map[int]string{
	//活动
	ACTIVITY_HAVA_RED_ENVELOPE:  "有红包",
	ACTIVITY_AI_UNITY:           "ai团",
	ACTIVITY_STAGING:            "分期",
	ACTIVITY_CAN_BE_BOOKED:      "可预约",
	ACTIVITY_MLC:                "mlc",
	ACTIVITY_ASSOCIATION_MEMBER: "协会会员",
}
var PriceScreen = map[int]string{
	//价格
	PRICE_BELOW_50:       "50以下",
	PRICE_BETEEN_51_100:  "51-100",
	PRICE_BETEEN_101_200: "101-200",
	PRICE_BETEEN_201_300: "201-300",
	PRICE_OVER_300:       "300以上",
}

var SortSearch = map[int]string{
	SORT_SALES_HIGH:      "Sales",
	SORT_PRICE_LOW:       "AvgPrice",
	SORT_EVALUATION_BEST: "CommentScore",
}

var SafeCodeSearch = map[int]string{
	//安全码
	SAFECODE_GRAY:   "灰码",
	SAFECODE_RED:    "红码",
	SAFECODE_YELLOW: "黄码",
	SAFECODE_GREEN:  "绿码",
	SAFECODE_BLACK:  "黑码",
}

type ArgsItemInfo4Es struct {
	ShopId   int
	ItemId   int
	ItemType int //项目类型
}

type ReplyItemInfo4Es struct {
	ItemName       string  //项目名称
	ItemShopPrice  float64 //项目再门店的价格
	ItemShopStatus int     //项目在门店的状态
}

type ArgsItemList struct {
	common.Paging       //分页
	CallType      int   // 调用类型:1-服务列表;2-附近
	IndustryId    int   //领域id
	BindId        []int //行业id
	Pid           int
	Cid           int     //城市id
	Did           int     //区id
	DistrictId    []int   //商圈id
	FundMode      int     //资金管理方式 0=未知 1=存管 2=保险
	IsCredit      int     //是否信用评级 0=否 1=是
	Active        []int   //参与活动类型 1=有红包 2=ai团 3=分期 4=可预约 5=mlc 6=协会会员
	Keywords      string  //搜索关键字
	ItemType      int     //项目类型
	Start         int     //列表开始位置
	Limit         int     //获取条数
	Lat           float64 //维度
	Lng           float64 //经度
	Price         int     // 价格范围
	Order         int     //排序 1=智能排序 2=销量最高 3=价格优先 4=评价最好 5=离我最近 6=最新门店 7=人均高到低 8=人均低到高
	Distance      float64 //距离，默认单位：米
}

type ShopItemList struct {
	ShopName   string  //门店名称
	BranchName string  //分店名称
	ShopPic    string  //门店照
	Flag       []int   //门店flag 1=有红包 2=ai团 3=分期 4=可预约 5=mlc 6=协会会员 7=有保险 8=有信用评估
	distance   float64 //距离
}

type ReplyItemList struct {
	Lists    []ShopItemList
	TotalNum int
}

type ArgsGetItemsBySsids struct {
	SsIds    []int //卡项在门店的ids
	ItemType int   //卡项类型 2=套餐
}

type ItemSingle struct {
	SingleId         int     //单项目id
	Num              int     //数量
	Name             string  //名称
	Discount         float64 //折扣
	SspId            int
	SpecNames        string
	PeriodOfValidity int //有效期，单位天
}

type ItemGoods struct {
	GoodsId  int     //产品id
	Name     string  //产品名称
	Price    float64 // 产品价格
	Discount float64 //折扣
}
type ItemBaseStruct struct {
	ItemId     int     //卡项id
	ItemName   string  //卡项名称
	ItemPrice  float64 //单项目价格
	ShopItemId int     //单项目在门店的id

	Price         float64 //面值
	RealPrice     float64 //真实售价
	ServicePeriod int     //保险周期 月
	Status        int     //卡项在门店的销售状态
}
type ShopInfoItemSinglesBase struct {
	ItemId     int     //项目ID(单项目,综合卡....)
	Name       string  //名称
	Sales      int     // 销量
	RealPrice  float64 //现价
	Price      float64 //标价
	MinPrice   float64 //最低价
	MaxPrice   float64 //最高价
	ShopStatus int     //在当前门店的状态 1=下架 2=上架 3=被总店禁用
	Ssid       int     // 套餐在门店的主键id
}
type ItemBase struct {
	ItemName      string  //卡项名称
	ShopId        int     //门店id
	ItemId        int     //卡项id
	SsId          int     //卡项在门店的id
	ImgId         int     //封面图片id
	Price         float64 //面值
	RealPrice     float64 //真实售价
	ServicePeriod int     //保险周期 月
	ValidCount    int     // 包含单项目的总次数
	ShopSales     int     // 门店销量
	ShopRealPrice float64 //门店真实售价
	DiscountType  int     //充值卡折扣类型
	Discount      float64 //充值卡折扣率
	Status        int     //卡项在门店的销售状态
	CableShopIds  []int   //适用门店ids

	Gives          []ItemSingle     //赠送的服务
	GiveSingleDesc []GiveSingleDesc //赠品描述
	Singles        []ItemSingle     //包含的服务
	Goods          []ItemGoods      //包含的产品
}

type SsId int

//查询九百岁卡项服务
type ArgsAppInfos struct {
	common.Paging
	Id         int     // 1单项目 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7充值卡 8身份卡
	Cid        int     //市id
	Did        int     //区id
	Lon        float64 //经度
	Lat        float64 //纬度
	Flag       int     //卡包进1   附近进0
	DistrictId int     //商圈Id
	BindId     int     //行业id
	IndustryId int     //领域id
	SortId     int     //排序规则  0智能排序 1销量最高 2价格优先 3评价最好
	Credit     int     //征信评级  默认开 0开  1关
	Insure     int     //投保商户 默认开 0开  1关
	Activity   string  //活动 默认0不参加   1AI团 2分期 3可预约 4有红包 5mic 多个逗号隔开
	KeyWord    string  //搜索关键字
	Num        int     //要查询多少条门店的项目 默认2条
	SafeCode   int     //安全码
	PriceSel   int     //价格 默认0不选择  1.50以下  2.5-100 3.101-200 4.201-300 5.300以上
}

type ReplyAppInfo struct {
	TotalNum int
	Lists    []AppInfo
}

type AppInfo struct {
	//门店基本信息
	BusId         int            //商户信息
	ShopId        int            //门店id
	ShopName      string         //门店名称
	BranchName    string         //分店名
	ShopAddress   string         //门店地址
	ShopImage     string         //门店图片
	ShopPhone     string         //门店电话
	IndustryId    int            //领域id
	MainBindId    int            //主行业id
	BindId        []int          //行业id集合
	BindNames     string         //行业名称
	Pid           int            //省id
	Cid           int            //市id
	Did           int            //区id
	DName         string         //区名称
	DistrictId    []int          //商圈id
	DistrictNames string         //商圈名称
	Lon           float64        //经度
	Lat           float64        //纬度
	FundMode      int            //资金管理方式 0=未知 1=存管 2=保险
	IsCredit      int            //是否评级 0=否 1=是
	Ctime         string         //分店审核通过时间
	Sales         int            //总销量
	CommentScore  float64        //平均评分
	Active        map[int]string //参与的活动，键id 值name
	AvgPrice      float64        //平均消费价格
	Distance      int            //距离 m
	SafeCode      int            //商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色

	//单项目信息
	Sub []SubItem
}

//单项目信息
type SubItem struct {
	ShopItemId      int
	ItemId          int
	ItemName        string
	ItemPrice       float64
	ServicePeriod   int     //保险周期，限时月份
	ServiceTime     int     //服务时长
	TotalNum        int     //总次数
	Price           float64 //标价
	ServiceDiscount float64 //服务折扣
	ProductDiscount float64 //商品折扣
}

//查询项目详情入参
type ArgsShopList struct {
	Lng      float64
	Lat      float64
	ItemType int // 1单项目 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7身份卡 8充值卡
	ItemId   int
}

/*type ReplyShopInfos struct {
	TotalNum int
	Lists []ReplyShopInfo
}

//根据ItemId查询门店信息
type ReplyShopInfo struct {
	//es
	ShopId     int       //门店id
	ShopName   string    //门店门店
	BranchName string    //分店名字
	BindIds    string    //多个行业id
	MainBindId  int   //主行业id
	Lng  float64  //经度
	Lat  float64  //纬度
	Distance float64   //距离
	BusinessHours string  //营业时间
	Contact       string  //负责人姓名
	CompanyName   string  //分店工商营业执照名称
	Status        int     //分店状态 0=待审核 1=审核失败 2=审核通过 3=已下架
	IndustryId    int     //分店所属领域

	//mysql
	IndusNames    string //多个行业名称
	Address string         //地址
	Phone   string         //手机号
	ImgPath string         //门店照
}*/

type ArgsGetItemsByShopId struct {
	common.Paging
	ItemType int // 类型
	ShopId   int
	OrderBy  string // 排序字段
}

type ReplyGetItemsByShopIdBase struct {
	ItemType      int     //类型:1单项目 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7充值卡 8身份卡
	ItemId        int     //单项目id
	Name          string  //名称
	ImgId         int     //封面图片id
	RealPrice     float64 //售价
	Price         float64 //标价
	Sales         int     // 销量
	ShopId        int     //门店id
	BusId         int
	SsId          int    //卡项在门店的id
	Discount      string //充值卡、身份卡折扣
	ServicePeriod int    // 服务周期
	ValidCount    int    // 包含单项目的总次数
}

type ReplyGetItemsByShopId struct {
	TotalNum    int
	Lists       []ReplyGetItemsByShopIdBase
	IndexImg    map[int]file.ReplyFileInfo
	DefaultImgs map[int]file.ReplyFileInfo
}

//活动 筛选返回
type ReplyActivityScreen struct {
	Id   int
	Name string
}

//价格 筛选返回
type ReplyPriceScreen struct {
	Id   int
	Name string
}

//卡项 筛选返回
type ReplyCardsScreen struct {
	Id   int
	Name string
}

//排序规格 筛选返回
type ReplySortScreen struct {
	Id   int
	Name string
}

type Args struct{}

//门店详情-推荐的单项目(上架的)入参
type ArgsGetRecommendSingles struct {
	TotalNum int // 展示数量，默认三条
	ShopId   int
}

//门店详情-推荐的单项目(上架的)出参
type ReplyGetRecommendSingles struct {
	Lists    []ReplyGetItemsByShopIdBase
	IndexImg map[int]file.ReplyFileInfo
}

//卡项收藏入参
type ArgsCollectItems struct {
	common.Utoken
	Uid      int //内部使用
	ItemType int // 1单项目 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7充值卡 8身份卡
	ItemId   int // 卡项id
	SsId     int //
}

//卡项收藏状态入参数
type ArgsCollectStatus struct {
	common.Utoken
	Uid      int //内部使用
	ItemType int // 1单项目 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7充值卡 8身份卡
	ItemId   int // 卡项id
	SsId     int //

}

//返回卡项收藏状态
type ReplyCollectStatus struct {
	CollectStatus int // 0:卡项未收藏   1：卡项服务已收藏
}

//获取收藏的卡项入参
type ArgsGetCollectItems struct {
	common.Utoken
	common.Paging
	Uid      int //内部使用
	ItemType int
	ShopId   int
}

//获取收藏的卡项出参
type ReplyGetCollectItems struct {
	TotalNum int
	Lists    []ReplyGetItemsByShopIdBase
	IndexImg map[int]file.ReplyFileInfo
}

//获取当月下架的卡项-出参
type ReplyGetItemAllXCradsNumByShopId struct {
	BusId                int
	DateTime             int64
	AllItemXCardNum      int //所有的卡项
	AllUnderItemXCardNum int //本月所有下架的卡项
}

//获取卡项下的商品信息--入参
type ArgsGetCardProductsSinglesInfo struct {
	common.Paging
	ItemId   int
	ItemType int // 2套餐 3综合卡 4限时卡 5限次卡 6限时限次卡 7身份卡 8充值卡
	ShopId   int

	ShopStatus string//单项目在门店的状态 1=下架 2=上架
	IsDel string//是否删除：0-否，1-是
}

//获取卡项下的商品信息--出参
type ReplyGetCardProductsInfo struct {
	Lists    []IncProductDetail2 //包含的商品
	TotalNum int
}

//预付卡包含的单项目-出参
type ReplyGetItemIncludeSingles struct {
	TotalNum int
	Lists    []IncSingleDetail2 //包含的单项目
}

//预付卡赠送的单项目-出参
type ReplyGetItemGiveSingles struct {
	TotalNum int
	Lists []IncSingleDetail2 //赠送的单项目
}

type ArgsGetItemDefaultImgs struct {
	ImgType  int //图片类型：0-全部；1-小图，2-大图
	ItemType int //卡项类型
}
type ReplyGetItemDefaultImgs struct {
	CardsDefaultPics      map[int]string
	CardsSmallDefaultPics map[int]string
}

//一些通用接口
type Item interface {
	//为es的shop_items文档，获取项目数据
	ItemInfo4Es(ctx context.Context, args *ArgsItemInfo4Es, reply *ReplyItemInfo4Es) error
	//根据筛选条件，获取项目列表
	ItemList(ctx context.Context, args *ArgsItemList, reply *ReplyItemList) error
	//根据卡项在门店的ids，获取卡项数据
	GetItemsBySsids(ctx context.Context, args *ArgsGetItemsBySsids, reply *map[SsId]ItemBase) error

	//根据条件查询九百岁
	GetInfos(ctx context.Context, args *ArgsAppInfos, reply *ReplyAppInfo) error
	//根据单项目id查询所有适用门店
	GetDetailById(ctx context.Context, args *ArgsShopList, reply *[]order.ReplyCableShopInfo) error

	// 门店拥有的项目-上架的项目
	GetItemsByShopId(ctx context.Context, args *ArgsGetItemsByShopId, reply *ReplyGetItemsByShopId) error

	//门店详情-推荐的单项目
	GetRecommendSingles(ctx context.Context, args *ArgsGetRecommendSingles, reply *ReplyGetRecommendSingles) error
	//卡项收藏入参
	CollectItems(ctx context.Context, args *ArgsCollectItems, reply *bool) error
	//获取用户收藏的卡项入参
	GetCollectItems(ctx context.Context, args *ArgsGetCollectItems, reply *ReplyGetCollectItems) error
	//卡项收藏状态
	GetCollectStatus(ctx context.Context, args *ArgsCollectStatus, reply *ReplyCollectStatus) error
	//获取当月下架的卡项
	GetItemAllXCradsNumByShopId(ctx context.Context, shopId *int, reply *ReplyGetItemAllXCradsNumByShopId) error

	//购买成功，设置卡项的销量
	IncrItemSales(ctx context.Context, orderSn *string, reply *bool) error
	//获取卡项下的商品信息
	GetCardProductsInfo(ctx context.Context, args *ArgsGetCardProductsSinglesInfo, reply *ReplyGetCardProductsInfo) error
	//预付卡包含的单项目
	GetItemIncludeSingles(ctx context.Context, args *ArgsGetCardProductsSinglesInfo, reply *ReplyGetItemIncludeSingles) error
	//预付卡赠送的单项目
	GetItemGiveSingles(ctx context.Context, args *ArgsGetCardProductsSinglesInfo, reply *ReplyGetItemGiveSingles) error
	//预付卡默认图片
	GetItemDefaultImgs(ctx context.Context, args *ArgsGetItemDefaultImgs, reply *ReplyGetItemDefaultImgs) error
}
