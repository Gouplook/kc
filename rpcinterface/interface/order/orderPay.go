//订单接获取支付信息
//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/29 16:22

package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	file2 "git.900sui.cn/kc/rpcinterface/interface/file"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

type ArgsGetPaySign struct {
	OrderSn      string // 订单号
	ChosePayType int    // 选择的支付方式
	Channel      int    // 客户端渠道
	OpenId       string // 微信openid 非必传
	AppId        string //微信Appid 非必传
	Version      string // 客户端版本
	CreateIP     string // 客户端真实ip
	StoreID      string // 门店号
	PayExtra     string //
	AccsplitFlag string // 分账标识 NO YES
	SignType     string // MD5
}

type ArgsGetPayInfo struct {
	OrderSn      string //订单号
	ChosePayType int    //选择的支付方式
	Channel      int    //客户端渠道
	OpenId       string // 微信openid 非必传
	AppId        string //微信Appid 非必传
}

type ReplyGetPayInfo struct {
	PayUrl  string //支付二维码链接或H5跳转地址
	PayData string //app和小程序的支付结果数据
	PaySign string //PaySign
}

//获取分账信息的返回数据结构体
type ReplyGetOrderSplitBill struct {
	BusId             int     //商户id
	InsuranceChannel  int     //保险渠道
	FundMode          int     //资金管理
	DepositRatio      float64 //留存比例
	DepositAmount     float64 //留存金额
	InsureAmount      float64 //保险费用
	RenewInsureAmount float64 //续保费用
	PlatformAmount    float64 //平台手续费
	BusAmount         float64 //商户应收起金额
}

//现金支付的参数
type ArgsCashPaySuc struct {
	common.BsToken         //商家token
	OrderSn        string  //订单号
	Amount         float64 //支付现金金额
	BackAmount     float64 //找零金额

}

//服务订单入参
type ArgsServiceOrderList struct {
	common.Paging
	common.BsToken
	CreateTimeStart string // 下单开始时间
	CreateTimeEnd   string // 下单结束时间
	ShopId          int    // 下单门店
	DateType        int    // 时间类型
	Status          string // ”“为全部，0=待支付 1=已支付 2=支付失败 3=订单关闭
	KeyWords        string // 关键字：订单编号，手机号码
	PayType         int    // 付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
	OrderType       int    //订单类型 1=单项目订单 2=卡项订单
	Uid             int    //用户ID-会员列表-服务订单使用
}

//单项目/卡包支付子订单数据
type SubOrderListsBase struct {
	Id                int     //卡项/单项目子订单
	CardName          string  // 卡包名
	SingleName        string  //单项目名称
	ImgId             int     // 单项目/卡项封面图片
	SspId             int     // 规格ID
	SspName           string  // 规格名
	Price             float64 // 卡项/单项目面值
	RealPrice         float64 // 卡项/单项目销售金额
	DiscountPrice     float64 //改价后的金额
	TotalAmount       float64 // 实际付款：如果改价则使用改价的价格
	Num               int     // 购买数量
	Mid               int     // 购买用户会员ID
	StaffIds          string  // 销售人员
	CraftStaffIds     string  // 手艺人
	SubOrderId        int     //所属订单ID@kc_single_order表中的id字段
	OrderType         int     //子订单类型 1=会员订单 2=游客订单
	Status            int     //子订单状态 0=待支付 1=已支付 2=支付失败 3=订单关闭
	CardPackageStatus int     //卡包状态 1=待消费 2=消费中  3=已完成 4=已关闭
}

//单项目/卡包支付订单数据
type ReplyServiceOrderListBase struct {
	OrderSn        string  // 支付订单编号
	PaySn          string  // 支付流水号-支付渠道返回支付流水号
	OrderId        int     //支付订单ID
	Uid            int     // 下单用户Uid
	PayStatus      int     //支付状态 0=待支付 1=已支付 2=支付失败 3=订单关闭
	TotalAmount    float64 // 支付订单总金额
	RealPrice      float64 // 支付订单实际支付金额
	DiscountPrice  float64 //优惠总金额
	OldTotalAmount float64 // 合计金额（优惠总金额+实际支付金额）
	OrderType      int     //支付订单类型 1=单项目订单 2=卡项订单
	PayType        int     // 支付付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
	CreateTime     int64   // 支付订单生成时间
	CreateTimeStr  string  // 支付订单生成时间
	PayTime        int64   // 支付付款时间
	PayTimeStr     string  // 支付付款时间
	BusId          int
	ShopId         int
	ShopName       string
	BranchName     string
	SubOrderLists  []SubOrderListsBase
	BuyerInfo      bus.ReplyGetUserLevelByUids //用户信息
}

//服务订单出参
type ReplyServiceOrderList struct {
	TotalNum   int
	Lists      []ReplyServiceOrderListBase
	StaffIndex []staff.ReplyGetListByStaffIds2 // 员工数据
	ImgIndex   map[int]file2.ReplyFileInfo     // 图片数据
	DefaultImgs map[int]file2.ReplyFileInfo
}

//获取商品订单列表
type ArgsProductOrder struct {
	common.Paging
	common.BsToken
	Uid       int
	Flag      bool //默认是 false就是查询商店的订单  true是查询用户的商品订单
	StartTime string
	EndTime   string
	DateType  int
	PayType   int
	KeyWord   string // 待完善
}

//获取商品订单列表返回
type ReplyProductOrder struct {
	TotalNum int
	Lists    []ProductOrder
	//Image map[int]file2.ReplyFileInfo
}
type ProductOrder struct {
	OrderSn        string `mapstructure:"order_sn"`   //商品订单
	OrderStatus    int    `mapstructure:"pay_status"` //订单状态
	PickUpGoodsStatus int
	ShopId         int    `mapstructure:"shop_id"`    //门店id
	OrderId        int    //支付订单ID
	ShopName       string
	BranchName     string
	CreateTime     int `mapstructure:"create_time"` //下单时间
	CreateTimeStr  string
	PayTime        int64                       // 支付付款时间
	PayTimeStr     string                      // 支付付款时间
	RealPrice      float64                     `mapstructure:"real_price"` //订单实际支付金额
	DiscountPrice  float64                     // 优惠的总金额
	OldTotalAmount float64                     // 合计金额（优惠总金额+实际支付金额）
	BuyerInfo      bus.ReplyGetUserLevelByUids //用户信息
	ProductDetails []ProductDetail             //商品订单明细
}

type ReplyOneProductDetail struct {
	OrderSn        string `mapstructure:"order_sn"`   //商品订单
	OrderStatus    int    `mapstructure:"pay_status"` //订单状态
	ShopId         int    `mapstructure:"shop_id"`    //门店id
	ShopName       string
	BranchName     string
	CreateTime     int `mapstructure:"create_time"` //下单时间
	CreateTimeStr  string
	PayTime        int64                       // 支付付款时间
	PayTimeStr     string                      // 支付付款时间
	RealPrice      float64                     `mapstructure:"real_price"` //订单实际支付金额
	DiscountPrice  float64                     // 优惠的总金额
	OldTotalAmount float64                     // 合计金额（优惠总金额+实际支付金额）
	BuyerInfo      bus.ReplyGetUserLevelByUids //用户信息
	ProductDetails []ProductDetail             //商品订单明细
	OrderSource    string                      //订单类型
	PaySn          string                      //支付流水号
	PayType        int                         //支付付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
}

type ReplyOneServerDetail struct {
	OrderSn        string  // 支付订单编号
	OrderId        int     //支付订单ID
	Uid            int     // 下单用户Uid
	PayStatus      int     //支付状态 0=待支付 1=已支付 2=支付失败 3=订单关闭
	TotalAmount    float64 // 支付订单总金额
	RealPrice      float64 // 支付订单实际支付金额
	DiscountPrice  float64 //优惠总金额
	OldTotalAmount float64 // 合计金额（优惠总金额+实际支付金额）
	OrderType      int     //支付订单类型 1=单项目订单 2=卡项订单
	CreateTime     int64   // 支付订单生成时间
	CreateTimeStr  string  // 支付订单生成时间
	PayTime        int64   // 支付付款时间
	PayTimeStr     string  // 支付付款时间
	BusId          int
	ShopId         int
	ShopName       string
	BranchName     string
	SubOrderLists  []SubOrderListsBase2
	BuyerInfo      bus.ReplyGetUserLevelByUids //用户信息
	PaySn          string                      // 支付流水号-支付渠道返回支付流水号
	OrderSource    string                      //订单类型
	PayType        int                         // 支付付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
}

type SubOrderListsBase2 struct {
	Id                int    //卡项/单项目子订单
	CardName          string // 卡包名
	SingleName        string //单项目名称
	ImgId             int    // 单项目/卡项封面图片
	ImgPath           string
	SspId             int     // 规格ID
	SspName           string  // 规格名
	Price             float64 // 卡项/单项目面值
	RealPrice         float64 // 卡项/单项目销售金额
	DiscountPrice     float64 //改价后的金额
	TotalAmount       float64 // 实际付款：如果改价则使用改价的价格
	Num               int     // 购买数量
	Mid               int     // 购买用户会员ID
	StaffIds          string  // 销售人员
	StaffNames        string
	CraftStaffIds     string // 手艺人
	CraftStaffNames   string
	SubOrderId        int //所属订单ID@kc_single_order表中的id字段
	OrderType         int //子订单类型 1=会员订单 2=游客订单
	Status            int //子订单状态 0=待支付 1=已支付 2=支付失败 3=订单关闭
	CardPackageStatus int //卡包状态 1=待消费 2=消费中  3=已完成 4=已关闭
}

type ProductDetail struct {
	Id int //商品子订单id
	ProductId     int     `mapstructure:"product_id"`
	DetailId      int     `mapstructure:"detail_id"`
	ProductName   string  `mapstructure:"product_name"`
	SpecName      string  `mapstructure:"detail_name"`
	RealPrice     float64 // 卡项/单项目销售金额
	DiscountPrice float64 `mapstructure:"discount_price"` //改动后的价格
	Num           int     `mapstructure:"num"`
	ImgId         int     `mapstructure:"img_id"`
	ImgPath       string
}

//获取用户提货单入参
type ArgsPickUpGoods struct {
	Uid    int
	Status int
	Lng    float64
	Lat    float64
	common.Paging
}

type ReplyPickUpGoods struct {
	TotalNum int
	Lists    []PickUpGoods
}
type PickUpGoods struct {
	OrderId int
	ShopId            int
	ShopName          string
	BranchName        string
	ShopImg           int
	ShopImgUrl        string
	ShopImgHash       string
	ShopAddress       string
	Distance          float64
	Longitude         float64
	Latitude          float64
	PickUpGoodsCode   string
	PickUpGoodsStatus int
	CreateTime        int
	CreateTimeStr     string
	ProductDetails    []ProductDetail
}

type ArgsGetOneServerDetail struct {
	common.BsToken
	OrderId int
}

// 统计用户复购率入参
type ArgsUserPurchaseRate struct {
	BusId int //
}
type ReplyUserPurchaseRate struct {
	UserPurchaseRate float64 // 会员复购率
}

type OrderPay interface {
	//获取支付签名
	GetPaySign(ctx context.Context, args *ArgsGetPaySign, reply *ReplyGetPayInfo) error
	//获取支付信息
	GetPayInfo(ctx context.Context, args *ArgsGetPayInfo, reply *ReplyGetPayInfo) error
	//支付成功异步任务具体业务处理
	PaySuc(ctx context.Context, orderSn *string, reply *bool) error
	//获取订单分账信息
	GetOrderSplitBill(ctx context.Context, orderSn string, reply *ReplyGetOrderSplitBill) error
	//现金支付
	CashPaySuc(ctx context.Context, args *ArgsCashPaySuc, reply *bool) error
	//获取支付状态信息 0=待支付 1=支付成功 2=支付失败
	QueryPayStatus(ctx context.Context, orderSn *string, reply *int) error
	//订单-获取服务订单列表
	GetServiceOrderList(ctx context.Context, args *ArgsServiceOrderList, reply *ReplyServiceOrderList) error
	//
	//获取商品订单信息
	GetProductOrderList(ctx context.Context, args *ArgsProductOrder, reply *ReplyProductOrder) error
	//获取用户提货单
	GetUserPickUpGoods(ctx context.Context, args *ArgsPickUpGoods, reply *ReplyPickUpGoods) error
	//订单超时未支付，关闭订单
	CloseOrder(ctx context.Context, orderId *int, reply *bool) error
	//获取一条商品订单详情信息
	GetOneProductDetail(ctx context.Context, args int, reply *ReplyOneProductDetail) error
	//获取一条服务订单详情信息
	GetOneServerDetail(ctx context.Context, args *ArgsGetOneServerDetail, reply *ReplyOneServerDetail) error
	// 统计用户复购率
	GetUsePurchaseRate(ctx context.Context, args *ArgsUserPurchaseRate, reply *ReplyUserPurchaseRate) error
}
