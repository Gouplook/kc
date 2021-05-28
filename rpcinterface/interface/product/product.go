package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/file"
)

const (
	IS_DEL_NO  = 0 //未删除
	IS_DEL_YES = 1 //已删除

	IS_BUS  = true  //是总店
	IS_SHOP = false //是门店

	//盘点状态
	CHECK_UNDER_WAY    = 1 //进行中
	CHECK_FINISHED     = 2 //已完成
	CHECK_CANCELLATION = 3 //作废

	//是否盘点锁定状态
	CHECK_LOCK_NO = 0		//盘点未锁定
	CHECK_LOCK_YES = 1		//盘点已锁定

	//出入库类型
	STOCK_OTHER_IN        = 1 //其他入库
	STOCK_CHECK_IN        = 2 //盘盈入库
	STOCK_CHECK_OUT       = 3 //盘亏出库
	STOCK_PUR_IN          = 4 //采购入库
	STOCK_SHOP_REQUIRE_IN = 5 //门店要货入库
	STOCK_BUS_REQUIRE_OUT = 6 //总部要货出库
	STOCK_OTHER_OUT       = 7 //其他出库

	//供应商是否已经启用
	IS_OPENED_YES = 0 //已经启用
	IS_OPENED_NO  = 1 //没有启用

	//单据状态码
	BILL_STATUS_AWAIT_CHECK = 1 //待审核
	BILL_STATUS_AWAIT_IN    = 2 //待入库
	BILL_STATUS_RETURN      = 3 //已驳回
	BILL_STATUS_FINISH      = 4 //已完成
	BILL_STATUS_CLOSE       = 5 //已关闭

	//上下架状态
	IS_UP   = 0 //下架
	IS_DOWN = 1 //上架

	//是否入库 出库 盘点  标识
	IS_IN    = 1 //入库 标识
	IS_OUT   = 2 //出库 标识
	IS_CHECK = 3 //盘点 标识

	//是否盘平
	IS_BALANCE_YES = 1 // 盘平
	IS_BALANCE_NO  = 2 // 没有平

	//
	GET_CHECK_INFOS = 1 //查看盘点单
	CONTINUE_CHECK  = 2 //继续 盘点单

	TIME_TODAY = 1 //今天
	TIME_WEEK  = 2 //近七天
	TIME_MONTH = 3 //本月

	//预警开关
	WARN_IS_OPEN  = 0
	WARN_IS_CLOSE = 1

	//预警状态
	WARN_STATUS_NORMAL = 0
	WARN_STATUS_LOW = 1

	//是否开启自定义预警数量
	WARN_CUSTOM_IS_OPEN  = 1
	WARN_CUSTOM_IS_CLOSE = 0

	//出入库类型
	STOCKTYPE_INT_OTHER = 1      //其他入库
	STOCKTYPE_INT_INV_PROFIT = 2	//盘盈入库
	STOCKTYPE_OUT_INV_LOSS = 3		//盘亏出库
	STOCKTYPE_INT_PUR = 4			//采购入库
	STOCKTYPE_INT_SHOP_REQUIRE = 5	//门店要货入库
	STOCKTYPE_OUT_BUS_REQUIRE = 6	//总部要货出库
	STOCKTYPE_OUT_OTHER = 7			//其他出库
	STOCKTYPE_OUT_BREAKAGE = 8			//报损出库
	STOCKTYPE_OUT_GET = 9			//领料出库
	STOCKTYPE_OUT_SALE = 10			//门店销售出库

	//单位
	//个、件、包、只、套、对、张、本、条、根、桶、次、片、瓶、盒、箱、袋
	UNIT_GE = 1		//个
	UNIT_JIAN = 2	//件
	UNIT_BAO = 3	//包
	UNIT_ZHI = 4	//只
	UNIT_TAO = 5	//套
	UNIT_DUI = 6	//对
	UNIT_ZHANG = 7	//张
	UNIT_BEN = 8	//本
	UNIT_TIAO = 9	//条
	UNIT_GEN = 10	//根
	UNIT_TONG = 11	//桶
	UNIT_CI = 12	//次
	UNIT_PIAN = 13	//片
	UNIT_PING = 14	//瓶
	UNIT_HE = 15	//盒
	UNIT_XIANG = 16	//箱
	UNIT_DAI = 17	//袋
)

var UnitMap = map[int]string {
	UNIT_GE : "个",
	UNIT_JIAN : "件",
	UNIT_BAO : "包",
	UNIT_ZHI : "只",
	UNIT_TAO : "套",
	UNIT_DUI : "对",
	UNIT_ZHANG : "张",
	UNIT_BEN : "本",
	UNIT_TIAO : "条",
	UNIT_GEN : "根",
	UNIT_TONG : "桶",
	UNIT_CI : "次",
	UNIT_PIAN : "片",
	UNIT_PING : "瓶",
	UNIT_HE : "盒",
	UNIT_XIANG : "箱",
	UNIT_DAI : "袋",
}

var TypeMap = map[int]string{
	STOCKTYPE_INT_OTHER:"其他入库",
	STOCKTYPE_INT_INV_PROFIT :"盘盈入库" ,
	STOCKTYPE_OUT_INV_LOSS : "盘亏出库" ,
	STOCKTYPE_INT_PUR : "采购入库" ,
	STOCKTYPE_INT_SHOP_REQUIRE : "门店要货入库" ,
	STOCKTYPE_OUT_BUS_REQUIRE :"总部要货出库" ,
	STOCKTYPE_OUT_OTHER : "其他出库" ,
	STOCKTYPE_OUT_BREAKAGE : "报损出库" ,
	STOCKTYPE_OUT_GET : "领料出库" ,
	STOCKTYPE_OUT_SALE : "门店销售出库" ,

}

var TypeMapOut = map[int]string{
	STOCKTYPE_OUT_OTHER : "其他出库" ,
	STOCKTYPE_OUT_BREAKAGE : "报损出库" ,
	STOCKTYPE_OUT_GET : "领料出库" ,
}

//商品分类信息 新增 入参
type ArgsCategoryAdd struct {
	common.BsToken        //商户信息
	Name           string //商品分类名称
}

//商品分类信息 修改 入参
type ArgsCategoryUpdate struct {
	common.BsToken        //商户信息
	Id             int    //商品分类id
	Name           string //商品分类名称
}

//商品分类信息 删除 入参
type ArgsCategoryDel struct {
	common.BsToken     //商户信息
	Id             int //商品分类id
}

//商品分类信息 查询 入参
type ArgsCategoryGet struct {
	common.Paging //分页信息
	BusId         int
}

//单条商品分类信息
type Category struct {
	Id         int    //商品分类id
	Name       string //商品分类名称
	CreateTime string //创建时间
}

//商品分类信息返回
type ReplyCategoryPage struct {
	TotalNum int        //总条数
	List     []Category //商品分类信息集合
}

//商品标签信息 新增 入参
type ArgsTagAdd struct {
	common.BsToken        //商户信息
	Name           string //商品分类名称
}

//商品标签信息 修改 入参
type ArgsTagUpdate struct {
	common.BsToken        //商户信息
	Id             int    //商品分类id
	Name           string //商品分类名称
}

//商品标签信息 删除 入参
type ArgsTagDel struct {
	common.BsToken     //商户信息
	Id             int //商品分类id
}

//商品标签信息 查询 入参
type ArgsTagGet struct {
	common.Paging //分页信息
	BusId         int
}

//单条商品标签信息
type Tag struct {
	Id         int    //商品分类id
	Name       string //商品分类名称
	CreateTime string //创建时间
}

//商品标签返回
type ReplyTagPage struct {
	TotalNum int   //总条数
	List     []Tag //标签信息集合
}

//商品规格名称
type ProductSpec struct {
	Id  int   //规格id
	Sub []int //子规格
}

//es 门店批量添加商品入参
type ArgsProductsEs struct {
	Ids    []int
	ShopId int
}

//es 门店加商品入参
type ArgsShopAddProductEs struct {
	Ids       []int
	ShopId    int
}

//es检索商品入参
type ArgsProductEsGet struct {
	common.Paging
	Id     int
	CateId int
	TagId  int
	BusId  int
	ShopId int
}

//es检索规格明细商品入参
type ArgsProductDetailEsGet struct {
	common.Paging
	Id     int
	CateId int
	Name   string
	BusId  int
	ShopId int
	Flag   bool //true 表示预警调用这个接口
}

//Es查询商品返回
type ReplyProductEs struct {
	List     []map[string]interface{}
	TotalNum int
}


//es添加商品入参
type ArgsAddProductEs struct {
	Id        int
	CateId    int
	Name      string
	TagIds    []int
	BusId     int
	ShopId    int
	Details []DetailEs
}
type DetailEs struct {
	SpecId	  string
	DetailId int
	Stock int
	UsableStock int
	IsCustom bool
	BarCode   string
}

//es删除商品入参
type ArgsDelProductEs struct {
	Ids       []int
	Deleted   int
	DetailIds []int
}

//商品新增 和 修改 共用
type AddProductInfo struct {
	Name     string  //商品名称
	BarCode  string  //条形码
	CateId   int     //商品分类id
	TagIds   string  //商品标签组合id  多个id,号隔开
	FirstImg string  //封面照片hash
	Imgs     string  //多张照片hash  多个hash,号隔开
	Cost     float64 //成本价
	Sell     float64 //售价
	Price    float64 //标价
	UnitId   int	 //单位id
}

//商品信息 新增 入参
type ArgsProductAdd struct {
	common.BsToken //商户信息
	Id             int //商品id
	AddProductInfo
	Specs  []ProductSpec   //商品规格详情
	Detail []ProductDetail //商品详情信息
}

//商品详情信息
type ProductDetail struct {
	SpecIds []int   //商品规格组合id
	BarCode string  //条形码
	Cost    float64 //商品成本
	Sell    float64 //商品售价
}

//商品信息 删除 入参
type ArgsProductDel struct {
	common.BsToken     //商户信息
	Id             int //商品id
}

//根据批量id查询商品信息
type ArgsProductGetByIds struct {
	Ids []int
}

//查询返回
type ReplyProductGetByIds struct {
	Id        int
	BusId     int
	Name      string
	ImgId 	  int //图片id
	SpecPrice string
	MinPrice  float64 //最低价
	MaxPrice  float64 //最高价
	IsDel     int //是否删除
}

//商品信息 批量删除 入参
type ArgsProductDelMore struct {
	common.BsToken       //商户信息
	Ids            []int //多个商品id
}

//商品信息 根据分类和标签查询 入参
type ArgsProductGet struct {
	common.BsToken
	common.Paging     //分页信息
	CategoryId    int //分类id
	//TagId         int //标签id
}

type Img struct {
	Hash string
	Url  string
}

//商品图片数据
type ProductImg struct {
	FirstImg Img   //封面图片地址
	MoreImg  []Img //相册图片
}

//商品信息返回
type ReplyProductPage struct {
	TotalNum int           //总条数
	List     []ProductBase //商品批量信息
}

//基础 商品信息
type ProductBase struct {
	Id         int     //商品id
	Name       string  //商品名称
	MinPrice   float64 //最低价
	MaxPrice   float64 //最高价
	CateId     int     //商品分类id
	CateName   string  //分类名称
	TagIds     string  //商品组合标签id  多个用,号隔开
	TagNames   string  //标签名称返回
	ImgId	int
	ImgUrl	string         //相册
	ImgHash	string
	Count       int  //在售门店数量
	TotalStock  int   //总库存
	TotalSales  int   //总销量
}

//查询 一条商品详情 入参
type ArgsProductOneGet struct {
	BusId int //商户信息
	ShopId int //门店id
	Id    int //商品id
}

//查询 一条详情 返回
type ReplyProductOne struct {

	//product
	Id     int           `mapstructure:"id"`        //商品明细id
	Name   string        `mapstructure:"name"`      //商品名称
	CateId int           `mapstructure:"cate_id"`   //分类id
	CateName string
	UnitId int			//单位id
	UnitName string		//单位名称
	TagIds string        `mapstructure:"tag_ids"`   //商品组合标签id  多个用,号隔开
	TagNames string
	MinPrice float64
	MaxPrice float64
	Specs  []ProductSpec `mapstructure:"spec_json"` //规格
	Price  float64       `mapstructure:"price"`     //标价

	//detail
	Detail  []DetailInfo //详情信息
	BarCode string       `mapstructure:"bar_code"` //条形码
	Cost    float64      `mapstructure:"cost"`     //成本价
	Sell    float64      `mapstructure:"sell"`     //售价
	DetailId int		//商品明细id

	//total stock
	TotalStock int //总库存
	TotalSales int //总销量

	//img
	ProductImg

	//规格map
	SpecList map[int]Spec
}

type Spec struct {
	Id int
	Name string
}

type DetailInfo struct {
	Id      int     `mapstructure:"id"`       //商品明细id
	SpecIds string  `mapstructure:"spec_ids"` //规格
	Cost    float64 `mapstructure:"cost"`     //成本
	Sell    float64 `mapstructure:"sell"`     //售价
	BarCode string  `mapstructure:"bar_code"` //条形码
	Stock   int     `mapstructure:"stock"`      //库存
	Sales   int     `mapstructure:"sales"`    //销量
}

//出入库商品明细 根据分类和名称查询 入参
type ArgsDetailGet struct {
	BusId         int    //商户信息
	common.Paging        //分页信息
	CateId        int    //分类id
	//Name          string //商品名称
}

//商品详情信息返回
type ReplyDetailPage struct {
	TotalNum int      //总条数
	List     []Detail //商品批量信息
}

//基础 商品明细
type DetailBase struct {
	Id         int    //商品明细id
	ProductId  int    //商品 id   //出入库添加商品需要传参
	Name       string //商品名称
	SpecIds    string //规格组合id
	SpecNames string
	BarCode    string //条形码
	CateId     int    //商品分类id
	CateName   string //商品分类名称
	UnitId int
	UnitName string
	ImgId int
	ImgPath string
}

//基础 商品明细
type DetailBase2 struct {
	Id         int    //商品明细id
	ProductId  int    //商品 id   //出入库添加商品需要传参
	Name       string //商品名称
	SpecIds    string //规格组合id
	SpecNames string
	BarCode    string //条形码
	CateId     int    //商品分类id
	CateName   string //商品分类名称
	UnitId int
	UnitName string
	ImgId int
	ImgPath string
	Stock int
	UsableStock int
	Sales int
}

//商品详情信息
type Detail struct {
	Id         int    //商品明细id
	ProductId  int    //商品 id   //出入库添加商品需要传参
	Name       string //商品名称
	SpecIds    string   //规格组合id
	SpecNames string
	BarCode    string //条形码
	CateId     int    //商品分类id
	CateName   string //商品分类名称
	UnitId     int
	UnitName   string
	ImgId int
	ImgPath string
	Stock int
	UsableStock int
	IsLock int   //是否被锁定
}

//总部添加预警入参

// 商品基础价格信息
type ProductPriceInfo struct {
	ProductId int     //商品 id
	Name      string  //名称
	MinPrice  float64 //最低价
	MaxPrice  float64 //最高价
}

//添加商品规格入参
type ArgsSpecAdd struct {
	common.BsToken
	Name     string //规格名称
	ParentId int    //父级规格id
}

//查询所有规格返回
type ReplySpecs struct {
	Id       int    `mapstructure:"id"`        //规格id
	Name     string `mapstructure:"name"`      //规格名称
	ParentId int    `mapstructure:"parent_id"` //父级规格id
}

//根据一级规格id
type ReplySpec struct {
	Id   int    `mapstructure:"id"`   //规格id
	Name string `mapstructure:"name"` //规格名称
}

//一级和二级规格
type GetSpecs struct {
	ReplySpec
	Sub []ReplySpecs
}

type ReplyGetSpecs struct {
	Specs []GetSpecs
}

//根据一级规格查询下属规格 传0 获取所有1级
type ArgsSpecGet struct {
	BusId    int
	ParentId int
}

type ReplyTag struct {
	Id   int    //标签id
	Name string //标签名称
}

type ReplyGetDetailById struct {
	Id int
	SpecIds string// 规格组合ID
	Name string // 商品name
	Sell float64 //售价
	Cost float64//成本价
}

//查询商品照片返回
type ReplyImage struct {
	ProductId int
	FileInfo file.ReplyFileInfo
}

//根据商品id查询商品规格和价格
type ReplySubServer struct {
	SubServer   []SubServer
	Specs []cards.Specs
}
type SubServer struct {
	SubServerIds string	`mapstructure:"spec_ids"`	 //规格id
	SspId        int    `mapstructure:"id"`		//子服务id
	Price        float64 `mapstructure:"sell"`
	Stock        int
}

//异步修改总库存和预警状态
type ArgsMqUpdateStock struct {
	BusId int
	ShopId int
	Type int    // 1是入库 2是出库 3是盘点
	Details []StockDetail
}

// 获取添加商品总数入参
type ArgsProductNum struct {
	BusId int
}
//返回企业/店铺 商品总数量
type ReplyProductNum struct {
	ProductNum int
}

//根据busId获取商品列表入参--rpc
type ArgsGetProductByBusId struct {
	BusId int
	ShopId int
	common.Paging
}
type GoodsInfo struct {
	GoodsId int		`mapstructure:"id"`//商品Id
	Name string		`mapstructure:"name"`//商品名称
}
//根据busId获取商品列表出参--rpc
type ReplyGetProductByBusId struct {
	GoodsInfos []GoodsInfo
	TotalNum int 	//总数
}

//根据busId或者shopId获取商品id入参 --rpc
type ArgsGetProductIds struct {
	BusId int
	ShopId int
	common.Paging
}

//根据busId或者shopId获取商品id出参 --rpc
type ReplyGetProductIds struct {
	ProductIds []int
	TotalNum int 	//总数
}

//判断分店是否添加总店所有商品
type ArgsIsShopProductEqBus struct {
	BusId int
	ShopId int
}

//判断分店下是否含有指定商品
type ArgsIsShopIncProducts struct {
	ShopId int
	ProductIds []int
}

type Product interface {
	//添加商品分类
	AddCategory(ctx context.Context, args *ArgsCategoryAdd, categoryId *int) error
	//更改商品类名
	UpdateCategory(ctx context.Context, args *ArgsCategoryUpdate, reply *bool) error
	//删除 一条 商品分类
	DelCategory(ctx context.Context, args *ArgsCategoryDel, reply *bool) error
	//查询 对应商户 的商品分类
	GetCategories(ctx context.Context, args *ArgsCategoryGet, reply *ReplyCategoryPage) error

	//添加商品标签
	AddTag(ctx context.Context, args *ArgsTagAdd, tagId *int) error
	//更改商品标签名
	UpdateTag(ctx context.Context, args *ArgsTagUpdate, reply *bool) error
	//删除 一条 商品标签
	DelTag(ctx context.Context, args *ArgsTagDel, reply *bool) error
	//查询 对应商户 的商品标签
	GetTags(ctx context.Context, args *ArgsTagGet, reply *ReplyTagPage) error
	//根据busId查询所有标签id和标签名字
	GetTagsByBusId(ctx context.Context, args *int, reply *[]ReplyTag) error

	//添加商品
	AddProduct(ctx context.Context, args *ArgsProductAdd, detailId *int) error
	//修改商品
	UpdateProduct(ctx context.Context, args *ArgsProductAdd, reply *bool) error
	//异步修改预警
	MqUpdateWarn(ctx context.Context, args int,reply *bool) error
	//删除一条 商品
	DelProduct(ctx context.Context, args *ArgsProductDel, reply *bool) error
	//根据ids 查询商品信息
	GetProductByIds(ctx context.Context, args *ArgsProductGetByIds, reply *[]ReplyProductGetByIds) error
	//异步删除商品
	MqDelProduct(ctx context.Context,args []int,reply *bool) error
	//批量删除  商品
	DelMoreProduct(ctx context.Context, args *ArgsProductDelMore, reply *bool) error
	//根据 商品分类 和 商品标签 查询商品
	GetProducts(ctx context.Context, args *ArgsProductGet, reply *ReplyProductPage) error
	//查询一条详情
	GetProductOne(ctx context.Context, args *ArgsProductOneGet, reply *ReplyProductOne) error
	//根据 商品分类 和 商品名称 查询商品详情
	GetDetail(ctx context.Context, args *ArgsStockGet, reply *ReplyDetailPage) error

	//添加商品规格
	AddSpec(ctx context.Context, args *ArgsSpecAdd, reply *int) error
	//根据busId查询所属
	GetSpecs(ctx context.Context, busId *int, reply *[]ReplySpec) error
	//根据一级规格id查询所有二级规格 //传0 获取所有一级规格
	GetSpecsById(ctx context.Context, args *ArgsSpecGet, reply *[]ReplySpec) error
	//根据detailIds查询售价-RPC调用
	GetByDetailIds(ctx context.Context, detailIds *[]int, reply *[]ReplyShopSell) error
	//根据DetailId获取商品规格-rpc
	GetSpecsByDeatailIds(ctx context.Context,detailIds *[]int,reply *ReplyGetSpecs)error
	//商品明细-rpc
	GetDetailByIds(ctx context.Context,detailIds *[]int, reply *[]ReplyGetDetailById)error
	//根据多个商品id查询照片
	GetImageByProductIds(ctx context.Context,productIds *[]int, reply *[]ReplyImage)error
	//根据商品id获取子规格商品
	GetDetailsByProductId(ctx context.Context, args *cards.ArgsSubServer, reply *ReplySubServer) error
	//购买成功，设置产品的销量和库存
	ChangeProductSalesAndStack(ctx context.Context, orderSn *string, reply *bool) error
	//异步修改总库存和预警
	MqUpdateStock (ctx context.Context,args *ArgsMqUpdateStock,reply *bool) error
	// 获取添加商品总数
	GetProductNum(ctx context.Context, args *ArgsProductNum, reply *ReplyProductNum) error
	//根据busid获取商品列表 --rpc用
	GetProductByBusId(ctx context.Context, args *ArgsGetProductByBusId, reply *ReplyGetProductByBusId )error
	//根据busId或者shopId获取商品id --rpc
	GetProductIds(ctx context.Context, args *ArgsGetProductIds, reply *ReplyGetProductIds)error
	//判断分店是否添加总店所有商品
	IsShopProductEqBus(ctx context.Context,args *ArgsIsShopProductEqBus,reply *bool)error
	//判断分店下是否含有指定商品
	IsShopIncProducts(ctx context.Context,args *ArgsIsShopIncProducts,reply *bool)error
}
