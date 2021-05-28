// @author yangzhiwu<578154898@qq.com>
// @date 20200403 13:28：11

package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/file"
)

// 单项目不同规格组合的价格
type SingleSpecPrice struct {
	SpecIds []int   // 规格组合id
	Price   float64 // 对应的价格
	SspId   int     // kc_single_spec_price表的主键id(规格组合id)
}

// 规格
type SingleSpec struct {
	SpecId int    // 规格id
	Name   string // 名称
}

// 单项目包含的规格ids
type SingleSpecIds struct {
	SpecId int
	Sub    []int
}

// 当项目基本信息
type SingleBase struct {
	Name           string  // 单项目名称
	BusId          int     // 商家id
	Sales          int     // 销售量
	SortDesc       string  // 单项目短描述
	BindId         int     // 所属分类id
	TagIds         []int   // 多个标签id
	RealPrice      float64 // 现价
	Price          float64 // 标价
	MinPrice       float64 // 最低价格
	MaxPrice       float64 // 最高价格
	ServiceTime    int     // 服务时长 单位：分钟
	Sex            int     // 适用性别 1:不限 2=女 3=男
	AgeBracket     []int   // 适用年龄ids AgeBracket
	TailorIndus    int     // 私人定制主类id
	TailorSubIndus []int   // 私人定制子类ids
	Subscribe      string  // 预约信息

	// Reminder      map[string]interface{} // 温馨提示
	// SingleContext map[string]interface{} //单项目内容

	Reminder      []ReminderInfo      // 温馨提示
	SingleContext []SingleContextInfo // 单项目内容
	ImgHash       string              // 封面图片hansh串
	Pictures      []string            // 相册图片
	SpecIds       []SingleSpecIds     // 包含的规格id
	SpecPrices    []SingleSpecPrice   // 不同规格组合的价格
	EffectImgs    []string            // 服务效果的图片hash串
	ToolsImgs     []string            // 仪器设备的图片哈市串

}

// 温馨提示
type ReminderInfo struct {
	ReminderName    string
	ReminderContent string
}

// 单项目内容
type SingleContextInfo struct {
	Content string
}

type Img struct {
	Hash string
	Url  string
}

// 单项目图片数据
type SingleImg struct {
	IndexPic   Img   // 封面图片地址
	AlbumPics  []Img // 相册图片
	EffectPics []Img // 服务效果图片
	ToolsPics  []Img // 仪器设备图片
}

// 单项目的会员优惠信息
type MemberLev struct {
	Level     int     // 等级
	LevelName string  // 等级名称
	Rebate    float64 // 折扣 一位小数 如8.5折
}

// 单项目详情数据
type SingleDetail struct {
	SingleId int // 项目id
	SingleBase
	SingleImg                        // 图片数据
	ShareLink     string             // 分享链接
	SexStr        string             // 适用性别名称
	AgeBracketStr []string           // 适用年龄名称
	TailorStr     string             // 订制分类名称
	TailorSubStr  []string           // 订制分类子项
	Specs         map[int]SingleSpec // 规格
	TagsStr       []TagInfo          // 标签名称
	SsId          int                // 单项目在门店的id@kc_shop_single.ss_id
	OnSale        int                // 是否销售中 0=不是 1=是
	MemberLev     MemberLev          // 会员折扣信息
	IsShop        int                // 是否适用所有门店  0= 否  1=是
}

// 单项目信息
type ArgsAddSingle struct {
	common.BsToken
	SingleBase
}

// 修改单项目使用的结构体
type ArgsEditSingle struct {
	common.BsToken
	SingleBase
	SingleId int // 项目id
}

// 获取单项目详情的入参
type ArgsGetSingleInfo struct {
	SingleId      int
	ShopId        int // 门店id  非必填
	Uid           int // 用户uid saas和康享包使用，非必填
	common.Utoken     // 前端当前登录用户的登录信息 ， 非必传
}

// 子店添加单项入参
type ArgsShopAddSingle struct {
	common.BsToken
	SingleId []int // 项目id
}

// 获取商家的单项目入参
type ArgsBusSinglePage struct {
	common.Paging
	BusId            int
	ShopId           int
	FilterShopHasAdd bool   // false-获取全部，true-过滤添加过的数据
	IsGround         string // 状态过滤：默认全部，1=下架 2=上架
	IsDel            string // 是否被删除 ""-全部，"0"-否，"1"-是
}

// 列表里面单项目展示的数据
type ListSingle struct {
	SingleId      int     // 单项目id
	Name          string  // 名称
	ImgId         int     // 封面图片id
	RealPrice     float64 // 现价
	Price         float64 // 标价
	MinPrice      float64 // 最低价
	MaxPrice      float64 // 最高价
	Sales         int     // 销量
	Click         int     // 点击数量
	SaleShopNum   int     // 在售门店数量
	Ctime         int     // 发布时间
	CtimeStr      string  // 发布时间字符串
	BindId        int     // 行业
	TagIds        string  // 标签
	TagIdsArr     []int   // 标签id数组
	ServiceTime   int     // 服务时长
	IsGround      int     // 在总店的是否上架
	IsDel         int     // 在总店删除状态
	ShopStatus    int     // 在子店铺的是否销售状态
	ShopDelStatus int     // 在店铺的删除状态
	ShopHasAdd    int     // 子店是否添加 0=否 1=是 只有在门店才有效
	ShopItemId    int     // 项目在门店的id
	SsId          int     // 在子店的id
	HasSpec       int     // 是否有多规格
}

// 获取商家的单项目返回数据
type ReplyBusSinglePage struct {
	TotalNum  int          // 总条数
	List      []ListSingle // 单项目数组
	TagNames  map[int]string
	IndexImgs map[int]string
}

// 子店修改价格入参
type ArgsShopChangePrice struct {
	common.BsToken
	SingleId  int               // 单项目id
	RealPrice float64           // 售价
	SpecPrice []SingleSpecPrice // 规格价格
	Name      string            // 单项目名称
}

// 总店上下架单项目
type ArgsDownUpSingle struct {
	common.BsToken
	SingleIds []int // 单项目ids
	OptType   uint8 // 操作类型 参考常量OPT_UP/OPT_DOWN
}

// 总店或者分店 删除单项目入参
type ArgsDelSingle struct {
	common.BsToken
	SingleIds []int // 单项目ids
}

// 获取商家的单项目入参
type ArgsShopSinglePage struct {
	common.Paging
	ShopId    int
	SingleIds []int
	Status    string // 状态 ""-全部，"1"=下架 ，"2"=上架， "3"=被总店禁用
	IsDel     string // 是否被删除 ""-全部，"0"-否，"1"-是
}

// 获取门店的单项目返回数据
type ReplyShopSinglePage struct {
	TotalNum  int          // 总条数
	List      []ListSingle // 单项目数组
	TagNames  map[int]string
	IndexImgs map[int]string
}

// 子店上下架单项目
type ArgsShopDownUpSingle struct {
	common.BsToken
	SsIds   []int // 单项目在子店的ids
	OptType uint8 // 操作类型 参考常量OPT_UP/OPT_DOWN
}

// 获取单项目门店价格入参
type ArgsGetShopSinglePrice struct {
	SsId  int // 单项目在子店的id
	SspId int // 单项目具体规格组合id
}

// 属性标签结构体
type Attr struct {
	Id   int //
	Name string
}

// 获取单项的所有属性标签返回数据
type ReplyGetAttrs struct {
	SexBracket []Attr         // 适用性别
	AgeBracket []Attr         // 适用年龄
	Tailor     []Attr         // 私人定制主类
	TailorSub  map[int][]Attr // 私人定制主类
}

// 单项目基础价格信息
type SinglePriceInfo struct {
	SingleId  int     // 单项目id
	Name      string  // 名称
	ImgId     int     // 封面图片
	Price     float64 // 标价
	RealPrice float64 // 现价
	MinPrice  float64 // 最低价
	MaxPrice  float64 // 最高价
}

// 简单的单项目返回信息
type SimpleSingleInfo struct {
	SingleId int    // 单项目id
	Name     string // 单项目名称
	HasSpec  int    // 是否有特殊规格
}

// SpecPricesBase 子规格组合价格
type SpecPricesBase struct {
	SspId    int     `mapstructure:"ssp_id"`    // 单项目规格ID
	SingleId int     `mapstructure:"single_id"` // 单项目ID
	SpecIds  string  `mapstructure:"spec_ids"`  // 单项目规格组合
	Price    float64 `mapstructure:"price"`     // 售价
	Selected bool    // 是否选中
}

// SignleStaBase SignleStaBase
type SignleStaBase struct {
	SingleId  int     // 单项目ID
	BindId    int     // 分类ID
	MaxPrice  float64 `mapstructure:"max_price"` // 商品价格()
	MinPrice  float64 `mapstructure:"min_price"`
	RealPrice float64 `mapstructure:"real_price"`
	Name      string  // 单项目名称
	Status    int
	ImgId     int `mapstructure:"img_id"`
	ImgUrl    string
	Ctime     int64 `mapstructure:"ctime"`
	CtimeStr  string
}

// ArgsGetSignlesByStaffID ArgsGetSignlesByStaffID
type ArgsGetSignlesByStaffID struct {
	// common.BsToken
	common.Paging
	ShopId    int
	First     bool
	GetAll    bool
	BindId    int
	StaffId   int
	Name      string
	SspIds    []int // 单项目规格IDs
	SingleIds []int // 单项目IDs
}

// ReplyGetSignlesByStaffID ReplyGetSignlesByStaffID
type ReplyGetSignlesByStaffID struct {
	TotalNum      int
	StaffId       int
	AssignService bool               // 是否指定服务
	List          []SignleStaBase    // 单项目
	SpecPrices    []SpecPricesBase   // 单项目子规格组合价格
	SpecIds       []SingleSpecIds    // 规格包含的子规格
	Specs         map[int]SingleSpec // 每个规格ID和name
}

type ArgsGetShopSingleBySingleIdsRpc struct {
	ShopId    int
	SingleIds []int
}

type ReplyGetShopSingleBySingleIdsRpc struct {
	List []SimpleSingleInfo
}

// 获取门店的单项目入参
type ArgsShopSingleByPage struct {
	common.Paging
	ShopId    int
	SingleIds []int
	Status    string // 状态 ""-全部，"1"=下架 ，"2"=上架， "3"=被总店禁用
	IsDel     string // 删除状态 ""-全部，"0"-否，"1"-是
}

// 获取门店的单项目返回
type ReplyShopSingle struct {
	TotalNum  int
	Lists     *[]SingleInfo
	SingleImg map[int]file.ReplyFileInfo
}
type SingleInfo struct {
	SingleId      int
	ImgId         int
	Name          string
	RealPrice     float64
	Price         float64
	ServiceTime   int
	MinPrice      string
	MaxPrice      string
	SsId          int // 单项目在门店的id
	ShopStatus    int // 项目在门店上下架的状态：1=下架 2=上架
	ShopDelStatus int // 项目在门店是否被删除：0-否，1-是
}

// 根据单项目在门店的id获取所有子项目入参
type ArgsSubServer struct {
	SingleId int
	ShopId   int
}

type ReplySubServer2 struct {
	SubServer *[]SubServer
	Specs     *[]Specs
	// SpecNameMap *map[string]interface{}
}

// 根据单项目id获取子规格服务返回
type ReplySubServer struct {
	SubServer []SubServer
	Specs     []Specs
}
type Specs struct {
	Id   int
	Name string
	Sub  []SubSpecs
}
type SubSpecs struct {
	Id   int
	Name string
}
type SubServer struct {
	SubServerIds string  `mapstructure:"spec_ids"` // 规格id
	SspId        int     `mapstructure:"ssp_id"`   // 子服务id
	Price        float64 `mapstructure:"price"`
}
type ReplyGetBySsidsRpc struct {
	SsId             int
	ShopId           int
	SingleId         int
	ChangedRealPrice float64
	Status           int
}

type ArgsGetShopSpecs struct {
	ShopId int
	SspIds []int
}

type ReplyGetShopSpecs struct {
	ShopId        int
	SingleId      int
	SsId          int
	SspId         int
	SspName       string
	Price         float64
	OriginalPrice float64 // 原价
	MinPrice      float64 `mapstructure:"changed_min_price"` // 规格最低价
	MaxPrice      float64 `mapstructure:"changed_max_price"` // 规格最高价

}

// 根据单项目在门店的id获取所有子项目入参
type ArgsSubSpecID struct {
	SspIds []int
	ShopId int
}

type ReplyGetBySspids struct {
	SspId    int
	SingleId int
	Price    float64
}

// 以下为rpc扩展
type ReplyCommonSingleSpec struct {
	SspId   int     // 组合规格ID
	SspName string  // 组合规格名
	Price   float64 // 组合规格价格
}

type ArgsCommonShopSingle struct {
	ShopId    int
	SingleIds []int
	SspIds    []int
	Status    int // 单项目在门店的状态 0=全部 1=下架 2=上架 3=被总店禁用
}

type ReplyCommonSingle struct {
	SingleId        int     // 单项目ID
	ImgId           int     // 封面ID
	ServiceTime     int     // 服务时长
	Price           float64 `mapstructure:"real_price"` // 单项目现价
	ChangedMinPrice float64 // 区间最低价
	ChangedMaxPrice float64 // 区间最高价
}

type ReplyCommonShopSingle struct {
	SingleId        int // 单项目ID
	SsId            int
	Price           float64 `mapstructure:"changed_real_price"` // 单项目现价
	SspId           int
	IsDel           int     // 删除状态：0-否，1-是
	Status          int     // 上下架状态 1=下架 2=上架 3=已被总店禁用
	ChangedMinPrice float64 // 区间最低价
	ChangedMaxPrice float64 // 区间最高价
}

type ReplyGetPriceByShopIdAndSsspId struct {
	SingleId int     // 单项目ID
	Price    float64 // 单项目规格价格
	SspId    int
}

// 根据门店id和单项目Id获取单项目数据
type ArgsGetSingleByShopIdAndSingleIds struct {
	ShopId    int
	SingleIds []int
}

// 九百岁首页精选服务入参
type ArgsGetSelectServices struct {
	Cid int
	Num int
	Lng float64
	Lat float64
}

// 九百岁首页精选服务返回值
type ReplyGetSelectServices struct {
	ImgId      int
	ImgPath    string
	SingleId   int
	SingleName string
	Ssid       int
	RealPrice  float64
	Price      float64
	Sales      int
	ShopId     int
}

// \\以下为rpc扩展

type Single interface {
	// 添加单项目
	AddSingle(ctx context.Context, single *ArgsAddSingle, singleId *int) error
	// 获取单项目详情数据
	GetSingleInfo(ctx context.Context, args *ArgsGetSingleInfo, reply *SingleDetail) error
	// 修改单项目
	EditSingle(ctx context.Context, single *ArgsEditSingle, reply *bool) error
	// 子店添加单项到本店
	ShopAddSingle(ctx context.Context, single *ArgsShopAddSingle, reply *bool) error
	// 子店修改单项目价格
	ShopChangePrice(ctx context.Context, args *ArgsShopChangePrice, reply *bool)
	/*//总店上下架单项目
	DownUpSingle(ctx context.Context, args *ArgsDownUpSingle, reply *bool) error*/
	// 总店 删除单项目
	DelSingle(ctx context.Context, args *ArgsDelSingle, reply *bool) error
	// 分店 删除项目
	DelShopSingle(ctx context.Context, args *ArgsDelSingle, reply *bool) error
	// 子店上下架单项目
	ShopDownUpSingle(ctx context.Context, args *ArgsShopDownUpSingle, reply *bool) error
	// 获取商家的单项目列表
	BusSinglePage(ctx context.Context, args *ArgsBusSinglePage, reply *ReplyBusSinglePage) error
	// 获取门店的单项目列表
	ShopSinglePage(ctx context.Context, args *ArgsShopSinglePage, reply *ReplyShopSinglePage) error
	// 获取单项目在门店的价格
	GetShopSinglePrice(ctx context.Context, args *ArgsGetShopSinglePrice, price *float64) error
	// 获取单项的所有属性标签
	GetAttrs(ctx context.Context, emptyStr string, reply *ReplyGetAttrs) error
	// 根据单项目id批量获取基础价格信息
	GetSinglePriceListsBySingleIds(ctx context.Context, singleIds []int, reply *map[int]SinglePriceInfo) error
	// 单项目详情-rpc
	GetSimpleSingleInfo(ctx context.Context, singleId *int, reply *SimpleSingleInfo) error
	// 批量查询单项目详情-rpc
	GetSimpleSingleInfos(ctx context.Context, singleIds *[]int, reply *[]SimpleSingleInfo) error

	// 根据手艺人ID获取关联的单项目
	GetSignlesByStaffId(ctx context.Context, args *ArgsGetSignlesByStaffID, reply *ReplyGetSignlesByStaffID) error
	// 获取门店的单项目-rpc内部调用
	GetShopSingleBySingleIdsRpc(ctx context.Context, args *ArgsGetShopSingleBySingleIdsRpc, reply *ReplyGetShopSingleBySingleIdsRpc) error
	// 根据门店id和单项目Id获取单项目数据
	GetSingleByShopIdAndSingleIds(ctx context.Context, args *ArgsGetSingleByShopIdAndSingleIds, reply *ReplyShopSingle) error

	// 根据门店id查询单项目服务
	GetSingleByShopIdAndTagId(ctx context.Context, args *ArgsShopSingleByPage, reply *ReplyShopSingle) error
	// 根据单项目id获取子规格服务
	GetSubServerBySingleId(ctx context.Context, args *ArgsSubServer, reply *ReplySubServer) error

	// 根据门店单项目ids 获取门店单项目数据-rpc内部调用
	GetBySsidsRpc(ctx context.Context, ssIds *[]int, replay *[]ReplyGetBySsidsRpc) error
	// 根据门店id和规格组合ids获取数据-rpc内部使用
	GetShopSpecs(ctx context.Context, args *ArgsGetShopSpecs, reply *[]ReplyGetShopSpecs) error
	// 根据规格ID获取规格数据
	GetSingleSpecBySspId(ctx context.Context, args *ArgsSubSpecID, reply *ReplySubServer2) error
	// 获取门店指定单项目规格的价格
	GetPriceByShopIdAndSsspId(ctx context.Context, args *ArgsCommonShopSingle, reply *[]ReplyGetPriceByShopIdAndSsspId) error

	// 根据sspIds获取对应的singleId -rpc内部使用 消费确认时检查sspid和singleid的匹配
	GetBySspids(ctx context.Context, args *[]int, reply *[]ReplyGetBySspids)

	// 九百岁首页精选服务
	GetSelectServices(ctx context.Context, args *ArgsGetSelectServices, reply *[]ReplyGetSelectServices) error
}
