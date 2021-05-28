package financial

import "context"

//确认消费时，根据付款时间释放银行的存管金额-入参
type ArgsConsumeRelaseDepositAssets struct {
	BusId        int
	RelaseAmount float64 //释放的存管金额
	PayTime      string  //卡项付款时间
	BankType     int     //存管银行渠道 1=上海银行 2=交通银行
}

//ReplyMonthTotalNumList 月存管商户数出参list
type ReplyMonthTotalNumList struct {
	List []ReplyMonthTotalNum `mapstructure:"list" json:"list"`
}

//ReplyMonthTotalNum 月存管商户数出参
type ReplyMonthTotalNum struct {
	DateMonth string `mapstructure:"dateMonth" json:"dateMonth"`
	BusNum    string `mapstructure:"busNum" json:"busNum"`
}

//总数-出参
type ReplyTotalNum struct {
	TotalNum int
}

//ReplyTotalAssetsMonthList 每月存管金额-出参list
type ReplyTotalAssetsMonthList struct {
	List []ReplyTotalAssetsMonth `mapstructure:"list" json:"list"`
}

//ReplyTotalAssetsMonth 每月存管金额-出参
type ReplyTotalAssetsMonth struct {
	DateMonth     string `mapstructure:"dateMonth" json:"dateMonth"`
	DepositAssets string `mapstructure:"depositAssets" json:"depositAssets"`
}

//总金额-出参
type ReplyTotalAssets struct {
	TotalAssets float64
}

//获取银行商户总数-入参
type ArgsGetBankBusMonthNum struct {
	BankType  int    `mapstructure:"bankType" form:"bankType" json:"bankType"`    //存管银行渠道
	StartDate string `mapstructure:"startDate" form:"startDate" json:"startDate"` //开始时间节点，格式："200601"
	EndDate   string `mapstructure:"endDate" form:"endDate" json:"endDate"`       //结束时间节点
}

//获取银行商户总数-入参
type ArgsGetBankBusNum struct {
	BankType int //存管银行渠道
}

//ArgsGetTopBank  获取银行top-入参
type ArgsGetTopBank struct {
	Limit     int    `mapstructure:"limit" form:"limit" json:"limit"`             //limit
	SortMode  int    `mapstructure:"sortMode" form:"sortMode" json:"sortMode"`    //排序方式：0 降序， 1 升序    默认0
	SortField string `mapstructure:"sortField" form:"sortField" json:"sortField"` //排序字段
}

//ReplyTopBankList 获取top银行-出参
type ReplyTopBankList struct {
	List []ReplyTopBank
}

//ReplyTopBank 获取top银行-出参
type ReplyTopBank struct {
	BankType      string `mapstructure:"bankType" json:"bankType"` //存管银行渠道 1=上海银行 2=交通银行
	BusNum        string `mapstructure:"busNum" json:"busNum"`
	DepositAssets string `mapstructure:"depositAssets" json:"depositAssets"`
}

//获取银行当前金额-入参
type ArgsGetBankCurrentAssets struct {
	BankType  int    //存管银行渠道
	StartDate string //开始时间节点，格式："200601"
	EndDate   string //结束时间节点
}

//获取银行累计金额-入参
type ArgsGetBankTotalAssets struct {
	BankType int //存管银行渠道
}

//获取银行商户总数-入参
type ArgsGetReportBankBusNum struct {
	BankType  int    `mapstructure:"bankType" form:"bankType" json:"bankType"`    //存管银行渠道
	StartDate string `mapstructure:"startDate" form:"startDate" json:"startDate"` //开始时间节点，格式："200601"
	EndDate   string `mapstructure:"endDate" form:"endDate" json:"endDate"`       //结束时间节点
}

type ReplyGetAllBank struct {
	BankType int
	BankName string
}

type Bank interface {
	//统计银行/保险-商户数量
	StatisticsBankOrInsuranceBusNum(ctx context.Context, busId *int, reply *bool) error
	//统计银行/保险-发卡金额(购卡成功后消费调用)
	StatisticsBankOrInsuranceSalesCardAssets(ctx context.Context, orderSn *string, reply *bool) error
	//统计银行-存管金额（购卡/复充成功后消费调用）
	StatisticsBankDepositAssets(ctx context.Context, orderSn *string, reply *bool) error
	//确认消费时，根据付款时间释放银行的存管金额
	ConsumeRelaseDepositAssetsRpc(ctx context.Context, args *ArgsConsumeRelaseDepositAssets, reply *bool) error

	//================================API接口================================
	//获取top银行商户总数
	GetTopBankBusNum(ctx context.Context, args *ArgsGetTopBank, reply *ReplyTopBankList) error
	//获取top银存管金额
	GetTopBankDepositAssets(ctx context.Context, args *ArgsGetTopBank, reply *ReplyTopBankList) error
	//获取月银行商户总数
	GetMonthBankBusNum(ctx context.Context, args *ArgsGetBankBusMonthNum, reply *ReplyMonthTotalNumList) error
	//获取银行商户总数
	GetBankBusNum(ctx context.Context, args *ArgsGetBankBusNum, reply *ReplyTotalNum) error
	//获取银行每月存管金额
	GetBankDepositAssetsMonth(ctx context.Context, args *ArgsGetBankCurrentAssets, reply *ReplyTotalAssetsMonthList) error
	//获取银行当前存管金额
	GetBankCurrentDepositAssets(ctx context.Context, args *ArgsGetBankCurrentAssets, reply *ReplyTotalAssets) error
	//获取银行累计存管金额
	GetBankTotalDepositAssets(ctx context.Context, args *ArgsGetBankTotalAssets, reply *ReplyTotalAssets) error
	//获取银行当前发卡金额
	GetBankCurrentSalesCardAssets(ctx context.Context, args *ArgsGetBankCurrentAssets, reply *ReplyTotalAssets) error
	//获取银行累计发卡金额
	GetBankTotalSalesCardAssets(ctx context.Context, args *ArgsGetBankTotalAssets, reply *ReplyTotalAssets) error
	//获取银行当前普惠金融业绩
	GetBankCurrentInclusiveFinanceAssets(ctx context.Context, args *ArgsGetBankCurrentAssets, reply *ReplyTotalAssets) error
	//获取银行累计普惠金融业绩
	GetBankTotalInclusiveFinanceAssets(ctx context.Context, args *ArgsGetBankTotalAssets, reply *ReplyTotalAssets) error
	//获取所有支持的银行
	GetAllBank(ctx context.Context, args *int, reply *[]ReplyGetAllBank) error
}
