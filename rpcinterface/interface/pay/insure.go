package pay

import "context"

type InsureAcc struct {
	AcctName   string //账户名
	AcctNo     string //账户号
	MerchantId string //子账户会员代码
}

//保险和续保账号
type ReplyGetInsureAcct struct {
	Insure      InsureAcc
	RenewInsure InsureAcc
}

type ArgsGetInsureAcct struct {
	PayChannel    int //支付渠道
	InsureChannel int //保险渠道
}

//保险
type Insure interface {
	//获取平台保险和续在不同支付渠道的开户账号
	GetInsureAcct(ctx context.Context, args *ArgsGetInsureAcct, reply *ReplyGetInsureAcct) error
}
