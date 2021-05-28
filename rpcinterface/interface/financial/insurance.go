/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/25 上午10:32

*******************************************/

package financial

import "context"

// 保险公司上月保费金额入参数
type ArgsInsAssetsMonth struct {
	TransNo string `json:"trans_no"` // 流水号
}

//保险公司累计保费金额
type ArgsInsAssetsTotal struct {
	TransNo string `json:"trans_no"` // 流水号
}

//保险公司上个月场景保险业务业绩
type ArgsSceneInsMonth struct {
	SceneInsuranceAssets float64 `json:"scene_insurance_assets"` //场景保险业务业绩
	InsuranceType        string  `json:"insurance_type"`         //保险渠道 1=长安保险 2=人保保险 3=安信保险
	DateMonth            string  `json:"date_month"`             //统计时间间隔，单位月，格式如201810
}

//保险公司累计场景保险业务业绩
type ArgsSceneInsTotal struct {
	SceneInsuranceAssets float64 `json:"scene_insurance_assets"` //场景保险业务业绩
	InsuranceType        string  `json:"insurance_type"`         //保险渠道 1=长安保险 2=人保保险 3=安信保险
}

// 获取保险公司商户统计入参数
type ArgsInsBusType struct {
	InsuranceType int // 保险渠道 1=长安保险 2=人保保险 3=安信保险
}

// 返回保险公司商户总数量
type ReplyBusNum struct {
	BusNum int //保险公司删铺总数量
}

// 获取保险公司保费金额入参
type ArgsGetAssetsMonth struct {
	InsuranceType int    // 保险渠道 1=长安保险 2=人保保险 3=安信保险
	StartMonth    string // 开始月份时间
	EndMonth      string // 结束月份时间
}

// 返回保险公司保费金额参数
type ReplySalesCardMonth struct {
	SalesCardMonth float64
}

// 获取保险公司累计保险金额入参
type ArgsGetAssetsAmount struct {
	InsuranceType int // 保险渠道 1=长安保险 2=人保保险 3=安信保险
}

// 返回保险公司累计保险金额
type ReplySalesCardAmount struct {
	SalesCardAmount float64
}

// 获取当月保险发卡金额入参数
type ArgsIssueCardMonth struct {
	InsuranceType int    // 保险渠道 1=长安保险 2=人保保险 3=安信保险
	StartMonth    string // 开始月份时间
	EndMonth      string // 结束月份时间
}

// 返回当月保险发卡金额
type ReplyIssueCardMonth struct {
	SalesCardMonth float64 // 月统计发卡金额
}

// 获取累计保险发卡金额统计
type ArgsIssueCardAmount struct {
	InsuranceType int // 保险渠道 1=长安保险 2=人保保险 3=安信保险
}

// 返回累计保险发卡金额统计
type ReplyIssueCardAmount struct {
	SalesCardAmount float64 // 保险类型累计发卡金额
}

//ArgsGetTopInsurance  获取保险top-入参
type ArgsGetTopInsurance struct {
	Limit     int    `mapstructure:"limit" form:"limit" json:"limit"`             //limit
	SortMode  int    `mapstructure:"sortMode" form:"sortMode" json:"sortMode"`    //排序方式：0 降序， 1 升序    默认0
	SortField string `mapstructure:"sortField" form:"sortField" json:"sortField"` //排序字段
}

//ReplyTopInsuranceList 获取top保险-出参
type ReplyTopInsuranceList struct {
	List []ReplyTopInsurance
}

//ReplyTopInsurance 获取top保险商户总数-出参
type ReplyTopInsurance struct {
	InsuranceType   string `mapstructure:"insuranceType" json:"insuranceType"` //保险渠道 1=长安保险 2=人保保险 3=安信保险
	BusNum          string `mapstructure:"busNum" json:"busNum"`
	InsuranceAssets string `mapstructure:"insuranceAssets" json:"insuranceAssets"`
}

type ReplyGetAllInsurance struct {
	InsuranceType int
	InsuranceName string
}

//ReplyGetMonthInsBusNumList 获取每月保险公司商户统计list
type ReplyGetMonthInsBusNumList struct {
	List []ReplyGetMonthInsBusNum `mapstructure:"list" json:"list"`
}

//ReplyGetMonthInsBusNum 获取每月保险公司商户统计
type ReplyGetMonthInsBusNum struct {
	DateMonth string `mapstructure:"dateMonth" json:"dateMonth"`
	BusNum    string `mapstructure:"busNum" json:"busNum"`
}

//ReplyGetMonthInsAssetsMonthList 每月存管金额-出参list
type ReplyGetMonthInsAssetsMonthList struct {
	List []ReplyGetMonthInsAssetsMonth `mapstructure:"list" json:"list"`
}

//ReplyGetMonthInsAssetsMonth 每月存管金额-出参
type ReplyGetMonthInsAssetsMonth struct {
	DateMonth       string `mapstructure:"dateMonth" json:"dateMonth"`
	InsuranceAssets string `mapstructure:"insuranceAssets" json:"insuranceAssets"`
}

type Insurance interface {
	//保险公司上月保费金额
	AddInsAssetsMonth(ctx context.Context, args *ArgsInsAssetsMonth, reply *bool) error
	//保险公司累计保费金额
	AddInsAssetsTotal(ctx context.Context, args *ArgsInsAssetsTotal, reply *bool) error

	//保险公司上个月场景保险业务业绩 ******预留暂时场景保险还没有做******
	//AddSceneInsMonth(ctx context.Context, args *ArgsSceneInsMonth, reply *bool) error
	//保险公司累计场景保险业务业绩
	//AddSceneInsTotal(ctx context.Context, args *ArgsSceneInsTotal, reply *bool) error

	//================================API接口================================
	// 获取top保险公司商户统计
	GetTopInsBusTotal(ctx context.Context, args *ArgsGetTopInsurance, reply *ReplyTopInsuranceList) error
	// 获取top保险公司累计保费金额
	GetTopInsAssetsAmount(ctx context.Context, args *ArgsGetTopInsurance, reply *ReplyTopInsuranceList) error
	// 获取保险公司商户统计
	GetInsBusTotal(ctx context.Context, args *ArgsInsBusType, reply *ReplyBusNum) error
	// 获取每月保险公司商户统计
	GetMonthInsBusNum(ctx context.Context, args *ArgsGetAssetsMonth, reply *ReplyGetMonthInsBusNumList) error
	// 获取每月保险公司保费金额
	GetMonthInsAssetsMonth(ctx context.Context, args *ArgsGetAssetsMonth, reply *ReplyGetMonthInsAssetsMonthList) error
	// 获取保险公司月保费金额
	GetInsAssetsMonth(ctx context.Context, args *ArgsGetAssetsMonth, reply *ReplySalesCardMonth) error
	// 获取保险公司累计保费金额
	GetInsAssetsAmount(ctx context.Context, args *ArgsGetAssetsAmount, reply *ReplySalesCardAmount) error
	// 获取当月保险发卡金额
	GetInsIssueCardMonth(ctx context.Context, args *ArgsIssueCardMonth, reply *ReplyIssueCardMonth) error
	// 获取累计保险发卡金额统计
	GetInsIssueCardAmount(ctx context.Context, args *ArgsIssueCardAmount, reply *ReplyIssueCardAmount) error
	//获取所有保险公司类型
	GetAllInsurance(ctx context.Context, args *int, reply *[]ReplyGetAllInsurance) error
}
