package financial

import "context"

type TpayData struct {
	BusNum   int    //商家数量
	Rate     string //支付费率
	TpayType int    //类型
	Assets   string //累计收单金额
}

//ArgsGetTopTpay  获取top第三方支付-入参
type ArgsGetTopTpay struct {
	Limit     int    `mapstructure:"limit" form:"limit" json:"limit"`             //limit
	SortMode  int    `mapstructure:"sortMode" form:"sortMode" json:"sortMode"`    //排序方式：0 降序， 1 升序    默认0
	SortField string `mapstructure:"sortField" form:"sortField" json:"sortField"` //排序字段
}

//ReplyTopTpayList 获取top第三方支付的商户总数-出参
type ReplyTopTpayList struct {
	List []ReplyTopTpay
}

//ReplyTopTpay 获取top第三方支付的商户总数-出参
type ReplyTopTpay struct {
	TpayType string `mapstructure:"tpayType" json:"tpayType"` //第三方支付类型 2=杉德支付 3=杭州建行直连 4=工行支付
	BusNum   string `mapstructure:"busNum" json:"busNum"`
	Assets   string `mapstructure:"assets" json:"assets"`
}

//ReplyGetTpayBusNumMonthList 获取每月商户统计list
type ReplyGetTpayBusNumMonthList struct {
	List []ReplyGetTpayBusNumMonth `mapstructure:"list" json:"list"`
}

//ReplyGetTpayBusNumMonth 获取每月商户统计
type ReplyGetTpayBusNumMonth struct {
	DateMonth string `mapstructure:"dateMonth" json:"dateMonth"`
	BusNum    string `mapstructure:"busNum" json:"busNum"`
}

//ReplyGetTpayAssetsMonthList 每月存管金额-出参list
type ReplyGetTpayAssetsMonthList struct {
	List []ReplyGetTpayAssetsMonth `mapstructure:"list" json:"list"`
}

//ReplyGetTpayAssetsMonth 每月存管金额-出参
type ReplyGetTpayAssetsMonth struct {
	DateMonth string `mapstructure:"dateMonth" json:"dateMonth"`
	Assets    string `mapstructure:"assets" json:"assets"`
}

//ArgsGetTpay 获取第三方支付-入参
type ArgsGetTpay struct {
	TpayType   int    // 第三方支付类型 2=杉德支付 3=杭州建行直连 4=工行支付
	StartMonth string // 开始月份时间
	EndMonth   string // 结束月份时间
}

type ReplyGetAllTpay struct {
	TpayType int
	TpayName string
}
type Tpay interface {
	//获取top第三方支付的商户总数
	GetTopTpayBusNum(ctx context.Context, tpayType *ArgsGetTopTpay, reply *ReplyTopTpayList) error
	//获取top第三方支付的总收单金额
	GetTopTpayAssetsTotal(ctx context.Context, tpayType *ArgsGetTopTpay, reply *ReplyTopTpayList) error

	//获取每月第三方支付的商户总数
	GetTpayBusNumMonth(ctx context.Context, tpayType *ArgsGetTpay, reply *ReplyGetTpayBusNumMonthList) error
	//获取每月第三方支付的总收单金额
	GetTpayAssetsMonth(ctx context.Context, tpayType *ArgsGetTpay, reply *ReplyGetTpayAssetsMonthList) error

	//获取单个第三方支付的数据
	GetOneTpayData(ctx context.Context, tpayType *int, reply *TpayData) error
	//获取全部第三方支付的数据
	GetAllData(ctx context.Context, tpayType *int, reply *[]TpayData) error
	//获取所有第三方支付类型
	GetAllTpay(ctx context.Context, args *int, reply *[]ReplyGetAllTpay) error
}
