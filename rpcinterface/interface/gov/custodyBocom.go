package gov

import (
	"context"
)

//InputParamsUploadImgs InputParamsUploadImgs
type InputParamsUploadImgs struct {
	ApplyFilesA  string `mapstructure:"applyFilesA" form:"applyFilesA" json:"applyFilesA"`    // ApplyFilesA 申请人身份证正反面
	ApplyFilesB  string `mapstructure:"applyFilesB" form:"applyFilesB" json:"applyFilesB"`    // ApplyFilesB 申请人身份证正反面
	MainFiles    string `mapstructure:"mainFiles" form:"mainFiles" json:"mainFiles"`          // MainFiles 营业执照
	CorporFilesA string `mapstructure:"corporFilesA" form:"corporFilesA" json:"corporFilesA"` // corporFilesA 法人身份证正反面
	CorporFilesB string `mapstructure:"corporFilesB" form:"corporFilesB" json:"corporFilesB"` // corporFilesB 法人身份证正反面
}

//OutputParamsUploadImgs OutputParamsUploadImgs
type OutputParamsUploadImgs struct {
	DocumentID string `mapstructure:"documentId" json:"DocumentId"` // Content
}

//交行开户申请入参
type InputParamsOpenAcc struct {
	IntentOpenOrgRegion	 string  	 `mapstructure:"intentOpenOrgRegion" form:"intentOpenOrgRegion" json:"intentOpenOrgRegion"`
	ThdId 				 string  	 `mapstructure:"thdId" form:"thdId" json:"thdId"`
	IntentOpenOrg 		 string  	 `mapstructure:"intentOpenOrg" form:"intentOpenOrg" json:"intentOpenOrg"`
	IntentOpenOrgName	 string  	 `mapstructure:"intentOpenOrgName" form:"intentOpenOrgName" json:"intentOpenOrgName"`
	OpenContactName 	 string  	 `mapstructure:"openContactName" form:"openContactName" json:"openContactName"`
	OpenContactIdType 	 string  	 `mapstructure:"openContactIdType" form:"openContactIdType" json:"openContactIdType"`
	OpenOntactIdNo 		 string  	 `mapstructure:"openOntactIdNo" form:"openOntactIdNo" json:"openOntactIdNo"`
	ExpireDate 			 string  	 `mapstructure:"expireDate" form:"expireDate" json:"expireDate"`
	DeadDate 			 string  	 `mapstructure:"deadDate" form:"deadDate" json:"deadDate"`
	OpenContactMobile 	 string  	 `mapstructure:"openContactMobile" form:"openContactMobile" json:"openContactMobile"`
	EmailAddr 			 string  	 `mapstructure:"emailAddr" form:"emailAddr" json:"emailAddr"`
	OpenContactDocId 	 string  	 `mapstructure:"openContactDocId" form:"openContactDocId" json:"openContactDocId"`
	OpenContactDocidList []BocomDoc	 `mapstructure:"openContactDocidList" form:"openContactDocidList" json:"openContactDocidList"`
	ReserveDate 		 string  	 `mapstructure:"reserveDate" form:"reserveDate" json:"reserveDate"`
	ReserveTime 		 string  	 `mapstructure:"reserveTime" form:"reserveTime" json:"reserveTime"`
	AccountCreateType 	 string  	 `mapstructure:"accountCreateType" form:"accountCreateType" json:"accountCreateType"`
	MainDocid 			 string  	 `mapstructure:"mainDocid" form:"mainDocid" json:"mainDocid"`
	MainDocidList 		 []BocomDoc  `mapstructure:"mainDocidList" form:"mainDocidList" json:"mainDocidList"`
	CorporDocid 		 string  	 `mapstructure:"corporDocid" form:"corporDocid" json:"corporDocid"`
	CorporDocidList 	 []BocomDoc	 `mapstructure:"corporDocidList" form:"corporDocidList" json:"corporDocidList"`
	MainIsLicense 		 string  	 `mapstructure:"mainIsLicense" form:"mainIsLicense" json:"mainIsLicense"`
	UnifiedNo 			 string  	 `mapstructure:"unifiedNo" form:"unifiedNo" json:"unifiedNo"`
	UnitName 			 string  	 `mapstructure:"unitName" form:"unitName" json:"unitName"`
	CertRegisterAddr 	 string  	 `mapstructure:"certRegisterAddr" form:"certRegisterAddr" json:"certRegisterAddr"`
	CertCorpName 		 string  	 `mapstructure:"certCorpName" form:"certCorpName" json:"certCorpName"`
	CertRegMoney 		 string  	 `mapstructure:"certRegMoney" form:"certRegMoney" json:"certRegMoney"`
	CertRegCurrency 	 string  	 `mapstructure:"certRegCurrency" form:"CertRegCurrency" json:"CertRegCurrency"`
	CertBusinessScope 	 string  	 `mapstructure:"certBusinessScope" form:"CertBusinessScope" json:"CertBusinessScope"`
	CertBeginDate 		 string  	 `mapstructure:"certBeginDate" form:"certBeginDate" json:"certBeginDate"`
	CertEndDate 		 string  	 `mapstructure:"certEndDate" form:"certEndDate" json:"certEndDate"`
	CorporMap 			 []CorporMap `mapstructure:"corporMap" form:"corporMap" json:"corporMap"`
	BusiType 			 string  	 `mapstructure:"BusiType" form:"cusiType" json:"busiType"`
	CorresponAddr 		 string  	 `mapstructure:"corresponAddr" form:"corresponAddr" json:"corresponAddr"`
	TraceNo 			 string  	 `mapstructure:"traceNo" form:"traceNo" json:"traceNo"`
	Hold1 				 string  	 `mapstructure:"hold1" form:"hold1" json:"hold1"`
	Hold2 				 string  	 `mapstructure:"hold2" form:"hold2" json:"hold2"`
	GovBocomtId 		 string	 	 `mapstructure:"govBocomtId" form:"govBocomtId" json:"govBocomtId"`
	MerchantId 			 string	 	 `mapstructure:"merchantId" form:"merchantId" json:"merchantId"`
}

//交行开户申请出参
type OutputParamsOpenAcc struct {
	ApplyNo string `json:"apply_no"` //申请号
}

//交行绑卡入参
type InputParamsBindBankCard struct{
	GovBocomtId 	string	`mapstructure:"govBocomtId" json:"govBocomtId"` 		//监管平台开户申请id
	BankCardNo 		string	`mapstructure:"bankCardNo" json:"bankCardNo"`			//结算卡号
	BankAccount 	string	`mapstructure:"bankAccount" json:"bankAccount"`			//结算账户名称
	AccountType 	string	`mapstructure:"accountType" json:"accountType"`			//账户类型 1=对公账户 2=个人账户
	BankAccountType string 	`mapstructure:"bankAccountType" json:"bankAccountType"`	//行内外标志 1=行内 2行外
}

//回调参数
type CallbackParams struct {
	Transaction string 		   `json:"TRANSACTION"`
	Head 		CallbackHead   `json:"HEAD"`
}

//回调Head参数
type CallbackHead struct {
	Sign 		string `json:"sign"`
	ServiceId 	string `json:"service_id"`
	Sender 		string `json:"sender"`
	Receiver 	string `json:"receiver"`
	MsgId 		string `json:"msg_id"`
	RandomKey 	string `json:"random_key"`
	TimeStamp 	string `json:"time_stamp"`
	Version 	string `json:"version"`
}

//开户回调传递信息
type OpenNotifyTransData struct {
	TransCode 		string  `json:"trans_code"`
	ApplyNo 		string  `json:"apply_no"`
	AccountStatus 	string  `json:"account_status"`
	AccountName 	string  `json:"account_name"`
	AccountNo 		string  `json:"account_no"`
}

//开户回调传递信息
type CashNotifyTransData struct {
	TxnAmt 		float64 `json:"txn_amt"`
	PayAc 		string  `json:"pay_ac"`
	PaeAc 		string  `json:"pae_ac"`
	PaeAcNme 	string  `json:"pae_ac_nme"`
	FailRs 		string  `json:"fail_rs"`
	TxnStatus 	string  `json:"txn_status"`
	OriTraceNo 	string  `json:"ori_trace_no"`
}

//CustodyBocom 银行管存[交通银行]
type CustodyBocom interface {
	UploadImgs(ctx context.Context, args *InputParamsUploadImgs, reply *OutputParamsUploadImgs) error
	OpenAccount(ctx context.Context, args *InputParamsUploadImgs, reply *OutputParamsUploadImgs) error
	BindBankCard(ctx context.Context, args *InputParamsUploadImgs, reply *OutputParamsUploadImgs) error
	OpenNotify(ctx context.Context, args *InputParamsUploadImgs, reply *OutputParamsUploadImgs) error
	CashNotify(ctx context.Context, args *InputParamsUploadImgs, reply *OutputParamsUploadImgs) error
}
