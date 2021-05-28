package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/file"
)

const (
	IS_GROUND_UP = 0
	IS_GROUND_DOWN   = 1
)

//分店 商品列表 添加 入参
type ArgsShopProductAdd struct {
	common.BsToken //商户信息
	ProductId int //商品id
}

//分店 商品列表 批量添加 入参
type ArgsShopProductMoreAdd struct {
	common.BsToken //商户信息
	ProductIds []int//商品ids
}

//分店 商品列表 查询 入参
type ArgsShopProductGet struct {
	ShopId        int //商户信息
	common.Paging     //分页信息
	CateId        int //分类Id
	IsFilterPutaway  bool  //
	//TagId         int //标签Id
}

//分店 商品列表 返回
type ReplyShopProductPage struct {
	TotalNum int               //总条数
	Lists    []ShopProductInfo //数据
}

//分店 商品列表
type ShopProductInfo struct {
	ProductId         int     //商品id
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
	TotalStock  int   //总库存
	TotalSales  int   //总销量
	IsGround      int //是否上架
	BidPrice float64	//标价
	Sell float64	//没有规格的时候有值
	DetailId     	int //商品明细id   没有规格的时候才有值
}

//分店商品上下架
type ArgsShopProductUpDown struct {
	common.BsToken     //门店信息
	ProductId      int //商品id
	Type           int //判断是上架还是下架  0上架 1下架
}

//分店商品批量上下架
type ArgsShopProductUpDowns struct {
	common.BsToken       //门店信息
	ProductIds     []int //商品id
	Type           int   //判断是上架还是下架  0上架 1下架
}

//分店 出入库 明细查询
type ArgsShopDetailGet struct {
	BusId         int
	ShopId        int    //门店信息
	common.Paging        //分页信息
	CateId        int    //分类id
	Name          string //产品 名称
}

//门店要货申请入参
type ArgsShopRequireAdd struct {
	common.BsToken
	common.Utoken
	Remark  string //备注
	Details []StockDetail
}

//门店要货申请查看一条详情
type ArgsRequireDetailGet struct {
	common.Paging
	Status    int    //单据状态 1待审核 2待入库 3已驳回 4已完成 5已关闭
	BillNum   string //申请单据号
	BillNumYh string //要货单据号
}

//门店要货详情返回
type ReplyRequireDetailPage struct {
	TotalNum int
	List     []RequireDetail
}
type RequireDetail struct {
	//基础 数据
	DetailBase

	//申请库
	ReqNum int //要货数量

	//出入库
	ReaNum int     //实发数量  //总店出库
	RecNum int     //实收数量  //门店入库
	Price  float64 //单价
	TotalNum float64 //总价
}

//门店要货申请查询入参
type ArgsShopRequireGet struct {
	Uid int
	BusId         int      // 总店查 用
	ShopId        int      //  分店查 用
	common.Paging          //分页信息
	Status        int      //要货状态 1待审核 2待入库 3已驳回 4已完成 5已关闭
	Time          []string //入库时间 开始 到 结束
	Flag          int      //1表示 今天  2表示 近七天  3表示 本月
}

//门店要货申请查询返回
type ReplyShopRequirePage struct {
	TotalNum int
	List     []ShopRequireInfo
}
type ShopRequireInfo struct {
	BillNum    string  //申请单据号
	BillNumYh  string  //要货单据号
	BillNumCk  string  //总店出库单据号
	ShopId     int     //门店id
	ShopName   string  //门店名称
	BranchName string  //分店名称
	Status     int     //单据状态 1待审核 2待入库 3已驳回 4已完成 5已关闭
	TotalMoney float64 //
	TypeId  int       //出入库类型id
	TypeName string
	Time       string  //入库时间
	CreateTime string  //申请时间
	Remark     string  //备注
	CheckTime  string  //审核时间
	Uid int
	Uname string
}

//取消要货申请
type ArgsRequireCancel struct {
	common.BsToken
	BillNum string
}

//要货申请入库完成
type ArgsRequireUpdate struct {
	common.BsToken
	common.Utoken
	BillNum string
	Time    string
	Remark  string
	Details []InRequireDetail
}
type InRequireDetail struct {
	DetailId  int     //明细id
	//ReaNum    int     //实发数量
	RecNum    int     //实收数量
	Price     float64 //实收单价
	ProductId int     //商品id
}

type ProductDetailBase struct {
	Id        int     `mapstructure:"id"` //商品明细id
	ProductId int     `mapstructure:"product_id"`
	SpecIds   string  `mapstructure:"spec_ids"` //规格
	Cost      float64 `mapstructure:"cost"`     //成本
	Sell      float64 `mapstructure:"sell"`     //售价
	Name      string
	Selected  bool
}

//ProductServiceBase ProductServiceBase
type ProductServiceBase struct {
	Id        int    // 商品ID
	CateId    int    // 分类ID
	MaxPrice  string // 商品价格
	MinPrice  string
	SpecPrice string // 规格价格
	Name      string // 商品name
	//ImgId int // 商品图片ID
	ImgUrl        string
	CreateTime    int64 `mapstructure:"create_time"`
	CreateTimeStr string
	//ParentSkus []ParentSkusBase
	SkuList []ProductDetailBase
}

// 门店详情-本店商品列表
type ArgsShopInfoProductList struct {
	common.Paging
	ShopId int
}
type ReplyShopInfoProductListBase struct {
	Id        int     // 门店商品主键
	Name      string  // 商品名
	ImgId     int     // 商品封面
	ProductId int     //商品id
	Num       int     //门店库存
	Sales     int     //门店销量
	IsGround  int     // 是否上架 0表示下架 1表示上架
	Price  float64			 // 门店商品标价/划线价格
	MinPrice  float64 // 门店商品最小规格售价
	MaxPrice  float64 // 门店商品最小规格售价
}
type ReplyShopInfoProductList struct {
	TotalNum  int
	Lists      []ReplyShopInfoProductListBase
	IndexImgs map[int]file.ReplyFileInfo
}

//门店修改商品价格
type ShopPriceUpdate struct {
	common.BsToken
	ProductId int
	SpecPrices []SpecPrice
	//Price float64
	RealPrice float64
}
//明细id和价格
type SpecPrice struct {
	DetailId int
	Sell float64
}

//根据门店id和detailIds查询门店售价 入参
type ArgsShopSellGet struct {
	ShopId int
	DetailIds []int
}

//根据门店id和detailIds查询门店售价 返回
type ReplyShopSell struct {
	DetailId int
	Price float64
}
//根据门店id和商品ID查询是否存在-RPC调用
type ArgsGetProductByIds struct {
	ShopId int
	Status string
	ProductIds []int
}
type ReplyGetProductByIds struct {
	ProductId int
	SpecPrice string `mapstructure:"spec_price"`
}
//根据多个商品id查询 价格区间
type ArgsGetImages struct {
	ShopId int
	ProductIds []int
}

//查询门店价格区间返回
type ReplySpecPrice struct {
	ProductId int
	MaxPrice string
	MinPrice string
	IsGround int
}

//获取门店的产品信息入参
type ArgsGetShopProdcuts struct {
	ShopId int //门店id
	DetailIds []int //产品明细id
}

//获取门店的产品信息返回数据
type ReplyGetShopProdcuts struct {
	ProductId int //产品id
	DetailId int //明细id
	Price float64 //标价
	RealPrice float64 //售价
	Name string //产品名称
	ImgId int //产品图片id
	DetailName string //规格组合名称
	IsDel int //是否被删除 0=正常 1=删除
	IsGround int //门店是否上架 0=下架 1=上架
	Stock int //库存数量
}

type ArgsAddProduct struct {
	ShopId int
	ProductIds []int
}

type ArgsGetBusProductUnderRate struct {
	ShopId int
	ProductId int
}

type ReplyGetBusProductUnderRate struct {
	BusId int
	DateTime int64
	AllProductNum int
	MonthUnderProductNum int
	UnderProductRate float64
}

type ShopProduct interface {

	//总部商品添加到门店
	AddProduct(ctx context.Context, args *ArgsShopProductAdd, id *int) error
	//异步回调 添加商品
	MqAddProduct(ctx context.Context,args *ArgsAddProduct,reply *bool) error
	//批量添加
	AddMoreProduct(ctx context.Context, args *ArgsShopProductMoreAdd, id *int) error
	//商品上下架
	UPDownProduct(ctx context.Context, args *ArgsShopProductUpDown, reply *bool) error
	//商品批量上下架
	MoreUpDownProduct(ctx context.Context, args *ArgsShopProductUpDowns, reply *bool) error
	//查询门店商品
	GetShopProduct(ctx context.Context, args *ArgsShopProductGet, reply *ReplyShopProductPage) error
	// 出入库添加商品 查询 根据分类和名称
	GetShopDetails(ctx context.Context, args *ArgsStockGet, reply *ReplyDetailPage) error
	//门店查询出入库明细
	GetShopInOutStock(ctx context.Context, args *ArgsInOutStockGet, reply *ReplyInOutStockPage) error

	/*//统一设置库存预警数量
	SetShopWarnNum(ctx context.Context, args *ArgsWarnNumSet, reply *bool) error
	//查询所有预警 商品明细
	GetShopWarn(ctx context.Context, args *ArgsWarnDetail, reply *ReplyWarnDetailPage) error
	//关闭或开启 明细商品 预警
	SetShopWarnDetail(ctx context.Context, args *ArgsDetailSet, reply *bool) error
	//自定义预警值
	SetShopCustom(ctx context.Context, args *ArgsCustom, reply *bool) error*/

	//门店要货申请
	AddShopRequire(ctx context.Context, args *ArgsShopRequireAdd, reply *int) error
	/*//门店一条要货申请详情查询
	GetRequireDetail(ctx context.Context, args *ArgsRequireDetailGet, reply *ReplyRequireDetailPage) error*/
	//门店要货申请查询 总部 门店 共用 接口
	GetShopRequire(ctx context.Context, args *ArgsShopRequireGet, reply *ReplyShopRequirePage) error
	//门店要货申请修改状态 //取消操作  待审核 修改为 已关闭
	CancelRequire(ctx context.Context, args *ArgsRequireCancel, reply *bool) error
	//门店要货申请修改状态 //入库操作  待入库 修改为 已完成
	InRequire(ctx context.Context, args *ArgsRequireUpdate, reply *bool) error

	// 附近-门店详情-本店商品
	ShopInfoProductList(ctx context.Context, args *ArgsShopInfoProductList, reply *ReplyShopInfoProductList) error
	//门店修改价格
	UpdateShopPrice(ctx context.Context, args *ShopPriceUpdate, reply *bool) error
	//根据门店id和detailIds 查询售价-RPC调用
	GetByShopIdAndDetailIds(ctx context.Context, args *ArgsShopSellGet, reply *[]ReplyShopSell) error
	//根据门店id和商品ID查询是否存在-RPC调用
	GetProductByIds(ctx context.Context,args *ArgsGetProductByIds,reply *[]ReplyGetProductByIds)error
	//根据多个商品id查询 价格区间
	GetShopSpecPrice(ctx context.Context,args *ArgsGetImages,reply *[]ReplySpecPrice)error
	//根据门店id和detailIds 获取门店产品信息
	GetShopProdcuts(ctx context.Context, args *ArgsGetShopProdcuts, reply *[]ReplyGetShopProdcuts) error

	//获取企业当月商品下架率
	GetBusProductUnderRate(ctx context.Context,args *ArgsGetBusProductUnderRate,reply *ReplyGetBusProductUnderRate)
}
