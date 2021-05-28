package gov

import (
	"context"
)

//InputParamsSendSms InputParamsSendSms
type InputParamsSendSms struct {
	Phone int64 `mapstructure:"phone" form:"phone" json:"phone"` // Phone 手机号
}

//OutputParamsSendSms OutputParamsSendSms
type OutputParamsSendSms struct {
	Content string `mapstructure:"content" json:"Content"` // Content 短信内容
}

//InputParamsOpenAccount InputParamsOpenAccount
type InputParamsOpenAccount struct {
	CompanyOrgNo       int                `mapstructure:"companyOrgNo" form:"companyOrgNo" json:"companyOrgNo"`
	MerchantID         string             `mapstructure:"merchantId" form:"merchantId" json:"merchantId"`
	GovBoscID          int                `mapstructure:"govBoscId" form:"govBoscId" json:"govBoscId"`
	CheckerIDType      int                `mapstructure:"checkerIdType" form:"checkerIdType" json:"checkerIdType"`
	CfoIDVideo         string             `mapstructure:"cfoIdVideo" form:"cfoIdVideo" json:"cfoIdVideo"`
	WorkAddressPhone   int64              `mapstructure:"workAddressPhone" form:"workAddressPhone" json:"workAddressPhone"`
	WorkStrentDoor     string             `mapstructure:"workStrentDoor" form:"workStrentDoor" json:"workStrentDoor"`
	WorkDist           string             `mapstructure:"workDist" form:"workDist" json:"workDist"`
	CheckerIDFromDate  int                `mapstructure:"checkerIdFromDate" form:"checkerIdFromDate" json:"checkerIdFromDate"`
	CfoIDBackImg       string             `mapstructure:"cfoIdBackImg" form:"cfoIdBackImg" json:"cfoIdBackImg"`
	CfoIDFrontImg      string             `mapstructure:"cfoIdFrontImg" form:"cfoIdFrontImg" json:"cfoIdFrontImg"`
	CfoMobile          int64              `mapstructure:"cfoMobile" form:"cfoMobile" json:"cfoMobile"`
	CfoIDLimitDate     int                `mapstructure:"cfoIdLimitDate" form:"cfoIdLimitDate" json:"cfoIdLimitDate"`
	CfoIDFromDate      int                `mapstructure:"cfoIdFromDate" form:"cfoIdFromDate" json:"cfoIdFromDate"`
	CfoName            string             `mapstructure:"cfoName" form:"cfoName" json:"cfoName"`
	CfoIDNo            int                `mapstructure:"cfoIdNo" form:"cfoIdNo" json:"cfoIdNo"`
	CorpIDFromDate     int                `mapstructure:"corpIdFromDate" form:"corpIdFromDate" json:"corpIdFromDate"`
	TransNo            int                `mapstructure:"transNo" form:"transNo" json:"transNo"`
	FlowNo             int64              `mapstructure:"flowNo" form:"flowNo" json:"flowNo"`
	TranType           int                `mapstructure:"tranType" form:"tranType" json:"tranType"`
	AgreementType      string             `mapstructure:"agreementType" form:"agreementType" json:"agreementType"`
	AgreementNumber    string             `mapstructure:"agreementNumber" form:"agreementNumber" json:"agreementNumber"`
	CompanyDesc        string             `mapstructure:"companyDesc" form:"companyDesc" json:"companyDesc"`
	CompanyName        string             `mapstructure:"companyName" form:"companyName" json:"companyName"`
	DelegationImg      string             `mapstructure:"delegationImg" form:"delegationImg" json:"delegationImg"`
	CompanyImg         string             `mapstructure:"companyImg" form:"companyImg" json:"companyImg"`
	ConfirmationDate   int                `mapstructure:"confirmationDate" form:"confirmationDate" json:"confirmationDate"`
	ScopeOfBus         string             `mapstructure:"scopeOfBus" form:"scopeOfBus" json:"scopeOfBus"`
	AgreementMoney     string             `mapstructure:"agreementMoney" form:"agreementMoney" json:"agreementMoney"`
	ComDibilityLimit   string             `mapstructure:"comDibilityLimit" form:"comDibilityLimit" json:"comDibilityLimit"`
	CustType           string             `mapstructure:"custType" form:"custType" json:"custType"`
	Industry           string             `mapstructure:"industry" form:"industry" json:"industry"`
	DepositHuman       string             `mapstructure:"depositHuman" form:"depositHuman" json:"depositHuman"`
	RegDist            string             `mapstructure:"regDist" form:"regDist" json:"regDist"`
	RegStrentDoor      string             `mapstructure:"regStrentDoor" form:"regStrentDoor" json:"regStrentDoor"`
	RegTePhone         string             `mapstructure:"regTePhone" form:"regTePhone" json:"regTePhone"`
	Account            string             `mapstructure:"account" form:"account" json:"account"`
	AcctBank           string             `mapstructure:"acctBank" form:"acctBank" json:"acctBank"`
	AppApplyType       string             `mapstructure:"appApplyType" form:"appApplyType" json:"appApplyType"`
	CfoType            string             `mapstructure:"cfoType" form:"cfoType" json:"cfoType"`
	LegPerID           string             `mapstructure:"legPerId" form:"legPerId" json:"legPerId"`
	CorpName           string             `mapstructure:"corpName" form:"corpName" json:"corpName"`
	CorpIDLimitDate    int                `mapstructure:"corpIdLimitDate" form:"corpIdLimitDate" json:"corpIdLimitDate"`
	CorpMobile         string             `mapstructure:"corpMobile" form:"corpMobile" json:"corpMobile"`
	CorpIDFrontImg     string             `mapstructure:"corpIdFrontImg" form:"corpIdFrontImg" json:"corpIdFrontImg"`
	CorpIDBackImg      string             `mapstructure:"corpIdBackImg" form:"corpIdBackImg" json:"corpIdBackImg"`
	CorpIDVideo        string             `mapstructure:"corpIdVideo" form:"corpIdVideo" json:"corpIdVideo"`
	CheckerIDCard      string             `mapstructure:"checkerIdCard" form:"checkerIdCard" json:"checkerIdCard"`
	CheckerName        string             `mapstructure:"checkerName" form:"checkerName" json:"checkerName"`
	CheckerIDLimitDate int                `mapstructure:"checkerIdLimitDate" form:"checkerIdLimitDate" json:"checkerIdLimitDate"`
	CheckerMobile      string             `mapstructure:"checkerMobile" form:"checkerMobile" json:"checkerMobile"`
	CheckerIDFrontImg  string             `mapstructure:"checkerIdFrontImg" form:"checkerIdFrontImg" json:"checkerIdFrontImg"`
	CheckerIDBackImg   string             `mapstructure:"checkerIdBackImg" form:"checkerIdBackImg" json:"checkerIdBackImg"`
	DynamicCode        string             `mapstructure:"dynamicCode" form:"dynamicCode" json:"dynamicCode"`
	EarningOwnerList   []EarningOwnerList `mapstructure:"earningOwnerList" form:"earningOwnerList" json:"earningOwnerList"`
}

//EarningOwnerList EarningOwnerList
type EarningOwnerList struct {
	EarningOwnerType      string `mapstructure:"earningOwnerType" form:"earningOwnerType" json:"earningOwnerType"`
	EarningOwnerName      string `mapstructure:"earningOwnerName" form:"earningOwnerName" json:"earningOwnerName"`
	EarningOwnerIDType    string `mapstructure:"earningOwnerIdType" form:"earningOwnerIdType" json:"earningOwnerIdType"`
	EarningOwnerIDNo      int64  `mapstructure:"earningOwnerIdNo" form:"earningOwnerIdNo" json:"earningOwnerIdNo"`
	EarningOwnerLimitDate int    `mapstructure:"earningOwnerLimitDate" form:"earningOwnerLimitDate" json:"earningOwnerLimitDate"`
	EarningOwnerCountry   string `mapstructure:"earningOwnerCountry" form:"earningOwnerCountry" json:"earningOwnerCountry"`
	EarningOwnerAddress   string `mapstructure:"earningOwnerAddress" form:"earningOwnerAddress" json:"earningOwnerAddress"`
	EarningOwnerTypeName  string `mapstructure:"earningOwnerTypeName" form:"earningOwnerTypeName" json:"earningOwnerTypeName"`
	EarningOwnerFrontImg  string `mapstructure:"earningOwnerFrontImg" form:"earningOwnerFrontImg" json:"earningOwnerFrontImg"`
	EarningOwnerBackImg   string `mapstructure:"earningOwnerBackImg" form:"earningOwnerBackImg" json:"earningOwnerBackImg"`
}

//OutputParamsOpenAccount OutputParamsOpenAccount
type OutputParamsOpenAccount struct {
}

//InputParamsActiveAccount InputParamsActiveAccount
type InputParamsActiveAccount struct {
	GovBoscID int `mapstructure:"govBoscId" form:"govBoscId" json:"govBoscId"`
}

//OutputParamsActiveAccount OutputParamsActiveAccount
type OutputParamsActiveAccount struct {
	Content string `mapstructure:"content" json:"Content"`
}

//InputParamsApplyContractSign InputParamsApplyContractSign
type InputParamsApplyContractSign struct {
	GovBoscID int `mapstructure:"govBoscId" form:"govBoscId" json:"govBoscId"`
}

//OutputParamsApplyContractSign OutputParamsApplyContractSign
type OutputParamsApplyContractSign struct {
	Content         string `mapstructure:"content" json:"Content"`
	ProjectCode     string `mapstructure:"projectCode" json:"ProjectCode"`
	ProjectCodeName string `mapstructure:"projectCodeName" json:"ProjectCodeName"`
	ContractNo      string `mapstructure:"contractNo" json:"ContractNo"`
}

//InputParamsContractSign InputParamsContractSign
type InputParamsContractSign struct {
	GovBoscID int `mapstructure:"govBoscId" form:"govBoscId" json:"govBoscId"`
	Captcha   int `mapstructure:"captcha" form:"captcha" json:"captcha"` // 短信验证码
}

//OutputParamsContractSign OutputParamsContractSign
type OutputParamsContractSign struct {
	Content string `mapstructure:"content" json:"Content"`
}

//InputParamsOpenAcctNotify InputParamsOpenAcctNotify
type InputParamsOpenAcctNotify struct {
	AgreementNumber int64 `mapstructure:"agreementNumber" form:"agreementNumber" json:"AgreementNumber"`
	AuditAllStat    int64 `mapstructure:"auditAllStat" form:"auditAllStat" json:"AuditAllStat"`
	// AgreementNumber int64 `mapstructure:"AgreementNumber" form:"AgreementNumber" json:"AgreementNumber"`
}

//OutputParamsOpenAcctNotify OutputParamsOpenAcctNotify
type OutputParamsOpenAcctNotify struct {
}

//CustodyBosc 银行管存[上海银行]
type CustodyBosc interface {
	SendSms(ctx context.Context, args *InputParamsSendSms, reply *OutputParamsSendSms) error
	OpenAccount(ctx context.Context, args *InputParamsOpenAccount, reply *OutputParamsOpenAccount) error
	ActiveAccount(ctx context.Context, args *InputParamsActiveAccount, reply *OutputParamsActiveAccount) error
	ApplyContractSign(ctx context.Context, args *InputParamsApplyContractSign, reply *OutputParamsApplyContractSign) error
	ContractSign(ctx context.Context, args *InputParamsContractSign, reply *OutputParamsContractSign) error
	OpenAcctNotify(ctx context.Context, args *InputParamsOpenAcctNotify, reply *OutputParamsOpenAcctNotify) error
}
