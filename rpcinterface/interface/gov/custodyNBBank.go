package gov

import (
	"context"
)

//InputParamsCustodyNBBankOpenAccount [开户申请 入参]
type InputParamsCustodyNBBankOpenAccount struct {
	NbcbID            string `mapstructure:"nbcbId" form:"nbcbId" json:"nbcbId"`                                  // 监管平台申请id
	MerchantID        string `mapstructure:"merchantId" form:"merchantId" json:"merchantId"`                      // 商家企业编号
	BusID             string `mapstructure:"busId" form:"busId" json:"busId"`                                     // 商家id
	ServiceID         string `mapstructure:"serviceId" form:"serviceId" json:"serviceId"`                         // 请求服务号
	ApplyNo           string `mapstructure:"applyNo" form:"applyNo" json:"applyNo"`                               // 申请编号
	ContactName       string `mapstructure:"contactName" form:"contactName" json:"contactName"`                   // 开户联系人姓名
	ContactIDType     string `mapstructure:"contactIdType" form:"contactIdType" json:"contactIdType"`             // 开户联系人证件种类
	ContactIDNo       string `mapstructure:"contactIdNo" form:"contactIdNo" json:"contactIdNo"`                   // 开户联系人证件号码
	ContactMobile     string `mapstructure:"contactMobile" form:"contactMobile" json:"contactMobile"`             // 开户联系人手机号
	CertType          string `mapstructure:"certType" form:"certType" json:"certType"`                            // 证件类型 只能为9.营业执照
	CertNo            string `mapstructure:"certNo" form:"certNo" json:"certNo"`                                  // 证件号码
	Name              string `mapstructure:"name" form:"name" json:"name"`                                        // 单位名称
	RegAddr           string `mapstructure:"regAddr" form:"regAddr" json:"regAddr"`                               // 注册地址
	Addr              string `mapstructure:"addr" form:"addr" json:"addr"`                                        // 办公地址
	Phone             string `mapstructure:"phone" form:"phone" json:"phone"`                                     // 办公电话
	LegalPersonName   string `mapstructure:"legalPersonName" form:"legalPersonName" json:"legalPersonName"`       // 法人姓名
	LegalPersonIDNo   string `mapstructure:"legalPersonIdNo" form:"legalPersonIdNo" json:"legalPersonIdNo"`       // 法人证件号码
	LegalPersonIDType string `mapstructure:"legalPersonIdType" form:"legalPersonIdType" json:"legalPersonIdType"` // 法人证件种类
	LegalPersonMobile string `mapstructure:"legalPersonMobile" form:"legalPersonMobile" json:"legalPersonMobile"` // 法人手机号
	SalesNo           string `mapstructure:"salesNo" form:"salesNo" json:"salesNo"`                               // 客户经理工号
}

//OutputParamsCustodyNBBankOpenAccount [开户申请 出参]
type OutputParamsCustodyNBBankOpenAccount struct {
	ApplyNo string `mapstructure:"applyNo" json:"applyNo"` // applyNo
	Message string `mapstructure:"message" json:"message"` // message 提示信息
}

//InputParamsCustodyNBBankBindingAccount [账户绑定 入参]
type InputParamsCustodyNBBankBindingAccount struct {
	ServiceID string `mapstructure:"serviceId" form:"serviceId" json:"serviceId"` // 请求服务号
	CusAc     string `mapstructure:"cusAc" form:"cusAc" json:"cusAc"`             // 存管账号
	AcNme150  string `mapstructure:"acNme150" form:"acNme150" json:"acNme150"`    // 存管账号名称
	BioFlg    string `mapstructure:"bioFlg" form:"bioFlg" json:"bioFlg"`          // 行内外标志 I-行内 (In首字母) O-行外 (Out首字母)
	CusAtr    string `mapstructure:"cusAtr" form:"cusAtr" json:"cusAtr"`          // 客户属性 P-个人 C-对公
	OppAc     string `mapstructure:"oppAc" form:"oppAc" json:"oppAc"`             // 绑定账号
	NBbankID  int    `mapstructure:"NBbankId" form:"NBbankId" json:"NBbankId"`    // 开户申请id@kc_nbbank.nbbank_id
}

//OutputParamsCustodyNBBankBindingAccount [账户绑定 出参]
type OutputParamsCustodyNBBankBindingAccount struct {
	Message string `mapstructure:"message" json:"message"` // message 提示信息
}

//InputParamsCustodyNBBankCashApply [资金提现 入参]
type InputParamsCustodyNBBankCashApply struct {
	ServiceID string `mapstructure:"serviceId" form:"serviceId" json:"serviceId"` // 请求服务号
	PayAc     string `mapstructure:"payAc" form:"payAc" json:"payAc"`             // 付款账号
	PayAcNm   string `mapstructure:"payAcNm" form:"payAcNm" json:"payAcNm"`       // 付款账号名称
	PaeAc     string `mapstructure:"paeAc" form:"paeAc" json:"paeAc"`             // 收款账号
	PaeAcNm   string `mapstructure:"paeAcNm" form:"paeAcNm" json:"paeAcNm"`       // 收款账号名称
	PaeBankno string `mapstructure:"paeBankno" form:"paeBankno" json:"paeBankno"` // 收款行号
	BioFlg    string `mapstructure:"bioFlg" form:"bioFlg" json:"bioFlg"`          // 行内外标志
	TxnAmt    string `mapstructure:"txnAmt" form:"txnAmt" json:"txnAmt"`          // 提现金额
	TxnRmk    string `mapstructure:"txnRmk" form:"txnRmk" json:"txnRmk"`          // 交易摘要
	PstMde    string `mapstructure:"pstMde" form:"pstMde" json:"pstMde"`          // 入账方式
	PayCcy    string `mapstructure:"payCcy" form:"payCcy" json:"payCcy"`          // 付款币种
	NBbankID  int    `mapstructure:"NBbankId" form:"NBbankId" json:"NBbankId"`    // 开户申请id@kc_nbbank.nbbank_id
}

//OutputParamsCustodyNBBankCashApply [资金提现 出参]
type OutputParamsCustodyNBBankCashApply struct {
	Message string `mapstructure:"message" json:"message"` // message 提示信息
}

// 请求报文头
// Sender    string `mapstructure:"sender" form:"sender" json:"sender"`          // 请求单位名称
// Receiver  string `mapstructure:"receiver" form:"receiver" json:"receiver"`    // 接收单位名称
// ServiceID string `mapstructure:"serviceId" form:"serviceId" json:"serviceId"` // 请求服务号
// MsgID     string `mapstructure:"msgId" form:"msgId" json:"msgId"`             // 报文标识号
// TimeStamp string `mapstructure:"timeStamp" form:"timeStamp" json:"timeStamp"` // 时间戳
// Version   string `mapstructure:"version" form:"version" json:"version"`       // 版本号
// Sign      string `mapstructure:"sign" form:"sign" json:"sign"`                // 加签域
// RandomKey string `mapstructure:"randomKey" form:"randomKey" json:"randomKey"` // 随机数加密域

//InputParamsCustodyNBBankOpenAccountNotify [开户状态通知 入参]
type InputParamsCustodyNBBankOpenAccountNotify struct {
	TransCode     string `mapstructure:"transCode" form:"transCode" json:"transCode"`             // 交易类型 001-预约账户生成 002-开户成功 003-开户失败/预约账号生成失败
	OriApplyNo    string `mapstructure:"oriApplyNo" form:"oriApplyNo" json:"oriApplyNo"`          // 原申请编号
	AccountNo     string `mapstructure:"accountNo" form:"accountNo" json:"accountNo"`             // 账号
	AccountName   string `mapstructure:"accountName" form:"accountName" json:"accountName"`       // 账户名称
	AccountStatus string `mapstructure:"accountStatus" form:"accountStatus" json:"accountStatus"` // 开户状态 0-初始状态 1-开户申请提交成功 2-开户申请提交失败 3-预约账号生4-开户成功 5-开户失败 S-开户流程结束
	FailRs        string `mapstructure:"failRs" form:"failRs" json:"failRs"`                      // 失败原因 开户状态为2、5时为必输
}

//OutputParamsCustodyNBBankOpenAccountNotify [开户状态通知 出参]
type OutputParamsCustodyNBBankOpenAccountNotify struct {
	Flag      string `mapstructure:"flag" json:"flag"`           // 成功标识
	Message   string `mapstructure:"message" json:"message"`     // 提示信息
	ApplyNo   string `mapstructure:"applyNo" json:"applyNo"`     // 申请编号
	AccountNo string `mapstructure:"accountNo" json:"accountNo"` // 账号
	OriMsgID  string `mapstructure:"oriMsgId" json:"oriMsgId"`   // 报文标识号
}

//InputParamsCustodyNBBankCashApplyNotify [资金提现通知 入参]
type InputParamsCustodyNBBankCashApplyNotify struct {
	PayAc      string `mapstructure:"payAc" form:"payAc" json:"payAc"`                // 付款账号
	PaeAc      string `mapstructure:"paeAc" form:"paeAc" json:"paeAc"`                // 收款账号
	PaeAcNm    string `mapstructure:"paeAcNm" form:"paeAcNm" json:"paeAcNm"`          // 收款账号名称
	OriApplyNo string `mapstructure:"oriApplyNo" form:"oriApplyNo" json:"oriApplyNo"` // 原申请编号
	TxnAmt     string `mapstructure:"txnAmt" form:"txnAmt" json:"txnAmt"`             // 提现金额
	TxnStatus  string `mapstructure:"txnStatus" form:"txnStatus" json:"txnStatus"`    // 交易状态
	FailRs     string `mapstructure:"failRs" form:"failRs" json:"failRs"`             // 失败原因
}

//OutputParamsCustodyNBBankCashApplyNotify [资金提现通知 出参]
type OutputParamsCustodyNBBankCashApplyNotify struct {
	Flag     string `mapstructure:"flag" json:"flag"`         // 成功标识
	Message  string `mapstructure:"message" json:"message"`   // 提示信息
	OriMsgID string `mapstructure:"oriMsgId" json:"oriMsgId"` // 报文标识号
	PayAc    string `mapstructure:"payAc" json:"payAc"`       // 付款账号
	TxnAmt   string `mapstructure:"txnAmt" json:"txnAmt"`     // 提现金额
}

//CustodyNBBank 银行管存[宁波银行]
type CustodyNBBank interface {
	OpenAccount(ctx context.Context, args *InputParamsCustodyNBBankOpenAccount, reply *OutputParamsCustodyNBBankOpenAccount) error
	BindingAccount(ctx context.Context, args *InputParamsCustodyNBBankBindingAccount, reply *OutputParamsCustodyNBBankBindingAccount) error
	CashApply(ctx context.Context, args *InputParamsCustodyNBBankCashApply, reply *OutputParamsCustodyNBBankCashApply) error
	OpenAccountNotify(ctx context.Context, args string, reply *OutputParamsCustodyNBBankOpenAccountNotify) error
	CashApplyNotify(ctx context.Context, args string, reply *OutputParamsCustodyNBBankCashApplyNotify) error
}
