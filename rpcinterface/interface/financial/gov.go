package financial

import "context"

type ArgsAnsyGovInfo struct {
	BankType int
	InsuranceType int
	TpayType int
}

type Gov interface {
	//同步监管平台的银行，保险公司，第三方支付
	AnsyGovInfo(ctx context.Context, args *ArgsAnsyGovInfo, reply *bool) error
}