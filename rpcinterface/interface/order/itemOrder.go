//卡项订单接口定义
//@author yangzhiwu<578154898@qq.com>
//@date 2020/7/27 16:53

package order

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//下单卡项数据结构
type ItemData struct {
	SsId        int     //目在门店的id
	StaffIds    []int   //销售ids
	Num         int     //购买数量
	ChangePrice float64 //手动改价后的单件价格
}

//生成卡项订单参数数据结构
type ItemOrderData struct {
	Items       []ItemData       //购买的项目
	Gives       []GiveSingleBase //购卡时商家赠送的项目
	OrderSource int              //订单渠道 使用channel
	ItemType    int              //项目类型 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡
}

//saas后台下单的入参数据
type ArgsSaasCreateItemOrder struct {
	common.BsToken
	common.Utoken //当前操作人员登录信息
	ItemOrderData
	Uid int //购买人uid

}

//saas后台下单返回数据结构
type ReplySaasCreateItemOrder struct {
	OrderSn     string    //订单号
	Ctime       int64     //下单时间
	TotalAmount float64   //订单金额
	PayTypes    []PayType //saas的付款方式
}

//前端用户下单的入参数据
type ArgsUserCreateItemOrder struct {
	common.Utoken
	ShopId int //门店id
	ItemOrderData
}

//前端用户下单返回的数据格式
type ReplyUserCreateItemOrder struct {
	OrderSn     string    //订单号
	TotalAmount float64   //订单金额
	PayTypes    []PayType //前端的付款方式
}

//订单包含的项目数据结构
type OrderItemData struct {
	ItemId        int     //项目id
	DetailId      int     //产品的明细od
	SspId         int     //单项目规格组合id
	Num           int     //购买数量
	Price         float64 // 面值
	Source        int     //订单来源
	DiscountPrice float64 //售价
}

//根据订单号，获取订单的详细 数据结构，目前先定义这些需要再补充
type ReplyGetOrderInfoByOrderSnRpc struct {
	OrderId          int             //支付总订单id
	BusId            int             //商家id
	ShopId           int             //门店id
	Uid              int             //用户uid
	TotalAmount      float64         //总支付金额
	RealPrice        float64         // 卡项实际支付金额
	PayStatus        int             //支付状态
	PayTime          int64           //支付时间
	PayType          int             //付款方式 1=支付宝 2=微信 3=现金 4=渠道原生支付
	PayChannel       int             //支付渠道 支付渠道 1=原生支付 2=杉德支付 3=杭州建行直连 4=工商银行
	CreateTime       int64           //下单时间
	OrderType        int             //1=单项目订单 2=卡项订单  3=商品订单
	CardOrderType    int             //卡项的类型 2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡
	Source           int             //订单来源 0=普通订单 1=一元体验订单 2=拼团订单 3=充值卡复充
	DepositAmount    float64         //存管金额
	DeposBankChannel int             //资金存管银行渠道 1=上海银行 2=交通银行
	FundMode         int             //资金管理方式 0=无管理方式 1=资金存管 2=保证保险
	InsuranceChannel int             //保险渠道 0=无保险 1=长安保险 2=宁波人保 3=上海安信保险
	OrderItems       []OrderItemData //包含的项目

}

// 获取卡项目购卡总数入参
type ArgsBuyCardNum struct {
	BusId int // 商铺ID

}
type ReplyBuyCardNum struct {
	RiskBusId       int     // 风控系统商铺的ID
	CardNum         int     // 购卡总数（卡项的，不包括单项目）
	SingleNum       int     // 不包括单项目总数
	SalesCardAssets float64 // 累计发（售）卡总额 (单位:万元)
}

//子订单信息
type OrderDetail struct {
	CardPackageId int     //卡包id
	ItemId        int     //项目id
	Name          string  //项目名称
	RealPrice     float64 //购卡金额
	DepositAmount float64 //存管金额
	ReleaseAmount float64 //已释放的留存金额
	RlogId        int     //充值记录id
}

//商家订单存管明细返回值-入参
type ArsgOrderDeposLists struct {
	common.Paging
	common.BsToken
	OrderType int //订单类型
}

type ReplyOrderDeposListsBase struct {
	Id                int           //自增id
	BusId             int           //企业id
	Uid               int           //下单人
	PorderId          int           //订单id
	OrderSource       int           //订单来源
	OrderType         int           //订单类型
	OrderSn           string        //订单号
	PayPrice          float64       //实际支付金额
	DepositRatio      float64       //留存比例
	DepositAmount     float64       //留存金额
	InsureAmount      float64       //保险费用
	RenewInsureAmount float64       //续保费用
	PaymentAmount     float64       //支付手续费
	PlatAmount        float64       //平台手续费
	PayTime           int           //支付时间
	PayTimeStr        string        //支付时间戳
	PaySn             string        //支付流水号
	ReleaseAmount     float64       //已释放的留存金额
	FundType          int           //资金管理方式 1=资金存管 2=保证保险
	ItemOrderInfo     []OrderDetail //项目详情信息
}

//商家订单存管明细返回值-出参
type ReplyOrderDeposLists struct {
	TotalNum int
	Lists    []ReplyOrderDeposListsBase
}

//卡包存管资金释放明细入参
type ArgsGetOrderDeposReleaseInfo struct {
	common.Paging
	common.BsToken
	OrderType     int //订单类型
	CardPackageId int //卡包id
	RlogId        int //充值记录id
}
type GetOrderDeposReleaseInfoBase struct {
	ActualConsumePrice float64 //实际消费金额
	ReleaseAmount      float64 //已释放的留存金额
	Ctime              int64   //消费时间
	CtimeStr           string  //消费时间戳
}

//卡包存管资金释放明细出参
type ReplyGetOrderDeposReleaseInfo struct {
	TotalNum int
	Lists    []GetOrderDeposReleaseInfoBase
}

type ArgsGetBusOrderListsForGov struct {
	BusId int
	common.Paging
}

type GetBusOrderListsForGovList struct {
	OrderSn            string //订单号
	CardType           int    //卡类型
	CardTypeName       string //卡类型名称
	RealPrice          string //付款金额
	ActualConsumePrice string //兑付金额
	DepositAmount      string //存管金额
	ReleaseAmount      string //释放金额
	CardPackageId      string //卡包id
	Ctime              string //付款时间

}

type ReplyGetBusOrderListsForGov struct {
	TotalNum int
	Lists    []GetBusOrderListsForGovList
}

type GetConsumeLogForGovList struct {
	ShopName           string //消费门店
	ActualConsumePrice string //实际消费金额
	Ctime              string //消费时间
	ShopId             string //门店id
}

type ReplyGetConsumeLogForGov struct {
	TotalNum int
	Lists    []GetConsumeLogForGovList
}

type ArgsGetConsumeLogForGov struct {
	CardPackageId   int //卡包id
	CardPackageType int //卡包类型
}

type ItemOrder interface {

	//SaaS创建卡项订单方法
	SaasCreateItemOrder(ctx context.Context, args *ArgsSaasCreateItemOrder, reply *ReplySaasCreateItemOrder) error
	//用户创建卡项订单
	UserCreateItemOrder(ctx context.Context, args *ArgsUserCreateItemOrder, reply *ReplyUserCreateItemOrder) error
	//根据订单号，获取订单的详细-rpc使用
	GetOrderInfoByOrderSnRpc(ctx context.Context, orderSn *string, reply *ReplyGetOrderInfoByOrderSnRpc) error
	// 获取卡项目购卡总数、（售）卡总额 、（消费）金额（对应的店铺）
	GetBuyCardNum(ctx context.Context, args *ArgsBuyCardNum, reply *ReplyBuyCardNum) error
	//获取商家订单存管明细数据
	GetOrderDeposLists(ctx context.Context, args *ArsgOrderDeposLists, reply *ReplyOrderDeposLists) error
	//获取卡包存管资金释放明细
	GetOrderDeposReleaseInfo(ctx context.Context, args *ArgsGetOrderDeposReleaseInfo, reply *ReplyGetOrderDeposReleaseInfo) error
	//根据商家id，获取商家卡包信息 - 监管平台直连接口
	GetBusOrderListsForGov(ctx context.Context, args *ArgsGetBusOrderListsForGov, reply *ReplyGetBusOrderListsForGov) error
	//获取卡包的消费记录 - 监管平台直连接口
	GetConsumeLogForGov(ctx context.Context, args *ArgsGetConsumeLogForGov, reply *ReplyGetConsumeLogForGov) error
}
