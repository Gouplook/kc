package gov

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//平安银行绑卡逻辑需要的数据入参
type ArgsPinganBindInfo struct {
	common.BsToken
}

//平安银行绑卡逻辑需要的返回数据
type ReplyPinganBindInfo struct {
	CompanyName      string //企业/商户名称
	CompanyCreditNo  string //企业/商户统一社会信用代码
	MemberName       string //认证会员名称
	MemberGlobalType string //认证会员证件类型
	MemberGlobalId   string //认证会员证件号
	Mobile           string //解绑手机号 //当绑卡成功时有值
}

//平安银行绑卡
type ArgsPinganBindAmount struct {
	common.BsToken
	BindAmountData
}

//平安银行绑卡验证
type ArgsCheckAmountData struct {
	common.BsToken
	CheckAmountData
}

//解绑提现账户
type ArgsUnbindRelateAcct struct {
	common.BsToken
	UnbindRelateAcctData
	Captcha string //验证码
}

type CustodyPingan interface {
	//平安银行绑卡
	PinganBindAmount(ctx context.Context, args ArgsPinganBindAmount, reply *bool) error
	//平安银行绑卡验证
	CheckAmount(ctx context.Context, args ArgsCheckAmountData, reply *bool) error
	//解绑提现账户
	UnbindRelateAcct(ctx context.Context, args ArgsUnbindRelateAcct, reply *bool) error
	//核对提现结果
	GetCashOutResult(ctx context.Context, args string, reply *bool) error
}
