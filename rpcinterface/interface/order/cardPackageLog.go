package order

import "git.900sui.cn/kc/rpcinterface/interface/common"

// 卡包消费记录
// @author liyang<654516092@qq.com>
// @date  2020/8/13 13:52

//公共卡包消费data数据
type CardPackageCommonConsumeData struct {
	//{"singleId":2,"num":1,"price":120,"singleName":"洗发","sspName":"黑色","sspId":1,"staffId":[1]}
	SingleId   int     `json:"singleId"`
	SingleName string  `json:"singleName"`
	SspId      int     `json:"sspId"`
	SspName    string  `json:"sspName"`
	Num        int     `json:"num"`
	Price      float64 `json:"price"`
	StaffId    []int   `json:"staffId"`
}

//综合卡/充值卡卡包消费data数据
type CardPackageCardConsumeData struct {
	SingleId   int     `json:"singleId"`
	SingleName string  `json:"singleName"`
	SspId      int     `json:"sspId"`
	SspName    string  `json:"sspName"`
	Type       int     `json:"type"` //类型 0=消费服务 1=消费产品
	Num        int     `json:"num"`
	Price      float64 `json:"price"`
	StaffId    []int   `json:"staffId"`
}

//返回卡包消费记录入参
type ArgsConsumeDataLog struct {
	RelationId int
	common.Paging
}

//返回卡包消费记录
type ReplyConsumeDataLog struct {
	TotalNum int
	Lists    []ReplyConsumeDataList
	Staff    map[int]ReplyConsumeStaff
}

//返回卡包消费记录详情LIST
type ReplyConsumeDataList struct {
	LogId           int `mapstructure:"id"`
	RelationId      int
	CardPackageId   int
	CardPackageSn   string `mapstructure:"card_package_card_sn"`
	CardPackageType int
	CardLogId       int
	BusId           int
	ShopId          int
	Uid             int
	ConsumeComp     int //是否消费完成
	Ctime           int64
	CtimeStr        string
	ConsumeData     ReplyConsumeData
	ConsumeDataConf []ReplyConsumeDataConf
}

//ConsumeStaff
type ReplyConsumeStaff struct {
	StaffId  int
	Name     string
	NickName string
}

//ConsumeData节点数据
type ReplyConsumeData struct {
	ConsumePrice       float64
	ActualConsumePrice float64
	ConsumeType        int
}

//ConsumeDataJson节点数据
type ReplyConsumeDataComplie struct {
	ConsumePrice       float64
	ActualConsumePrice float64
	ConfData           string
}

//ConsumeDataConf 节点数据
type ReplyConsumeDataConf struct {
	SingleId   int
	SingleName string
	SspId      int
	SspName    string
	Type       int
	Num        int
	Price      float64
	StaffId    []int
}

//充值记录
type ArgsRechangeLog struct {
	common.Paging
	RelationId int
}
type RechangeLogBase struct {
	Id            int //充值订单记录id
	BusId         int
	ShopId        int
	ShopName      string  //分店门店名称
	BranchName    string  //分店名称
	CardPackageId int     //卡包id
	RechargeSn    string  //充值订单编号
	RealPrice     float64 //充值金额
	Price         float64 //充值面值
	Disaccount    float64 //折扣率
	Status        int     //卡状态：0=可消费 1=消费完毕
	RechargeType  int     //充值类型  0=购卡充值 1=复充值
	PayChannel    int     //支付渠道 1=原生支付 2=杉德支付 3=建行直连
	PayTime       int64   //付款时间
	PayTimeStr    string
	//ConsumeData     ReplyConsumeData
}
type ReplyRechangeLog struct {
	TotalNum int
	Lists    []RechangeLogBase
}

//根据充值编号获取充值记录
type ArgsGetRcardRechargeBySn struct {
	RechargeSn string //充值记录编号
}
type ReplyGetRcardRechargeBySnBase struct {
	Id               int //充值订单记录id
	BusId            int
	ShopId           int
	SubOrderId       int     //订单ID@kc_card_order_card表中的id字段
	CardPackageId    int     //卡包id
	RechargeSn       string  //充值订单编号
	RealPrice        float64 //充值金额
	PayTime          int64   //付款时间
	DepositRatio     float64 //留存比例
	DepositAmount    float64 //留存金额
	FundMode         int     //资金管理方式 0=无管理方式 1=资金存管 2=保证保险
	InsuranceChannel int     //保险渠道 0=无保险 1=长安保险 2=宁波人保 3=上海安信保险
	DeposBankChannel int     //该笔订单存管银行渠道 1=上海银行 2=交通银行
}
type ReplyGetRcardRechargeBySn struct {
	Lists []ReplyGetRcardRechargeBySnBase
}
